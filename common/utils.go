package common

import (
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ChallengerRegistry"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ProverRegistry"
	op_common "github.com/witnesschain-com/operator-cli/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func IsProverRegistered(prover common.Address, operator common.Address, proverRegistry *ProverRegistry.ProverRegistry) bool {
	registered, err := proverRegistry.IsRegisteredProver(&bind.CallOpts{}, prover, operator)
	op_common.CheckError(err, "Error checking if prover is already registered")
	return registered
}

func IsValidChallenger(challenger common.Address, challengerRegistry *ChallengerRegistry.ChallengerRegistry) bool {
	valid, err := challengerRegistry.IsValidChallenger(&bind.CallOpts{}, challenger)
	op_common.CheckError(err, "Error checking if challenger is already registered")
	return valid
}
