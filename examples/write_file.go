package main

import (
	"fmt"
	"github.com/kochavalabs/mazzaroth-proto/pb"
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
)

// Write a transaction XDR file with 100 TXs and an equicalent proto file for
// comparison
func main() {
	signature := [64]byte{1, 2, 3}
	address := [32]byte{6, 7, 8}
	// contract := []byte{1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	channelID := [32]byte{3, 4, 5}
	txXdr := xdr.Transaction{
		Signature: signature,
		Address:   address,
		Action:    xdr.Action{},
	}

	txPb := pb.Transaction{
		Signature: signature[:],
		Address:   address[:],
		Action: &pb.Action{
			ChannelId: channelID[:],
			Nonce:     4,
		},
	}

	fmt.Println(txXdr)
	fmt.Println(txPb)
}
