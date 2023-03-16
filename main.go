package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"myclient/myclient"
	"net/http"
	"time"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	"github.com/gin-gonic/gin"
)

const (
	ADMIN1_CRYPTO_PATH = "./crypto-config/node1/admin/admin1"
	ADMIN2_CRYPTO_PATH = "./crypto-config/node1/admin/admin2"
	ADMIN3_CRYPTO_PATH = "./crypto-config/node1/admin/admin3"
	ADMIN4_CRYPTO_PATH = "./crypto-config/node1/admin/admin4"
	ADMIN1_NAME        = "admin1"
	ADMIN2_NAME        = "admin2"
	ADMIN3_NAME        = "admin3"
	ADMIN4_NAME        = "admin4"

	USER_CRYPTO_PATH          = "./crypto-config/users"
	SDK_CONFIG_PATH           = "./sdk-configs"
	NFT_BYTE_CODE_PATH        = "contract/erc721/erc721.7z"
	MULTI_BYTE_CODE_PATH      = "contract/erc1155/erc1155.7z"
	NFT_ASSET_CONTRACT_NAME   = "DDE_NFTAsset"
	MULTI_ASSET_CONTRACT_NAME = "DDE_MULTIAsset_TEST11"
)

var (
	eventTransferCallBackUrl string
)

type RespContractEventInfo struct {
	BlockHeight uint64   `json:"blockheight"`
	TxId        string   `json:"txid"`
	Topic       string   `json:"topic"`
	EventData   []string `json:"eventdata"`
}

type userConfig struct {
	UserName      string `json:"name"`
	UserCrytoPath string `json:"path"`
}

type newNFTAssetContract struct {
	// ByteCodePath      string `json:"codepath"`
	Version           string `json:"version"`
	Name              string `json:"name"`
	Symbol            string `json:"symbol"`
	TokenURI          string `json:"tokenuri"`
	OperatorAdminAddr string `json:"opadmin"`
}

type newMultiAssetContract struct {
	// ByteCodePath      string `json:"codepath"`
	Version           string   `json:"version"`
	URI               string   `json:"uri"`
	OperatorAdminAddr []string `json:"opadmin"`
}

type updateAssetContract struct {
	ByteCodePath string `json:"codepath"`
	Version      string `json:"version"`
}

type mintNFTAssetInfo struct {
	Id string `json:"id"`
	To string `json:"to"`
}

type transferNFTAssetInfo struct {
	Id string `json:"id"`
	To string `json:"to"`
}

type newEventCallbackUrl struct {
	EventCallbackUrl string `json:"callbackurl"`
	EventTopic       string `json:"callbackurl"`
}

func getNumAdminPkUsers(c *gin.Context) {
	data := gin.H{
		"number": myclient.NumAdminPkUsers(),
	}
	c.JSON(http.StatusOK, data)
}

func getAssetContractName(c *gin.Context) {
	data := gin.H{
		"name": NFT_ASSET_CONTRACT_NAME,
	}
	c.JSON(http.StatusOK, data)
}

func getCurrentPK(c *gin.Context) {
	data := gin.H{
		"pubkey": myclient.CurClientPK(),
	}
	c.JSON(http.StatusOK, data)
}

func getCurrentAddr(c *gin.Context) {
	data := gin.H{
		"address": myclient.CurClientEVMAddr(),
	}
	c.JSON(http.StatusOK, data)
}

func getHeight(c *gin.Context) {
	txid := c.Param("txid")
	height, err := myclient.GetBlockHeightByTxId(txid)

	if err != nil {
		data := gin.H{
			"txid":  height,
			"error": err.Error(),
		}
		c.JSON(http.StatusProcessing, data)
	} else {
		data := gin.H{
			"txid":   height,
			"height": height,
		}
		c.JSON(http.StatusOK, data)
	}
}

func getNFTAssetAdmins(c *gin.Context) {

	data := gin.H{
		"admins": myclient.GetNFTAssetAdmins(NFT_ASSET_CONTRACT_NAME),
	}
	c.JSON(http.StatusOK, data)
}

func getAssetName(c *gin.Context) {
	data := gin.H{
		"name": myclient.GetNFTAssetName(NFT_ASSET_CONTRACT_NAME),
	}
	c.JSON(http.StatusOK, data)
}

func getAssetSymbol(c *gin.Context) {
	data := gin.H{
		"symbol": myclient.GetNFTAssetSymbol(NFT_ASSET_CONTRACT_NAME),
	}
	c.JSON(http.StatusOK, data)
}

func getAssetTokenURI(c *gin.Context) {
	id := c.Param("id")
	data := gin.H{
		"tokenuri": myclient.GetNFTAssetTokenURI(NFT_ASSET_CONTRACT_NAME, id),
	}
	c.JSON(http.StatusOK, data)
}

func getAssetOwner(c *gin.Context) {
	id := c.Param("id")

	data := gin.H{
		"owner": myclient.GetNFTAssetOwner(NFT_ASSET_CONTRACT_NAME, id),
	}
	c.JSON(http.StatusOK, data)
}

func getAssetBalance(c *gin.Context) {
	address := c.Param("address")
	data := gin.H{
		"balance": myclient.GetNFTAssetBalance(NFT_ASSET_CONTRACT_NAME, address),
	}
	c.JSON(http.StatusOK, data)
}

func getNFTAssetSender(c *gin.Context) {
	data := gin.H{
		"sender": myclient.GetNFTAssetSender(NFT_ASSET_CONTRACT_NAME),
	}
	c.JSON(http.StatusOK, data)
}

func postNewUser(c *gin.Context) {
	name := c.Param("name")

	myclient.GenerateUser(name)
	buildConf(USER_CRYPTO_PATH, name, SDK_CONFIG_PATH)

	data := gin.H{
		"message": "user " + name + " generated",
	}
	c.JSON(http.StatusCreated, data)
}

func postBuildAssetContract(c *gin.Context) {
	var contractInfo newNFTAssetContract

	if err := c.BindJSON(&contractInfo); err != nil {
		return
	}

	err := myclient.CreateNFTAsset(NFT_ASSET_CONTRACT_NAME, NFT_BYTE_CODE_PATH, contractInfo.Version, contractInfo.Name, contractInfo.Symbol, contractInfo.TokenURI, contractInfo.OperatorAdminAddr)
	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusForbidden, data)
	} else {
		c.JSON(http.StatusCreated, contractInfo)
	}
}

func postUpdateAssetContract(c *gin.Context) {
	version := c.Param("version")

	err := myclient.UpdateAsset(NFT_ASSET_CONTRACT_NAME, NFT_BYTE_CODE_PATH, version)
	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusForbidden, data)
	} else {
		data := gin.H{
			"message": NFT_ASSET_CONTRACT_NAME + " updated to " + version,
		}
		c.JSON(http.StatusOK, data)
	}
}

func postFreezeAssetContract(c *gin.Context) {

	err := myclient.FreezeAsset(NFT_ASSET_CONTRACT_NAME)
	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusForbidden, data)
	} else {
		data := gin.H{
			"message": NFT_ASSET_CONTRACT_NAME + ":" + " Freezed",
		}
		c.JSON(http.StatusOK, data)
	}
}

func postUnfreezeAssetContract(c *gin.Context) {

	err := myclient.UnfreezeAsset(NFT_ASSET_CONTRACT_NAME)
	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusForbidden, data)
	} else {
		data := gin.H{
			"message": NFT_ASSET_CONTRACT_NAME + ":" + " Unfreezed",
		}
		c.JSON(http.StatusOK, data)
	}
}

func postRevokeAssetContract(c *gin.Context) {

	err := myclient.RevokeAsset(NFT_ASSET_CONTRACT_NAME)

	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusForbidden, data)
	} else {
		data := gin.H{
			"message": NFT_ASSET_CONTRACT_NAME + ":" + " Revoked",
		}
		c.JSON(http.StatusOK, data)
	}

}

func postAlterNFTAssetAdmins(c *gin.Context) {

	var admininfo newAdminsInfo

	if err := c.BindJSON(&admininfo); err != nil {
		return
	}

	txId, err := myclient.AlterNFTAssetAdmins(NFT_ASSET_CONTRACT_NAME, admininfo.Adminlist)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusForbidden, data)
	} else {
		data := gin.H{
			"txid":      txId,
			"adminlist": admininfo.Adminlist,
		}
		c.JSON(http.StatusOK, data)
	}
}

func postNFTAssetName(c *gin.Context) {
	name := c.Param("name")

	txId, err := myclient.SetNFTAssetTokenName(NFT_ASSET_CONTRACT_NAME, name)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		data := gin.H{
			"txid": txId,
			"name": name,
		}
		c.JSON(http.StatusCreated, data)
	}
}

func postNFTAssetSymbol(c *gin.Context) {
	symbol := c.Param("symbol")

	txId, err := myclient.SetNFTAssetTokenSymbol(NFT_ASSET_CONTRACT_NAME, symbol)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		data := gin.H{
			"txid":   txId,
			"symbol": symbol,
		}
		c.JSON(http.StatusCreated, data)
	}
}

type newUriLinkInfo struct {
	Link string `json:"link"`
}

func postNFTAssetUri(c *gin.Context) {
	var linkinfo newUriLinkInfo

	if err := c.BindJSON(&linkinfo); err != nil {
		return
	}

	txId, err := myclient.SetNFTAssetTokenURI(NFT_ASSET_CONTRACT_NAME, linkinfo.Link)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		data := gin.H{
			"txid":    txId,
			"urilink": linkinfo.Link,
		}
		c.JSON(http.StatusCreated, data)
	}
}

func postMintAsset(c *gin.Context) {
	var mintinfo mintNFTAssetInfo

	if err := c.BindJSON(&mintinfo); err != nil {
		return
	}

	txId, err := myclient.MintNFTAsset(NFT_ASSET_CONTRACT_NAME, mintinfo.To, mintinfo.Id, true)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		data := gin.H{
			"txid": txId,
			"id":   mintinfo.Id,
			"to":   mintinfo.To,
		}
		c.JSON(http.StatusCreated, data)
	}
}

func postMintAssetAsync(c *gin.Context) {
	var mintinfo mintNFTAssetInfo

	if err := c.BindJSON(&mintinfo); err != nil {
		return
	}

	txId, err := myclient.MintNFTAsset(NFT_ASSET_CONTRACT_NAME, mintinfo.To, mintinfo.Id, false)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		data := gin.H{
			"txid": txId,
			"id":   mintinfo.Id,
			"to":   mintinfo.To,
		}
		c.JSON(http.StatusCreated, data)
	}
}

func postTransferAsset(c *gin.Context) {
	var transinfo transferNFTAssetInfo

	if err := c.BindJSON(&transinfo); err != nil {
		return
	}

	txId, err := myclient.TransferNFTAsset(NFT_ASSET_CONTRACT_NAME, transinfo.To, transinfo.Id, true)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		data := gin.H{
			"txid": txId,
			"id":   transinfo.Id,
			"to":   transinfo.To,
		}
		c.JSON(http.StatusOK, data)
	}
}

func postTransferAssetAsync(c *gin.Context) {
	var transinfo transferNFTAssetInfo

	if err := c.BindJSON(&transinfo); err != nil {
		return
	}

	txId, err := myclient.TransferNFTAsset(NFT_ASSET_CONTRACT_NAME, transinfo.To, transinfo.Id, false)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		data := gin.H{
			"txid": txId,
			"id":   transinfo.Id,
			"to":   transinfo.To,
		}
		c.JSON(http.StatusOK, data)
	}
}

func postSetClient(c *gin.Context) {
	username := c.Param("username")

	// myclient.SetChainClientWithSDKConf(SDK_CONFIG_PATH, username)
	myclient.CurrentUser(username)

	data := gin.H{
		"name": username,
	}

	c.JSON(http.StatusOK, data)
}

func postSetCallBackUrl(c *gin.Context) {

	var callbackurl newEventCallbackUrl

	if err := c.BindJSON(&callbackurl); err != nil {
		return
	}

	eventTransferCallBackUrl = callbackurl.EventCallbackUrl
	callbackurl.EventTopic = "transfer"

	c.JSON(http.StatusOK, callbackurl)
}

func eventService(contractname string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, err := myclient.SubscribeContractEvent(ctx, 0, -1, contractname, "transfer")

	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {
		case event, ok := <-c:
			if !ok {
				fmt.Println("chan is close!")
				return
			}
			if event == nil {
				log.Fatalln("require not nil")
			}
			contractEventInfo, ok := event.(*common.ContractEventInfo)
			if !ok {
				log.Fatalln("require true")
			}
			fmt.Printf("recv contract event [%d] => %+v\n", contractEventInfo.BlockHeight, contractEventInfo)

			if eventTransferCallBackUrl != "" {
				callbackdata := RespContractEventInfo{contractEventInfo.BlockHeight, contractEventInfo.TxId, contractEventInfo.Topic, contractEventInfo.EventData}
				callbackMarshalled, err := json.Marshal(callbackdata)
				if err != nil {
					log.Fatalln("impossible to marshall teacher: %s", err)
				}
				responseBody := bytes.NewBuffer(callbackMarshalled)

				resp, err := http.Post(eventTransferCallBackUrl, "application/json", responseBody)
				if err != nil {
					log.Fatalln("An Error Occured %v", err)
				}
				defer resp.Body.Close()
				//Read the response body
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatalln(err)
				}
				sb := string(body)
				log.Printf(sb)
			}

		case <-ctx.Done():
			return
		}
	}

}

func getMultiAssetContractName(c *gin.Context) {
	data := gin.H{
		"name": MULTI_ASSET_CONTRACT_NAME,
	}
	c.JSON(http.StatusOK, data)
}

type multiAssetBalanceInfo struct {
	Owner string `form:"owner"`
	Id    string `form:"id"`
}

func getMultiAssetBalance(c *gin.Context) {
	var balanceInfo multiAssetBalanceInfo

	if err := c.ShouldBind(&balanceInfo); err != nil {
		return
	}

	// fmt.Printf("id: %v, owner: %v", balanceInfo.Id, balanceInfo.Owner)

	data := gin.H{
		"balance": myclient.GetMultiAssetBalance(MULTI_ASSET_CONTRACT_NAME, balanceInfo.Owner, balanceInfo.Id),
	}
	c.JSON(http.StatusOK, data)
}

type batchMultiAssetBalanceInfo struct {
	Owners []string `form:"owners"`
	Ids    []string `form:"ids"`
}

func getBatchMultiAssetBalance(c *gin.Context) {
	var balanceInfo batchMultiAssetBalanceInfo

	if err := c.BindJSON(&balanceInfo); err != nil {
		return
	}

	data := gin.H{
		"balance": myclient.GetBatchMultiAssetBalance(MULTI_ASSET_CONTRACT_NAME, balanceInfo.Owners, balanceInfo.Ids),
	}
	c.JSON(http.StatusOK, data)
}

func getMultiAssetUri(c *gin.Context) {
	id := c.Param("id")

	data := gin.H{
		"tokenuri": myclient.GetMultiAssetUri(MULTI_ASSET_CONTRACT_NAME, id),
	}
	c.JSON(http.StatusOK, data)
}

func getMultiAdmins(c *gin.Context) {

	data := gin.H{
		"admins": myclient.GetMultiAdmins(MULTI_ASSET_CONTRACT_NAME),
	}
	c.JSON(http.StatusOK, data)
}

func postMultiAssetUri(c *gin.Context) {
	var linkinfo newUriLinkInfo

	if err := c.BindJSON(&linkinfo); err != nil {
		return
	}

	txId, err := myclient.SetMultiAssetTokenURI(MULTI_ASSET_CONTRACT_NAME, linkinfo.Link)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		data := gin.H{
			"txid":    txId,
			"urilink": linkinfo.Link,
		}
		c.JSON(http.StatusCreated, data)
	}
}

type mintMultiAssetInfo struct {
	To     string `json:"to"`
	Id     string `json:"id"`
	Amount string `json:"amount"`
}

func postMintMultiAsset(c *gin.Context) {
	var mintinfo mintMultiAssetInfo

	if err := c.BindJSON(&mintinfo); err != nil {
		return
	}

	txId, err := myclient.MintMultiAsset(MULTI_ASSET_CONTRACT_NAME, mintinfo.To, mintinfo.Id, mintinfo.Amount, true)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		data := gin.H{
			"txid":   txId,
			"id":     mintinfo.Id,
			"amount": mintinfo.Amount,
			"to":     mintinfo.To,
		}
		c.JSON(http.StatusOK, data)
	}
}

func postMintMultiAssetAsync(c *gin.Context) {
	var mintinfo mintMultiAssetInfo

	if err := c.BindJSON(&mintinfo); err != nil {
		return
	}

	txId, err := myclient.MintMultiAsset(MULTI_ASSET_CONTRACT_NAME, mintinfo.To, mintinfo.Id, mintinfo.Amount, false)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		data := gin.H{
			"txid":   txId,
			"id":     mintinfo.Id,
			"amount": mintinfo.Amount,
			"to":     mintinfo.To,
		}
		c.JSON(http.StatusOK, data)
	}
}

type mintBatchMultiAssetInfo struct {
	To      string   `json:"to"`
	Ids     []string `json:"ids"`
	Amounts []string `json:"amounts"`
}

func postBatchMintMultiAsset(c *gin.Context) {
	var mintinfo mintBatchMultiAssetInfo

	if err := c.BindJSON(&mintinfo); err != nil {
		return
	}

	txId, err := myclient.BatchMintMultiAsset(MULTI_ASSET_CONTRACT_NAME, mintinfo.To, mintinfo.Ids, mintinfo.Amounts, true)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		data := gin.H{
			"txid":    txId,
			"ids":     mintinfo.Ids,
			"amounts": mintinfo.Amounts,
			"to":      mintinfo.To,
		}
		c.JSON(http.StatusCreated, data)
	}
}

func postBatchMintMultiAssetAsync(c *gin.Context) {
	var mintinfo mintBatchMultiAssetInfo

	if err := c.BindJSON(&mintinfo); err != nil {
		return
	}

	txId, err := myclient.BatchMintMultiAsset(MULTI_ASSET_CONTRACT_NAME, mintinfo.To, mintinfo.Ids, mintinfo.Amounts, false)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		data := gin.H{
			"txid":    txId,
			"ids":     mintinfo.Ids,
			"amounts": mintinfo.Amounts,
			"to":      mintinfo.To,
		}
		c.JSON(http.StatusCreated, data)
	}
}

type transferMultiInfo struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Id     string `json:"id"`
	Amount string `json:"amount"`
}

func postTransferMultiAsset(c *gin.Context) {
	var transinfo transferMultiInfo

	if err := c.BindJSON(&transinfo); err != nil {
		return
	}

	txId, err := myclient.TransferMultiAsset(MULTI_ASSET_CONTRACT_NAME, transinfo.From, transinfo.To,
		transinfo.Id, transinfo.Amount, true)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		data := gin.H{
			"txid": txId,
			"from": transinfo.From,
			"to":   transinfo.To,
			"id":   transinfo.Id,
		}
		c.JSON(http.StatusOK, data)
	}
}

func postTransferMultiAssetAsync(c *gin.Context) {
	var transinfo transferMultiInfo

	if err := c.BindJSON(&transinfo); err != nil {
		return
	}

	txId, err := myclient.TransferMultiAsset(MULTI_ASSET_CONTRACT_NAME, transinfo.From, transinfo.To,
		transinfo.Id, transinfo.Amount, false)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		data := gin.H{
			"txid": txId,
			"from": transinfo.From,
			"to":   transinfo.To,
			"id":   transinfo.Id,
		}
		c.JSON(http.StatusOK, data)
	}
}

type batchTransferMultiInfo struct {
	From    string   `json:"from"`
	To      string   `json:"to"`
	Ids     []string `json:"ids"`
	Amounts []string `json:"amounts"`
}

func postBatchTransferMultiAsset(c *gin.Context) {
	var transinfo batchTransferMultiInfo

	if err := c.BindJSON(&transinfo); err != nil {
		return
	}

	txId, err := myclient.BatchTransferMultiAsset(MULTI_ASSET_CONTRACT_NAME, transinfo.From, transinfo.To,
		transinfo.Ids, transinfo.Amounts, true)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		data := gin.H{
			"txid":    txId,
			"from":    transinfo.From,
			"to":      transinfo.To,
			"ids":     transinfo.Ids,
			"amounts": transinfo.Amounts,
		}
		c.JSON(http.StatusOK, data)
	}
}

func postBatchTransferMultiAssetAsync(c *gin.Context) {
	var transinfo batchTransferMultiInfo

	if err := c.BindJSON(&transinfo); err != nil {
		return
	}

	txId, err := myclient.BatchTransferMultiAsset(MULTI_ASSET_CONTRACT_NAME, transinfo.From, transinfo.To,
		transinfo.Ids, transinfo.Amounts, false)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		data := gin.H{
			"txid":    txId,
			"from":    transinfo.From,
			"to":      transinfo.To,
			"ids":     transinfo.Ids,
			"amounts": transinfo.Amounts,
		}
		c.JSON(http.StatusOK, data)
	}
}

type newAdminsInfo struct {
	Adminlist []string `json:"adminlist"`
}

func postAlterMultiAssetAdmins(c *gin.Context) {

	var admininfo newAdminsInfo

	if err := c.BindJSON(&admininfo); err != nil {
		return
	}

	txId, err := myclient.AlterMultiAssetAdmins(MULTI_ASSET_CONTRACT_NAME, admininfo.Adminlist)
	if err != nil {
		data := gin.H{
			"txid":  txId,
			"error": err.Error(),
		}
		c.JSON(http.StatusForbidden, data)
	} else {
		data := gin.H{
			"txid":      txId,
			"adminlist": admininfo.Adminlist,
		}
		c.JSON(http.StatusOK, data)
	}
}

func postBuildMultiAssetContract(c *gin.Context) {
	var contractInfo newMultiAssetContract

	if err := c.BindJSON(&contractInfo); err != nil {
		return
	}

	err := myclient.CreateMultiAsset(MULTI_ASSET_CONTRACT_NAME, MULTI_BYTE_CODE_PATH, contractInfo.Version, contractInfo.URI, contractInfo.OperatorAdminAddr)
	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusForbidden, data)
	} else {
		c.JSON(http.StatusCreated, contractInfo)
	}
}

func postUpdateMultiAssetContract(c *gin.Context) {
	version := c.Param("version")

	err := myclient.UpdateAsset(MULTI_ASSET_CONTRACT_NAME, MULTI_BYTE_CODE_PATH, version)
	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusForbidden, data)
	} else {
		data := gin.H{
			"message": MULTI_ASSET_CONTRACT_NAME + " updated to " + version,
		}
		c.JSON(http.StatusOK, data)
	}
}

func postFreezeMultiAssetContract(c *gin.Context) {

	err := myclient.FreezeAsset(MULTI_ASSET_CONTRACT_NAME)
	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusForbidden, data)
	} else {
		data := gin.H{
			"message": MULTI_ASSET_CONTRACT_NAME + ":" + " Freezed",
		}
		c.JSON(http.StatusOK, data)
	}
}

func postUnfreezeMultiAssetContract(c *gin.Context) {

	err := myclient.UnfreezeAsset(MULTI_ASSET_CONTRACT_NAME)
	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusForbidden, data)
	} else {
		data := gin.H{
			"message": MULTI_ASSET_CONTRACT_NAME + ":" + " Unfreezed",
		}
		c.JSON(http.StatusOK, data)
	}
}

func postRevokeMultiAssetContract(c *gin.Context) {

	err := myclient.RevokeAsset(MULTI_ASSET_CONTRACT_NAME)

	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusForbidden, data)
	} else {
		data := gin.H{
			"message": MULTI_ASSET_CONTRACT_NAME + ":" + " Revoked",
		}
		c.JSON(http.StatusOK, data)
	}

}

func atService() {
	router := gin.Default()
	router.GET("/admin/number", getNumAdminPkUsers)
	router.POST("/user/new/:name", postNewUser)
	router.GET("/user/pk", getCurrentPK)
	router.GET("/user/addr", getCurrentAddr)
	router.POST("/client/:username", postSetClient)
	router.GET("/height/:txid", getHeight)

	erc721 := router.Group("/nft")
	{
		erc721.GET("/contract/name", getAssetContractName)
		erc721.GET("/contract/asset/admins", getNFTAssetAdmins)
		erc721.GET("/contract/asset/name", getAssetName)
		erc721.GET("/contract/asset/symbol", getAssetSymbol)
		erc721.GET("/contract/asset/uri/:id", getAssetTokenURI)
		erc721.GET("/contract/asset/sender", getNFTAssetSender)
		erc721.GET("/contract/asset/owner/:id", getAssetOwner)
		erc721.GET("/contract/asset/balance/:address", getAssetBalance)

		erc721.POST("/user/callbackurl", postSetCallBackUrl)
		erc721.POST("/contract/asset/admins", postAlterNFTAssetAdmins)
		erc721.POST("/contract/asset/uri", postNFTAssetUri)
		erc721.POST("/contract/asset/name/:name", postNFTAssetName)
		erc721.POST("/contract/asset/symbol/:symbol", postNFTAssetSymbol)
		// router.POST("/contract/name/:contractname", postAssetContractName)
		erc721.POST("/contract/build", postBuildAssetContract)
		erc721.POST("/contract/update/:version", postUpdateAssetContract)
		erc721.POST("/contract/freeze", postFreezeAssetContract)
		erc721.POST("/contract/unfreeze", postUnfreezeAssetContract)
		// erc721.POST("/contract/revoke", postRevokeAssetContract)
		erc721.POST("/contract/asset/mint", postMintAsset)
		erc721.POST("/contract/asset/async/mint", postMintAssetAsync)
		erc721.POST("/contract/asset/transfer", postTransferAsset)
		erc721.POST("/contract/asset/async/transfer", postTransferAssetAsync)
	}

	erc1155 := router.Group("/multi")
	{
		erc1155.GET("/contract/name", getMultiAssetContractName)
		erc1155.GET("/contract/asset/balance", getMultiAssetBalance)
		// erc1155.GET("/contract/asset/batchbalance", getBatchMultiAssetBalance)
		erc1155.GET("/contract/asset/uri/:id", getMultiAssetUri)
		erc1155.GET("/contract/asset/admins", getMultiAdmins)

		erc1155.POST("/contract/build", postBuildMultiAssetContract)
		erc1155.POST("/contract/update/:version", postUpdateMultiAssetContract)
		erc1155.POST("/contract/freeze", postFreezeMultiAssetContract)
		erc1155.POST("/contract/unfreeze", postUnfreezeMultiAssetContract)
		// erc1155.POST("/contract/revoke", postRevokeMultiAssetContract)

		erc1155.POST("/contract/asset/uri", postMultiAssetUri)
		erc1155.POST("/contract/asset/mint", postMintMultiAsset)
		erc1155.POST("/contract/asset/async/mint", postMintMultiAssetAsync)
		erc1155.POST("/contract/asset/batchmint", postBatchMintMultiAsset)
		erc1155.POST("/contract/asset/async/batchmint", postBatchMintMultiAssetAsync)
		erc1155.POST("/contract/asset/transfer", postTransferMultiAsset)
		erc1155.POST("/contract/asset/async/transfer", postTransferMultiAssetAsync)
		erc1155.POST("/contract/asset/batchtransfer", postBatchTransferMultiAsset)
		erc1155.POST("/contract/asset/async/batchtransfer", postBatchTransferMultiAssetAsync)
		erc1155.POST("/contract/asset/admins", postAlterMultiAssetAdmins)
	}

	router.Run("localhost:7890")
}

func main() {

	rand.Seed(time.Now().UnixNano())

	fmt.Println("-------------- prepare --------------")
	prepareAdminUsers()
	buildConf(ADMIN1_CRYPTO_PATH, ADMIN1_NAME, SDK_CONFIG_PATH)
	buildConf(ADMIN2_CRYPTO_PATH, ADMIN2_NAME, SDK_CONFIG_PATH)
	buildConf(ADMIN3_CRYPTO_PATH, ADMIN3_NAME, SDK_CONFIG_PATH)
	buildConf(ADMIN4_CRYPTO_PATH, ADMIN4_NAME, SDK_CONFIG_PATH)
	myclient.SetChainClientWithSDKConf(SDK_CONFIG_PATH, ADMIN1_NAME)

	// go eventService(NFT_ASSET_CONTRACT_NAME)

	atService()
}

func prepareAdminUsers() {
	myclient.SetAdminPkUser(ADMIN1_NAME, "./crypto-config/node1/admin/admin1")
	myclient.SetAdminPkUser(ADMIN2_NAME, "./crypto-config/node1/admin/admin2")
	myclient.SetAdminPkUser(ADMIN3_NAME, "./crypto-config/node1/admin/admin3")
	myclient.SetAdminPkUser(ADMIN4_NAME, "./crypto-config/node1/admin/admin4")

	fmt.Println("admin number: %d", myclient.NumAdminPkUsers())
}

func buildConf(keyPath, keyName, confPath string) {
	var nodeConf []myclient.PkSdkNodeConf
	nodeItem := new(myclient.PkSdkNodeConf)
	nodeItem.NodeAddr = "127.0.0.1:12301"
	nodeItem.ConnCnt = 10
	nodeConf = append(nodeConf, *nodeItem)

	nodeItem.NodeAddr = "127.0.0.1:12302"
	nodeItem.ConnCnt = 10
	nodeConf = append(nodeConf, *nodeItem)

	nodeItem.NodeAddr = "127.0.0.1:12303"
	nodeItem.ConnCnt = 10
	nodeConf = append(nodeConf, *nodeItem)

	nodeItem.NodeAddr = "127.0.0.1:12304"
	nodeItem.ConnCnt = 10
	nodeConf = append(nodeConf, *nodeItem)

	dest := confPath + "/sdk_config_pk_" + keyName + ".yml"
	gotSDKConfPath, err := myclient.SetSDKConf("chain1", keyPath, keyName, confPath, nodeConf)
	if gotSDKConfPath != dest || err != nil {
		fmt.Errorf("sdk conf path = %s, expected %s", gotSDKConfPath, dest)
	}

	_, err = myclient.ReadSDKConf(confPath, keyName)
	if err != nil {
		fmt.Errorf("read sdk config failed")
	}

	readFile, err := ioutil.ReadFile(dest)
	if err != nil {
		fmt.Errorf("target config file not found")
	} else {
		fmt.Println(string(readFile))
	}

}
