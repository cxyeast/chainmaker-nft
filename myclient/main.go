package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"myclient/myclient"
	"net/http"
	"time"

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

	USER_CRYPTO_PATH = "./crypto-config/users"
	SDK_CONFIG_PATH  = "./sdk-configs"
	BYTE_CODE_PATH   = "contract/erc721.7z"
)

type userConfig struct {
	UserName      string `json:"name"`
	UserCrytoPath string `json:"path"`
}

type newAssetContract struct {
	// ByteCodePath      string `json:"codepath"`
	Version           string `json:"version"`
	Name              string `json:"name"`
	Symbol            string `json:"symbol"`
	TokenURI          string `json:"tokenuri"`
	OperatorAdminAddr string `json:"opadmin"`
}

type updateAssetContract struct {
	ByteCodePath string `json:"codepath"`
	Version      string `json:"version"`
}

type mintInfo struct {
	Tokenid string `json:"tokenid"`
	Address string `json:"address"`
}

type transferInfo struct {
	Tokenid string `json:"tokenid"`
	To      string `json:"to"`
}

func getNumAdminPkUsers(c *gin.Context) {
	data := gin.H{
		"number": myclient.NumAdminPkUsers(),
	}
	c.JSON(http.StatusOK, data)
}

func getAssetContractName(c *gin.Context) {
	data := gin.H{
		"name": myclient.GetAssetContractName(),
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

func getAssetName(c *gin.Context) {
	data := gin.H{
		"name": myclient.GetName(),
	}
	c.JSON(http.StatusOK, data)
}

func getAssetSymbol(c *gin.Context) {
	data := gin.H{
		"symbol": myclient.GetSymbol(),
	}
	c.JSON(http.StatusOK, data)
}

func getAssetTokenURI(c *gin.Context) {
	data := gin.H{
		"tokenuri": myclient.GetTokenURI(),
	}
	c.JSON(http.StatusOK, data)
}

func getAssetOwner(c *gin.Context) {
	id := c.Param("tokenid")

	data := gin.H{
		"owner": myclient.GetOwner(id),
	}
	c.JSON(http.StatusOK, data)
}

func getAssetBalance(c *gin.Context) {
	address := c.Param("address")
	data := gin.H{
		"balance": myclient.GetBalance(address),
	}
	c.JSON(http.StatusOK, data)
}

func getAssetSender(c *gin.Context) {
	data := gin.H{
		"sender": myclient.GetSender(),
	}
	c.JSON(http.StatusOK, data)
}

func postNewUser(c *gin.Context) {
	name := c.Param("name")

	myclient.GeneratePKFile(USER_CRYPTO_PATH, name)
	buildConf(USER_CRYPTO_PATH, name, SDK_CONFIG_PATH)

	data := gin.H{
		"message": "user " + name + " generated",
	}
	c.JSON(http.StatusCreated, data)
}

func postAssetContractName(c *gin.Context) {
	contractname := c.Param("contractname")

	myclient.SetAssetContractName(contractname)

	data := gin.H{
		"name": contractname,
	}

	c.JSON(http.StatusCreated, data)
}

func postBuildAssetContract(c *gin.Context) {
	var contractInfo newAssetContract

	if err := c.BindJSON(&contractInfo); err != nil {
		return
	}

	err := myclient.CreateNFTAsset(BYTE_CODE_PATH, contractInfo.Version, contractInfo.Name, contractInfo.Symbol, contractInfo.TokenURI, contractInfo.OperatorAdminAddr)
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

	err := myclient.UpdateNFTAsset(BYTE_CODE_PATH, version)
	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusForbidden, data)
	} else {
		data := gin.H{
			"message": myclient.GetAssetContractName() + " updated to " + version,
		}
		c.JSON(http.StatusOK, data)
	}

}

func postFreezeAssetContract(c *gin.Context) {

	err := myclient.FreezeNFTAsset()
	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusForbidden, data)
	} else {
		data := gin.H{
			"message": myclient.GetAssetContractName() + ":" + " Freezed",
		}
		c.JSON(http.StatusOK, data)
	}
}

func postUnfreezeAssetContract(c *gin.Context) {

	err := myclient.UnfreezeNFTAsset()
	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusForbidden, data)
	} else {
		data := gin.H{
			"message": myclient.GetAssetContractName() + ":" + " Unfreezed",
		}
		c.JSON(http.StatusOK, data)
	}
}

func postRevokeAssetContract(c *gin.Context) {

	err := myclient.RevokeNFTAsset()

	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusForbidden, data)
	} else {
		data := gin.H{
			"message": myclient.GetAssetContractName() + ":" + " Revoked",
		}
		c.JSON(http.StatusOK, data)
	}

}

func postMintAsset(c *gin.Context) {
	var mintinfo mintInfo

	if err := c.BindJSON(&mintinfo); err != nil {
		return
	}

	err := myclient.MintAsset(mintinfo.Address, mintinfo.Tokenid)
	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		c.JSON(http.StatusCreated, mintinfo)
	}
}

func postTransferAsset(c *gin.Context) {
	var transinfo transferInfo

	if err := c.BindJSON(&transinfo); err != nil {
		return
	}

	err := myclient.TransferAsset(transinfo.To, transinfo.Tokenid)
	if err != nil {
		data := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
	} else {
		c.JSON(http.StatusOK, transinfo)
	}
}

func postSetClient(c *gin.Context) {
	username := c.Param("username")

	myclient.SetChainClientWithSDKConf(SDK_CONFIG_PATH, username)

	data := gin.H{
		"name": username,
	}

	c.JSON(http.StatusOK, data)
}

func main() {

	rand.Seed(time.Now().UnixNano())

	fmt.Println("-------------- prepare --------------")
	prepareAdminUsers()
	buildConf(ADMIN1_CRYPTO_PATH, ADMIN1_NAME, SDK_CONFIG_PATH)
	buildConf(ADMIN2_CRYPTO_PATH, ADMIN2_NAME, SDK_CONFIG_PATH)
	buildConf(ADMIN3_CRYPTO_PATH, ADMIN3_NAME, SDK_CONFIG_PATH)
	buildConf(ADMIN4_CRYPTO_PATH, ADMIN4_NAME, SDK_CONFIG_PATH)

	router := gin.Default()

	router.GET("/admin/number", getNumAdminPkUsers)
	router.GET("/contract/name", getAssetContractName)
	router.GET("/user/pk", getCurrentPK)
	router.GET("/user/addr", getCurrentAddr)
	router.GET("/contract/asset/name", getAssetName)
	router.GET("/contract/asset/symbol", getAssetSymbol)
	router.GET("/contract/asset/tokenuri", getAssetTokenURI)
	router.GET("/contract/asset/sender", getAssetSender)
	router.GET("/contract/asset/owner/:tokenid", getAssetOwner)
	router.GET("/contract/asset/balance/:address", getAssetBalance)

	router.POST("/user/new/:name", postNewUser)
	router.POST("/contract/name/:contractname", postAssetContractName)
	router.POST("/client/:username", postSetClient)
	router.POST("/contract/build", postBuildAssetContract)
	router.POST("/contract/update/:version", postUpdateAssetContract)
	router.POST("/contract/freeze", postFreezeAssetContract)
	router.POST("/contract/unfreeze", postUnfreezeAssetContract)
	router.POST("/contract/revoke", postRevokeAssetContract)
	router.POST("/contract/asset/mint", postMintAsset)
	router.POST("/contract/asset/transfer", postTransferAsset)

	router.Run("localhost:7890")
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
