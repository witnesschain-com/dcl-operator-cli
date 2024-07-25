package operator_commands

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ChallengerRegistry"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ProverRegistry"
	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
	op_common "github.com/witnesschain-com/operator-cli/common"
)

func GetProverSignature(client *ethclient.Client, proverRegistry *ProverRegistry.ProverRegistry, proverAddress common.Address, vault *keystore.Vault, operatorAddress common.Address, salt [32]byte, expiry *big.Int) []byte {
	digestHash, err := proverRegistry.CalculateProverRegistrationDigestHash(&bind.CallOpts{}, proverAddress, operatorAddress, salt, expiry)
	op_common.CheckError(err, "Digest hash calculation failed")

	fullSignature, err := vault.SignData(digestHash[:], apitypes.DataTyped.Mime)
	op_common.CheckError(err, "unable to sign prover hash")

	return fullSignature
}

func GetChallengerSignature(client *ethclient.Client, challengerRegistry *ChallengerRegistry.ChallengerRegistry, challengerAddress common.Address, vault *keystore.Vault, operatorAddress common.Address, salt [32]byte, expiry *big.Int) []byte {
	digestHash, err := challengerRegistry.CalculateChallengerRegistrationDigestHash(&bind.CallOpts{}, challengerAddress, operatorAddress, salt, expiry)
	op_common.CheckError(err, "Digest hash calculation failed")

	fullSignature, err := vault.SignData(digestHash[:], apitypes.DataTyped.Mime)
	op_common.CheckError(err, "unable to sign challenger hash")

	return fullSignature
}
