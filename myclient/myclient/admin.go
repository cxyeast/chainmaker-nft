package myclient

import (
	"fmt"
	"io/ioutil"

	"chainmaker.org/chainmaker/common/v2/crypto/asym"
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

	key, err := asym.PrivateKeyFromPEM(keyPem, nil)
	if err != nil {
		return fmt.Errorf("get private key failed, %s", err)
	}

	pubKey := key.PublicKey()
	memberInfo, err := pubKey.String()
	if err != nil {
		return fmt.Errorf("get pubkey info failed, %s", err)
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
