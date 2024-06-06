package operator_commands

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	wc_common "github.com/witnesschain-com/dcl-operator-cli/common"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ProverRegistry"
)

func GetProverSignature(client *ethclient.Client, proverRegistry *ProverRegistry.ProverRegistry, proverAddress common.Address, proverPrivateKey *ecdsa.PrivateKey, operatorAddress common.Address, salt [32]byte, expiry *big.Int) []byte {
	digestHash, err := proverRegistry.CalculateProverRegistrationDigestHash(&bind.CallOpts{}, proverAddress, operatorAddress, salt, expiry)
	wc_common.CheckError(err, "Digest hash calculation failed")

	signature, err := crypto.Sign(digestHash[:], proverPrivateKey)
	wc_common.CheckError(err, "Signing the digest hash failed")

	v := new(big.Int).SetBytes(signature[64:])
	v.Add(v, big.NewInt(27))

	// Construct the full signature (r, s, v)
	fullSignature := append(signature[:64], v.Bytes()...)

	return fullSignature
}
