package myclient

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// SdkPkConfig SdkPkConfig
type SdkPkConfig struct {
	ChainClient PkChainClientConf `yaml:"chain_client"`
}

// PkChainClientConf pk chain client config
type PkChainClientConf struct {
	ChainId             string          `yaml:"chain_id"`
	UserSignKeyFilePath string          `yaml:"user_sign_key_file_path"`
	Crypto              CryptoHashConf  `yaml:"crypto"`
	AuthType            string          `yaml:"auth_type"`
	EnableNormalKey     bool            `yaml:"enable_normal_key"`
	Nodes               []PkSdkNodeConf `yaml:"nodes"`
	Archive             ArchiveConf     `yaml:"archive"`
	RpcClient           RpcClientConf   `yaml:"rpc_client"`
}

// CryptoHashConf crypto hash conf
type CryptoHashConf struct {
	Hash string `yaml:"hash"`
}

// PkSdkNodeConf pk sdk node conf
type PkSdkNodeConf struct {
	NodeAddr string `yaml:"node_addr"`
	ConnCnt  int    `yaml:"conn_cnt"`
}

// ArchiveConf archive conf
type ArchiveConf struct {
	Type      string `yaml:"type"`
	Dest      string `yaml:"dest"`
	SecretKey string `yaml:"secret_key"`
}

// RpcClientConf rpc client conf
type RpcClientConf struct {
	MaxReceiveMessageSize int `yaml:"max_receive_message_size"`
	MaxSendMessageSize    int `yaml:"max_send_message_size"`
}

func ReadSDKConf(path, name string) (*SdkPkConfig, error) {
	dest := path + "/sdk_config_pk_" + name + ".yml"
	conf := new(SdkPkConfig)
	bcFile, err := ioutil.ReadFile(dest)
	if err != nil {
		return nil, err
	}
	_ = yaml.Unmarshal(bcFile, conf)
	return conf, nil
}

func SetSDKConf(chainid, keyPath, name, confPath string, nodeConf []PkSdkNodeConf) (string, error) {

	keySourceFile := keyPath + "/" + name + ".key"
	conf := &SdkPkConfig{
		ChainClient: PkChainClientConf{
			ChainId:             chainid,
			UserSignKeyFilePath: keySourceFile,
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
			Nodes: nodeConf,
		},
	}

	sdkFile, err := yaml.Marshal(conf)
	if err != nil {
		return "", err
	}

	dest := confPath + "/sdk_config_pk_" + name + ".yml"
	err = ioutil.WriteFile(dest, sdkFile, 0644)
	if err != nil {
		return "", err
	}

	return dest, nil
}
