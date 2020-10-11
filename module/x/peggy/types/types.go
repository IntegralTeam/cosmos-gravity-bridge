package types

import (
	"fmt"
	"math/big"
	"sort"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

// BridgeValidator represents the validator data in the Ethereum bridge MultiSig set.
type BridgeValidator struct {
	Power           uint64          `json:"power"`
	EthereumAddress EthereumAddress `json:"ethereum_address"`
}

func (b BridgeValidator) ValidateBasic() error {
	if b.Power == 0 {
		return sdkerrors.Wrap(ErrEmpty, "power")
	}
	if b.EthereumAddress.IsEmpty() {
		return sdkerrors.Wrap(ErrEmpty, "address")
	}
	return nil
}

// BridgeValidators is the sorted set of validator data for Ethereum bridge MultiSig set
type BridgeValidators []BridgeValidator

func (b BridgeValidators) Sort() {
	sort.SliceIsSorted(b, func(i, j int) bool {
		if b[i].Power == b[j].Power {
			// Secondary sort on eth address in case powers are equal
			return b[i].EthereumAddress.LessThan(b[j].EthereumAddress)
		}
		return b[i].Power < b[j].Power
	})
}

func (b BridgeValidators) HasDuplicates() bool {
	m := make(map[EthereumAddress]struct{}, len(b))
	for i := range b {
		m[b[i].EthereumAddress] = struct{}{}
	}
	return len(m) != len(b)
}

// GetPowers returns only the power values for all members
func (b BridgeValidators) GetPowers() []uint64 {
	r := make([]uint64, len(b))
	for i := range b {
		r[i] = b[i].Power
	}
	return r
}

func (b BridgeValidators) ValidateBasic() error {
	if len(b) == 0 {
		return ErrEmpty
	}
	for i := range b {
		if err := b[i].ValidateBasic(); err != nil {
			return sdkerrors.Wrapf(err, "member %d", i)
		}
	}
	if b.HasDuplicates() {
		return sdkerrors.Wrap(ErrDuplicate, "addresses")
	}
	return nil
}

// Valset is the Ethereum Bridge Multsig Set
type Valset struct {
	Nonce   UInt64Nonce      `json:"nonce"`
	Members BridgeValidators `json:"members"`
}

func NewValset(nonce UInt64Nonce, members BridgeValidators) Valset {
	members.Sort()
	return Valset{Nonce: nonce, Members: members}
}

func (v Valset) GetCheckpoint() []byte {
	// TODO replace hardcoded "foo" here with a getter to retrieve the correct PeggyID from the store
	// this will work for now because 'foo' is the test Peggy ID we are using
	var peggyIDString = "foo"

	// The go-ethereum ABI encoder *only* encodes function calls and then it only encodes
	// function calls for which you provide an ABI json just like you would get out of the
	// solidity compiler with your compiled contract.
	// You are supposed to compile your contract, use abigen to generate an ABI , import
	// this generated go module and then use for that for all testing and development.
	// This abstraction layer is more trouble than it's worth, because we don't want to
	// encode a function call at all, but instead we want to emulate a Solidity encode operation
	// which has no equal available from go-ethereum.
	//
	// In order to work around this absurd series of problems we have to manually write the below
	// 'function specification' that will encode the same arguments into a function call. We can then
	// truncate the first several bytes where the call name is encoded to finally get the equal of the
	const checkpointAbiJSON = `[{
	  "inputs": [
	    {
	      "internalType": "bytes32",
	      "name": "_peggyId",
	      "type": "bytes32"
	    },
	    {
	      "internalType": "bytes32",
	      "name": "_checkpoint",
	      "type": "bytes32"
	    },
	    {
	      "internalType": "uint256",
	      "name": "_valsetNonce",
	      "type": "uint256"
	    },
	    {
	      "internalType": "address[]",
	      "name": "_validators",
	      "type": "address[]"
	    },
	    {
	      "internalType": "uint256[]",
	      "name": "_powers",
	      "type": "uint256[]"
	    }
	  ],
	  "name": "checkpoint",
	  "outputs": [
	    {
	      "internalType": "bytes32",
	      "name": "",
	      "type": "bytes32"
	    }
	  ],
	  "stateMutability": "pure",
	  "type": "function"
	}]`
	// Solidity abi.Encode() call.
	// error case here should not occur outside of testing since the above is a constant
	contractAbi, abiErr := abi.JSON(strings.NewReader(checkpointAbiJSON))
	if abiErr != nil {
		panic("Bad ABI constant!")
	}
	peggyIDBytes := []uint8(peggyIDString)
	// the contract argument is not a arbitrary length array but a fixed length 32 byte
	// array, therefore we have to utf8 encode the string (the default in this case) and
	// then copy the variable length encoded data into a fixed length array. This function
	// will panic if peggyId is too long to fit in 32 bytes
	var peggyID [32]uint8
	copy(peggyID[:], peggyIDBytes[:])
	checkpointBytes := []uint8("checkpoint")
	var checkpoint [32]uint8
	copy(checkpoint[:], checkpointBytes[:])

	memberAddresses := make([]EthereumAddress, len(v.Members))
	convertedPowers := make([]*big.Int, len(v.Members))
	for i, m := range v.Members {
		memberAddresses[i] = m.EthereumAddress
		convertedPowers[i] = big.NewInt(int64(m.Power))
	}
	// the word 'checkpoint' needs to be the same as the 'name' above in the checkpointAbiJson
	// but other than that it's a constant that has no impact on the output. This is because
	// it gets encoded as a function name which we must then discard.
	bytes, packErr := contractAbi.Pack("checkpoint", peggyID, checkpoint, big.NewInt(int64(v.Nonce.Uint64())), memberAddresses, convertedPowers)

	// this should never happen outside of test since any case that could crash on encoding
	// should be filtered above.
	if packErr != nil {
		panic(fmt.Sprintf("Error packing checkpoint! %s/n", packErr))
	}

	// we hash the resulting encoded bytes discarding the first 4 bytes these 4 bytes are the constant
	// method name 'checkpoint'. If you where to replace the checkpoint constant in this code you would
	// then need to adjust how many bytes you truncate off the front to get the output of abi.encode()
	hash := crypto.Keccak256Hash(bytes[4:])
	return hash.Bytes()
}
