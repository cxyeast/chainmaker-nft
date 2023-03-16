/**
  Copyright (C) BABEC. All rights reserved.

  SPDX-License-Identifier: Apache-2.0
*/

/**
  @title ERC-1155 Multi Token Standard
  @dev See https://eips.ethereum.org/EIPS/eip-1155
  Note: The ERC-165 identifier for this interface is 0xd9b67a26.
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"chainmaker.org/chainmaker/common/v2/crypto"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sandbox"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
	"chainmaker.org/chainmaker/contract-utils/address"
	"chainmaker.org/chainmaker/contract-utils/safemath"
)

var (
	maps = mapManage{}
)

const (
	// use for approve true
	trueString = "1"
	// use for approve false
	falseString = "0"
)
const (
	keyAdminAddress   = "adminAddress"
	keyUri            = "uri"
	paramAdminAddress = "adminAddress"
	paramAmount       = "amount"
	paramAmounts      = "amounts"
	paramApproved     = "approved"
	paramData         = "data"
	paramFrom         = "from"
	paramTo           = "to"
	paramId           = "id"
	paramIds          = "ids"
	paramIdStart      = "idStart"
	paramOperator     = "operator"
	paramOwner        = "owner"
	paramUri          = "uri"
)

var _ IERC1155 = (*ERC1155Contract)(nil)

// ERC1155Contract erc1155的一种实现
type ERC1155Contract struct {
}

// InitContract used to deploy contract
func (e *ERC1155Contract) InitContract() protogo.Response {
	args := sdk.Instance.GetArgs()
	adminAddress := args[paramAdminAddress]
	uri := string(args[paramUri])
	var adminAddressStr string
	if len(adminAddress) == 0 {
		adminAddressStr, _ = sdk.Instance.Origin()
	} else {
		adminAddressStr = string(adminAddress)
	}
	adminAddresses := strings.Split(adminAddressStr, ",")
	resp := e.initContract(adminAddresses, uri)
	if resp.Status != sdk.OK {
		sdk.Instance.Warnf(resp.Message)
	}
	return resp
}

func (e *ERC1155Contract) initContract(adminAddresses []string, uri string) protogo.Response {
	if !address.IsValidAddress(adminAddresses...) {
		return sdk.Error("address error")
	}
	contractInfo := maps.getErc1155Map()
	_ = contractInfo.Set([]string{keyUri}, []byte(uri))

	adminAddressByte, _ := json.Marshal(adminAddresses)
	err := contractInfo.Set([]string{keyAdminAddress}, adminAddressByte)
	if err != nil {
		return sdk.Error("set admin address of contractInfo failed")
	}
	err = contractInfo.Set([]string{"userCount"}, []byte("0"))
	if err != nil {
		return sdk.Error("set user count of contractInfo failed")
	}
	sdk.Instance.EmitEvent("AlterAdminAddress", adminAddresses)
	return sdk.Success([]byte("Init contract success"))
}

// UpgradeContract used to upgrade contract
func (e *ERC1155Contract) UpgradeContract() protogo.Response {
	return sdk.Success([]byte("Upgrade contract success"))
}

// nolint
// InvokeContract used to invoke user contract
func (e *ERC1155Contract) InvokeContract(method string) (resp protogo.Response) {
	args := sdk.Instance.GetArgs()
	if len(method) == 0 {
		return sdk.Error("method of param should not be empty")
	}
	defer func() {
		if resp.Status != sdk.OK {
			sdk.Instance.Warnf(resp.Message)
		}
	}()

	switch method {
	case "Mint":
		to := string(args[paramTo])
		idStr := string(args[paramId])
		data := string(args[paramData])
		amountStr := string(args[paramAmount])
		to = address.GetCleanAddr(to)

		id, flag1 := safemath.ParseSafeUint256(idStr)
		if !flag1 {
			return sdk.Error("param error")
		}
		amount, flag2 := safemath.ParseSafeUint256(amountStr)
		if !flag2 {
			return sdk.Error("param error")
		}

		return e.Mint(Address(to), id, amount, data)
	case "MintBatchNormal":
		to := string(args[paramTo])
		ids := string(args[paramIds])
		amounts := string(args[paramAmounts])
		data := string(args[paramData])
		to = address.GetCleanAddr(to)

		return e.MintBatchNormal(Address(to), ids, amounts, data)
	case "MintBatchNft":
		to := string(args[paramTo])
		idStartStr := string(args[paramIdStart])
		data := string(args[paramData])
		amountStr := string(args[paramAmount])
		to = address.GetCleanAddr(to)

		idStart, flag1 := safemath.ParseSafeUint256(idStartStr)
		if !flag1 {
			return sdk.Error("param idStart error")
		}
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			return sdk.Error("param amount error")
		}

		return e.MintBatchNft(Address(to), idStart, amount, data)
	case "AlterAdminAddress":
		address := args[paramAdminAddress]
		var addresses []string
		if len(address) != 0 {
			addresses = strings.Split(string(address), ",")
		}

		return e.AlterAdminAddress(addresses)
	case "Uri":
		idStr := string(args[paramId])
		id, _ := safemath.ParseSafeUint256(idStr)

		return e.GetUri(id)
	case "GetAdmins":
		return e.GetAdminAddress()
	case "SetUri":
		uri := string(args[paramUri])

		return e.SetUri(uri)
	case "OwnerOf":
		idStr := string(args[paramId])

		id, ok := safemath.ParseSafeUint256(idStr)
		if !ok {
			return sdk.Error("param error")
		}

		addr := e.OwnerOf(id)
		return sdk.Success([]byte(addr))
	case "BalanceOf":
		owner := string(args[paramOwner])
		idStr := string(args[paramId])
		owner = address.GetCleanAddr(owner)

		id, err := safemath.ParseSafeUint256(idStr)
		if !err {
			return sdk.Error("parse id failed")
		}

		return e.BalanceOf(Address(owner), id)
	case "BalanceOfBatch":
		ownersStr := string(args[paramOwner])
		idStrs := string(args[paramId])

		o := strings.Split(ownersStr, ",")
		ids := strings.Split(idStrs, ",")

		if len(ids) != len(o) {
			return sdk.Error("param error")
		}
		var tokenIds = make([]*safemath.SafeUint256, 0)
		var owners = make([]Address, 0)
		for i, id := range ids {
			tokenId, err := safemath.ParseSafeUint256(id)
			if !err {
				return sdk.Error("parse id failed")
			}
			tokenIds = append(tokenIds, tokenId)
			owners = append(owners, Address(o[i]))
		}

		return e.BalanceOfBatch(owners, tokenIds)
	case "SetApprovalForAll":
		approvedStr := string(args[paramApproved])
		operatorStr := string(args[paramOperator])
		operatorStr = address.GetCleanAddr(operatorStr)

		approved, err := strconv.ParseBool(approvedStr)
		if err != nil {
			return sdk.Error("parse approved failed" + err.Error())
		}
		senderStr, _ := sdk.Instance.Sender()

		return e.SetApprovalForAll(Address(senderStr), Address(operatorStr), approved)
	case "IsApprovedForAll":
		owner := string(args[paramOwner])
		operator := string(args[paramOperator])

		owner = address.GetCleanAddr(owner)
		operator = address.GetCleanAddr(operator)

		return e.IsApprovedForAll(Address(owner), Address(operator))
	case "SafeTransferFrom":
		from := string(args[paramFrom])
		to := string(args[paramTo])
		idStr := string(args[paramId])
		amountStr := string(args[paramAmount])
		data := string(args[paramData])

		from = address.GetCleanAddr(from)
		to = address.GetCleanAddr(to)
		id, ok := safemath.ParseSafeUint256(idStr)
		if !ok {
			return sdk.Error("parse id failed")
		}
		amount, ok := safemath.ParseSafeUint256(amountStr)
		if !ok {
			return sdk.Error("parse amount failed")
		}

		return e.SafeTransferFrom(Address(from), Address(to), id, amount, data)
	case "SafeBatchTransferFrom":
		from := string(args[paramFrom])
		to := string(args[paramTo])
		idsStr := string(args[paramIds])
		amountsStr := string(args[paramAmounts])
		data := string(args[paramData])

		from = address.GetCleanAddr(from)
		to = address.GetCleanAddr(to)
		idArr := strings.Split(idsStr, ",")
		amountArr := strings.Split(amountsStr, ",")
		if len(amountArr) != len(idArr) || len(amountArr) == 0 {
			return sdk.Error("param error")
		}

		var tokenIds = make([]*safemath.SafeUint256, 0)
		var amounts = make([]*safemath.SafeUint256, 0)
		for i, id := range idArr {
			tokenId, err := safemath.ParseSafeUint256(id)
			if !err {
				return sdk.Error("parse id failed")
			}
			amount, err := safemath.ParseSafeUint256(amountArr[i])
			if !err {
				return sdk.Error("parse amount failed")
			}
			tokenIds = append(tokenIds, tokenId)
			amounts = append(amounts, amount)
		}
		return e.SafeBatchTransferFrom(Address(from), Address(to), tokenIds, amounts, data)
	default:
		return sdk.Error("Invalid method" + method)
	}
}

// GetUri 查询tokenId的uri
func (e *ERC1155Contract) GetUri(tokenId *safemath.SafeUint256) protogo.Response {
	contractInfo := maps.getErc1155Map()
	baseURI, _ := contractInfo.Get([]string{keyUri})
	uri := string(baseURI) + "/" + tokenId.ToString()
	return sdk.Success([]byte(uri))
}

// SetUri 设置uri
func (e *ERC1155Contract) SetUri(data string) protogo.Response {
	if !e.senderIsAdmin() {
		return sdk.Error("No permission")
	}
	contractInfo := maps.getErc1155Map()
	err := contractInfo.Set([]string{keyUri}, []byte(data))
	if err != nil {
		sdk.Instance.Warnf("Set uri failed, err:%s", err)
		return sdk.Error("Set uri failed, err:" + err.Error())
	}
	return successData(data)
}

// AlterAdminAddress 修改管理员
func (e *ERC1155Contract) AlterAdminAddress(adminAddress []string) protogo.Response {
	if len(adminAddress) == 0 {
		return sdk.Error("adminAddress of param should not be empty")
	}
	if !e.senderIsAdmin() {
		return sdk.Error("sender is not admin")
	}
	if !address.IsValidAddress(adminAddress...) {
		return sdk.Error("address format error")
	}
	identityInfo := maps.getErc1155Map()
	adminAddressByte, _ := json.Marshal(adminAddress)
	err := identityInfo.Set([]string{keyAdminAddress}, adminAddressByte)
	if err != nil {
		return sdk.Error("alter admin address of identityInfo failed." + err.Error())
	}
	sdk.Instance.EmitEvent("AlterAdminAddress", adminAddress)
	return successNormal()
}

func (e *ERC1155Contract) GetAdminAddress() protogo.Response {
	identityInfo := maps.getErc1155Map()
	adminAddressByte, _ := identityInfo.Get([]string{keyAdminAddress})
	return sdk.Success(adminAddressByte)
}

// SafeTransferFrom 交易
func (e *ERC1155Contract) SafeTransferFrom(from, to Address, id *safemath.SafeUint256, value *safemath.SafeUint256, data string) protogo.Response { // nolint
	if !address.IsValidAddress(string(from), string(to)) {
		return sdk.Error("address format error")
	}
	err := e.transferCore(from, to, id, value)
	if err != nil {
		return sdk.Error(err.Error())
	}

	// emit event
	sdk.Instance.EmitEvent("SafeTransferFrom", []string{string(from), string(to), id.ToString(), value.ToString()})
	return successNormal()
}

// SafeBatchTransferFrom 批量交易
func (e *ERC1155Contract) SafeBatchTransferFrom(from, to Address, ids, values []*safemath.SafeUint256, data string) protogo.Response { // nolint
	if len(ids) != len(values) {
		return sdk.Error("param len error")
	}
	if !address.IsValidAddress(string(from), string(to)) {
		return sdk.Error("address format error")
	}
	for i, id := range ids {
		value := values[i]
		err := e.transferCore(from, to, id, value)
		if err != nil {
			return sdk.Error(err.Error())
		}
	}
	// emit event
	idsStr, _ := json.Marshal(ids)
	valuesStr, _ := json.Marshal(values)
	sdk.Instance.EmitEvent("SafeTransferFrom", []string{string(from), string(to), string(idsStr), string(valuesStr)})

	return successNormal()
}

func (e *ERC1155Contract) transferCore(from Address, to Address, id *safemath.SafeUint256, value *safemath.SafeUint256) error { // nolint
	if !value.GTE(safemath.SafeUintOne) {
		return errors.New("error value " + value.ToString())
	}
	sender, _ := sdk.Instance.Origin()
	isApprovedOrOwner, err := e.isApprovedOrOwner(from, Address(sender), id)
	if err != nil {
		return fmt.Errorf("check isApprovedOrOwner failed, err:%s", err)
	}
	if !isApprovedOrOwner {
		return errors.New("caller is not token owner or approved")
	}
	// query balance
	balanceInfo := maps.getBalanceMap()

	balanceFromByte, _ := balanceInfo.Get([]string{string(from), id.ToString()})
	balanceToByte, _ := balanceInfo.Get([]string{string(to), id.ToString()})

	balanceFrom, _ := safemath.ParseSafeUint256(string(balanceFromByte))
	balanceTo, _ := safemath.ParseSafeUint256(string(balanceToByte))
	if !balanceFrom.GTE(value) {
		return errors.New("insufficient balance from:" + balanceFrom.ToString())
	}
	balanceFrom, _ = safemath.SafeSub(balanceFrom, value)
	balanceTo, _ = safemath.SafeAdd(balanceTo, value)

	// set balance
	_ = balanceInfo.Set([]string{string(from), id.ToString()}, []byte(balanceFrom.ToString()))
	_ = balanceInfo.Set([]string{string(to), id.ToString()}, []byte(balanceTo.ToString()))
	// set owner only nft valid
	_ = sdk.Instance.PutState("owner", id.ToString(), string(to))

	// set tx number
	tokenInfo := maps.getTokenMap()
	tokenTxCountByte, _ := tokenInfo.Get([]string{id.ToString()})
	tokenTxCount, _ := safemath.ParseSafeUint256(string(tokenTxCountByte))
	tokenTxCount, _ = safemath.SafeAdd(tokenTxCount, safemath.SafeUintOne)
	_ = tokenInfo.Set([]string{id.ToString()}, []byte(tokenTxCount.ToString()))
	return nil
}

func (e *ERC1155Contract) isApprovedOrOwner(owner, sender Address, tokenId *safemath.SafeUint256) (bool, error) {
	// check owner
	if owner == sender {
		return true, nil
	}

	// check operatorApprove
	resp := e.IsApprovedForAll(owner, sender)
	if resp.Status != sdk.OK {
		return false, errors.New(resp.Message)
	}
	if string(resp.Payload) == trueString {
		return true, nil
	}
	return false, nil
}

// BalanceOf 余额
func (e *ERC1155Contract) BalanceOf(owner Address, id *safemath.SafeUint256) protogo.Response {
	balanceInfo := maps.getBalanceMap()
	balanceByte, err := balanceInfo.Get([]string{string(owner), id.ToString()})
	if err != nil {
		return sdk.Error("get balance error. " + err.Error())
	}
	balance, _ := safemath.ParseSafeUint256(string(balanceByte))
	return sdk.Success([]byte(balance.ToString()))
}

// BalanceOfBatch 批量查询余额
func (e *ERC1155Contract) BalanceOfBatch(owners []Address, ids []*safemath.SafeUint256) protogo.Response {
	if len(owners) != len(ids) {
		return sdk.Error("param len error")
	}
	balanceInfo := maps.getBalanceMap()
	var balances []string
	for i, owner := range owners {
		id := ids[i]
		balanceByte, err := balanceInfo.Get([]string{string(owner), id.ToString()})
		if err != nil {
			return sdk.Error("get balance error. " + err.Error())
		}
		balance, _ := safemath.ParseSafeUint256(string(balanceByte))
		balances = append(balances, balance.ToString())
	}
	// 1,2,0,3,3
	balancesJson, _ := json.Marshal(balances)
	return sdk.Success(balancesJson)
}

// SetApprovalForAll 授权
func (e *ERC1155Contract) SetApprovalForAll(sender, operator Address, approved bool) protogo.Response {
	// check param
	if !address.IsValidAddress(string(operator)) {
		return sdk.Error("operator address invalid")
	}
	if sender == operator {
		return sdk.Error("approve to caller")
	}

	var approvedStr string
	if approved {
		approvedStr = trueString
	} else {
		approvedStr = falseString
	}

	// save approval
	approveInfo := maps.getApproveMap()
	err := approveInfo.Set([]string{string(sender), string(operator)}, []byte(approvedStr))
	if err != nil {
		return sdk.Error("set ApprovalForAll error, " + err.Error())
	}

	sdk.Instance.EmitEvent("ApprovalForAll", []string{string(sender), string(operator), approvedStr})

	return successNormal()
}

// IsApprovedForAll 查询权限
func (e *ERC1155Contract) IsApprovedForAll(owner, operator Address) protogo.Response {
	approveInfo := maps.getApproveMap()

	val, err := approveInfo.Get([]string{string(owner), string(operator)})
	if err != nil {
		return sdk.Error(fmt.Sprintf("get approved val from approve info failed, err:%s", err))
	}
	if string(val) == trueString {
		return sdk.Success([]byte(trueString))
	}
	return sdk.Success([]byte(falseString))
}

// MintBatchNormal 批量发行
func (e *ERC1155Contract) MintBatchNormal(to Address, ids, amounts string, data string) protogo.Response {
	if len(ids) == 0 || len(amounts) == 0 {
		return sdk.Error("param empty")
	}
	idStrArr := strings.Split(ids, ",")
	amountStrArr := strings.Split(amounts, ",")
	if len(idStrArr) != len(amountStrArr) {
		return sdk.Error("param err")
	}
	var idsArr []*safemath.SafeUint256
	var amountArr []*safemath.SafeUint256
	for i, idStr := range idStrArr {
		amountStr := amountStrArr[i]
		id, flag := safemath.ParseSafeUint256(idStr)
		if !flag {
			return sdk.Error("param id error")
		}
		amount, flag := safemath.ParseSafeUint256(amountStr)
		if !flag {
			return sdk.Error("param amount error")
		}
		idsArr = append(idsArr, id)
		amountArr = append(amountArr, amount)
	}
	return e.MintBatch(to, idsArr, amountArr, data)
}

// MintBatchNft 批量发行nft
func (e *ERC1155Contract) MintBatchNft(to Address, idStart *safemath.SafeUint256, amount int, data string) protogo.Response { // nolint
	// check param
	if amount <= 0 {
		return sdk.Error("nft amount mast large than 0")
	}
	if !e.senderIsAdmin() {
		return sdk.Error("No permission")
	}
	if !address.IsValidAddress(string(to)) {
		return sdk.Error("param error")
	}
	if !idStart.GTE(safemath.SafeUintOne) {
		return sdk.Error("param idStart error")
	}
	if len(e.OwnerOf(idStart)) > 0 {
		return sdk.Error("exists id " + idStart.ToString())
	}

	// make param
	var tokenIdArr []*safemath.SafeUint256
	var amountArr []*safemath.SafeUint256
	var tokenIdStart = idStart

	for i := 0; i < amount; i++ {
		tokenIdArr = append(tokenIdArr, tokenIdStart)
		amountArr = append(amountArr, safemath.SafeUintOne)
		tokenIdStart, _ = safemath.SafeAdd(tokenIdStart, safemath.SafeUintOne)
	}

	return e.MintBatch(to, tokenIdArr, amountArr, data)
}

// OwnerOf 查询owner
func (e *ERC1155Contract) OwnerOf(tokenId *safemath.SafeUint256) Address {
	val, _ := sdk.Instance.GetState("owner", tokenId.ToString())
	return Address(val)
}

// Mint 发行
func (e *ERC1155Contract) Mint(to Address, id, amount *safemath.SafeUint256, data string) protogo.Response {
	if !address.IsValidAddress(string(to)) {
		return sdk.Error("param to error")
	}
	if !e.senderIsAdmin() {
		return sdk.Error("No permission")
	}
	if !id.GTE(safemath.SafeUintZero) {
		return sdk.Error("param id error")
	}
	if !address.IsValidAddress(string(to)) {
		return sdk.Error("param error")
	}
	// query balance
	balanceInfo := maps.getBalanceMap()
	balanceByte, err := balanceInfo.Get([]string{string(to), id.ToString()})
	if err != nil {
		return sdk.Error(err.Error())
	}
	balance, _ := safemath.ParseSafeUint256(string(balanceByte))
	balance, _ = safemath.SafeAdd(balance, amount)

	// set balance
	_ = balanceInfo.Set([]string{string(to), id.ToString()}, []byte(balance.ToString()))
	// set owner only nft valid
	_ = sdk.Instance.PutState("owner", id.ToString(), string(to))
	// emit event
	sdk.Instance.EmitEvent("Mint", []string{string(to), id.ToString(), amount.ToString()})
	// set the number of token transactions
	tokenInfo := maps.getTokenMap()
	_ = tokenInfo.Set([]string{id.ToString()}, []byte("0"))

	return successNormal()
}

// MintBatch 批量发行实现
func (e *ERC1155Contract) MintBatch(to Address, ids, amounts []*safemath.SafeUint256, data string) protogo.Response {
	// check param
	balanceInfo := maps.getBalanceMap()
	tokenInfo := maps.getTokenMap()
	idStrArr := make([]string, 0)
	amountStrArr := make([]string, 0)
	for i, id := range ids {
		amount := amounts[i]
		// query balance
		balanceByte, err := balanceInfo.Get([]string{string(to), id.ToString()})
		if err != nil {
			return sdk.Error(err.Error())
		}
		balance, _ := safemath.ParseSafeUint256(string(balanceByte))
		balance, _ = safemath.SafeAdd(balance, amount)

		// set balance
		_ = balanceInfo.Set([]string{string(to), id.ToString()}, []byte(balance.ToString()))
		// set owner only nft valid
		_ = sdk.Instance.PutState("owner", id.ToString(), string(to))
		// set the number of token transactions
		_ = tokenInfo.Set([]string{id.ToString()}, []byte("0"))

		idStrArr = append(idStrArr, id.ToString())
		amountStrArr = append(amountStrArr, amount.ToString())
	}
	// emit event
	i, _ := json.Marshal(idStrArr)
	a, _ := json.Marshal(amountStrArr)
	sdk.Instance.EmitEvent("MintBatch", []string{string(to), string(i), string(a)})

	return successNormal()
}
func (e *ERC1155Contract) senderIsAdmin() bool {
	sender, _ := sdk.Instance.Origin()
	contractInfo := maps.getErc1155Map()
	adminAddressByte, err := contractInfo.Get([]string{keyAdminAddress})
	if len(adminAddressByte) == 0 || err != nil {
		sdk.Instance.Warnf("Get adminAddress failed, err:%s", err)
		return false
	}
	var adminAddress []string
	_ = json.Unmarshal(adminAddressByte, &adminAddress)
	for _, addr := range adminAddress {
		if addr == sender {
			return true
		}
	}
	return false
}

type mapManage struct {
}

func (m *mapManage) getApproveMap() *sdk.StoreMap {
	sm, _ := sdk.NewStoreMap("approve", 2, crypto.HASH_TYPE_SHA256)
	return sm
}

func (m *mapManage) getErc1155Map() *sdk.StoreMap {
	sm, _ := sdk.NewStoreMap("fxtoonErc1155", 1, crypto.HASH_TYPE_SHA256)
	return sm
}

func (m *mapManage) getBalanceMap() *sdk.StoreMap {
	sm, _ := sdk.NewStoreMap("balance", 2, crypto.HASH_TYPE_SHA256)
	return sm
}

func (m *mapManage) getTokenMap() *sdk.StoreMap {
	sm, _ := sdk.NewStoreMap("token", 1, crypto.HASH_TYPE_SHA256)
	return sm
}

func successNormal() protogo.Response {
	return sdk.Success([]byte("ok"))
}
func successData(data string) protogo.Response {
	return sdk.Success([]byte(data))
}

func main() {
	err := sandbox.Start(new(ERC1155Contract))
	if err != nil {
		log.Fatal(err)
	}
}
