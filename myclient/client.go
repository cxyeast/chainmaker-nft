/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package myclient

import (
	"context"
	"fmt"

	"chainmaker.org/chainmaker/common/v2/crypto"
	"chainmaker.org/chainmaker/common/v2/crypto/asym"
	"chainmaker.org/chainmaker/contract-utils/safemath"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
)

var (
	curClient *sdk.ChainClient
)

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

func CurrentUser(name string) error {

	if curClient == nil {
		return fmt.Errorf("client miss")
	}
	// fmt.Println("user name: ", name)

	seckey, err := hget(name, "seckey")
	if err != nil {
		return err
	}
	// fmt.Println("seckey: ", seckey)

	// pubkey, err := hget(name, "pubkey")
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("pubkey: ", pubkey)

	privateKey, err := asym.PrivateKeyFromPEM([]byte(seckey), nil)
	if err != nil {
		fmt.Println("PrivateKeyFromPEM", err)
		return err
	}

	err = curClient.ChangeSigner(privateKey, nil, crypto.HASH_TYPE_SHA256)
	if err != nil {
		fmt.Println("ChangeSigner", err)
		return err
	}

	return nil
}

func CreateNFTAsset(assetContractName, byteCodePath, version, name, symbol, tokenURI, operatorAdminAddr string) error {
	if curClient == nil {
		return fmt.Errorf("client miss")
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

	return createContract(curClient, assetContractName, version, byteCodePath, kvs)
}

func strByteArrConv(strArr []string) []byte {
	bytearr := make([]byte, 0)
	for i, str := range strArr {
		// fmt.Println(i, str)
		// fmt.Println([]byte(str))
		if i > 0 {
			bytearr = append(bytearr, ([]byte(","))...)
		}
		bytearr = append(bytearr, ([]byte(str))...)
	}
	return bytearr
}

func CreateMultiAsset(assetContractName, byteCodePath, version, uri string, operatorAdminAddr []string) error {
	if curClient == nil {
		return fmt.Errorf("client miss")
	}

	opAdminBytes := strByteArrConv(operatorAdminAddr)

	kvs := []*common.KeyValuePair{
		{
			Key:   "adminAddress",
			Value: opAdminBytes,
		},
		{
			Key:   "uri",
			Value: []byte(uri),
		},
	}

	return createContract(curClient, assetContractName, version, byteCodePath, kvs)
}

func UpdateAsset(assetContractName, byteCodePath, version string) error {
	if curClient == nil {
		return fmt.Errorf("client miss")
	}

	if PubkeyIsAdmin() == false {
		return fmt.Errorf("not chain admin")
	}

	return updateContract(curClient, assetContractName, version, byteCodePath)
}

func FreezeAsset(assetContractName string) error {
	if curClient == nil {
		return fmt.Errorf("client miss")
	}

	if PubkeyIsAdmin() == false {
		return fmt.Errorf("not chain admin")
	}

	return freezeContract(curClient, assetContractName)
}

func UnfreezeAsset(assetContractName string) error {
	if curClient == nil {
		return fmt.Errorf("client miss")
	}

	if PubkeyIsAdmin() == false {
		return fmt.Errorf("not chain admin")
	}

	return unfreezeContract(curClient, assetContractName)
}

func RevokeAsset(assetContractName string) error {
	if curClient == nil {
		return fmt.Errorf("client miss")
	}

	if PubkeyIsAdmin() == false {
		return fmt.Errorf("not chain admin")
	}

	return revokeContract(curClient, assetContractName)
}

func GetNFTAssetAdmins(assetContractName string) string {
	if curClient == nil {
		return ""
	}

	admins, err := queryUserContract(curClient, assetContractName, "GetAdmins", nil)
	if err != nil {
		return ""
	}
	return admins
}

func GetNFTAssetName(assetContractName string) string {
	if curClient == nil {
		return ""
	}

	name, err := queryUserContract(curClient, assetContractName, "name", nil)
	if err != nil {
		return ""
	}
	return name
}

func GetNFTAssetSymbol(assetContractName string) string {
	if curClient == nil {
		return ""
	}

	symbol, err := queryUserContract(curClient, assetContractName, "symbol", nil)
	if err != nil {
		return ""
	}
	return symbol
}

func GetNFTAssetTokenURI(assetContractName, id string) string {
	if curClient == nil {
		return ""
	}

	kvs := []*common.KeyValuePair{
		{
			Key:   "tokenId",
			Value: []byte(id),
		},
	}

	tokenURI, err := queryUserContract(curClient, assetContractName, "tokenURI", kvs)
	if err != nil {
		return ""
	}
	return tokenURI
}

func SetNFTAssetTokenName(assetContractName, name string) (string, error) {
	if curClient == nil {
		return "", fmt.Errorf("client miss")
	}

	kvs := []*common.KeyValuePair{
		{
			Key:   "name",
			Value: []byte(name),
		},
	}

	return invokeUserContract(curClient, assetContractName, "setname", kvs, true)
}

func SetNFTAssetTokenSymbol(assetContractName, symbol string) (string, error) {
	if curClient == nil {
		return "", fmt.Errorf("client miss")
	}

	kvs := []*common.KeyValuePair{
		{
			Key:   "symbol",
			Value: []byte(symbol),
		},
	}

	return invokeUserContract(curClient, assetContractName, "setsymbol", kvs, true)
}

func SetNFTAssetTokenURI(assetContractName, link string) (string, error) {
	if curClient == nil {
		return "", fmt.Errorf("client miss")
	}

	kvs := []*common.KeyValuePair{
		{
			Key:   "uri",
			Value: []byte(link),
		},
	}

	return invokeUserContract(curClient, assetContractName, "setTokenURI", kvs, true)
}

func GetNFTAssetOwner(assetContractName, id string) string {
	if curClient == nil {
		return ""
	}

	kvs := []*common.KeyValuePair{
		{
			Key:   "tokenId",
			Value: []byte(id),
		},
	}

	owner, err := queryUserContract(curClient, assetContractName, "ownerOf", kvs)
	if err != nil {
		return ""
	}
	return owner
}

func GetNFTAssetBalance(assetContractName, account string) string {
	if curClient == nil {
		return ""
	}

	kvs := []*common.KeyValuePair{
		{
			Key:   "account",
			Value: []byte(account),
		},
	}

	balance, err := queryUserContract(curClient, assetContractName, "balanceOf", kvs)
	if err != nil {
		return ""
	}
	return balance
}

func CheckNFTAssetAdmin(assetContractName, addr string) string {
	if curClient == nil {
		return ""
	}

	kvs := []*common.KeyValuePair{
		{
			Key:   "operator",
			Value: []byte(addr),
		},
	}

	check, err := queryUserContract(curClient, assetContractName, "checkadmin", kvs)
	if err != nil {
		return ""
	}
	return check
}

func SetNFTAssetAdmin(assetContractName, addr string) (string, error) {
	if curClient == nil {
		return "", fmt.Errorf("client miss")
	}

	kvs := []*common.KeyValuePair{
		{
			Key:   "operator",
			Value: []byte(addr),
		},
	}

	return invokeUserContract(curClient, assetContractName, "setadmin", kvs, true)
}

func GetNFTAssetSender(assetContractName string) string {
	if curClient == nil {
		return ""
	}

	sender, err := queryUserContract(curClient, assetContractName, "sender", nil)
	if err != nil {
		return ""
	}
	return sender
}

func AlterNFTAssetAdmins(assetContractName string, adminlist []string) (string, error) {
	if curClient == nil {
		return "", fmt.Errorf("client miss")
	}

	adminlistBytes := strByteArrConv(adminlist)

	kvs := []*common.KeyValuePair{
		{
			Key:   "adminAddress",
			Value: adminlistBytes,
		},
	}

	return invokeUserContract(curClient, assetContractName, "AlterAdminAddress", kvs, true)
}

func MintNFTAsset(assetContractName, toEvmAddr, id string, sync bool) (string, error) {
	if curClient == nil {
		return "", fmt.Errorf("client miss")
	}

	_, err := safemath.ParseSafeUint256(id)
	if !err {
		return "", fmt.Errorf("param error, must be uint256")
	}

	kvs := []*common.KeyValuePair{
		{
			Key:   "to",
			Value: []byte(toEvmAddr),
		},
		{
			Key:   "tokenId",
			Value: []byte(id),
		},
	}

	return invokeUserContract(curClient, assetContractName, "mint", kvs, sync)
}

func TransferNFTAsset(assetContractName, toEvmAddr, id string, sync bool) (string, error) {
	if curClient == nil {
		return "", fmt.Errorf("client miss")
	}

	if CurClientEVMAddr() == "" {
		return "", fmt.Errorf("client address not found")
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
			Value: []byte(id),
		},
	}

	return invokeUserContract(curClient, assetContractName, "transferFrom", kvs, sync)
}

func SubscribeContractEvent(ctx context.Context, startBlock, endBlock int64,
	contractName, topic string) (<-chan interface{}, error) {
	if curClient == nil {
		return nil, fmt.Errorf("client miss")
	}

	return curClient.SubscribeContractEvent(ctx, startBlock, endBlock, contractName, topic)
}

func GetBatchMultiAssetBalance(assetContractName string, owners, ids []string) string {
	if curClient == nil {
		return ""
	}

	ownersBytes := strByteArrConv(owners)
	idsBytes := strByteArrConv(ids)

	kvs := []*common.KeyValuePair{
		{
			Key:   "owner",
			Value: ownersBytes,
		},
		{
			Key:   "id",
			Value: idsBytes,
		},
	}

	balances, err := queryUserContract(curClient, assetContractName, "BalanceOfBatch", kvs)
	if err != nil {
		return ""
	}
	return balances
}

func GetMultiAssetBalance(assetContractName, owner, id string) string {
	if curClient == nil {
		return ""
	}

	kvs := []*common.KeyValuePair{
		{
			Key:   "owner",
			Value: []byte(owner),
		},
		{
			Key:   "id",
			Value: []byte(id),
		},
	}

	balance, err := queryUserContract(curClient, assetContractName, "BalanceOf", kvs)
	if err != nil {
		return ""
	}
	return balance
}

func GetMultiAssetUri(assetContractName, id string) string {
	if curClient == nil {
		return ""
	}

	kvs := []*common.KeyValuePair{
		{
			Key:   "id",
			Value: []byte(id),
		},
	}

	balance, err := queryUserContract(curClient, assetContractName, "Uri", kvs)
	if err != nil {
		return ""
	}
	return balance
}

func SetMultiAssetTokenURI(assetContractName, link string) (string, error) {
	if curClient == nil {
		return "", fmt.Errorf("client miss")
	}

	kvs := []*common.KeyValuePair{
		{
			Key:   "uri",
			Value: []byte(link),
		},
	}

	return invokeUserContract(curClient, assetContractName, "SetUri", kvs, true)
}

func GetMultiAdmins(assetContractName string) string {
	if curClient == nil {
		return ""
	}

	admins, err := queryUserContract(curClient, assetContractName, "GetAdmins", nil)
	if err != nil {
		return ""
	}
	return admins
}

func MintMultiAsset(assetContractName, to, id, amount string, sync bool) (string, error) {
	if curClient == nil {
		return "", fmt.Errorf("client miss")
	}

	_, err := safemath.ParseSafeUint256(id)
	if !err {
		return "", fmt.Errorf("param error, id must be uint256")
	}

	_, err = safemath.ParseSafeUint256(amount)
	if !err {
		return "", fmt.Errorf("param error, amount must be uint256")
	}

	kvs := []*common.KeyValuePair{
		{
			Key:   "to",
			Value: []byte(to),
		},
		{
			Key:   "id",
			Value: []byte(id),
		},
		{
			Key:   "amount",
			Value: []byte(amount),
		},
	}

	return invokeUserContract(curClient, assetContractName, "Mint", kvs, sync)
}

func BatchMintMultiAsset(assetContractName, to string, ids, amounts []string, sync bool) (string, error) {
	if curClient == nil {
		return "", fmt.Errorf("client miss")
	}

	for _, id := range ids {
		_, err := safemath.ParseSafeUint256(id)
		if !err {
			return "", fmt.Errorf("param error, id must be uint256")
		}
	}

	for _, amount := range amounts {
		_, err := safemath.ParseSafeUint256(amount)
		if !err {
			return "", fmt.Errorf("param error, amount must be uint256")
		}
	}

	idsbytes := strByteArrConv(ids)
	amountsbytes := strByteArrConv(amounts)

	kvs := []*common.KeyValuePair{
		{
			Key:   "to",
			Value: []byte(to),
		},
		{
			Key:   "ids",
			Value: idsbytes,
		},
		{
			Key:   "amounts",
			Value: amountsbytes,
		},
	}

	return invokeUserContract(curClient, assetContractName, "MintBatchNormal", kvs, sync)
}

func TransferMultiAsset(assetContractName, from, to, tokenid, amount string, sync bool) (string, error) {
	if curClient == nil {
		return "", fmt.Errorf("client miss")
	}

	kvs := []*common.KeyValuePair{
		{
			Key:   "from",
			Value: []byte(from),
		},
		{
			Key:   "to",
			Value: []byte(to),
		},
		{
			Key:   "id",
			Value: []byte(tokenid),
		},
		{
			Key:   "amount",
			Value: []byte(amount),
		},
	}

	return invokeUserContract(curClient, assetContractName, "SafeTransferFrom", kvs, sync)
}

func BatchTransferMultiAsset(assetContractName, from, to string, ids, amounts []string, sync bool) (string, error) {
	if curClient == nil {
		return "", fmt.Errorf("client miss")
	}

	idsBytes := strByteArrConv(ids)
	amountsBytes := strByteArrConv(amounts)

	kvs := []*common.KeyValuePair{
		{
			Key:   "from",
			Value: []byte(from),
		},
		{
			Key:   "to",
			Value: []byte(to),
		},
		{
			Key:   "ids",
			Value: idsBytes,
		},
		{
			Key:   "amounts",
			Value: amountsBytes,
		},
	}

	return invokeUserContract(curClient, assetContractName, "SafeBatchTransferFrom", kvs, sync)
}

func AlterMultiAssetAdmins(assetContractName string, adminlist []string) (string, error) {
	if curClient == nil {
		return "", fmt.Errorf("client miss")
	}

	adminlistBytes := strByteArrConv(adminlist)

	kvs := []*common.KeyValuePair{
		{
			Key:   "adminAddress",
			Value: adminlistBytes,
		},
	}

	return invokeUserContract(curClient, assetContractName, "AlterAdminAddress", kvs, true)
}

func GetBlockHeightByTxId(txid string) (uint64, error) {
	if curClient == nil {
		return 0, fmt.Errorf("client miss")
	}

	height, err := curClient.GetBlockHeightByTxId(txid)

	return height, err
}
