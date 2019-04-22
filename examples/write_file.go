package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/kochavalabs/mazzaroth-proto/pb"
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"os"
)

const (
	protoFile = "pbout"
	xdrFile   = "xdrout"
	txWritten = 10000
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Write a transaction XDR file with txWritten TXs and an equicalent proto file
// for comparison
func main() {
	signature := [64]byte{1, 2, 3}
	address := [32]byte{6, 7, 8}
	contract := []byte{1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	channelID := [32]byte{3, 4, 5}

	txXdr := xdr.Transaction{
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

	txPb := pb.Transaction{
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

	pF, err := os.Create(protoFile)
	check(err)
	defer pF.Close()

	xF, err := os.Create(xdrFile)
	check(err)
	defer xF.Close()

	for i := 0; i < txWritten; i++ {
		pbBytes, _ := proto.Marshal(&txPb)
		pF.Write(pbBytes)

		xdr.Marshal(xF, txXdr)
	}
}
