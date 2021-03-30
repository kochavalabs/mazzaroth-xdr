package xdr

import (
	"crypto"
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

// IdFromPublicKey : from generic crypto.PublicKey interface{} we try to extract an ID
func IdFromPublicKey(pk crypto.PublicKey) (ID, error) {
	bbytes, ok := pk.([]byte)
	if ok == false {
		return ID{}, errors.New("public key not a slice of bytes")
	}
	return IdFromSlice(bbytes)
}
