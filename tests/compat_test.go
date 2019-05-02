package crypto

import (
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/kochavalabs/mazzaroth-xdr/xdr_old"
	"reflect"
	"testing"
)

var signature = [64]byte{1, 2, 3}
var address = [32]byte{6, 7, 8}
var contract = []byte{1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
var channelID = [32]byte{3, 4, 5}

func getXdrOld() xdr_old.Transaction {
	return xdr_old.Transaction{
		Signature: signature,
		Address:   address,
		Action: xdr_old.Action{
			ChannelId: channelID,
			Nonce:     4,
			Category: xdr_old.ActionCategory{
				Type: xdr_old.ActionCategoryTypeUpdate,
				Update: &xdr_old.Update{
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

func TestOldVsNew(t *testing.T) {
	bytes, _ := getXdr().MarshalBinary()
	oldBytes, _ := getXdrOld().MarshalBinary()

	if !reflect.DeepEqual(bytes, oldBytes) {
		t.Errorf("Got %d, want %d", bytes, oldBytes)
	}
}
