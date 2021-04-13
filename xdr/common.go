package xdr

import (
	"crypto"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"errors"
	"reflect"
)

// fromSlice32 converts a byte slice to a [32]byte
func fromSlice32(slice []byte) ([32]byte, error) {
	if len(slice) != 32 {
		return [32]byte{}, errors.New("slice was not the correct length")
	}
	var array [32]byte
	copy(array[:], slice)
	return array, nil
}

// fromSlice64 converts a byte slice to a [64]byte
func fromSlice64(slice []byte) ([64]byte, error) {
	if len(slice) != 64 {
		return [64]byte{}, errors.New("slice was not the correct length")
	}
	var array [64]byte
	copy(array[:], slice)
	return array, nil
}

// IDFromSlice gets an Id from a byte slice.
func IDFromSlice(slice []byte) (ID, error) {
	return fromSlice32(slice)
}

// IDFromHexString : from hexadecimal string to ID
func IDFromHexString(hexString string) (ID, error) {
	bb, err := hex.DecodeString(hexString)
	if err != nil {
		return ID{}, err
	}
	return IDFromSlice(bb)
}

// SignatureFromSlice gets a Signature from a byte slice.
func SignatureFromSlice(slice []byte) (Signature, error) {
	return fromSlice64(slice)
}

// SignatureFromHexString : from hexadecimal string to Signature
func SignatureFromHexString(hexString string) (Signature, error) {
	bb, err := hex.DecodeString(hexString)
	if err != nil {
		return Signature{}, err
	}
	return SignatureFromSlice(bb)
}

// HashFromSlice gets a Hash from a byte slice.
func HashFromSlice(slice []byte) (Hash, error) {
	return fromSlice32(slice)
}

// HashFromHexString : from hexadecimal string to Hash
func HashFromHexString(hexString string) (Hash, error) {
	bb, err := hex.DecodeString(hexString)
	if err != nil {
		return Hash{}, err
	}
	return HashFromSlice(bb)
}

// Hash32FromSlice gets a Hash32 from a byte slice.
func Hash32FromSlice(slice []byte) (Hash32, error) {
	return fromSlice32(slice)
}

// Hash32FromHexString : from hexadecimal string to Hash32
func Hash32FromHexString(hexString string) (Hash32, error) {
	bb, err := hex.DecodeString(hexString)
	if err != nil {
		return Hash32{}, err
	}
	return Hash32FromSlice(bb)
}

// Hash64FromSlice gets a Hash64 from a byte slice.
func Hash64FromSlice(slice []byte) (Hash64, error) {
	return fromSlice64(slice)
}

// Hash64FromHexString : from hexadecimal string to Hash64
func Hash64FromHexString(hexString string) (Hash64, error) {
	bb, err := hex.DecodeString(hexString)
	if err != nil {
		return Hash64{}, err
	}
	return Hash64FromSlice(bb)
}

// IDFromPublicKey : from generic crypto.PublicKey interface{} we try to extract an ID
func IDFromPublicKey(pk crypto.PublicKey) (ID, error) {
	bbytes, ok := pk.(ed25519.PublicKey)
	if ok == false {
		return ID{}, errors.New("public key not a slice of bytes")
	}
	return IDFromSlice(bbytes)
}

// MarshalJSON : when marshaling to JSON we want ID to be represented as a hex-encoded string, not as an array of uint8
func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(id[:]))
}

// UnmarshalJSON : when marshaling to JSON we want ID to be represented as a hex-encoded string, not as an array of uint8
func (id *ID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	tempID, err := IDFromHexString(s)
	if err != nil {
		return &json.MarshalerError{Type: reflect.TypeOf(id), Err: err}
	}
	copy(id[:], tempID[:])
	return nil
}

// MarshalJSON : when marshaling to JSON we want signature to be represented as a hex-encoded string, not as an array of uint8
func (sig Signature) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(sig[:]))
}

// UnmarshalJSON : when marshaling to JSON we want Signature to be represented as a hex-encoded string, not as an array of uint8
func (sig *Signature) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	tempSig, err := SignatureFromHexString(s)
	if err != nil {
		return &json.MarshalerError{Type: reflect.TypeOf(sig), Err: err}
	}
	copy(sig[:], tempSig[:])
	return nil
}

// MarshalJSON : when marshaling to JSON we want Hash to be represented as a hex-encoded string, not as an array of uint8
func (h Hash) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(h[:]))
}

// UnmarshalJSON : when marshaling to JSON we want Hash to be represented as a hex-encoded string, not as an array of uint8
func (h *Hash) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	tempHash, err := HashFromHexString(s)
	if err != nil {
		return &json.MarshalerError{Type: reflect.TypeOf(h), Err: err}
	}
	copy(h[:], tempHash[:])
	return nil
}

// MarshalJSON : when marshaling to JSON we want Hash32 to be represented as a hex-encoded string, not as an array of uint8
func (h Hash32) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(h[:]))
}

// UnmarshalJSON : when marshaling to JSON we want Hash32 to be represented as a hex-encoded string, not as an array of uint8
func (h *Hash32) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	tempHash, err := Hash32FromHexString(s)
	if err != nil {
		return &json.MarshalerError{Type: reflect.TypeOf(h), Err: err}
	}
	copy(h[:], tempHash[:])
	return nil
}

// MarshalJSON : when marshaling to JSON we want Hash64 to be represented as a hex-encoded string, not as an array of uint8
func (h Hash64) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(h[:]))
}

// UnmarshalJSON : when marshaling to JSON we want Hash64 to be represented as a hex-encoded string, not as an array of uint8
func (h *Hash64) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	tempHash, err := Hash64FromHexString(s)
	if err != nil {
		return &json.MarshalerError{Type: reflect.TypeOf(h), Err: err}
	}
	copy(h[:], tempHash[:])
	return nil
}
