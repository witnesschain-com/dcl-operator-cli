package common

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	DefaultOpProverConfig     string = "operator-prover-config.json"
	DefaultOpChallengerConfig string = "operator-challenger-config.json"
)

type ChainConfig struct {
	ProverRegistryAddress     common.Address
	ChallengerRegistryAddress common.Address
	ChainID                   big.Int
	BlockExplorer             string
}

var BlueOrangutan = ChainConfig{
	ChallengerRegistryAddress: common.HexToAddress("0xeFFE8c100029F71924554aEd382f1919ecA6b203"),
	ProverRegistryAddress:     common.HexToAddress("0x91013d3CecE055603D8b1EE7DCB1f670f480fe24"),
	ChainID:                   *big.NewInt(1237146866),
	BlockExplorer:             "https://blue-orangutan-blockscout.eu-north-2.gateway.fm",
}

var Holesky = ChainConfig{
	ProverRegistryAddress:     common.HexToAddress(""),
	ChallengerRegistryAddress: common.HexToAddress(""),
	ChainID:                   *big.NewInt(17000),
	BlockExplorer:             "https://holesky.etherscan.io",
}

var WitnesschainMainnet = ChainConfig{
	ProverRegistryAddress:     common.HexToAddress(""),
	ChallengerRegistryAddress: common.HexToAddress(""),
	BlockExplorer:             "https://explorer.witnesschain.com",
	ChainID:                   *big.NewInt(1702448187),
}

var EthMainnet = ChainConfig{
	ProverRegistryAddress:     common.HexToAddress(""),
	ChallengerRegistryAddress: common.HexToAddress(""),
	ChainID:                   *big.NewInt(1),
	BlockExplorer:             "https://etherscan.io",
}

var NetworkConfig = map[string]ChainConfig{
	"1237146866": BlueOrangutan,
	"17000":      Holesky,
	"1702448187": WitnesschainMainnet,
	"1":          EthMainnet,
}
