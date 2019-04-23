package xdr

import (
	"errors"
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

// IdFromSlice gets an Id from a byte slice.
func IdFromSlice(slice []byte) (Id, error) {
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
