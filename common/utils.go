package dcl_common

import (
	"errors"

	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ChallengerRegistry"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ProverRegistry"
	op_common "github.com/witnesschain-com/operator-cli/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

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

func IsOperatorWhitelisted(operator common.Address, registryInterface interface{}) bool {
	var active bool
	var err error
	switch registry := registryInterface.(type) {
	case *ProverRegistry.ProverRegistry:
		active, err = registry.IsWhiteListed(&bind.CallOpts{}, operator)
	case *ChallengerRegistry.ChallengerRegistry:
		active, err = registry.IsAllowlisted(&bind.CallOpts{}, operator)
	default:
		err = errors.New("invalid registry type")
	}

	op_common.CheckError(err, "Error checking if operator is whitelisted")
	return active
}
