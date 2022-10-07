package blockchain

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

var (
	SDK             *fabsdk.FabricSDK
	ChannelName     = "emrchannel"
	ChaincodeName   = "emrcc"
	Org             = "Org1"
	User            = "Admin"
	ConfigPath      = "./application/config.yaml"
	TargetEndpoint1 = "peer0.org1.medical.com"
)

func Init() {
	var err error
	// initialization sdk
	SDK, err = fabsdk.New(config.FromFile(ConfigPath))
	if err != nil {
		panic(err.Error())
		return
	}
}

func ChannelQuery(fcn string, args [][]byte) (channel.Response, error) {
	ctx := SDK.ChannelContext(ChannelName, fabsdk.WithOrg(Org), fabsdk.WithUser(User))

	// Create client object
	cli, err := channel.New(ctx)
	if err != nil {
		return channel.Response{}, err
	}
	return cli.Query(channel.Request{
		ChaincodeID: ChaincodeName,
		Fcn:         fcn,
		Args:        args,
	}, channel.WithTargetEndpoints(TargetEndpoint1))

}
func ChannelExecute(fcn string, args [][]byte) (channel.Response, error) {
	ctx := SDK.ChannelContext(ChannelName, fabsdk.WithOrg(Org), fabsdk.WithUser(User))
	cli, err := channel.New(ctx)
	if err != nil {
		return channel.Response{}, err
	}
	return cli.Execute(channel.Request{
		ChaincodeID: ChaincodeName,
		Fcn:         fcn,
		Args:        args,
	}, channel.WithTargetEndpoints(TargetEndpoint1))

}
