package main

import (
    _ "net/http/pprof"
    "hgeekl.com/gkicoin/src/gkicoin"
    "hgeekl.com/gkicoin/src/util/logging"
    "hgeekl.com/gkicoin/src/visor"
)


var (
	// Version of the node. Can be set by -ldflags
	Version = "1.24.1"
	// Commit ID. Can be set by -ldflags
	Commit = ""
	// Branch name. Can be set by -ldflags
	Branch = ""
	// ConfigMode (possible values are "", "STANDALONE_CLIENT").
	// This is used to change the default configuration.
	// Can be set by -ldflags
	ConfigMode = ""

	logger = logging.MustGetLogger("main")

	// GenesisSignatureStr hex string of genesis signature
	GenesisSignatureStr = "7ad7e072f55d27b473323aa0595b23a955b67eb3d72ca16535dd22997e901a05433b3b999bdf86aee65e6f1a19c1fac12b6fc9608a7cc3148243e9b3d3344ade01"
	// GenesisAddressStr genesis address string
	GenesisAddressStr = "3Tck4EJBAGL1uo6xsAmYaw1vrJxi5pjdHP"
	// BlockchainPubkeyStr pubic key string
	BlockchainPubkeyStr = "039b12f3269545c8257339896d6eda09b76ae539a9b720cac45594af03697dadf3"
	// BlockchainSeckeyStr empty private key string
	BlockchainSeckeyStr = ""

	// GenesisTimestamp genesis block create unix time
	GenesisTimestamp uint64 = 1568250863
	// GenesisCoinVolume represents the coin capacity
	GenesisCoinVolume uint64 = 100000000000000

	// DefaultConnections the default trust node addresses
	DefaultConnections = []string{
        "47.104.22.149:5858",
		"47.56.171.18:5858",
		"47.56.9.56:5858",
	}
)

func main() {
    // get node config
    nodeConfig := gkicoin.NewNodeConfig(ConfigMode, gkicoin.NodeParameters{
        GenesisSignatureStr: GenesisSignatureStr,
        GenesisAddressStr:   GenesisAddressStr,
        GenesisCoinVolume:   GenesisCoinVolume,
        GenesisTimestamp:    GenesisTimestamp,
        BlockchainPubkeyStr: BlockchainPubkeyStr,
        BlockchainSeckeyStr: BlockchainSeckeyStr,
        DefaultConnections:  DefaultConnections,
        PeerListURL:         "http://block.hgeekl.com/s/gkipeers.txt",
        Port:                5858,
        WebInterfacePort:    5959,
        DataDirectory:       "$HOME/.gkicoin",
        ProfileCPUFile:      "gkicoin.prof",
    })

    // create a new fiber coin instance
    coin := gkicoin.NewCoin(
        gkicoin.Config{
            Node: *nodeConfig,
            Build: visor.BuildInfo{
                Version: Version,
                Commit:  Commit,
                Branch:  Branch,
            },
        },
        logger,
    )

    // parse config values
    coin.ParseConfig()

    // run fiber coin node
    coin.Run()
}
