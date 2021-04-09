package xdr

import (
	"crypto"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
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

// SignatureFromSlice gets a Signature from a byte slice.
func SignatureFromSlice(slice []byte) (Signature, error) {
	return fromSlice64(slice)
}

// HashFromSlice gets a Hash from a byte slice.
func HashFromSlice(slice []byte) (Hash, error) {
	return fromSlice32(slice)
}

// IDFromPublicKey : from generic crypto.PublicKey interface{} we try to extract an ID
func IDFromPublicKey(pk crypto.PublicKey) (ID, error) {
	bbytes, ok := pk.([]byte)
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
	bb, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	if len(bb) != len(id) {
		return &json.MarshalerError{Type: reflect.TypeOf(id), Err: fmt.Errorf("length of decoded string incorrect: %d, should be %d", len(bb), len(id))}
	}
	copy(id[:], bb)
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
	bb, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	if len(bb) != len(sig) {
		return &json.MarshalerError{Type: reflect.TypeOf(sig), Err: fmt.Errorf("length of decoded string incorrect: %d, should be %d", len(bb), len(sig))}
	}
	copy(sig[:], bb)
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
	bb, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	if len(bb) != len(h) {
		return &json.MarshalerError{Type: reflect.TypeOf(h), Err: fmt.Errorf("length of decoded string incorrect: %d, should be %d", len(bb), len(h))}
	}
	copy(h[:], bb)
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
	bb, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	if len(bb) != len(h) {
		return &json.MarshalerError{Type: reflect.TypeOf(h), Err: fmt.Errorf("length of decoded string incorrect: %d, should be %d", len(bb), len(h))}
	}
	copy(h[:], bb)
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
	bb, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	if len(bb) != len(h) {
		return &json.MarshalerError{Type: reflect.TypeOf(h), Err: fmt.Errorf("length of decoded string incorrect: %d, should be %d", len(bb), len(h))}
	}
	copy(h[:], bb)
	return nil
}
