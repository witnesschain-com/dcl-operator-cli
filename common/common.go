package common

import (
	"math/big"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ChallengerRegistry"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ProverRegistry"
	op_common "github.com/witnesschain-com/operator-cli/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	RPC                       string
	GasPrice                  int
}

var BlueOrangutan = ChainConfig{
	ChallengerRegistryAddress: common.HexToAddress("0xeFFE8c100029F71924554aEd382f1919ecA6b203"),
	ProverRegistryAddress:     common.HexToAddress("0x91013d3CecE055603D8b1EE7DCB1f670f480fe24"),
	ChainID:                   *big.NewInt(1237146866),
	BlockExplorer:             "https://blue-orangutan-blockscout.eu-north-2.gateway.fm",
	RPC:                       "https://blue-orangutan-rpc.eu-north-2.gateway.fm",
	GasPrice: -1,
}

var Holesky = ChainConfig{
	ProverRegistryAddress:     common.HexToAddress(""),
	ChallengerRegistryAddress: common.HexToAddress(""),
	ChainID:                   *big.NewInt(17000),
	BlockExplorer:             "https://holesky.etherscan.io",
	GasPrice: 0,
}

var WitnesschainMainnet = ChainConfig{
	ProverRegistryAddress:     common.HexToAddress("0xCdb30BE21A44fB111A48661ECc755B34a41C4e82"),
	ChallengerRegistryAddress: common.HexToAddress("0x49220De1c883D9358e3cCe0aB42304460216Dc2C"),
	BlockExplorer:             "https://explorer.witnesschain.com",
	RPC:                       "https://rpc.witnesschain.com",
	ChainID:                   *big.NewInt(1702448187),
	GasPrice: -1,
}

var EthMainnet = ChainConfig{
	ProverRegistryAddress:     common.HexToAddress(""),
	ChallengerRegistryAddress: common.HexToAddress(""),
	ChainID:                   *big.NewInt(1),
	BlockExplorer:             "https://etherscan.io",
	GasPrice: 0,
}

var NetworkConfig = map[string]ChainConfig{
	"1237146866": BlueOrangutan,
	"17000":      Holesky,
	"1702448187": WitnesschainMainnet,
	"1":          EthMainnet,
}

func IsProverRegistered(prover common.Address, operator common.Address, proverRegistry *ProverRegistry.ProverRegistry) bool {
	registered, err := proverRegistry.IsRegisteredProver(&bind.CallOpts{}, prover, operator)
	op_common.CheckError(err, "Error checking if watchtower is already registered")
	return registered
}

func IsValidChallenger(challenger common.Address, challengerRegistry *ChallengerRegistry.ChallengerRegistry) bool {
	valid, err := challengerRegistry.IsValidChallenger(&bind.CallOpts{}, challenger)
	op_common.CheckError(err, "Error checking if challenger is already registered")
	return valid
}
