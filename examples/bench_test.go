package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/kochavalabs/mazzaroth-proto/pb"
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"testing"
)

var signature = [64]byte{1, 2, 3}
var address = [32]byte{6, 7, 8}
var contract = []byte{1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
var channelID = [32]byte{3, 4, 5}

func getPb() *pb.Transaction {
	return &pb.Transaction{
		Signature: signature[:],
		Address:   address[:],
		Action: &pb.Action{
			ChannelId: channelID[:],
			Nonce:     4,
			ActionCategory: &pb.Action_Update{
				Update: &pb.Update{
					Contract: contract,
				},
			},
		},
	}
}

func getXdr() xdr.Transaction {
	return xdr.Transaction{
		Signature: signature,
		Address:   address,
		Action: xdr.Action{
			ChannelId: channelID,
			Nonce:     4,
			Category: xdr.ActionCategory{
				Type: xdr.ActionCategoryTypeUpdate,
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

func BenchmarkProtoSerialize(b *testing.B) {
	txPb := getPb()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(txPb)
	}
}

func BenchmarkXdrDeserialize(b *testing.B) {
	txXdr := getXdr()
	bytes, _ := txXdr.MarshalBinary()
	for i := 0; i < b.N; i++ {
		txXdr.UnmarshalBinary(bytes)
	}
}

func BenchmarkProtoDeserialize(b *testing.B) {
	txPb := getPb()
	bytes, _ := proto.Marshal(txPb)
	dest := &pb.Transaction{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(bytes, dest)
	}
}
