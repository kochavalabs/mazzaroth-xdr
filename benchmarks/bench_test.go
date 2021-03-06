package main

import (
	"testing"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
)

var signature = [64]byte{1, 2, 3}
var address = xdr.ID{6, 7, 8}
var contract = []byte{1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
var channelID = [32]byte{3, 4, 5}

func getXdr() xdr.Transaction {
	return xdr.Transaction{
		Signature: signature,
		Signer: xdr.Authority{
			Type:   xdr.AuthorityTypeNONE,
			Origin: &address,
		},
		Action: xdr.Action{
			ChannelID: channelID,
			Nonce:     4,
			Category: xdr.ActionCategory{
				Type: xdr.ActionCategoryTypeUPDATE,
				Update: &xdr.Update{
					Contract: contract,
				},
			},
		},
	}
}

func BenchmarkXdrSerialize(b *testing.B) {
	txXdr := getXdr()
	for i := 0; i < b.N; i++ {
		txXdr.MarshalBinary()
	}
}

func BenchmarkXdrDeserialize(b *testing.B) {
	txXdr := getXdr()
	bytes, _ := txXdr.MarshalBinary()
	for i := 0; i < b.N; i++ {
		txXdr.UnmarshalBinary(bytes)
	}
}
