package main

import (
	"encoding/json"
	"testing"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
)

var signature = [64]byte{1, 2, 3}
var address = xdr.ID{6, 7, 8}
var contract = []byte{1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
var channelID = [32]byte{3, 4, 5}

func getXdr() xdr.Transaction {
	return xdr.Transaction{
		Sender:    address,
		Signer:    address,
		Signature: signature,
		Data: &xdr.Data{
			ChannelID: channelID,
			Nonce:     4,
			Category: xdr.Category{
				Type: xdr.CategoryTypeCONTRACT,
				Contract: &xdr.Contract{
					ContractBytes: contract,
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

func BenchmarkJSONSerialize(b *testing.B) {
	txXdr := getXdr()
	for i := 0; i < b.N; i++ {
		json.Marshal(txXdr)
	}
}

func BenchmarkJSONDeserialize(b *testing.B) {
	txXdr := getXdr()
	bytes, _ := json.Marshal(txXdr)
	for i := 0; i < b.N; i++ {
		json.Unmarshal(bytes, &xdr.Transaction{})
	}
}
