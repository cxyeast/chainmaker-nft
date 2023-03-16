ERC721 Token Standard:
https://eips.ethereum.org/EIPS/eip-721

## The description of methods are below:
## 1. InitContract
### args:
#### key1: name(optional)
#### value1: string
#### key2: symbol(optional)
#### value2: string
#### key3: tokenURI(optional)
#### value3: string
#### example:
```json
{"tokenURI":"http://chainmaker.org.cn"}
```

## 2. name
### args: no args
### response example: "erc721"

## 3. symbol
### args: no args
### response example: "erc721X"

## 4. balanceOf
### args:
#### key1: "account"
#### value1: string
#### example:
```json
{"account":"ec47ae0f0d6a0e952c240383d70ab43b19997a9f"}
```
### response example: "0"

## 5. ownerOf
### args:
#### key1: "tokenId"
#### value1: string
#### example:
```json
{"tokenId":"111111111111111111111112"}
```
### response example: "ec47ae0f0d6a0e952c240383d70ab43b19997a9f"

## 6. mint
### args:
#### key1: "to"
#### value1: string
#### key2: "tokenId"
#### value2: string
#### key3: "metadata"
#### value3: bytes
#### example:
```json
{"to":"ec47ae0f0d6a0e952c240383d70ab43b19997a9f", "tokenId":"111111111111111111111112", "metadata": "url:https://chainmaker.org.cn"}
```
#### resp exampl: "mint success"

## 7. tokenURI
### args:
#### key1: "tokenId"
#### value1: string
#### example:
```json
{"tokenId":"111111111111111111111112"}
```
#### resp exampl: "http://chainmaker.org.cn/111111111111111111111112"

## 8. tokenMetadata
### args:
#### key1: "tokenId"
#### value1: string
#### example:
```json
{"tokenId":"111111111111111111111112"}
```
#### resp exampl: "url:http://chainmaker.org.cn/111111111111111111111112"

## 9. tokenLatestTxInfo
### args:
#### key1: "tokenId"
#### value1: string
#### example:
```json
{"tokenId":"111111111111111111111112"}
```
#### resp exampl: 
```json
{"TxId":"17262429164a0e82ca17c10d4d4bc2b11be6c7c1e9cd4d6db287a8a4f3f2e2e5","BlockHeight":79,"From":"0000000000000000000000000000000000000000","To":"ec47ae0f0d6a0e952c240383d70ab43b19997a9f","Timestamp":"1668060470"}
```

## 10. accountTokens
### args:
#### key1: "account"
#### value1: string
#### example:
```json
{"account":"ec47ae0f0d6a0e952c240383d70ab43b19997a9f"}
```
#### resp exampl:
```json
{"Account":"ec47ae0f0d6a0e952c240383d70ab43b19997a9f","Tokens":["111111111111111111111112","111111111111111111111113"]}
```

## 11. approve
### args:
#### key1: "to"
#### value1: string
#### key2: "tokenId"
#### value2: string
#### example:
```json
{"to":"a04f7895de24f61807a729be230f03da8c0eef42", "tokenId":"111111111111111111111112"}
```
#### resp exampl: "approve success"
### event:
#### topic: approve
#### data: owner, to, tokenId
#### example:
```json
["ec47ae0f0d6a0e952c240383d70ab43b19997a9f","a04f7895de24f61807a729be230f03da8c0eef42","111111111111111111111112"]
```

## 12. getApprove
### args:
#### key1: "tokenId"
#### value1: string
#### example:
```json
{"tokenId":"111111111111111111111112"}
```
#### resp exampl: "ec47ae0f0d6a0e952c240383d70ab43b19997a9f"

## 13. transferFrom
### args:
#### key1: "from"
#### value1: string
#### key2: "to"
#### value2: string
#### key3: "tokenId"
#### value2: string
#### example:
```json
{"from":"ec47ae0f0d6a0e952c240383d70ab43b19997a9f", "to":"a04f7895de24f61807a729be230f03da8c0eef42", "tokenId":"111111111111111111111112"}
```
#### resp exampl: "transfer success"
### event:
#### topic: transfer
#### data: from, to, tokenId
#### example:
```json
["ec47ae0f0d6a0e952c240383d70ab43b19997a9f","a04f7895de24f61807a729be230f03da8c0eef42","111111111111111111111112"]
```

## Test

### 部署合约
```sh
./cmc client contract user create --contract-name=erc721 --version=1.0 --sync-result=true --sdk-conf-path=./testdata/sdk_config_solo.yml --byte-code-path=./testdata/erc721.7z --runtime-type=DOCKER_GO --admin-crt-file-paths=./testdata/crypto-config/wx-org.chainmaker.org/user/admin1/admin1.sign.crt --admin-key-file-paths=./testdata/crypto-config/wx-org.chainmaker.org/user/admin1/admin1.sign.key --params="{\"name\":\"huanletoken\", \"symbol\":\"hlt\", \"tokenURI\":\"https://chainmaker.org.cn\"}"
```

### 查询name
#### 验证Case1：
部署合约时如果没有指定erc721的name，默认的name为空，需要验证name为空
#### 验证Case2：
部署合约时如果指定了name参数，在这儿获取时验证name是否和部署合约时指定的name一致
```sh
./cmc client contract user invoke --contract-name=erc721 --method=name --sync-result=true --sdk-conf-path=./testdata/sdk_config_solo.yml
```

### 查询symbol
#### 验证Case1：
部署合约时如果没有指定erc721的symbol，默认的symbol为空，这儿需要验证symbol为空
#### 验证Case2：
部署合约时如果指定了symbol参数，在这儿获取时验证symbol是否和部署合约时指定的symbol一致
```sh
./cmc client contract user invoke --contract-name=erc721 --method=symbol --sync-result=true --sdk-conf-path=./testdata/sdk_config_solo.yml
```

### 查询tokenURI
#### 验证Case1：
验证返回的tokenURI是否为安装合约时指定的tokenURI+'/'+tokenId
```sh
./cmc client contract user invoke --contract-name=erc721test --method=tokenURI --sync-result=true --sdk-conf-path=./testdata/sdk_config_solo.yml --params="{\"tokenId\":\"111111111111111111111112\"}"
```

### 查询账户nft数量
#### 验证Case1：
部署合约后所有账户默认的nft数量为0，这儿需要验证账户默认nft数量是否为0
```sh
./cmc client contract user invoke --contract-name=erc721 --method=balanceOf --sync-result=true --sdk-conf-path=./testdata/sdk_config_solo.yml --params="{\"account\":\"ec47ae0f0d6a0e952c240383d70ab43b19997a9f\"}"
```

### 查询nft所属账户
#### 验证Case1：
部署合约后如果nft不存在，查询nft所属账户会报错，这儿需要验证nft不存在时的错误情况
```sh
./cmc client contract user invoke --contract-name=erc721 --method=ownerOf --sync-result=true --sdk-conf-path=./testdata/sdk_config_solo.yml --params="{\"tokenId\":\"111111111111111111111112\"}"
```

### 发行nft
#### 验证Case1：
发行nft后需要验证账户nft数量是否增加1
#### 验证Case2：
发行nft后需要验证nft所属账户是否正确
```sh
./cmc client contract user invoke --contract-name=erc721test --method=mint --sync-result=true --sdk-conf-path=./testdata/sdk_config_solo.yml --params="{\"to\":\"ec47ae0f0d6a0e952c240383d70ab43b19997a9f\", \"tokenId\":\"111111111111111111111112\", \"metadata\":\"url:http://chainmaker.org.cn/\"}"
```

### 查询token metadata信息
#### 验证Case1：
这儿验证查询到的metadata是否和mint时传递的一致
```sh
./cmc client contract user invoke --contract-name=erc721 --method=tokenMetadata --sync-result=true --sdk-conf-path=./testdata/sdk_config_solo.yml --params="{\"tokenId\":\"111111111111111111111112\"}"
```

### 查询account tokens信息
#### 验证Case1：
验证账户下是否包含了所有发行的nft
```sh
./cmc client contract user invoke --contract-name=erc721test --method=accountTokens --sync-result=true --sdk-conf-path=./testdata/sdk_config_solo.yml --params="{\"account\":\"ec47ae0f0d6a0e952c240383d70ab43b19997a9f\"}"
```

### 查询token 最近一笔交易信息
#### 验证Case1：
验证token最近一笔的交易信息是否正确
```sh
./cmc client contract user invoke --contract-name=erc721test --method=tokenLatestTxInfo --sync-result=true --sdk-conf-path=./testdata/sdk_config_solo.yml --params="{\"tokenId\":\"111111111111111111111112\"}"
```

### 获取授权信息
#### 验证Case1：
如果nft没有进行过授权，查询到的授权信息应为空
```sh
./cmc client contract user invoke --contract-name=erc721 --method=getApprove --sync-result=true --sdk-conf-path=./testdata/sdk_config_solo.yml --params="{\"tokenId\":\"111111111111111111111112\"}"
```

### 授权
#### 验证Case1：
授权后需要验证授权信息是否正确
```sh
./cmc client contract user invoke --contract-name=erc721 --method=approve --sync-result=true --sdk-conf-path=./testdata/sdk_config_solo.yml --params="{\"to\":\"a04f7895de24f61807a729be230f03da8c0eef42\", \"tokenId\":\"111111111111111111111112\"}"
```

### 根据授权转账
#### 验证Case1：
转账后需要验证授权信息是否发生了变化
```sh
./cmc client contract user invoke --contract-name=erc721 --method=transferFrom --sync-result=true --sdk-conf-path=./testdata/sdk_config_solo.yml --params="{\"from\":\"ec47ae0f0d6a0e952c240383d70ab43b19997a9f\", \"to\":\"a04f7895de24f61807a729be230f03da8c0eef42\", \"tokenId\":\"111111111111111111111112\"}"
```
