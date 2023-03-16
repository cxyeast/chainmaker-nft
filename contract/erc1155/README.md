# erc1155合约
ERC1155 Token Standard:
https://eips.ethereum.org/EIPS/eip-1155

用于批量发行nft和token，包含了erc20和erc721的功能。

接口详情请查看接口：IERC1155
此处主要说明和标准合约的区别在于：多了接口：MintBatchNft，用户批量发行nft，有一些nft的校验。


## cmc使用示例

命令行工具使用示例

```sh
# 注： 
# 6fa2bac1c8f5c2c7225c0ec5a3f9262dd315d962为用户2的地址即sdk_config_admin2.yml配置中的sign
# 08cd36c7be843d70bfc585ccd20e101e8bb8bc60为用户1的地址即sdk_config.yml配置中的sign
# 实际执行下方命令时，请自行替换相应的地址
# 采用如下命令计算地址：默认使用`ethereum`地址
./cmc address  -h
./cmc address cert-to-addr ./testdata/admin1.sign.crt
./cmc address pk-to-addr ./testdata/crypto-config/node1/admin/admin1/admin1.pem

function Erc1155() {
    echo
    echo "安装合约 erc1155，设置管理员地址，若不手动设置，则为创建合约的人，管理员可 mint"
    ./cmc client contract user create \
    --contract-name=erc1155 \
    --version=1.0 \
    --sync-result=true \
    --sdk-conf-path=./testdata/sdk_config.yml \
    --byte-code-path=./testdata/erc1155.7z \
    --runtime-type=DOCKER_GO \
    --admin-crt-file-paths=./testdata/crypto-config/wx-org.chainmaker.org/user/admin1/admin1.sign.crt \
    --admin-key-file-paths=./testdata/crypto-config/wx-org.chainmaker.org/user/admin1/admin1.sign.key \
    --params="{\"uri\":\"https://www.fxtoon.com/nft/{tokenId}.json\"}"
  echo
  echo "set admin"
  ./cmc client contract user invoke \
  --contract-name=fxtoonErc1155 \
  --method=AlterAdminAddress \
  --sdk-conf-path=./testdata/sdk_config.yml \
  --gas-limit=100000000 \
  --params="{\"adminAddress\":\"5fa92a33364dd5ce26a9814a6aceb240bd6bf083,08cd36c7be843d70bfc585ccd20e101e8bb8bc60\"}" \
  --sync-result=true \
  --result-to-string=true
  echo
  echo "Mint erc721 给admin2发送token value为1 "
  ./cmc client contract user invoke \
  --contract-name=fxtoonErc1155 \
  --method=Mint \
  --sdk-conf-path=./testdata/sdk_config.yml \
  --gas-limit=100000000 \
  --params="{\"to\":\"6fa2bac1c8f5c2c7225c0ec5a3f9262dd315d962\",\"id\":\"1\",\"amount\":\"1\"}" \
  --sync-result=true \
  --result-to-string=true
  echo
  echo "Mint erc20 给admin2发送token value为2 "
  ./cmc client contract user invoke \
  --contract-name=fxtoonErc1155 \
  --method=Mint \
  --sdk-conf-path=./testdata/sdk_config.yml \
  --gas-limit=100000000 \
  --params="{\"to\":\"6fa2bac1c8f5c2c7225c0ec5a3f9262dd315d962\",\"id\":\"2\",\"amount\":\"2\"}" \
  --sync-result=true \
  --result-to-string=true
  echo
  echo "MintBatchNft admin2 2 3 4"
  ./cmc client contract user invoke \
  --contract-name=fxtoonErc1155 \
  --method=MintBatchNft \
  --sdk-conf-path=./testdata/sdk_config.yml \
  --gas-limit=100000000 \
  --params="{\"to\":\"6fa2bac1c8f5c2c7225c0ec5a3f9262dd315d962\",\"amount\":\"3\",\"idStart\":\"3\"}" \
  --sync-result=true \
  --result-to-string=true
  echo
  echo "BalanceOf admin1 token1"
  ./cmc client contract user get \
  --contract-name=fxtoonErc1155 \
  --method=BalanceOf \
  --sdk-conf-path=./testdata/sdk_config.yml \
  --params="{\"owner\":\"08cd36c7be843d70bfc585ccd20e101e8bb8bc60\",\"id\":\"1\"}" \
  --result-to-string=true
  echo
  echo "BalanceOf admin2 token1"
  ./cmc client contract user get \
  --contract-name=fxtoonErc1155 \
  --method=BalanceOf \
  --sdk-conf-path=./testdata/sdk_config.yml \
  --params="{\"owner\":\"6fa2bac1c8f5c2c7225c0ec5a3f9262dd315d962\",\"id\":\"1\"}" \
  --result-to-string=true
  echo
  echo "SafeTransferFrom token 1 admin1 to admin2 err"
  ./cmc client contract user invoke \
  --contract-name=fxtoonErc1155 \
  --method=SafeTransferFrom \
  --sdk-conf-path=./testdata/sdk_config.yml \
  --gas-limit=100000000 \
  --params="{\"from\":\"08cd36c7be843d70bfc585ccd20e101e8bb8bc60\",\"to\":\"6fa2bac1c8f5c2c7225c0ec5a3f9262dd315d962\",\"id\":\"1\",\"amount\":\"1\"}" \
  --sync-result=true \
  --result-to-string=true
  echo
  echo "SafeTransferFrom token 1 admin2 to admin1"
  ./cmc client contract user invoke \
  --contract-name=fxtoonErc1155 \
  --method=SafeTransferFrom \
  --sdk-conf-path=./testdata/sdk_config_admin2.yml \
  --gas-limit=100000000 \
  --params="{\"from\":\"6fa2bac1c8f5c2c7225c0ec5a3f9262dd315d962\",\"to\":\"08cd36c7be843d70bfc585ccd20e101e8bb8bc60\",\"id\":\"1\",\"amount\":\"1\"}" \
  --sync-result=true \
  --result-to-string=true
  echo
  echo "BalanceOfBatch admin2 admin1 token1"
  ./cmc client contract user get \
  --contract-name=fxtoonErc1155 \
  --method=BalanceOfBatch \
  --sdk-conf-path=./testdata/sdk_config.yml \
  --params="{\"owner\":\"08cd36c7be843d70bfc585ccd20e101e8bb8bc60,6fa2bac1c8f5c2c7225c0ec5a3f9262dd315d962\",\"id\":\"1,1\"}" \
  --result-to-string=true
  echo
  echo "SafeTransferFrom token 1 admin1 to admin2, err sender not owner"
  ./cmc client contract user invoke \
  --contract-name=fxtoonErc1155 \
  --method=SafeTransferFrom \
  --sdk-conf-path=./testdata/sdk_config_admin2.yml \
  --gas-limit=100000000 \
  --params="{\"from\":\"08cd36c7be843d70bfc585ccd20e101e8bb8bc60\",\"to\":\"6fa2bac1c8f5c2c7225c0ec5a3f9262dd315d962\",\"id\":\"1\",\"amount\":\"1\"}" \
  --sync-result=true \
  --result-to-string=true
  echo
  echo "SetApprovalForAll admin1 to admin2 no sender"
  ./cmc client contract user invoke \
  --contract-name=fxtoonErc1155 \
  --method=SetApprovalForAll \
  --sdk-conf-path=./testdata/sdk_config.yml \
  --gas-limit=100000000 \
  --params="{\"operator\":\"6fa2bac1c8f5c2c7225c0ec5a3f9262dd315d962\",\"approved\":\"true\"}" \
  --sync-result=true \
  --result-to-string=true
  echo
  echo "SetApprovalForAll admin1 to admin2 with sender"
  ./cmc client contract user invoke \
  --contract-name=fxtoonErc1155 \
  --method=SetApprovalForAll \
  --sdk-conf-path=./testdata/sdk_config.yml \
  --gas-limit=100000000 \
  --params="{\"operator\":\"6fa2bac1c8f5c2c7225c0ec5a3f9262dd315d962\",\"sender\":\"08cd36c7be843d70bfc585ccd20e101e8bb8bc60\",\"approved\":\"true\",\"sign\":\"3046022100a4afcc04659d5f8bbf831daf0e6e98c08520cf95a482b58ccd5f804b195ff430022100f71d1ff8c9cdc6f90570d88dc3daa39c3dbd158b2bc1ace0b40cae054c3a9444\"}" \
  --sync-result=true \
  --result-to-string=true
  echo
  echo "IsApprovedForAll admin1 to admin2"
  ./cmc client contract user get \
  --contract-name=fxtoonErc1155 \
  --method=IsApprovedForAll \
  --sdk-conf-path=./testdata/sdk_config.yml \
  --params="{\"operator\":\"6fa2bac1c8f5c2c7225c0ec5a3f9262dd315d962\",\"owner\":\"08cd36c7be843d70bfc585ccd20e101e8bb8bc60\"}" \
  --result-to-string=true
  echo
  echo "SafeTransferFrom token 1 admin1 to admin2"
  ./cmc client contract user invoke \
  --contract-name=fxtoonErc1155 \
  --method=SafeTransferFrom \
  --sdk-conf-path=./testdata/sdk_config_admin2.yml \
  --gas-limit=100000000 \
  --params="{\"from\":\"08cd36c7be843d70bfc585ccd20e101e8bb8bc60\",\"to\":\"6fa2bac1c8f5c2c7225c0ec5a3f9262dd315d962\",\"id\":\"1\",\"amount\":\"1\"}" \
  --sync-result=true \
  --result-to-string=true
  echo
  echo "Uri 1"
  ./cmc client contract user get \
  --contract-name=fxtoonErc1155 \
  --method=Uri \
  --sdk-conf-path=./testdata/sdk_config.yml \
  --params="{\"id\":\"1\"}" \
  --result-to-string=true
  echo
  echo "OwnerOf 1"
  ./cmc client contract user get \
  --contract-name=fxtoonErc1155 \
  --method=OwnerOf \
  --sdk-conf-path=./testdata/sdk_config.yml \
  --params="{\"id\":\"1\"}" \
  --result-to-string=true
}
```