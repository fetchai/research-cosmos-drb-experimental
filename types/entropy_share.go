package types

import (
	"errors"
	"fmt"

	"github.com/tendermint/tendermint/crypto"
)

// For event switch in entropy generator
const (
	EventComputedEntropy = "EventComputedEntropy"
	// TODO: Check this is ok with mcl
	maxEntropyShareSize = 256
	GenesisHeight       = int64(0)
)

type ThresholdSignature = []byte

type ComputedEntropy struct {
	Height         int64
	GroupSignature ThresholdSignature
}

func (ce *ComputedEntropy) IsEmpty() bool {
	return ce.GroupSignature == nil || len(ce.GroupSignature) == 0
}

//-----------------------------------------------------------------------------
// Wrappers for signing entropy message

type CanonicalEntropyShare struct {
	Height         int64
	SignerAddress  crypto.Address
	SignatureShare string
	ChainID        string
}

func CanonicalizeEntropyShare(chainID string, entropy *EntropyShare) CanonicalEntropyShare {
	return CanonicalEntropyShare{
		Height:         entropy.Height,
		SignerAddress:  entropy.SignerAddress,
		SignatureShare: entropy.SignatureShare,
		ChainID:        chainID,
	}
}

//-----------------------------------------------------------------------------

type EntropyShare struct {
	Height         int64          `json:"height"`
	SignerAddress  crypto.Address `json:"signer"`
	SignatureShare string         `json:"entropy_signature"`
	Signature      []byte         `json:"signature"`
}

// ValidateBasic performs basic validation.
func (entropy *EntropyShare) ValidateBasic() error {
	if entropy.Height < GenesisHeight+1 {
		return errors.New("invalid Height")
	}

	if len(entropy.SignerAddress) != crypto.AddressSize {
		return fmt.Errorf("expected ValidatorAddress size to be %d bytes, got %d bytes",
			crypto.AddressSize,
			len(entropy.SignerAddress),
		)
	}
	if len(entropy.SignatureShare) == 0 {
		return errors.New("signature is missing")
	}
	if len(entropy.SignatureShare) > maxEntropyShareSize {
		return fmt.Errorf("signature is too big (max: %d)", maxEntropyShareSize)
	}
	return nil
}

// String returns a string representation of EntropyShare
func (entropy EntropyShare) String() string {
	return entropy.StringIndented("")
}

// StringIndented returns a string representation of the EntropyShare
func (entropy EntropyShare) StringIndented(indent string) string {
	return fmt.Sprintf(`EntropySignatureShare{
%s  %v/%v/%v
%s}`,
		indent, entropy.Height, entropy.SignerAddress, entropy.SignatureShare,
		indent)
}

func (entropy EntropyShare) Copy() EntropyShare {
	return EntropyShare{
		Height:         entropy.Height,
		SignerAddress:  entropy.SignerAddress,
		SignatureShare: entropy.SignatureShare,
		Signature:      entropy.Signature,
	}
}

//-----------------------------------------------------------
// These methods are for Protobuf Compatibility

// Marshal returns the amino encoding.
func (entropy *EntropyShare) Marshal() ([]byte, error) {
	return cdc.MarshalBinaryBare(entropy)
}

// MarshalTo calls Marshal and copies to the given buffer.
func (entropy *EntropyShare) MarshalTo(data []byte) (int, error) {
	bs, err := entropy.Marshal()
	if err != nil {
		return -1, err
	}
	return copy(data, bs), nil
}

// Unmarshal deserializes from amino encoded form.
func (entropy *EntropyShare) Unmarshal(bs []byte) error {
	return cdc.UnmarshalBinaryBare(bs, entropy)
}

// For signing with private key
func (entropy *EntropyShare) SignBytes(chainID string) []byte {
	bz, err := cdc.MarshalBinaryLengthPrefixed(CanonicalizeEntropyShare(chainID, entropy))
	if err != nil {
		panic(err)
	}
	return bz
}
