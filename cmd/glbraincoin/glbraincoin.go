package main

import (
	_ "net/http/pprof"

	"github.com/skycoin/skycoin/src/skycoin"
	"github.com/skycoin/skycoin/src/util/logging"
	"github.com/skycoin/skycoin/src/visor"
)

var (
	// Version of the node. Can be set by -ldflags
	Version = "0.24.1"
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
	GenesisSignatureStr = "dcf9ba232462953cd98aea3db47bdbeefac2a762723883fe45e2e8db31d6531c7f59047a5ad148be9fc7d5adf9951ce9b81af4c4dc4892a9e1fc31ab63085d9500"
	// GenesisAddressStr genesis address string
	GenesisAddressStr = "qhtJc1xjNdSvGB2aeCSbPYE5FyNTy5WMpG"
	// BlockchainPubkeyStr pubic key string
	BlockchainPubkeyStr = "032977eb6a998c42d9396b1f9052c4bb9a3fafa693d2eaf29ca2083b50ca86d52b"
	// BlockchainSeckeyStr empty private key string
	BlockchainSeckeyStr = ""

	// GenesisTimestamp genesis block create unix time
	GenesisTimestamp uint64 = 1534413563
	// GenesisCoinVolume represents the coin capacity
	GenesisCoinVolume uint64 = 300e12

	// DefaultConnections the default trust node addresses
	DefaultConnections = []string{
	"144.76.115.79:9100",
"144.76.115.79:9101",
"144.76.115.79:9102",
"144.76.115.80:9100",
"144.76.115.81:9100",}
)

func main() {
	// get node config
	nodeConfig := skycoin.NewNodeConfig(ConfigMode, skycoin.NodeParameters{
		GenesisSignatureStr: GenesisSignatureStr,
		GenesisAddressStr:   GenesisAddressStr,
		GenesisCoinVolume:   GenesisCoinVolume,
		GenesisTimestamp:    GenesisTimestamp,
		BlockchainPubkeyStr: BlockchainPubkeyStr,
		BlockchainSeckeyStr: BlockchainSeckeyStr,
		DefaultConnections:  DefaultConnections,
		PeerListURL:         "https://glbrain.com/peers.txt",
		Port:                9100,
		WebInterfacePort:    9120,
		DataDirectory:       "$HOME/.glbrain",
		ProfileCPUFile:      "skycoin.prof",
	})

	// create a new fiber coin instance
	coin := skycoin.NewCoin(
		skycoin.Config{
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
