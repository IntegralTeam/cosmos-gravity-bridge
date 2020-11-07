package types

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/crypto"
)

const signaturePrefix = "\x19Ethereum Signed Message:\n32"

func NewEthereumSignature(hash []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	if privateKey == nil {
		return nil, sdkerrors.Wrap(ErrEmpty, "private key")
	}
	protectedHash := crypto.Keccak256Hash(append([]uint8(signaturePrefix), hash...))
	return crypto.Sign(protectedHash.Bytes(), privateKey)
}

func ValidateEthereumSignature(hash []byte, signature []byte, ethAddress string) error {
	if len(signature) < 65 {
		return sdkerrors.Wrap(ErrInvalid, "signature too short")
	}
	// To verify signature
	// - use crypto.SigToPub to get the public key
	// - use crypto.PubkeyToAddress to get the address
	// - compare this to the address given.

	// for backwards compatibility reasons  the V value of an Ethereum sig is presented
	// as 27 or 28, internally though it should be a 0-3 value due to changed formats.
	// It seems that go-ethereum expects this to be done before sigs actually reach it's
	// internal validation functions. In order to comply with this requirement we check
	// the sig an dif it's in standard format we correct it. If it's in go-ethereum's expected
	// format already we make no changes.
	//
	// We could attempt to break or otherwise exit early on obviously invalid values for this
	// byte, but that's a task best left to go-ethereum
	if signature[64] == 27 || signature[64] == 28 {
		signature[64] -= 27
	}

	protectedHash := crypto.Keccak256Hash(append([]uint8(signaturePrefix), hash...))

	pubkey, err := crypto.SigToPub(protectedHash.Bytes(), signature)
	if err != nil {
		return sdkerrors.Wrap(err, "signature to public key")
	}

	addr := crypto.PubkeyToAddress(*pubkey)

	if addr.Hex() != ethAddress {
		return sdkerrors.Wrap(ErrInvalid, "signature not matching")
	}

	return nil
}

// SignType defines what has been signed by an orchestrator
type SignType byte

const (
	SignTypeUnknown SignType = 0
	// signed confirmations on cosmos for Ethereum side
	SignTypeOrchestratorSignedMultiSigUpdate SignType = 1
	SignTypeOrchestratorSignedWithdrawBatch  SignType = 2
)

var signTypeToNames = map[SignType]string{
	SignTypeOrchestratorSignedMultiSigUpdate: "orchestrator_signed_multisig_update",
	SignTypeOrchestratorSignedWithdrawBatch:  "orchestrator_signed_withdraw_batch",
}

// AllSignTypes types that are signed with by the bridge multisig set
var AllSignTypes = []SignType{SignTypeOrchestratorSignedMultiSigUpdate, SignTypeOrchestratorSignedWithdrawBatch}

func IsSignType(s SignType) bool {
	for _, v := range AllSignTypes {
		if s == v {
			return true
		}
	}
	return false
}

func SignTypeFromName(s string) (SignType, bool) {
	for _, v := range AllSignTypes {
		name, ok := signTypeToNames[v]
		if ok && name == s {
			return v, true
		}
	}
	return SignTypeUnknown, false
}
func ToSignTypeNames(s ...SignType) []string {
	r := make([]string, len(s))
	for i := range s {
		r[i] = s[i].String()
	}
	return r
}

func (t SignType) String() string {
	return signTypeToNames[t]
}

func (t SignType) Bytes() []byte {
	return []byte{byte(t)}
}

func (t SignType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", t.String())), nil
}

func (t *SignType) UnmarshalJSON(input []byte) error {
	if string(input) == `""` {
		return nil
	}
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return err
	}
	c, exists := SignTypeFromName(s)
	if !exists {
		return sdkerrors.Wrap(ErrUnknown, "claim type")
	}
	*t = c
	return nil
}
