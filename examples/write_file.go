package main

import (
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"os"
)

const (
	xdrFile   = "xdrout"
	txWritten = 10000
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Write a transaction XDR file with txWritten TXs
func main() {
	signature := [64]byte{1, 2, 3}
	address := [32]byte{6, 7, 8}
	contract := []byte{1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	channelID := [32]byte{3, 4, 5}

	txXdr := xdr.Transaction{
		Signature: signature,
		Address:   address,
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

	xF, err := os.Create(xdrFile)
	check(err)
	defer xF.Close()

	for i := 0; i < txWritten; i++ {
		xdr.Marshal(xF, txXdr)
	}
}
