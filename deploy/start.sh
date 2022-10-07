#!/bin/bash

# 本脚本从头构建一个区块链网络
# 请确保 cryptogen 和 configtxgen 这两个可执行文件已经被正确安装
# 创建一个通道 emrchannel

echo "1. Environmental cleaning"
mkdir -p config
mkdir -p crypto-config
rm -fr config/*
rm -fr crypto-config/*

echo "Clean up"

echo "2. Generate certificate and start block information"
cryptogen generate --config=./crypto-config.yaml
configtxgen -profile OneOrgOrdererGenesis -outputBlock ./config/genesis.block

echo "3. Blockchain: start"
docker-compose up -d        # 按照docker-compose.yaml的配置启动区块链网络并在后台运行
echo "Waiting for node startup to complete, waiting for 10 seconds"
sleep 10                    # 启动整个区块链网络需要一点时间，所以此处等待10s，让区块链网络完全启动

# 四、生成通道(这个动作会创建一个创世交易，也是该通道的创世交易)
echo "4. Generate TX file for channel"
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./config/emrchannel.tx -channelID emrchannel

# 五、在区块链上按照刚刚生成的TX文件去创建通道
# 该操作和上面操作不一样的是，这个操作会写入区块链
echo "5. Create channels on the blockchain according to TX files"
docker exec cli peer channel create -o orderer.medical.com:7050 -c emrchannel -f /etc/hyperledger/config/emrchannel.tx

# 六、让节点去加入到通道
echo "6. Let the node join the channel"
docker exec cli peer channel join -b emrchannel.block

# 七、链码安装
# -n 是链码的名字，可以自己随便设置
# -v 就是版本号，就是composer的bna版本
# -p 是目录，目录是基于cli这个docker里面的$GOPATH相对的
# 此处安装的是示例链码，后续课程会自己编写
echo "7. Installation chain code"
docker exec cli peer chaincode install -n emrcc -v 1.0.0 -l golang -p github.com/medical-system/chaincode

#八、实例化链码
#-n 对应前文安装链码的名字 其实就是composer network start bna名字
#-v 为版本号，相当于composer network start bna名字@版本号
#-C 是通道，在fabric的世界，一个通道就是一条不同的链，composer并没有很多提现这点，composer提现channel也就在于多组织时候的数据隔离和沟通使用
#-c 为传参，传入init参数
echo "8. Instantiated chain code"
docker exec cli peer chaincode instantiate -o orderer.medical.com:7050 -C emrchannel -n emrcc -l golang -v 1.0.0 -c '{"Args":["init"]}'

# Conduct chain code interaction to verify whether the chain code is correctly installed and whether the blockchain network can work normally
#docker exec cli peer chaincode invoke -C emrchannel -n emrcc -c '{"Args":["queryInformation","patient","akanancy","2021-08-24 14:54:12"]}'
#docker exec cli peer chaincode invoke -C emrchannel -n emrcc -c '{"Args":["authorizeDoctor","akanancy","doctor11"]}'
#docker exec cli peer chaincode invoke -C emrchannel -n emrcc -c '{"Args":["authorizeDoctor","evayam","dototrkkk"]}'
#docker exec cli peer chaincode invoke -C emrchannel -n emrcc -c '{"Args":["enterData","doctor11","akanancy","a blood sugar level of 200 milligrams per deciliter (mg/dL)","Type 2 diabetes","Meglitinides and sulfonylureas","2021-08-24 15:11:12"]}'
#docker exec cli peer chaincode invoke -C emrchannel -n emrcc -c '{"Args":["enterData","doctor11","akanancy","skin and the whites of your eyes have jaundice (a yellowish color)","liver cancer","Targeted Drug Therapy for Liver Cancer","2021-08-24 15:11:12"]}'
#docker exec cli peer chaincode invoke -C emrchannel -n emrcc -c '{"Args":["enterData","dototrkkk","evayam","random blood sugar > 200 mg/dL, hemoglobin A1c > 6.5","Type 1 diabetes","take Metformin and insulin","2021-08-24 15:11:12"]}'
#docker exec cli peer chaincode invoke -C emrchannel -n emrcc -c '{"Args":["queryInformation","doctor","doctor11","akanancy","2021-08-24 14:58:12"]}'
#docker exec cli peer chaincode invoke -C emrchannel -n emrcc -c '{"Args":["queryInformation","doctor","dototrkkk","evayam","2021-08-24 14:58:12"]}'
#docker exec cli peer chaincode invoke -C emrchannel -n emrcc -c '{"Args":["queryCases","researcher","res1312","diabetes","2021-08-24 15:18:24"]}'
#docker exec cli peer chaincode invoke -C emrchannel -n emrcc -c '{"Args":["queryCases","regulator","regu002","cancer","2021-08-24 15:18:24"]}'
#docker exec cli peer chaincode query -C emrchannel -n emrcc -c '{"Args":["showRecords"]}'


