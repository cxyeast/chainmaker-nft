package myclient

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"chainmaker.org/chainmaker/common/v2/crypto"
	"chainmaker.org/chainmaker/common/v2/crypto/asym"
	"chainmaker.org/chainmaker/pb-go/v2/accesscontrol"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	sdkutils "chainmaker.org/chainmaker/sdk-go/v2/utils"
	"chainmaker.org/chainmaker/utils/v2"
	"github.com/hokaccha/go-prettyjson"
)

var adminPkUsers map[string]*pkUserInfo

type pkUserInfo struct {
	keyPem     []byte
	memberInfo string
}

func AdminExisted(username string) bool {
	if adminPkUsers == nil {
		return false
	}
	_, ok := adminPkUsers[username]
	return ok
}

func SetAdminPkUser(username, keyPath string) error {
	if adminPkUsers == nil {
		adminPkUsers = make(map[string]*pkUserInfo)
	}

	keyPem, err := ioutil.ReadFile(keyPath + "/" + username + ".key")
	if err != nil {
		return fmt.Errorf("read key file failed, %s", err)
	}

	err = hset(username, "seckey", string(keyPem))
	if err != nil {
		return err
	}

	key, err := asym.PrivateKeyFromPEM(keyPem, nil)
	if err != nil {
		return fmt.Errorf("get private key failed, %s", err)
	}

	pubKey := key.PublicKey()
	memberInfo, err := pubKey.String()
	if err != nil {
		return fmt.Errorf("get pubkey info failed, %s", err)
	}

	err = hset(username, "pubkey", memberInfo)
	if err != nil {
		return err
	}

	pUserInfo := new(pkUserInfo)
	pUserInfo.keyPem = keyPem
	pUserInfo.memberInfo = memberInfo
	adminPkUsers[username] = pUserInfo

	return nil
}

func GetAdminPkUser(username string) (string, string, error) {
	pUserInfo := adminPkUsers[username]
	if pUserInfo == nil {
		return "", "", fmt.Errorf("non-existed")
	}

	return string(pUserInfo.keyPem), pUserInfo.memberInfo, nil
}

func DelAdminPkUser(username string) {
	if AdminExisted(username) {
		delete(adminPkUsers, username)
	}
}

func NumAdminPkUsers() int {
	if adminPkUsers == nil {
		return 0
	} else {
		return len(adminPkUsers)
	}
}

func ClearAdminPkUsers() {
	if adminPkUsers != nil {
		for name, _ := range adminPkUsers {
			delete(adminPkUsers, name)
		}
	}
}

func PubkeyIsAdmin() bool {
	if adminPkUsers != nil {
		for name, _ := range adminPkUsers {
			if CurClientPK() == adminPkUsers[name].memberInfo {
				return true
			}
		}
	}
	return false
}

func makeFile(filePath, fileName string) error {
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return fmt.Errorf("make dir failed, %s", err.Error())
	}

	f, err := os.Create(filepath.Join(filePath, fileName))
	if err != nil {
		return fmt.Errorf("create file failed, %s", err.Error())
	}

	if err = f.Close(); err != nil {
		return err
	}

	return nil
}

func printPrettyJson(data interface{}) {
	output, err := prettyjson.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
}

func CheckProposalRequestResp(resp *common.TxResponse, needContractResult bool) error {
	if resp.Code != common.TxStatusCode_SUCCESS {
		if resp.Message == "" {
			resp.Message = resp.Code.String()
		}
		return errors.New(resp.Message)
	}

	if needContractResult && resp.ContractResult == nil {
		return fmt.Errorf("contract result is nil")
	}

	if resp.ContractResult != nil && resp.ContractResult.Code != 0 {
		return errors.New(resp.ContractResult.Message)
	}

	return nil
}

func GenerateUser(username string) error {
	keyType := crypto.AsymAlgoMap[strings.ToUpper(crypto.CRYPTO_ALGO_ECC_P256)]

	privKeyPEM, pubKeyPEM, err := asym.GenerateKeyPairPEM(keyType)
	if err != nil {
		return fmt.Errorf("generate key pair failed, %s", err.Error())
	}

	err = hset(username, "pubkey", pubKeyPEM)
	if err != nil {
		return err
	}

	err = hset(username, "seckey", privKeyPEM)
	if err != nil {
		return err
	}

	digest := sha256.Sum256([]byte("Blockchain"))

	sign, err := asym.Sign(privKeyPEM, digest[:])
	if err != nil {
		return fmt.Errorf("Sign Error!")
	}

	ok, err := asym.Verify(pubKeyPEM, digest[:], sign)
	if err != nil || !ok {
		return fmt.Errorf("Verifying Error!")
	}

	// if err = makeFile(filePath, keyName); err != nil {
	// 	return fmt.Errorf("make key file failed, %s", err.Error())
	// }

	// if err = ioutil.WriteFile(
	// 	filepath.Join(filePath, keyName), []byte(privKeyPEM), 0600,
	// ); err != nil {
	// 	return fmt.Errorf("write key file failed, %s", err.Error())
	// }

	// if err = makeFile(filePath, pemName); err != nil {
	// 	return fmt.Errorf("create pubkey file failed, %s", err.Error())
	// }

	// if err = ioutil.WriteFile(
	// 	filepath.Join(filePath, pemName), []byte(pubKeyPEM), 0600,
	// ); err != nil {
	// 	return fmt.Errorf("write pubkey file failed, %s", err.Error())
	// }

	return nil
}

func CalcEVMAddr(path, name string) (string, error) {

	target := path + "/" + name + ".pem"

	pkBlock, err := ioutil.ReadFile(target)
	if err != nil {
		return "", err
	}

	addrEvm, err := sdk.GetEVMAddressFromPKPEM(string(pkBlock), crypto.HASH_TYPE_SHA256)
	if err != nil {
		return "", fmt.Errorf("generate addr failed, %s", err.Error())
	}

	return addrEvm, nil
}

func getAdminEndorsers(hashType crypto.HashType, payload *common.Payload) ([]*common.EndorsementEntry, error) {
	var endorsers []*common.EndorsementEntry

	if len(adminPkUsers) == 0 {
		return endorsers, fmt.Errorf("no admin users available!")
	}

	for _, pkUserInfo := range adminPkUsers {
		entry, err := sdkutils.MakeEndorser("", crypto.HASH_TYPE_SHA256, accesscontrol.MemberType_PUBLIC_KEY, pkUserInfo.keyPem,
			[]byte(pkUserInfo.memberInfo), payload)
		if err != nil {
			return nil, err
		}
		endorsers = append(endorsers, entry)
	}
	return endorsers, nil
}

func getSingleEndorser(hashType crypto.HashType, payload *common.Payload, signKeyPath string) ([]*common.EndorsementEntry, error) {
	var endorsers []*common.EndorsementEntry

	entry, err := sdkutils.MakePkEndorserWithPath(signKeyPath, hashType, "", payload)
	if err != nil {
		return nil, err
	}
	endorsers = append(endorsers, entry)

	return endorsers, nil
}

func createContract(client *sdk.ChainClient, assetContractName, version, byteCodePath string, kvs []*common.KeyValuePair) error {

	if PubkeyIsAdmin() == false {
		return fmt.Errorf("not chain admin")
	}

	payload, err := client.CreateContractCreatePayload(assetContractName, version, byteCodePath, common.RuntimeType_DOCKER_GO, kvs)
	if err != nil {
		return err
	}

	endorsers, err := getAdminEndorsers(client.GetHashType(), payload)
	if err != nil {
		return err
	}

	resp, err := client.SendContractManageRequest(payload, endorsers, 5, true)
	if err != nil {
		return err
	}

	err = CheckProposalRequestResp(resp, true)
	if err != nil {
		return err
	}

	return nil
}

func updateContract(client *sdk.ChainClient, assetContractName, version, byteCodePath string) error {

	if PubkeyIsAdmin() == false {
		return fmt.Errorf("not chain admin")
	}

	payload, err := client.CreateContractUpgradePayload(assetContractName, version, byteCodePath, common.RuntimeType_DOCKER_GO, nil)
	if err != nil {
		return err
	}

	endorsers, err := getAdminEndorsers(client.GetHashType(), payload)
	if err != nil {
		return err
	}

	resp, err := client.SendContractManageRequest(payload, endorsers, 5, true)
	if err != nil {
		return err
	}

	err = CheckProposalRequestResp(resp, true)
	if err != nil {
		return err
	}

	return nil
}

func freezeContract(client *sdk.ChainClient, assetContractName string) error {

	if PubkeyIsAdmin() == false {
		return fmt.Errorf("not chain admin")
	}

	payload, err := client.CreateContractFreezePayload(assetContractName)
	if err != nil {
		return err
	}

	endorsers, err := getAdminEndorsers(client.GetHashType(), payload)
	if err != nil {
		return err
	}

	resp, err := client.SendContractManageRequest(payload, endorsers, 5, true)
	if err != nil {
		return err
	}

	err = CheckProposalRequestResp(resp, true)
	if err != nil {
		return err
	}

	return nil
}

func unfreezeContract(client *sdk.ChainClient, assetContractName string) error {

	if PubkeyIsAdmin() == false {
		return fmt.Errorf("not chain admin")
	}

	payload, err := client.CreateContractUnfreezePayload(assetContractName)
	if err != nil {
		return err
	}

	endorsers, err := getAdminEndorsers(client.GetHashType(), payload)
	if err != nil {
		return err
	}

	resp, err := client.SendContractManageRequest(payload, endorsers, 5, true)
	if err != nil {
		return err
	}

	err = CheckProposalRequestResp(resp, true)
	if err != nil {
		return err
	}

	return nil
}

func revokeContract(client *sdk.ChainClient, assetContractName string) error {

	if PubkeyIsAdmin() == false {
		return fmt.Errorf("not chain admin")
	}

	payload, err := client.CreateContractRevokePayload(assetContractName)
	if err != nil {
		return err
	}

	endorsers, err := getAdminEndorsers(client.GetHashType(), payload)
	if err != nil {
		return err
	}

	resp, err := client.SendContractManageRequest(payload, endorsers, 5, true)
	if err != nil {
		return err
	}

	err = CheckProposalRequestResp(resp, true)
	if err != nil {
		return err
	}

	return nil
}

func queryUserContract(client *sdk.ChainClient, assetContractName, method string, kvs []*common.KeyValuePair) (string, error) {

	resp, err := client.QueryContract(assetContractName, method, kvs, -1)
	if err != nil {
		return "", err
	}

	err = CheckProposalRequestResp(resp, true)
	if err != nil {
		return "", err
	}
	return string(resp.ContractResult.Result), nil
}

func invokeUserContract(client *sdk.ChainClient, contractName, method string,
	kvs []*common.KeyValuePair, withSyncResult bool) (string, error) {

	txId := utils.GetTimestampTxId()
	// height, err := client.GetBlockHeightByTxId(txId)
	// fmt.Println("1st txId: %d, block height: %v, err: %v", txId, height, err)
	resp, err := client.InvokeContract(contractName, method, txId, kvs, -1, withSyncResult)
	if err != nil {
		return txId, err
	}

	if resp.Code != common.TxStatusCode_SUCCESS {
		return "", fmt.Errorf("invoke contract failed, [code:%d]/[msg:%s]\n", resp.Code, resp.Message)
	}

	if !withSyncResult {
		// fmt.Printf("invoke contract success, resp: [code:%d]/[msg:%s]/[txId:%s]\n", resp.Code, resp.Message, resp.ContractResult.Result)
	} else {
		fmt.Printf("invoke contract success, resp: [code:%d]/[msg:%s]/[contractResult:%s]\n", resp.Code, resp.Message, resp.ContractResult)
	}

	// height, err = client.GetBlockHeightByTxId(txId)
	// fmt.Println("2nd txId: %d, block height: %v, err: %v", txId, height, err)

	return txId, nil
}
