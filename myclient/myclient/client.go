/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package myclient

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"chainmaker.org/chainmaker/common/v2/crypto"
	"chainmaker.org/chainmaker/common/v2/crypto/asym"
	"chainmaker.org/chainmaker/pb-go/v2/accesscontrol"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	sdkutils "chainmaker.org/chainmaker/sdk-go/v2/utils"
)

var (
	assetContractName string
	curClient         *sdk.ChainClient
)

func SetAssetContractName(name string) {
	assetContractName = name
}

func GetAssetContractName() string {
	return assetContractName
}

func GeneratePKFile(filePath, fileName string) error {
	keyName := fmt.Sprintf("%s.key", fileName)
	pemName := fmt.Sprintf("%s.pem", fileName)
	keyType := crypto.AsymAlgoMap[strings.ToUpper(crypto.CRYPTO_ALGO_ECC_P256)]

	privKeyPEM, pubKeyPEM, err := asym.GenerateKeyPairPEM(keyType)
	if err != nil {
		return fmt.Errorf("generate key pair failed, %s", err.Error())
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

	if err = makeFile(filePath, keyName); err != nil {
		return fmt.Errorf("make key file failed, %s", err.Error())
	}

	if err = ioutil.WriteFile(
		filepath.Join(filePath, keyName), []byte(privKeyPEM), 0600,
	); err != nil {
		return fmt.Errorf("write key file failed, %s", err.Error())
	}

	if err = makeFile(filePath, pemName); err != nil {
		return fmt.Errorf("create pubkey file failed, %s", err.Error())
	}

	if err = ioutil.WriteFile(
		filepath.Join(filePath, pemName), []byte(pubKeyPEM), 0600,
	); err != nil {
		return fmt.Errorf("write pubkey file failed, %s", err.Error())
	}

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

func CurClientPK() string {

	if curClient == nil {
		return ""
	}

	keyPemStr, _ := curClient.GetPublicKey().String()

	return keyPemStr
}

func CurClientEVMAddr() string {

	if curClient == nil {
		return ""
	}

	keyPemStr, _ := curClient.GetPublicKey().String()
	addrEvm, _ := sdk.GetEVMAddressFromPKPEM(keyPemStr, crypto.HASH_TYPE_SHA256)

	return addrEvm
}

func SetChainClientWithSDKConf(confPath, name string) error {
	dest := confPath + "/sdk_config_pk_" + name + ".yml"
	client, err := sdk.NewChainClient(
		sdk.WithConfPath(dest),
	)
	if err != nil {
		curClient = nil
	} else {
		curClient = client
	}

	return err
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

func CreateNFTAsset(byteCodePath, version, name, symbol, tokenURI, operatorAdminAddr string) error {
	if curClient == nil {
		return fmt.Errorf("client miss")
	}

	if PubkeyIsAdmin() == false {
		return fmt.Errorf("not chain admin")
	}

	kvs := []*common.KeyValuePair{
		{
			Key:   "name",
			Value: []byte(name),
		},
		{
			Key:   "symbol",
			Value: []byte(symbol),
		},
		{
			Key:   "tokenURI",
			Value: []byte(tokenURI),
		},
		{
			Key:   "admin",
			Value: []byte(operatorAdminAddr),
		},
	}

	payload, err := curClient.CreateContractCreatePayload(assetContractName, version, byteCodePath, common.RuntimeType_DOCKER_GO, kvs)
	if err != nil {
		return err
	}

	endorsers, err := getAdminEndorsers(curClient.GetHashType(), payload)
	if err != nil {
		return err
	}

	resp, err := curClient.SendContractManageRequest(payload, endorsers, 5, true)
	if err != nil {
		return err
	}

	err = CheckProposalRequestResp(resp, true)
	if err != nil {
		return err
	}

	return nil

}

func UpdateNFTAsset(byteCodePath, version string) error {
	if curClient == nil {
		return fmt.Errorf("client miss")
	}

	if PubkeyIsAdmin() == false {
		return fmt.Errorf("not chain admin")
	}

	payload, err := curClient.CreateContractUpgradePayload(assetContractName, version, byteCodePath, common.RuntimeType_DOCKER_GO, nil)
	if err != nil {
		return err
	}

	endorsers, err := getAdminEndorsers(curClient.GetHashType(), payload)
	if err != nil {
		return err
	}

	resp, err := curClient.SendContractManageRequest(payload, endorsers, 5, true)
	if err != nil {
		return err
	}

	err = CheckProposalRequestResp(resp, true)
	if err != nil {
		return err
	}

	return nil

}

func FreezeNFTAsset() error {
	if curClient == nil {
		return fmt.Errorf("client miss")
	}

	if PubkeyIsAdmin() == false {
		return fmt.Errorf("not chain admin")
	}

	payload, err := curClient.CreateContractFreezePayload(assetContractName)
	if err != nil {
		return err
	}

	endorsers, err := getAdminEndorsers(curClient.GetHashType(), payload)
	if err != nil {
		return err
	}

	resp, err := curClient.SendContractManageRequest(payload, endorsers, 5, true)
	if err != nil {
		return err
	}

	err = CheckProposalRequestResp(resp, true)
	if err != nil {
		return err
	}

	return nil

}

func UnfreezeNFTAsset() error {
	if curClient == nil {
		return fmt.Errorf("client miss")
	}

	if PubkeyIsAdmin() == false {
		return fmt.Errorf("not chain admin")
	}

	payload, err := curClient.CreateContractUnfreezePayload(assetContractName)
	if err != nil {
		return err
	}

	endorsers, err := getAdminEndorsers(curClient.GetHashType(), payload)
	if err != nil {
		return err
	}

	resp, err := curClient.SendContractManageRequest(payload, endorsers, 5, true)
	if err != nil {
		return err
	}

	err = CheckProposalRequestResp(resp, true)
	if err != nil {
		return err
	}

	return nil

}

func RevokeNFTAsset() error {
	if curClient == nil {
		return fmt.Errorf("client miss")
	}

	if PubkeyIsAdmin() == false {
		return fmt.Errorf("not chain admin")
	}

	payload, err := curClient.CreateContractRevokePayload(assetContractName)
	if err != nil {
		return err
	}

	endorsers, err := getAdminEndorsers(curClient.GetHashType(), payload)
	if err != nil {
		return err
	}

	resp, err := curClient.SendContractManageRequest(payload, endorsers, 5, true)
	if err != nil {
		return err
	}

	err = CheckProposalRequestResp(resp, true)
	if err != nil {
		return err
	}

	return nil

}

func GetName() string {
	name, err := queryUserContract("name", nil)
	if err != nil {
		return ""
	}
	return name
}

func GetSymbol() string {
	symbol, err := queryUserContract("symbol", nil)
	if err != nil {
		return ""
	}
	return symbol
}

func GetTokenURI() string {
	tokenURI, err := queryUserContract("tokenURI", nil)
	if err != nil {
		return ""
	}
	return tokenURI
}

func GetOwner(tokenId string) string {

	kvs := []*common.KeyValuePair{
		{
			Key:   "tokenId",
			Value: []byte(tokenId),
		},
	}

	owner, err := queryUserContract("ownerOf", kvs)
	if err != nil {
		return ""
	}
	return owner
}

func GetBalance(address string) string {

	kvs := []*common.KeyValuePair{
		{
			Key:   "account",
			Value: []byte(address),
		},
	}

	balance, err := queryUserContract("balanceOf", kvs)
	if err != nil {
		return ""
	}
	return balance
}

func CheckAdmin(addr string) string {
	kvs := []*common.KeyValuePair{
		{
			Key:   "operator",
			Value: []byte(addr),
		},
	}

	check, err := queryUserContract("checkadmin", kvs)
	if err != nil {
		return ""
	}
	return check
}

func SetAdmin(addr string) error {
	kvs := []*common.KeyValuePair{
		{
			Key:   "operator",
			Value: []byte(addr),
		},
	}

	return invokeUserContract(assetContractName, "setadmin", "", kvs, true)
}

func GetSender() string {
	sender, err := queryUserContract("sender", nil)
	if err != nil {
		return ""
	}
	return sender
}

func MintAsset(toEvmAddr, tokenId string) error {

	kvs := []*common.KeyValuePair{
		{
			Key:   "to",
			Value: []byte(toEvmAddr),
		},
		{
			Key:   "tokenId",
			Value: []byte(tokenId),
		},
	}

	return invokeUserContract(assetContractName, "mint", "", kvs, true)

}

func TransferAsset(toEvmAddr, tokenId string) error {
	if curClient == nil {
		return fmt.Errorf("client miss")
	}

	if CurClientEVMAddr() == "" {
		return fmt.Errorf("client address not found")
	}

	kvs := []*common.KeyValuePair{
		{
			Key:   "from",
			Value: []byte(CurClientEVMAddr()),
		},
		{
			Key:   "to",
			Value: []byte(toEvmAddr),
		},
		{
			Key:   "tokenId",
			Value: []byte(tokenId),
		},
	}

	return invokeUserContract(assetContractName, "transferFrom", "", kvs, true)
}

func queryUserContract(method string, kvs []*common.KeyValuePair) (string, error) {
	if curClient == nil {
		return "", fmt.Errorf("client miss")
	}

	resp, err := curClient.QueryContract(assetContractName, method, kvs, -1)
	if err != nil {
		return "", err
	}

	err = CheckProposalRequestResp(resp, true)
	if err != nil {
		return "", err
	}
	return string(resp.ContractResult.Result), nil
}

func invokeUserContract(contractName, method, txId string,
	kvs []*common.KeyValuePair, withSyncResult bool) error {

	if curClient == nil {
		return fmt.Errorf("client miss")
	}

	resp, err := curClient.InvokeContract(contractName, method, txId, kvs, -1, withSyncResult)
	if err != nil {
		return err
	}

	if resp.Code != common.TxStatusCode_SUCCESS {
		return fmt.Errorf("invoke contract failed, [code:%d]/[msg:%s]\n", resp.Code, resp.Message)
	}

	if !withSyncResult {
		fmt.Printf("invoke contract success, resp: [code:%d]/[msg:%s]/[txId:%s]\n", resp.Code, resp.Message, resp.ContractResult.Result)
	} else {
		fmt.Printf("invoke contract success, resp: [code:%d]/[msg:%s]/[contractResult:%s]\n", resp.Code, resp.Message, resp.ContractResult)
	}

	return nil
}
