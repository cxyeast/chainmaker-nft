/*
  Copyright (C) BABEC. All rights reserved.
  Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

  SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"chainmaker.org/chainmaker/common/v2/crypto"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sandbox"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
	"chainmaker.org/chainmaker/contract-utils/address"
	"chainmaker.org/chainmaker/contract-utils/safemath"
)

// ERC721 standard interface
type ERC721 interface {
	// @notice Count all NFTs assigned to an owner
	// @dev NFTs assigned to the zero address are considered invalid, and this
	//  function throws for queries about the zero address.
	// @param _owner An address for whom to query the balance
	// @return The number of NFTs owned by `_owner`, possibly zero
	balanceOf(owner string) protogo.Response

	// @notice Find the owner of an NFT
	// @dev NFTs assigned to zero address are considered invalid, and queries
	//  about them do throw.
	// @param _tokenId The identifier for an NFT
	// @return The address of the owner of the NFT
	ownerOf(tokenId *safemath.SafeUint256) protogo.Response

	// @notice Transfers the ownership of an NFT from one address to another address
	// @dev Throws unless `msg.sender` is the current owner, an authorized
	//  operator, or the approved address for this NFT. Throws if `_from` is
	//  not the current owner. Throws if `_to` is the zero address. Throws if
	//  `_tokenId` is not a valid NFT. When transfer is complete, this function
	//  checks if `_to` is a smart contract (code size > 0). If so, it calls
	//  `onERC721Received` on `_to` and throws if the return value is not
	//  `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))`.
	// @param _from The current owner of the NFT
	// @param _to The new owner
	// @param _tokenId The NFT to transfer
	// @param data Additional data with no specified format, sent in call to `_to`
	safeTransferFrom(from, to string, tokenId *safemath.SafeUint256, data []byte) protogo.Response

	// @notice Transfer ownership of an NFT -- THE CALLER IS RESPONSIBLE
	//  TO CONFIRM THAT `_to` IS CAPABLE OF RECEIVING NFTS OR ELSE
	//  THEY MAY BE PERMANENTLY LOST
	// @dev Throws unless `msg.sender` is the current owner, an authorized
	//  operator, or the approved address for this NFT. Throws if `_from` is
	//  not the current owner. Throws if `_to` is the zero address. Throws if
	//  `_tokenId` is not a valid NFT.
	// @param _from The current owner of the NFT
	// @param _to The new owner
	// @param _tokenId The NFT to transfer
	transferFrom(from, to string, tokenId *safemath.SafeUint256) protogo.Response

	// @notice Change or reaffirm the approved address for an NFT
	// @dev The zero address indicates there is no approved address.
	//  Throws unless `msg.sender` is the current NFT owner, or an authorized
	//  operator of the current owner.
	// @param _approved The new approved NFT controller
	// @param _tokenId The NFT to approve
	approve(address string, tokenId *safemath.SafeUint256) protogo.Response

	// @notice Enable or disable approval for a third party ("operator") to manage
	//  all of `msg.sender`'s assets
	// @dev Emits the ApprovalForAll event. The contract MUST allow
	//  multiple operators per owner.
	// @param _operator Address to add to the set of authorized operators
	// @param _approved True if the operator is approved, false to revoke approval
	setApprovalForAll(operator string, approved bool) protogo.Response

	// @notice Get the approved address for a single NFT
	// @dev Throws if `_tokenId` is not a valid NFT.
	// @param _tokenId The NFT to find the approved address for
	// @return The approved address for this NFT, or the zero address if there is none
	getApproved(tokenId *safemath.SafeUint256) protogo.Response

	// @notice Query if an address is an authorized operator for another address
	// @param _owner The address that owns the NFTs
	// @param _operator The address that acts on behalf of the owner
	// @return True if `_operator` is an approved operator for `_owner`, false otherwise
	isApprovedForAll(owner, operator string) protogo.Response
}

// ERC165 interface
type ERC165 interface {
	// @notice Query if a contract implements an interface
	// @param interfaceID The interface identifier, as specified in ERC-165
	// @dev Interface identification is specified in ERC-165. This function
	//  uses less than 30,000 gas.
	// @return `true` if the contract implements `interfaceID` and
	//  `interfaceID` is not 0xffffffff, `false` otherwise
	supportsInterface(interfaceID [4]byte) protogo.Response
}

// ERC721TokenReceiver the ERC-165 identifier for this interface is 0x150b7a02.
type ERC721TokenReceiver interface {
	/// @notice Handle the receipt of an NFT
	/// @dev The ERC721 smart contract calls this function on the recipient
	///  after a `transfer`. This function MAY throw to revert and reject the
	///  transfer. Return of other than the magic value MUST result in the
	///  transaction being reverted.
	///  Note: the contract address is always the message sender.
	/// @param _operator The address which called `safeTransferFrom` function
	/// @param _from The address which previously owned the token
	/// @param _tokenId The NFT identifier which is being transferred
	/// @param _data Additional data with no specified format
	/// @return `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))`
	///  unless throwing
	onERC721Received(operator, from string, tokenId safemath.SafeUint256, data []byte) protogo.Response
}

// ERC721Metadata is the metadata of the nfts
// @title ERC-721 Non-Fungible Token Standard, optional metadata extension
// @dev See https://eips.ethereum.org/EIPS/eip-721
//
//	Note: the ERC-165 identifier for this interface is 0x5b5e139f.
type ERC721Metadata interface /* is ERC721 */ {
	// @notice A descriptive name for a collection of NFTs in this contract
	name() protogo.Response

	// @notice An abbreviated name for NFTs in this contract
	symbol() protogo.Response

	// @notice A distinct Uniform Resource Identifier (URI) for a given asset.
	// @dev Throws if `_tokenId` is not a valid NFT. URIs are defined in RFC
	//  3986. The URI may point to a JSON file that conforms to the "ERC721
	//  Metadata JSON Schema".
	tokenURI(tokenId *safemath.SafeUint256) protogo.Response
}

// IERC721 interface
type IERC721 interface {
	ERC721
	ERC721Metadata
	ERC165
	ERC721TokenReceiver
	InitContract() protogo.Response    // return "Init contract success"
	UpgradeContract() protogo.Response // return "Upgrade contract success"
}

const (
	erc721InfoMapName      = "erc721"
	balanceInfoMapName     = "balanceInfo"
	accountMapName         = "accountInfo"
	tokenOwnerMapName      = "tokenIds"
	tokenInfoMapName       = "tokenInfo"
	tokenApproveMapName    = "tokenApprove"
	operatorApproveMapName = "operatorApprove"
	trueString             = "1"
	falseString            = "0"
)

var _ IERC721 = (*ERC721Contract)(nil)

// ERC721Contract contract for erc721
type ERC721Contract struct {
}

type tokenLatestTxInfo struct {
	TxId        string
	BlockHeight uint64
	From        string
	To          string
	Timestamp   string
}

type accountTokens struct {
	Account string
	Tokens  []string
}

func (c *ERC721Contract) AlterAdminAddress(adminAddress []string) protogo.Response {
	if len(adminAddress) == 0 {
		return sdk.Error("adminAddress of param should not be empty")
	}
	if !c.senderIsAdmin() {
		return sdk.Error("sender is not admin")
	}

	if !address.IsValidAddress(adminAddress...) {
		return sdk.Error("address format error")
	}
	erc721Info, err := sdk.NewStoreMap(erc721InfoMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error("new storeMap of erc721Info failed")
	}

	adminAddressByte, _ := json.Marshal(adminAddress)
	err = erc721Info.Set([]string{"admin"}, adminAddressByte)
	if err != nil {
		return sdk.Error("alter admin address of identityInfo failed." + err.Error())
	}
	sdk.Instance.EmitEvent("AlterAdminAddress", adminAddress)
	return sdk.Success([]byte("ok"))
}

func (c *ERC721Contract) GetAdminAddress() protogo.Response {
	erc721Info, err := sdk.NewStoreMap(erc721InfoMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error("new storeMap of erc721Info failed")
	}
	adminAddressByte, _ := erc721Info.Get([]string{"admin"})
	return sdk.Success(adminAddressByte)
}

func (c *ERC721Contract) senderIsAdmin() bool {
	sender, _ := sdk.Instance.Origin()
	erc721Info, _ := sdk.NewStoreMap(erc721InfoMapName, 1, crypto.HASH_TYPE_SHA256)
	adminAddressByte, err := erc721Info.Get([]string{"admin"})
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

// InitContract install contract func
func (c *ERC721Contract) InitContract() protogo.Response {
	err := c.updateErc721Info()
	if err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success([]byte("Init contract success"))
}

// UpgradeContract upgrade contract func
func (c *ERC721Contract) UpgradeContract() protogo.Response {
	// err := c.updateErc721Info()
	// if err != nil {
	// 	return sdk.Error(err.Error())
	// }
	return sdk.Success([]byte("Upgrade contract success"))
}

// UpgradeContract upgrade contract func
func (c *ERC721Contract) updateErc721Info() error {
	args := sdk.Instance.GetArgs()
	// name, symbol and decimal are optional
	name := args["name"]
	symbol := args["symbol"]
	tokenURI := args["tokenURI"]
	adminAddress := args["adminAddress"]
	erc721Info, err := sdk.NewStoreMap(erc721InfoMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return fmt.Errorf("new storeMap of erc721Info failed, err:%s", err)
	}
	if len(name) > 0 {
		err = erc721Info.Set([]string{"name"}, name)
		if err != nil {
			return fmt.Errorf("set name of erc721Info failed, err:%s", err)
		}
	}
	if len(symbol) > 0 {
		err = erc721Info.Set([]string{"symbol"}, symbol)
		if err != nil {
			return fmt.Errorf("set symbol of erc721Info failed, err:%s", err)
		}
	}
	if len(tokenURI) > 0 {
		err = erc721Info.Set([]string{"tokenURI"}, tokenURI)
		if err != nil {
			return fmt.Errorf("set tokenURI of erc721Info failed, err:%s", err)
		}
	}

	var adminAddressStr string
	if len(adminAddress) == 0 {
		adminAddressStr, _ = sdk.Instance.Origin()
	} else {
		adminAddressStr = string(adminAddress)
	}
	adminAddresses := strings.Split(adminAddressStr, ",")
	adminAddressByte, _ := json.Marshal(adminAddresses)
	err = erc721Info.Set([]string{"admin"}, adminAddressByte)
	if err != nil {
		return fmt.Errorf("set admin address of contractInfo failed")
	}
	sdk.Instance.EmitEvent("AlterAdminAddress", adminAddresses)

	return nil
}

// InvokeContract the entry func of invoke contract func
// nolint: gocyclo
func (c *ERC721Contract) InvokeContract(method string) protogo.Response {
	if len(method) == 0 {
		return sdk.Error("method of param should not be empty")
	}
	args := sdk.Instance.GetArgs()
	switch method {
	case "AlterAdminAddress":
		address := args["adminAddress"]
		var addresses []string
		if len(address) != 0 {
			addresses = strings.Split(string(address), ",")
		}

		return c.AlterAdminAddress(addresses)
	case "GetAdmins":
		return c.GetAdminAddress()
	case "balanceOf":
		account := string(args["account"])
		if len(account) == 0 {
			return sdk.Error("Param account should not be empty")
		}
		return c.balanceOf(account)
	case "ownerOf":
		tokenIdStr := string(args["tokenId"])
		tokenId, ok := safemath.ParseSafeUint256(tokenIdStr)
		if !ok {
			return sdk.Error("invalid tokenId")
		}
		return c.ownerOf(tokenId)
	case "approve":
		tokenIdStr := string(args["tokenId"])
		tokenId, ok := safemath.ParseSafeUint256(tokenIdStr)
		if !ok {
			return sdk.Error("invalid tokenId")
		}

		to := string(args["to"])
		return c.approve(to, tokenId)
	case "getApprove":
		tokenIdStr := string(args["tokenId"])
		tokenId, ok := safemath.ParseSafeUint256(tokenIdStr)
		if !ok {
			return sdk.Error("invalid tokenId")
		}
		minted, err := c.minted(tokenId)
		if err != nil {
			return sdk.Error(err.Error())
		}
		if !minted {
			return sdk.Error("ERC721: invalid token ID")
		}
		return c.getApproved(tokenId)
	case "transferFrom":
		from := string(args["from"])
		to := string(args["to"])
		tokenIdStr := string(args["tokenId"])
		tokenId, ok := safemath.ParseSafeUint256(tokenIdStr)
		if !ok {
			return sdk.Error("Parse tokenId failed")
		}
		return c.transferFrom(from, to, tokenId)
	case "safeTransferFrom":
		from := string(args["from"])
		to := string(args["to"])
		tokenIdStr := string(args["tokenId"])
		tokenId, ok := safemath.ParseSafeUint256(tokenIdStr)
		if !ok {
			return sdk.Error("Parse tokenId failed")
		}
		return c.safeTransferFrom(from, to, tokenId, nil)
	// below methods are optional
	case "name":
		return c.name()
	case "setname":
		name := string(args["name"])
		return c.setname(name)
	case "symbol":
		return c.symbol()
	case "setsymbol":
		symbol := string(args["symbol"])
		return c.setsymbol(symbol)
	case "tokenURI":
		tokenIdStr := string(args["tokenId"])
		tokenId, ok := safemath.ParseSafeUint256(tokenIdStr)
		if !ok {
			return sdk.Error("Parse tokenId failed")
		}
		return c.tokenURI(tokenId)
	case "setTokenURI":
		uri := string(args["uri"])
		return c.setTokenURI(uri)
	case "tokenMetadata":
		tokenIdStr := string(args["tokenId"])
		tokenId, ok := safemath.ParseSafeUint256(tokenIdStr)
		if !ok {
			return sdk.Error("Parse tokenId failed")
		}
		return c.tokenMetadata(tokenId)
	case "tokenLatestTxInfo":
		tokenIdStr := string(args["tokenId"])
		tokenId, ok := safemath.ParseSafeUint256(tokenIdStr)
		if !ok {
			return sdk.Error("Parse tokenId failed")
		}
		return c.tokenLatestTxInfo(tokenId)
	case "accountTokens":
		account := string(args["account"])
		if !address.IsValidAddress(account) {
			return sdk.Error("invalid account")
		}
		return c.accountTokens(account)
	case "mint":
		to := string(args["to"])
		tokenIdStr := string(args["tokenId"])
		metadata := args["metadata"]
		tokenId, ok := safemath.ParseSafeUint256(tokenIdStr)
		if !ok {
			return sdk.Error("Parse tokenId failed")
		}
		return c.mint(metadata, to, tokenId)
	default:
		return sdk.Error("Invalid method")
	}
}

// balanceOf return token count of the account
func (c *ERC721Contract) balanceOf(account string) protogo.Response {
	if !address.IsValidAddress(account) {
		return sdk.Error("ERC721: balanceOf from the invalid address")
	}
	if address.IsZeroAddress(account) {
		return sdk.Error("ERC721: address zero is not a valid owner")
	}
	balanceInfo, err := sdk.NewStoreMap(balanceInfoMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error(fmt.Sprintf("New storeMap of balanceInfo failed, err:%s", err))
	}

	balanceCount, err := c.getBalance(balanceInfo, account)
	if err != nil {
		return sdk.Error(fmt.Sprintf("Get balance failed, err:%s", err))
	}
	return sdk.Success([]byte(balanceCount.ToString()))
}

func (c *ERC721Contract) ownerOf(tokenId *safemath.SafeUint256) protogo.Response {
	tokenIds, err := sdk.NewStoreMap(tokenOwnerMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error(fmt.Sprintf("New storeMap of tokenIds failed, err:%s", err))
	}

	owner, err := tokenIds.Get([]string{tokenId.ToString()})
	if err != nil {
		return sdk.Error(err.Error())
	}

	return sdk.Success(owner)
}

func (c *ERC721Contract) approve(to string, tokenId *safemath.SafeUint256) protogo.Response {
	// check owner
	resp := c.ownerOf(tokenId)
	if resp.Status != sdk.OK {
		return resp
	}
	owner := string(resp.Payload)
	if owner == to {
		return sdk.Error("approval to current owner")
	}
	// check approve info
	sender, err := sdk.Instance.Sender()
	if err != nil {
		return sdk.Error(fmt.Sprintf("get sender failed, err:%s", err))
	}
	if owner != sender {
		resp = c.isApprovedForAll(owner, sender)
		if resp.Status != sdk.OK {
			return sdk.Error(resp.Message)
		}
		if string(resp.Payload) == falseString {
			return sdk.Error("ERC721: approve caller is not token owner or approved for all")
		}
	}
	tokenApproveInfo, err := sdk.NewStoreMap(tokenApproveMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error(fmt.Sprintf("new storeMap of tokenApproveMap failed, err:%s", err))
	}
	err = tokenApproveInfo.Set([]string{tokenId.ToString()}, []byte(to))
	if err != nil {
		return sdk.Error(fmt.Sprintf("set owner failed, err:%s", err))
	}

	sdk.Instance.EmitEvent("approve", []string{owner, to, tokenId.ToString()})

	return sdk.Success([]byte("approve success"))
}

func (c *ERC721Contract) getApproved(tokenId *safemath.SafeUint256) protogo.Response {
	tokenApproveInfo, err := sdk.NewStoreMap(tokenApproveMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error(err.Error())
	}
	approveTo, err := tokenApproveInfo.Get([]string{tokenId.ToString()})
	if err != nil {
		return sdk.Error(err.Error())
	}

	return sdk.Success(approveTo)
}

func (c *ERC721Contract) setApprovalForAll(operator string, approved bool) protogo.Response {
	sender, err := sdk.Instance.Sender()
	if err != nil {
		return sdk.Error(err.Error())
	}
	if sender == operator {
		return sdk.Error("ERC721: approve to caller")
	}
	operatorApproveInfo, err := sdk.NewStoreMap(operatorApproveMapName, 2, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error(fmt.Sprintf("new storemap of operatorApprove failed, err:%s", err))
	}
	var approvedStr string
	if approved {
		approvedStr = trueString
	} else {
		approvedStr = falseString
	}
	err = operatorApproveInfo.Set([]string{sender, operator}, []byte(approvedStr))
	if err != nil {
		return sdk.Error(fmt.Sprintf("set operator approve failed, err:%s", err))
	}
	sdk.Instance.EmitEvent("ApprovalForAll", []string{sender, operator, approvedStr})

	return sdk.Success([]byte("setApprovalForAll success"))
}

func (c *ERC721Contract) isApprovedForAll(owner, sender string) protogo.Response {
	operatorApproveInfo, err := sdk.NewStoreMap(operatorApproveMapName, 2, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error(err.Error())
	}
	val, err := operatorApproveInfo.Get([]string{owner, sender})
	if err != nil {
		return sdk.Error(fmt.Sprintf("get approved val from approve info failed, err:%s", err))
	}
	if string(val) == trueString {
		return sdk.Success([]byte(trueString))
	}

	return sdk.Success([]byte(falseString))
}

func (c *ERC721Contract) transferFrom(from, to string, tokenId *safemath.SafeUint256) protogo.Response {
	sender, err := sdk.Instance.Sender()
	if err != nil {
		return sdk.Error(fmt.Sprintf("get sender failed, err:%s", err))
	}
	isApprovedOrOwner, err := c.isApprovedOrOwner(sender, tokenId)
	if err != nil {
		return sdk.Error(fmt.Sprintf("check isApprovedOrOwner failed, err:%s", err))
	}
	if !isApprovedOrOwner {
		return sdk.Error("ERC721: caller is not token owner or approved")
	}
	return c.transfer(from, to, tokenId)
}

// todo: add receiver check
func (c *ERC721Contract) safeTransferFrom(from, to string, tokenId *safemath.SafeUint256,
	data []byte) protogo.Response {
	return c.transferFrom(from, to, tokenId)
}

// name is optional
func (c *ERC721Contract) name() protogo.Response {
	erc721Info, err := sdk.NewStoreMap(erc721InfoMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error(fmt.Sprintf("new storeMap of erc721Info failed, err:%s", err))
	}
	name, err := erc721Info.Get([]string{"name"})
	if err != nil {
		return sdk.Error(fmt.Sprintf("get name from erc721Info failed, err:%s", err))
	}
	return sdk.Success(name)
}

func (c *ERC721Contract) setname(name string) protogo.Response {
	if !c.senderIsAdmin() {
		return sdk.Error("sender is not admin")
	}
	erc721Info, err := sdk.NewStoreMap(erc721InfoMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error(fmt.Sprintf("new storeMap of erc721Info failed, err:%s", err))
	}
	err = erc721Info.Set([]string{"name"}, []byte(name))
	if err != nil {
		return sdk.Error(fmt.Sprintf("set name from erc721Info failed, err:%s", err))
	}
	return sdk.Success([]byte(name))
}

// symbol is optional
func (c *ERC721Contract) symbol() protogo.Response {

	erc721Info, err := sdk.NewStoreMap(erc721InfoMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error(fmt.Sprintf("new storeMap of erc721Info failed, err:%s", err))
	}
	symbol, err := erc721Info.Get([]string{"symbol"})
	if err != nil {
		return sdk.Error(fmt.Sprintf("get symbol from erc721Info failed, err:%s", err))
	}
	return sdk.Success(symbol)
}

func (c *ERC721Contract) setsymbol(symbol string) protogo.Response {
	if !c.senderIsAdmin() {
		return sdk.Error("sender is not admin")
	}
	erc721Info, err := sdk.NewStoreMap(erc721InfoMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error(fmt.Sprintf("new storeMap of erc721Info failed, err:%s", err))
	}
	err = erc721Info.Set([]string{"symbol"}, []byte(symbol))
	if err != nil {
		return sdk.Error(fmt.Sprintf("set symbol from erc721Info failed, err:%s", err))
	}
	return sdk.Success([]byte(symbol))
}

// tokenURI is optional
func (c *ERC721Contract) tokenURI(tokenId *safemath.SafeUint256) protogo.Response {

	erc721Info, err := sdk.NewStoreMap(erc721InfoMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error(fmt.Sprintf("new storeMap of erc721Info failed, err:%s", err))
	}
	baseURI, err := erc721Info.Get([]string{"tokenURI"})
	if err != nil {
		return sdk.Error(fmt.Sprintf("get tokenURI from erc721Info failed, err:%s", err))
	}
	uri := string(baseURI) + "/" + tokenId.ToString()
	return sdk.Success([]byte(uri))
}

func (c *ERC721Contract) setTokenURI(data string) protogo.Response {
	if !c.senderIsAdmin() {
		return sdk.Error("sender is not admin")
	}
	erc721Info, err := sdk.NewStoreMap(erc721InfoMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error(fmt.Sprintf("new storeMap of erc721Info failed, err:%s", err))
	}
	err = erc721Info.Set([]string{"tokenURI"}, []byte(data))
	if err != nil {
		sdk.Instance.Warnf("Set uri failed, err:%s", err)
		return sdk.Error("Set uri failed, err:" + err.Error())
	}
	return sdk.Success([]byte(data))
}

func (c *ERC721Contract) tokenMetadata(tokenId *safemath.SafeUint256) protogo.Response {
	ti, err := sdk.NewStoreMap(tokenInfoMapName, 2, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error(fmt.Sprintf("new store map of token info failed, err:%s", err))
	}

	metadata, err := ti.Get([]string{tokenId.ToString(), "metadata"})
	if err != nil {
		return sdk.Error(fmt.Sprintf("set metadata of erc721Info failed, err:%s", err))
	}

	sdk.Instance.Debugf("tokenMetadata is %s", string(metadata))

	return sdk.Success(metadata)
}
func (c *ERC721Contract) tokenLatestTxInfo(tokenId *safemath.SafeUint256) protogo.Response {
	ti, err := sdk.NewStoreMap(tokenInfoMapName, 2, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error(fmt.Sprintf("new store map of token info failed, err:%s", err))
	}

	latestTxInfo, err := ti.Get([]string{tokenId.ToString(), "latestTxInfo"})
	if err != nil {
		return sdk.Error(fmt.Sprintf("get latest tx info failed, err:%s", err))
	}

	sdk.Instance.Debugf("tokenLatestTxInfo is %s", latestTxInfo)

	return sdk.Success(latestTxInfo)
}
func (c *ERC721Contract) accountTokens(account string) protogo.Response {
	am, err := sdk.NewStoreMap(accountMapName, 2, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error(fmt.Sprintf("new store map of token info failed, err:%s", err))
	}
	rs, err := am.NewStoreMapIteratorPrefixWithKey([]string{account})
	if err != nil {
		return sdk.Error(fmt.Sprintf("new store map iterator of project info failed, err: %s", err))
	}
	ats := &accountTokens{
		Account: account,
		Tokens:  make([]string, 0),
	}
	for {
		if !rs.HasNext() {
			break
		}
		var item string
		item, _, _, err = rs.Next()
		if err != nil {
			return sdk.Error(fmt.Sprintf("iterator next failed, err: %s", err))
		}
		itemId := strings.TrimPrefix(strings.TrimPrefix(item, accountMapName), account)
		if len(itemId) == 0 {
			return sdk.Error("invalid itemId")
		}
		ats.Tokens = append(ats.Tokens, itemId)
	}

	var atsBytes []byte
	atsBytes, err = json.Marshal(ats)
	if err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success(atsBytes)
}

func (c *ERC721Contract) mint(metadata []byte, to string, tokenId *safemath.SafeUint256) protogo.Response {
	if !address.IsValidAddress(to) {
		return sdk.Error("mint to invalid address")
	}
	if address.IsZeroAddress(to) {
		return sdk.Error("ERC721: mint to the zero address")
	}

	minted, err := c.minted(tokenId)
	if err != nil {
		return sdk.Error(err.Error())
	}
	if minted {
		return sdk.Error("duplicated token")
	}

	if !c.senderIsAdmin() {
		return sdk.Error("sender is not admin")
	}

	err = c.increaseTokenCountByOne(to)
	if err != nil {
		return sdk.Error(err.Error())
	}
	err = c.setTokenOwner(to, tokenId)
	if err != nil {
		return sdk.Error(err.Error())
	}

	err = c.setAccountToken(address.ZeroAddr, to, tokenId)
	if err != nil {
		return sdk.Error(err.Error())
	}

	err = c.setMetadata(tokenId, metadata)
	if err != nil {
		return sdk.Error(err.Error())
	}

	err = c.setTokenLatestTxInfo(tokenId, address.ZeroAddr, to)
	if err != nil {
		return sdk.Error(err.Error())
	}

	return sdk.Success([]byte("mint success"))
}

func (c *ERC721Contract) setAccountToken(from, to string, tokenId *safemath.SafeUint256) error {
	am, err := sdk.NewStoreMap(accountMapName, 2, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return fmt.Errorf("new store map of token info failed, err:%s", err)
	}

	err = am.Set([]string{to, tokenId.ToString()}, []byte(trueString))
	if err != nil {
		return fmt.Errorf("setAccountToken failed, err:%s", err)
	}
	if address.IsZeroAddress(from) {
		return nil
	}
	err = am.Del([]string{from, tokenId.ToString()})
	if err != nil {
		return err
	}
	return nil
}

func (c *ERC721Contract) setMetadata(tokenId *safemath.SafeUint256, metadata []byte) error {
	ti, err := sdk.NewStoreMap(tokenInfoMapName, 2, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return fmt.Errorf("new store map of token info failed, err:%s", err)
	}
	sdk.Instance.Debugf("setMetadata to %s", metadata)
	if len(metadata) > 0 {
		err = ti.Set([]string{tokenId.ToString(), "metadata"}, metadata)
		if err != nil {
			return fmt.Errorf("set metadata of erc721Info failed, err:%s", err)
		}
	}
	return nil
}

func (c *ERC721Contract) minted(tokenId *safemath.SafeUint256) (bool, error) {
	resp := c.ownerOf(tokenId)
	if resp.Status != sdk.OK {
		return false, fmt.Errorf(resp.Message)
	}
	owner := string(resp.Payload)

	return len(owner) > 0 && !address.IsZeroAddress(owner), nil
}

/**
 * @dev Returns whether `spender` is allowed to manage `tokenId`.
 *
 * Requirements:
 *
 * - `tokenId` must exist.
 */
func (c *ERC721Contract) isApprovedOrOwner(sender string, tokenId *safemath.SafeUint256) (bool, error) {
	// check owner
	response := c.ownerOf(tokenId)
	if response.Status != sdk.OK {
		return false, fmt.Errorf(response.Message)
	}
	owner := string(response.Payload)
	if owner == sender {
		return true, nil
	}

	// check operatorApprove
	resp := c.isApprovedForAll(owner, sender)
	if resp.Status != sdk.OK {
		return false, fmt.Errorf(resp.Message)
	}
	if string(resp.Payload) == trueString {
		return true, nil
	}

	// check tokenApprove
	resp = c.getApproved(tokenId)
	if resp.Status != sdk.OK {
		return false, fmt.Errorf(resp.Message)
	}

	return string(resp.Payload) == sender, nil
}

func (c *ERC721Contract) getBalance(balanceInfo *sdk.StoreMap, account string) (balance *safemath.SafeUint256,
	err error) {
	balanceBytes, err := balanceInfo.Get([]string{account})
	if err != nil {
		return nil, fmt.Errorf("get balance failed, err:%s", err)
	}
	balance, ok := safemath.ParseSafeUint256(string(balanceBytes))
	if !ok {
		return nil, fmt.Errorf("balance bytes invalid")
	}

	return balance, nil
}

func (c *ERC721Contract) transfer(from, to string, tokenId *safemath.SafeUint256) protogo.Response {
	response := c.ownerOf(tokenId)
	if response.Status != sdk.OK {
		return response
	}
	owner := string(response.Payload)
	if owner != from {
		return sdk.Error("ERC721: transfer from incorrect owner")
	}

	if !address.IsValidAddress(to) {
		return sdk.Error("ERC20: transfer to the invalid address")
	}
	if address.IsZeroAddress(to) {
		return sdk.Error("ERC20: transfer to the zero address")
	}
	// delete token approve
	tokenApproveInfo, err := sdk.NewStoreMap(tokenApproveMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return sdk.Error(fmt.Sprintf("New storeMap of balanceInfo failed, err:%s", err))
	}
	err = tokenApproveInfo.Del([]string{tokenId.ToString()})
	if err != nil {
		return sdk.Error(fmt.Sprintf("delete token approve failed, err:%s", err))
	}

	// update "from" balance count
	err = c.decreaseTokenCountByOne(from)
	if err != nil {
		return sdk.Error(err.Error())
	}

	// update "to" balance count
	err = c.increaseTokenCountByOne(to)
	if err != nil {
		return sdk.Error(err.Error())
	}

	// update token owner
	err = c.setTokenOwner(to, tokenId)
	if err != nil {
		return sdk.Error(err.Error())
	}

	err = c.setAccountToken(from, to, tokenId)
	if err != nil {
		return sdk.Error(err.Error())
	}

	err = c.setTokenLatestTxInfo(tokenId, from, to)
	if err != nil {
		return sdk.Error(err.Error())
	}

	sdk.Instance.EmitEvent("transfer", []string{from, to, tokenId.ToString()})

	return sdk.Success([]byte("transfer success"))
}

func (c *ERC721Contract) setTokenLatestTxInfo(tokenId *safemath.SafeUint256, from, to string) error {
	ti, err := sdk.NewStoreMap(tokenInfoMapName, 2, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return fmt.Errorf("new store map of token info failed, err:%s", err)
	}
	txId, err := sdk.Instance.GetTxId()
	if err != nil {
		return fmt.Errorf("get tx id failed, err:%s", err)
	}
	blockHeight, err := sdk.Instance.GetBlockHeight()
	if err != nil {
		return fmt.Errorf("get block height failed, err:%s", err)
	}
	timeStamp, err := sdk.Instance.GetTxTimeStamp()
	if err != nil {
		return fmt.Errorf("get tx timestamp failed, err:%s", err)
	}
	tkTxInfo := &tokenLatestTxInfo{
		TxId:        txId,
		BlockHeight: uint64(blockHeight),
		From:        from,
		To:          to,
		Timestamp:   timeStamp,
	}

	latestInfoBytes, err := json.Marshal(tkTxInfo)
	if err != nil {
		return fmt.Errorf("json marshal tokenLatestTxInfo failed, err:%s", err)
	}
	sdk.Instance.Debugf("setTokenLatestTxInfo to %s", string(latestInfoBytes))
	err = ti.Set([]string{tokenId.ToString(), "latestTxInfo"}, latestInfoBytes)
	if err != nil {
		return fmt.Errorf("set latestTxInfo of token failed, err:%s", err)
	}
	return nil
}

func (c *ERC721Contract) increaseTokenCountByOne(account string) error {
	balanceInfo, err := sdk.NewStoreMap(balanceInfoMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return fmt.Errorf("new storeMap of balanceInfo failed, err:%s", err)
	}
	originTokenCount, err := c.getBalance(balanceInfo, account)
	if err != nil {
		return fmt.Errorf("get token count failed, err:%s", err)
	}
	newTokenCount, ok := safemath.SafeAdd(originTokenCount, safemath.SafeUintOne)
	if !ok {
		return fmt.Errorf("balance count of from is overflow")
	}
	err = balanceInfo.Set([]string{account}, []byte(newTokenCount.ToString()))
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (c *ERC721Contract) decreaseTokenCountByOne(account string) error {
	balanceInfo, err := sdk.NewStoreMap(balanceInfoMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return fmt.Errorf("new storeMap of balanceInfo failed, err:%s", err)
	}
	originTokenCount, err := c.getBalance(balanceInfo, account)
	if err != nil {
		return fmt.Errorf("get token count failed, err:%s", err)
	}
	newTokenCount, ok := safemath.SafeSub(originTokenCount, safemath.SafeUintOne)
	if !ok {
		return fmt.Errorf("token count of account is overflow")
	}
	err = balanceInfo.Set([]string{account}, []byte(newTokenCount.ToString()))
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (c *ERC721Contract) setTokenOwner(to string, tokenId *safemath.SafeUint256) error {
	tokenOwnerInfo, err := sdk.NewStoreMap(tokenOwnerMapName, 1, crypto.HASH_TYPE_SHA256)
	if err != nil {
		return fmt.Errorf("new storeMap of tokenOwner failed, err:%s", err)
	}
	err = tokenOwnerInfo.Set([]string{tokenId.ToString()}, []byte(to))
	if err != nil {
		return fmt.Errorf("set token owner failed, err:%s", err)
	}
	return nil
}

func (c *ERC721Contract) supportsInterface(interfaceID [4]byte) protogo.Response {
	panic("implement me")
}

func (c *ERC721Contract) onERC721Received(operator, from string, tokenId safemath.SafeUint256,
	data []byte) protogo.Response {
	panic("implement me")
}

func main() {
	err := sandbox.Start(new(ERC721Contract))
	if err != nil {
		sdk.Instance.Errorf(err.Error())
	}
}
