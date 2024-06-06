package dcl_common

import (
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ProverRegistry"
	op_common "github.com/witnesschain-com/operator-cli/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func IsOperatorWhitelisted(operator common.Address, proverRegistry *ProverRegistry.ProverRegistry) bool {
	active, err := proverRegistry.IsWhiteListed(&bind.CallOpts{}, operator)
	op_common.CheckError(err, "Error checking if operator is whitelisted")
	return active
}

func IsProverRegistered(prover common.Address, operator common.Address, proverRegistry *ProverRegistry.ProverRegistry) bool {
	registered, err := proverRegistry.IsRegisteredProver(&bind.CallOpts{}, prover, operator)
	op_common.CheckError(err, "Error checking if watchtower is already registered")
	return registered
}
