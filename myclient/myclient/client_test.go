package myclient

import (
	"crypto/sha256"

	"chainmaker.org/chainmaker/common/v2/crypto/asym"
	"testing"
)

func TestAddPkUsersFirstly(t *testing.T) {

	// add admin1
	wantPkUserKey := "./crypto-config/node1/admin/admin1/admin1.key"
	wantNum := 1
	wantExisted := true
	SetAdminPkUser("admin1", wantPkUserKey)
	keyPem, memberInfo, _ := GetAdminPkUser("admin1")
	digest := sha256.Sum256([]byte("Blockchain"))

	sign, err := asym.Sign(keyPem, digest[:])
	ok, err := asym.Verify(memberInfo, digest[:], sign)
	if err != nil || !ok {
		t.Errorf("PK/SK Signature Verifying Error!")
	}

	gotNum := NumAdminPkUsers()
	gotExisted := AdminExisted("admin1")
	if gotNum != wantNum {
		t.Errorf("NumAdminPkUsers = %d; want %d", gotNum, wantNum)
	}

	if gotExisted != wantExisted {
		t.Errorf("gotExisted = %v; wantExisted %v", gotExisted, wantExisted)
	}

	// add admin2
	wantPkUserKey = "./crypto-config/node1/admin/admin2/admin2.key"
	wantNum = 2
	wantExisted = true
	SetAdminPkUser("admin2", wantPkUserKey)
	keyPem, memberInfo, _ = GetAdminPkUser("admin2")

	sign, err = asym.Sign(keyPem, digest[:])
	ok, err = asym.Verify(memberInfo, digest[:], sign)
	if err != nil || !ok {
		t.Errorf("PK/SK Signature Verifying Error!")
	}

	gotNum = NumAdminPkUsers()
	gotExisted = AdminExisted("admin2")
	if gotNum != wantNum {
		t.Errorf("NumAdminPkUsers = %d; want %d", gotNum, wantNum)
	}

	if gotExisted != wantExisted {
		t.Errorf("gotExisted = %v; wantExisted %v", gotExisted, wantExisted)
	}

	// add admin3
	wantPkUserKey = "./crypto-config/node1/admin/admin3/admin3.key"
	wantNum = 3
	wantExisted = true
	SetAdminPkUser("admin3", wantPkUserKey)
	keyPem, memberInfo, _ = GetAdminPkUser("admin3")

	sign, err = asym.Sign(keyPem, digest[:])
	ok, err = asym.Verify(memberInfo, digest[:], sign)
	if err != nil || !ok {
		t.Errorf("PK/SK Signature Verifying Error!")
	}

	gotNum = NumAdminPkUsers()
	gotExisted = AdminExisted("admin3")
	if gotNum != wantNum {
		t.Errorf("NumAdminPkUsers = %d; want %d", gotNum, wantNum)
	}

	if gotExisted != wantExisted {
		t.Errorf("gotExisted = %v; wantExisted %v", gotExisted, wantExisted)
	}

	// add admin4
	wantPkUserKey = "./crypto-config/node1/admin/admin4/admin4.key"
	wantNum = 4
	wantExisted = true
	SetAdminPkUser("admin4", wantPkUserKey)
	keyPem, memberInfo, _ = GetAdminPkUser("admin4")

	sign, err = asym.Sign(keyPem, digest[:])
	ok, err = asym.Verify(memberInfo, digest[:], sign)
	if err != nil || !ok {
		t.Errorf("PK/SK Signature Verifying Error!")
	}

	gotNum = NumAdminPkUsers()
	gotExisted = AdminExisted("admin4")
	if gotNum != wantNum {
		t.Errorf("NumAdminPkUsers = %d; want %d", gotNum, wantNum)
	}

	if gotExisted != wantExisted {
		t.Errorf("gotExisted = %v; wantExisted %v", gotExisted, wantExisted)
	}

	// add admin5
	wantPkUserKey = "./crypto-config/node1/admin/admin5/admin5.key"
	wantNum = 5
	wantExisted = true
	SetAdminPkUser("admin5", wantPkUserKey)
	keyPem, memberInfo, _ = GetAdminPkUser("admin5")

	sign, err = asym.Sign(keyPem, digest[:])
	ok, err = asym.Verify(memberInfo, digest[:], sign)
	if err != nil || !ok {
		t.Errorf("PK/SK Signature Verifying Error!")
	}

	gotNum = NumAdminPkUsers()
	gotExisted = AdminExisted("admin5")
	if gotNum != wantNum {
		t.Errorf("NumAdminPkUsers = %d; want %d", gotNum, wantNum)
	}

	if gotExisted != wantExisted {
		t.Errorf("gotExisted = %v; wantExisted %v", gotExisted, wantExisted)
	}
}

func TestDelPkUsers(t *testing.T) {
	wantNum := 4
	DelAdminPkUser("admin5")
	gotNum := NumAdminPkUsers()
	if gotNum != wantNum {
		t.Errorf("NumAdminPkUsers = %d; want %d", gotNum, wantNum)
	}
}

func TestClearPkUsers(t *testing.T) {
	wantNum := 0
	ClearAdminPkUsers()
	gotNum := NumAdminPkUsers()
	if gotNum != wantNum {
		t.Errorf("NumAdminPkUsers = %d; want %d", gotNum, wantNum)
	}
}

func TestAddPkUsersSecondly(t *testing.T) {

	// add admin1
	wantPkUserKey := "./crypto-config/node1/admin/admin1/admin1.key"
	wantNum := 1
	wantExisted := true
	SetAdminPkUser("admin1", wantPkUserKey)
	keyPem, memberInfo, _ := GetAdminPkUser("admin1")
	digest := sha256.Sum256([]byte("Blockchain"))

	sign, err := asym.Sign(keyPem, digest[:])
	ok, err := asym.Verify(memberInfo, digest[:], sign)
	if err != nil || !ok {
		t.Errorf("PK/SK Signature Verifying Error!")
	}

	gotNum := NumAdminPkUsers()
	gotExisted := AdminExisted("admin1")
	if gotNum != wantNum {
		t.Errorf("NumAdminPkUsers = %d; want %d", gotNum, wantNum)
	}

	if gotExisted != wantExisted {
		t.Errorf("gotExisted = %v; wantExisted %v", gotExisted, wantExisted)
	}

	// add admin2
	wantPkUserKey = "./crypto-config/node1/admin/admin2/admin2.key"
	wantNum = 2
	wantExisted = true
	SetAdminPkUser("admin2", wantPkUserKey)
	keyPem, memberInfo, _ = GetAdminPkUser("admin2")

	sign, err = asym.Sign(keyPem, digest[:])
	ok, err = asym.Verify(memberInfo, digest[:], sign)
	if err != nil || !ok {
		t.Errorf("PK/SK Signature Verifying Error!")
	}

	gotNum = NumAdminPkUsers()
	gotExisted = AdminExisted("admin2")
	if gotNum != wantNum {
		t.Errorf("NumAdminPkUsers = %d; want %d", gotNum, wantNum)
	}

	if gotExisted != wantExisted {
		t.Errorf("gotExisted = %v; wantExisted %v", gotExisted, wantExisted)
	}

	// add admin3
	wantPkUserKey = "./crypto-config/node1/admin/admin3/admin3.key"
	wantNum = 3
	wantExisted = true
	SetAdminPkUser("admin3", wantPkUserKey)
	keyPem, memberInfo, _ = GetAdminPkUser("admin3")

	sign, err = asym.Sign(keyPem, digest[:])
	ok, err = asym.Verify(memberInfo, digest[:], sign)
	if err != nil || !ok {
		t.Errorf("PK/SK Signature Verifying Error!")
	}

	gotNum = NumAdminPkUsers()
	gotExisted = AdminExisted("admin3")
	if gotNum != wantNum {
		t.Errorf("NumAdminPkUsers = %d; want %d", gotNum, wantNum)
	}

	if gotExisted != wantExisted {
		t.Errorf("gotExisted = %v; wantExisted %v", gotExisted, wantExisted)
	}

	// add admin4
	wantPkUserKey = "./crypto-config/node1/admin/admin4/admin4.key"
	wantNum = 4
	wantExisted = true
	SetAdminPkUser("admin4", wantPkUserKey)
	keyPem, memberInfo, _ = GetAdminPkUser("admin4")

	sign, err = asym.Sign(keyPem, digest[:])
	ok, err = asym.Verify(memberInfo, digest[:], sign)
	if err != nil || !ok {
		t.Errorf("PK/SK Signature Verifying Error!")
	}

	gotNum = NumAdminPkUsers()
	gotExisted = AdminExisted("admin4")
	if gotNum != wantNum {
		t.Errorf("NumAdminPkUsers = %d; want %d", gotNum, wantNum)
	}

	if gotExisted != wantExisted {
		t.Errorf("gotExisted = %v; wantExisted %v", gotExisted, wantExisted)
	}
}

func TestAdmin1SDKConf(t *testing.T) {
	var nodeConf []PkSdkNodeConf
	nodeItem := new(PkSdkNodeConf)
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

	gotSDKConfPath, err := SetSDKConf("chain1", "./crypto-config/node1/admin/admin1", "admin1.key", "./sdk-configs", "sdk_config_pk_admin1.yml", nodeConf)
	if gotSDKConfPath != "./sdk-configs/sdk_config_pk_admin1.yml" {
		t.Errorf("sdk conf path = %s, expected %s", gotSDKConfPath, "./sdk-configs/sdk_config_pk_admin1.yml")
	}
	gotSDKConf, err := ReadSDKConf("./sdk-configs", "sdk_config_pk_admin1.yml")
	if err != nil {
		t.Errorf("read sdk config failed")
	}

	wantSDKConf := &SdkPkConfig{
		ChainClient: PkChainClientConf{
			ChainId:             "chain1",
			UserSignKeyFilePath: "./crypto-config/node1/admin/admin1/admin1.key",
			AuthType:            "public",
			EnableNormalKey:     false,
			Crypto: CryptoHashConf{
				Hash: "SHA256",
			},
			Archive: ArchiveConf{
				Type:      "mysql",
				Dest:      "root:123456:localhost:3306",
				SecretKey: "xxx",
			},
			RpcClient: RpcClientConf{
				MaxReceiveMessageSize: 16,
				MaxSendMessageSize:    16,
			},
			Nodes: []PkSdkNodeConf{
				{
					NodeAddr: "127.0.0.1:12301",
					ConnCnt:  10,
				},
				{
					NodeAddr: "127.0.0.1:12302",
					ConnCnt:  10,
				},
				{
					NodeAddr: "127.0.0.1:12303",
					ConnCnt:  10,
				},
				{
					NodeAddr: "127.0.0.1:12304",
					ConnCnt:  10,
				},
			},
		},
	}

	if wantSDKConf.ChainClient.ChainId != gotSDKConf.ChainClient.ChainId {
		t.Errorf("ChainId = %s; want %s", gotSDKConf.ChainClient.ChainId, wantSDKConf.ChainClient.ChainId)
	}

	if wantSDKConf.ChainClient.UserSignKeyFilePath != gotSDKConf.ChainClient.UserSignKeyFilePath {
		t.Errorf("UserSignKeyFilePath = %s; want %s", gotSDKConf.ChainClient.UserSignKeyFilePath, wantSDKConf.ChainClient.UserSignKeyFilePath)
	}

	if wantSDKConf.ChainClient.AuthType != gotSDKConf.ChainClient.AuthType {
		t.Errorf("AuthType = %s; want %s", gotSDKConf.ChainClient.AuthType, wantSDKConf.ChainClient.AuthType)
	}

	if wantSDKConf.ChainClient.EnableNormalKey != gotSDKConf.ChainClient.EnableNormalKey {
		t.Errorf("EnableNormalKey = %v; want %v", gotSDKConf.ChainClient.EnableNormalKey, wantSDKConf.ChainClient.EnableNormalKey)
	}

	if len(wantSDKConf.ChainClient.Nodes) != len(gotSDKConf.ChainClient.Nodes) {
		t.Errorf("nodes len = %d; want %d", len(gotSDKConf.ChainClient.Nodes), len(wantSDKConf.ChainClient.Nodes))
	} else {
		for i := 0; i < len(wantSDKConf.ChainClient.Nodes); i++ {
			if wantSDKConf.ChainClient.Nodes[i].NodeAddr != gotSDKConf.ChainClient.Nodes[i].NodeAddr {
				t.Errorf("got node addr = %s; want node addr %s", gotSDKConf.ChainClient.Nodes[i].NodeAddr, wantSDKConf.ChainClient.Nodes[i].NodeAddr)
			}
			if wantSDKConf.ChainClient.Nodes[i].ConnCnt != gotSDKConf.ChainClient.Nodes[i].ConnCnt {
				t.Errorf("got node ConnCnt = %d; want node ConnCnt %d", gotSDKConf.ChainClient.Nodes[i].ConnCnt, wantSDKConf.ChainClient.Nodes[i].ConnCnt)
			}
		}
	}

	if (wantSDKConf.ChainClient.Archive.Type != gotSDKConf.ChainClient.Archive.Type) ||
		(wantSDKConf.ChainClient.Archive.Dest != gotSDKConf.ChainClient.Archive.Dest) ||
		(wantSDKConf.ChainClient.Archive.SecretKey != gotSDKConf.ChainClient.Archive.SecretKey) {
		t.Errorf("got archive Type = %s; want archive Type %s", gotSDKConf.ChainClient.Archive.Type, wantSDKConf.ChainClient.Archive.Type)
		t.Errorf("got archive Dest = %s; want archive Dest %s", gotSDKConf.ChainClient.Archive.Dest, wantSDKConf.ChainClient.Archive.Dest)
		t.Errorf("got archive SecretKey = %s; want archive SecretKey %s", gotSDKConf.ChainClient.Archive.SecretKey, wantSDKConf.ChainClient.Archive.SecretKey)
	}

	if (wantSDKConf.ChainClient.RpcClient.MaxReceiveMessageSize != gotSDKConf.ChainClient.RpcClient.MaxReceiveMessageSize) ||
		(wantSDKConf.ChainClient.RpcClient.MaxSendMessageSize != gotSDKConf.ChainClient.RpcClient.MaxSendMessageSize) {
		t.Errorf("got archive Type = %d; want RpcClient MaxReceiveMessageSize %d", gotSDKConf.ChainClient.RpcClient.MaxReceiveMessageSize, wantSDKConf.ChainClient.RpcClient.MaxReceiveMessageSize)
		t.Errorf("got archive Dest = %d; want RpcClient MaxSendMessageSize %d", gotSDKConf.ChainClient.RpcClient.MaxSendMessageSize, wantSDKConf.ChainClient.RpcClient.MaxSendMessageSize)
	}
}

func TestAdmin2SDKConf(t *testing.T) {
	var nodeConf []PkSdkNodeConf
	nodeItem := new(PkSdkNodeConf)
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

	gotSDKConfPath, err := SetSDKConf("chain1", "./crypto-config/node1/admin/admin2", "admin2.key", "./sdk-configs", "sdk_config_pk_admin2.yml", nodeConf)
	if gotSDKConfPath != "./sdk-configs/sdk_config_pk_admin2.yml" {
		t.Errorf("sdk conf path = %s, expected %s", gotSDKConfPath, "./sdk-configs/sdk_config_pk_admin2.yml")
	}
	gotSDKConf, err := ReadSDKConf("./sdk-configs", "sdk_config_pk_admin2.yml")
	if err != nil {
		t.Errorf("read sdk config failed")
	}

	wantSDKConf := &SdkPkConfig{
		ChainClient: PkChainClientConf{
			ChainId:             "chain1",
			UserSignKeyFilePath: "./crypto-config/node1/admin/admin2/admin2.key",
			AuthType:            "public",
			EnableNormalKey:     false,
			Crypto: CryptoHashConf{
				Hash: "SHA256",
			},
			Archive: ArchiveConf{
				Type:      "mysql",
				Dest:      "root:123456:localhost:3306",
				SecretKey: "xxx",
			},
			RpcClient: RpcClientConf{
				MaxReceiveMessageSize: 16,
				MaxSendMessageSize:    16,
			},
			Nodes: []PkSdkNodeConf{
				{
					NodeAddr: "127.0.0.1:12301",
					ConnCnt:  10,
				},
				{
					NodeAddr: "127.0.0.1:12302",
					ConnCnt:  10,
				},
				{
					NodeAddr: "127.0.0.1:12303",
					ConnCnt:  10,
				},
				{
					NodeAddr: "127.0.0.1:12304",
					ConnCnt:  10,
				},
			},
		},
	}

	if wantSDKConf.ChainClient.ChainId != gotSDKConf.ChainClient.ChainId {
		t.Errorf("ChainId = %s; want %s", gotSDKConf.ChainClient.ChainId, wantSDKConf.ChainClient.ChainId)
	}

	if wantSDKConf.ChainClient.UserSignKeyFilePath != gotSDKConf.ChainClient.UserSignKeyFilePath {
		t.Errorf("UserSignKeyFilePath = %s; want %s", gotSDKConf.ChainClient.UserSignKeyFilePath, wantSDKConf.ChainClient.UserSignKeyFilePath)
	}

	if wantSDKConf.ChainClient.AuthType != gotSDKConf.ChainClient.AuthType {
		t.Errorf("AuthType = %s; want %s", gotSDKConf.ChainClient.AuthType, wantSDKConf.ChainClient.AuthType)
	}

	if wantSDKConf.ChainClient.EnableNormalKey != gotSDKConf.ChainClient.EnableNormalKey {
		t.Errorf("EnableNormalKey = %v; want %v", gotSDKConf.ChainClient.EnableNormalKey, wantSDKConf.ChainClient.EnableNormalKey)
	}

	if len(wantSDKConf.ChainClient.Nodes) != len(gotSDKConf.ChainClient.Nodes) {
		t.Errorf("nodes len = %d; want %d", len(gotSDKConf.ChainClient.Nodes), len(wantSDKConf.ChainClient.Nodes))
	} else {
		for i := 0; i < len(wantSDKConf.ChainClient.Nodes); i++ {
			if wantSDKConf.ChainClient.Nodes[i].NodeAddr != gotSDKConf.ChainClient.Nodes[i].NodeAddr {
				t.Errorf("got node addr = %s; want node addr %s", gotSDKConf.ChainClient.Nodes[i].NodeAddr, wantSDKConf.ChainClient.Nodes[i].NodeAddr)
			}
			if wantSDKConf.ChainClient.Nodes[i].ConnCnt != gotSDKConf.ChainClient.Nodes[i].ConnCnt {
				t.Errorf("got node ConnCnt = %d; want node ConnCnt %d", gotSDKConf.ChainClient.Nodes[i].ConnCnt, wantSDKConf.ChainClient.Nodes[i].ConnCnt)
			}
		}
	}

	if (wantSDKConf.ChainClient.Archive.Type != gotSDKConf.ChainClient.Archive.Type) ||
		(wantSDKConf.ChainClient.Archive.Dest != gotSDKConf.ChainClient.Archive.Dest) ||
		(wantSDKConf.ChainClient.Archive.SecretKey != gotSDKConf.ChainClient.Archive.SecretKey) {
		t.Errorf("got archive Type = %s; want archive Type %s", gotSDKConf.ChainClient.Archive.Type, wantSDKConf.ChainClient.Archive.Type)
		t.Errorf("got archive Dest = %s; want archive Dest %s", gotSDKConf.ChainClient.Archive.Dest, wantSDKConf.ChainClient.Archive.Dest)
		t.Errorf("got archive SecretKey = %s; want archive SecretKey %s", gotSDKConf.ChainClient.Archive.SecretKey, wantSDKConf.ChainClient.Archive.SecretKey)
	}

	if (wantSDKConf.ChainClient.RpcClient.MaxReceiveMessageSize != gotSDKConf.ChainClient.RpcClient.MaxReceiveMessageSize) ||
		(wantSDKConf.ChainClient.RpcClient.MaxSendMessageSize != gotSDKConf.ChainClient.RpcClient.MaxSendMessageSize) {
		t.Errorf("got archive Type = %d; want RpcClient MaxReceiveMessageSize %d", gotSDKConf.ChainClient.RpcClient.MaxReceiveMessageSize, wantSDKConf.ChainClient.RpcClient.MaxReceiveMessageSize)
		t.Errorf("got archive Dest = %d; want RpcClient MaxSendMessageSize %d", gotSDKConf.ChainClient.RpcClient.MaxSendMessageSize, wantSDKConf.ChainClient.RpcClient.MaxSendMessageSize)
	}
}

func TestAdmin3SDKConf(t *testing.T) {
	var nodeConf []PkSdkNodeConf
	nodeItem := new(PkSdkNodeConf)
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

	gotSDKConfPath, err := SetSDKConf("chain1", "./crypto-config/node1/admin/admin3", "admin3.key", "./sdk-configs", "sdk_config_pk_admin3.yml", nodeConf)
	if gotSDKConfPath != "./sdk-configs/sdk_config_pk_admin3.yml" {
		t.Errorf("sdk conf path = %s, expected %s", gotSDKConfPath, "./sdk-configs/sdk_config_pk_admin3.yml")
	}
	gotSDKConf, err := ReadSDKConf("./sdk-configs", "sdk_config_pk_admin3.yml")
	if err != nil {
		t.Errorf("read sdk config failed")
	}

	wantSDKConf := &SdkPkConfig{
		ChainClient: PkChainClientConf{
			ChainId:             "chain1",
			UserSignKeyFilePath: "./crypto-config/node1/admin/admin3/admin3.key",
			AuthType:            "public",
			EnableNormalKey:     false,
			Crypto: CryptoHashConf{
				Hash: "SHA256",
			},
			Archive: ArchiveConf{
				Type:      "mysql",
				Dest:      "root:123456:localhost:3306",
				SecretKey: "xxx",
			},
			RpcClient: RpcClientConf{
				MaxReceiveMessageSize: 16,
				MaxSendMessageSize:    16,
			},
			Nodes: []PkSdkNodeConf{
				{
					NodeAddr: "127.0.0.1:12301",
					ConnCnt:  10,
				},
				{
					NodeAddr: "127.0.0.1:12302",
					ConnCnt:  10,
				},
				{
					NodeAddr: "127.0.0.1:12303",
					ConnCnt:  10,
				},
				{
					NodeAddr: "127.0.0.1:12304",
					ConnCnt:  10,
				},
			},
		},
	}

	if wantSDKConf.ChainClient.ChainId != gotSDKConf.ChainClient.ChainId {
		t.Errorf("ChainId = %s; want %s", gotSDKConf.ChainClient.ChainId, wantSDKConf.ChainClient.ChainId)
	}

	if wantSDKConf.ChainClient.UserSignKeyFilePath != gotSDKConf.ChainClient.UserSignKeyFilePath {
		t.Errorf("UserSignKeyFilePath = %s; want %s", gotSDKConf.ChainClient.UserSignKeyFilePath, wantSDKConf.ChainClient.UserSignKeyFilePath)
	}

	if wantSDKConf.ChainClient.AuthType != gotSDKConf.ChainClient.AuthType {
		t.Errorf("AuthType = %s; want %s", gotSDKConf.ChainClient.AuthType, wantSDKConf.ChainClient.AuthType)
	}

	if wantSDKConf.ChainClient.EnableNormalKey != gotSDKConf.ChainClient.EnableNormalKey {
		t.Errorf("EnableNormalKey = %v; want %v", gotSDKConf.ChainClient.EnableNormalKey, wantSDKConf.ChainClient.EnableNormalKey)
	}

	if len(wantSDKConf.ChainClient.Nodes) != len(gotSDKConf.ChainClient.Nodes) {
		t.Errorf("nodes len = %d; want %d", len(gotSDKConf.ChainClient.Nodes), len(wantSDKConf.ChainClient.Nodes))
	} else {
		for i := 0; i < len(wantSDKConf.ChainClient.Nodes); i++ {
			if wantSDKConf.ChainClient.Nodes[i].NodeAddr != gotSDKConf.ChainClient.Nodes[i].NodeAddr {
				t.Errorf("got node addr = %s; want node addr %s", gotSDKConf.ChainClient.Nodes[i].NodeAddr, wantSDKConf.ChainClient.Nodes[i].NodeAddr)
			}
			if wantSDKConf.ChainClient.Nodes[i].ConnCnt != gotSDKConf.ChainClient.Nodes[i].ConnCnt {
				t.Errorf("got node ConnCnt = %d; want node ConnCnt %d", gotSDKConf.ChainClient.Nodes[i].ConnCnt, wantSDKConf.ChainClient.Nodes[i].ConnCnt)
			}
		}
	}

	if (wantSDKConf.ChainClient.Archive.Type != gotSDKConf.ChainClient.Archive.Type) ||
		(wantSDKConf.ChainClient.Archive.Dest != gotSDKConf.ChainClient.Archive.Dest) ||
		(wantSDKConf.ChainClient.Archive.SecretKey != gotSDKConf.ChainClient.Archive.SecretKey) {
		t.Errorf("got archive Type = %s; want archive Type %s", gotSDKConf.ChainClient.Archive.Type, wantSDKConf.ChainClient.Archive.Type)
		t.Errorf("got archive Dest = %s; want archive Dest %s", gotSDKConf.ChainClient.Archive.Dest, wantSDKConf.ChainClient.Archive.Dest)
		t.Errorf("got archive SecretKey = %s; want archive SecretKey %s", gotSDKConf.ChainClient.Archive.SecretKey, wantSDKConf.ChainClient.Archive.SecretKey)
	}

	if (wantSDKConf.ChainClient.RpcClient.MaxReceiveMessageSize != gotSDKConf.ChainClient.RpcClient.MaxReceiveMessageSize) ||
		(wantSDKConf.ChainClient.RpcClient.MaxSendMessageSize != gotSDKConf.ChainClient.RpcClient.MaxSendMessageSize) {
		t.Errorf("got archive Type = %d; want RpcClient MaxReceiveMessageSize %d", gotSDKConf.ChainClient.RpcClient.MaxReceiveMessageSize, wantSDKConf.ChainClient.RpcClient.MaxReceiveMessageSize)
		t.Errorf("got archive Dest = %d; want RpcClient MaxSendMessageSize %d", gotSDKConf.ChainClient.RpcClient.MaxSendMessageSize, wantSDKConf.ChainClient.RpcClient.MaxSendMessageSize)
	}
}

func TestAdmin4SDKConf(t *testing.T) {
	var nodeConf []PkSdkNodeConf
	nodeItem := new(PkSdkNodeConf)
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

	gotSDKConfPath, err := SetSDKConf("chain1", "./crypto-config/node1/admin/admin4", "admin4.key", "./sdk-configs", "sdk_config_pk_admin4.yml", nodeConf)
	if gotSDKConfPath != "./sdk-configs/sdk_config_pk_admin4.yml" {
		t.Errorf("sdk conf path = %s, expected %s", gotSDKConfPath, "./sdk-configs/sdk_config_pk_admin4.yml")
	}

	gotSDKConf, err := ReadSDKConf("./sdk-configs", "sdk_config_pk_admin4.yml")
	if err != nil {
		t.Errorf("read sdk config failed")
	}

	wantSDKConf := &SdkPkConfig{
		ChainClient: PkChainClientConf{
			ChainId:             "chain1",
			UserSignKeyFilePath: "./crypto-config/node1/admin/admin4/admin4.key",
			AuthType:            "public",
			EnableNormalKey:     false,
			Crypto: CryptoHashConf{
				Hash: "SHA256",
			},
			Archive: ArchiveConf{
				Type:      "mysql",
				Dest:      "root:123456:localhost:3306",
				SecretKey: "xxx",
			},
			RpcClient: RpcClientConf{
				MaxReceiveMessageSize: 16,
				MaxSendMessageSize:    16,
			},
			Nodes: []PkSdkNodeConf{
				{
					NodeAddr: "127.0.0.1:12301",
					ConnCnt:  10,
				},
				{
					NodeAddr: "127.0.0.1:12302",
					ConnCnt:  10,
				},
				{
					NodeAddr: "127.0.0.1:12303",
					ConnCnt:  10,
				},
				{
					NodeAddr: "127.0.0.1:12304",
					ConnCnt:  10,
				},
			},
		},
	}

	if wantSDKConf.ChainClient.ChainId != gotSDKConf.ChainClient.ChainId {
		t.Errorf("ChainId = %s; want %s", gotSDKConf.ChainClient.ChainId, wantSDKConf.ChainClient.ChainId)
	}

	if wantSDKConf.ChainClient.UserSignKeyFilePath != gotSDKConf.ChainClient.UserSignKeyFilePath {
		t.Errorf("UserSignKeyFilePath = %s; want %s", gotSDKConf.ChainClient.UserSignKeyFilePath, wantSDKConf.ChainClient.UserSignKeyFilePath)
	}

	if wantSDKConf.ChainClient.AuthType != gotSDKConf.ChainClient.AuthType {
		t.Errorf("AuthType = %s; want %s", gotSDKConf.ChainClient.AuthType, wantSDKConf.ChainClient.AuthType)
	}

	if wantSDKConf.ChainClient.EnableNormalKey != gotSDKConf.ChainClient.EnableNormalKey {
		t.Errorf("EnableNormalKey = %v; want %v", gotSDKConf.ChainClient.EnableNormalKey, wantSDKConf.ChainClient.EnableNormalKey)
	}

	if len(wantSDKConf.ChainClient.Nodes) != len(gotSDKConf.ChainClient.Nodes) {
		t.Errorf("nodes len = %d; want %d", len(gotSDKConf.ChainClient.Nodes), len(wantSDKConf.ChainClient.Nodes))
	} else {
		for i := 0; i < len(wantSDKConf.ChainClient.Nodes); i++ {
			if wantSDKConf.ChainClient.Nodes[i].NodeAddr != gotSDKConf.ChainClient.Nodes[i].NodeAddr {
				t.Errorf("got node addr = %s; want node addr %s", gotSDKConf.ChainClient.Nodes[i].NodeAddr, wantSDKConf.ChainClient.Nodes[i].NodeAddr)
			}
			if wantSDKConf.ChainClient.Nodes[i].ConnCnt != gotSDKConf.ChainClient.Nodes[i].ConnCnt {
				t.Errorf("got node ConnCnt = %d; want node ConnCnt %d", gotSDKConf.ChainClient.Nodes[i].ConnCnt, wantSDKConf.ChainClient.Nodes[i].ConnCnt)
			}
		}
	}

	if (wantSDKConf.ChainClient.Archive.Type != gotSDKConf.ChainClient.Archive.Type) ||
		(wantSDKConf.ChainClient.Archive.Dest != gotSDKConf.ChainClient.Archive.Dest) ||
		(wantSDKConf.ChainClient.Archive.SecretKey != gotSDKConf.ChainClient.Archive.SecretKey) {
		t.Errorf("got archive Type = %s; want archive Type %s", gotSDKConf.ChainClient.Archive.Type, wantSDKConf.ChainClient.Archive.Type)
		t.Errorf("got archive Dest = %s; want archive Dest %s", gotSDKConf.ChainClient.Archive.Dest, wantSDKConf.ChainClient.Archive.Dest)
		t.Errorf("got archive SecretKey = %s; want archive SecretKey %s", gotSDKConf.ChainClient.Archive.SecretKey, wantSDKConf.ChainClient.Archive.SecretKey)
	}

	if (wantSDKConf.ChainClient.RpcClient.MaxReceiveMessageSize != gotSDKConf.ChainClient.RpcClient.MaxReceiveMessageSize) ||
		(wantSDKConf.ChainClient.RpcClient.MaxSendMessageSize != gotSDKConf.ChainClient.RpcClient.MaxSendMessageSize) {
		t.Errorf("got archive Type = %d; want RpcClient MaxReceiveMessageSize %d", gotSDKConf.ChainClient.RpcClient.MaxReceiveMessageSize, wantSDKConf.ChainClient.RpcClient.MaxReceiveMessageSize)
		t.Errorf("got archive Dest = %d; want RpcClient MaxSendMessageSize %d", gotSDKConf.ChainClient.RpcClient.MaxSendMessageSize, wantSDKConf.ChainClient.RpcClient.MaxSendMessageSize)
	}
}

func TestGeneratePKPair(t *testing.T) {
	err := GeneratePKFile("./", "testKey")
	if err != nil {
		t.Errorf("generate pk file failed, err:=%v", err)
	}
}

func TestEVMAddr(t *testing.T) {
	keyPemStr, err := CalcEVMAddr("./crypto-config/node1/admin/admin1/admin1.pem")
	if err != nil {
		t.Errorf("calc EVM addr failed, err:=%v, keyPemStr=%v", err, keyPemStr)
	}

	if len(keyPemStr) != 40 {
		t.Errorf("EVM addr must be 40, in fact it is %d", len(keyPemStr))
	}
}
