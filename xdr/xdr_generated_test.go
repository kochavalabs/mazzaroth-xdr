package xdr

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_API_JSONMarshal(t *testing.T) {
	tests := []struct {
		responseType ResponseType
		object       interface{}
		expected     string
	}{
		{ResponseTypeUNKNOWN, nil, `{"type":0,"data":""}`},
		{ResponseTypeTRANSACTION, Transaction{Action: Action{Nonce: 1}}, `{"type":1,"data":{"signature":"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","signer":{"type":0,"data":""},"action":{"address":"0000000000000000000000000000000000000000000000000000000000000000","channelID":"0000000000000000000000000000000000000000000000000000000000000000","nonce":1,"blockExpirationNumber":0,"category":{"type":0,"data":""}}}}`},
		{ResponseTypeTRANSACTIONLIST, []Transaction{}, `{"type":2,"data":[]}`},
		{ResponseTypeRECEIPT, Receipt{}, `{"type":3,"data":{"status":0,"stateRoot":"0000000000000000000000000000000000000000000000000000000000000000","result":"","statusInfo":""}}`},
		{ResponseTypeRECEIPTLIST, []Receipt{}, `{"type":4,"data":[]}`},
		{ResponseTypeBLOCK, Block{}, `{"type":5,"data":{"header":{"blockHeight":0,"transactionHeight":0,"consensusSequenceNumber":0,"transactionsMerkleRoot":"0000000000000000000000000000000000000000000000000000000000000000","transactionsReceiptRoot":"0000000000000000000000000000000000000000000000000000000000000000","stateRoot":"0000000000000000000000000000000000000000000000000000000000000000","previousHeader":"0000000000000000000000000000000000000000000000000000000000000000","status":0},"transactions":null}}`},
		{ResponseTypeBLOCKLIST, []Block{}, `{"type":6,"data":[]}`},
		{ResponseTypeBLOCKHEADER, BlockHeader{}, `{"type":7,"data":{"blockHeight":0,"transactionHeight":0,"consensusSequenceNumber":0,"transactionsMerkleRoot":"0000000000000000000000000000000000000000000000000000000000000000","transactionsReceiptRoot":"0000000000000000000000000000000000000000000000000000000000000000","stateRoot":"0000000000000000000000000000000000000000000000000000000000000000","previousHeader":"0000000000000000000000000000000000000000000000000000000000000000","status":0}}`},
		{ResponseTypeBLOCKHEADERLIST, []BlockHeader{}, `{"type":8,"data":[]}`},
		{ResponseTypeCHANNEL, ChannelConfig{}, `{"type":9,"data":{"owner":"0000000000000000000000000000000000000000000000000000000000000000","admins":null}}`},
		{ResponseTypeCHANNELLIST, []ChannelConfig{}, `{"type":10,"data":[]}`},
		{ResponseTypeACCOUNT, Account{}, `{"type":11,"data":{"alias":"","transactionCount":0,"authorizedAccounts":null}}`},
		{ResponseTypeHEIGHT, BlockHeight{}, `{"type":12,"data":{"height":0}}`},
		{ResponseTypeABI, Abi{}, `{"type":13,"data":{"version":"","functions":null}}`},
		{ResponseTypeABI, Abi{Version: "1.0", Functions: []FunctionSignature{{FunctionName: "One"}, {FunctionName: "Two"}, {FunctionName: "Three"}, {FunctionName: "Four"}, {FunctionName: "Five"}}}, `{"type":13,"data":{"version":"1.0","functions":[{"functionType":0,"functionName":"One","parameters":null,"returns":null},{"functionType":0,"functionName":"Two","parameters":null,"returns":null},{"functionType":0,"functionName":"Three","parameters":null,"returns":null},{"functionType":0,"functionName":"Four","parameters":null,"returns":null},{"functionType":0,"functionName":"Five","parameters":null,"returns":null}]}}`},
	}

	for i, tt := range tests {
		response, err := NewResponse(tt.responseType, tt.object)
		if err != nil {
			t.Errorf("Error creating Response for type %d: %v", i, err)
			continue
		}

		// Marshal Response to JSON
		respBytes, err := json.Marshal(response)
		if err != nil {
			t.Errorf("Error marshalling response to JSON for type %d: %v", i, err)
			continue
		}

		require.Equal(t, tt.expected, string(respBytes), "Expected %s got %s", tt.expected, respBytes)
	}
}

func Test_API_JSONUnmarshal(t *testing.T) {
	tests := []struct {
		responseType ResponseType
		data         string
		expected     Response
	}{
		{ResponseTypeUNKNOWN, `{"type":0,"data":""}`, Response{Type: ResponseTypeUNKNOWN}},
		{ResponseTypeTRANSACTION, `{"type":1,"data":{"signature":"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","signer":{"type":0,"data":""},"action":{"address":"0000000000000000000000000000000000000000000000000000000000000000","channelID":"0000000000000000000000000000000000000000000000000000000000000000","nonce":1,"blockExpirationNumber":0,"category":{"type":0,"data":""}}}}`,
			Response{Type: ResponseTypeTRANSACTION, Transaction: &Transaction{Action: Action{Nonce: 1}}}},
		{ResponseTypeTRANSACTIONLIST, `{"type":2,"data":[{"type":1,"data":{"signature":"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","signer":{"type":0,"data":""},"action":{"address":"0000000000000000000000000000000000000000000000000000000000000000","channelID":"0000000000000000000000000000000000000000000000000000000000000000","nonce":1,"blockExpirationNumber":0,"category":{"type":0,"data":""}}}}]}`,
			Response{Type: ResponseTypeTRANSACTIONLIST, Transactions: &[]Transaction{{}}}},
		{ResponseTypeRECEIPT, `{"type":3,"data":{"status":0,"stateRoot":"0000000000000000000000000000000000000000000000000000000000000000","result":"","statusInfo":""}}`, Response{Type: ResponseTypeRECEIPT, Receipt: &Receipt{}}},
		{ResponseTypeRECEIPTLIST, `{"type":4,"data":[{"status":0,"stateRoot":"0000000000000000000000000000000000000000000000000000000000000000","result":"","statusInfo":""}]}`, Response{Type: ResponseTypeRECEIPTLIST, Receipts: &[]Receipt{{}}}},
		{ResponseTypeBLOCK, `{"type":5,"data":{"header":{"blockHeight":0,"transactionHeight":0,"consensusSequenceNumber":0,"transactionsMerkleRoot":"0000000000000000000000000000000000000000000000000000000000000000","transactionsReceiptRoot":"0000000000000000000000000000000000000000000000000000000000000000","stateRoot":"0000000000000000000000000000000000000000000000000000000000000000","previousHeader":"0000000000000000000000000000000000000000000000000000000000000000","status":0},"transactions":null}}`,
			Response{Type: ResponseTypeBLOCK, Block: &Block{}}},
		{ResponseTypeBLOCKLIST, `{"type":6,"data":[{"type":1,"data":{"signature":"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","signer":{"type":0,"data":""},"action":{"address":"0000000000000000000000000000000000000000000000000000000000000000","channelID":"0000000000000000000000000000000000000000000000000000000000000000","nonce":1,"blockExpirationNumber":0,"category":{"type":0,"data":""}}}}]}`,
			Response{Type: ResponseTypeBLOCKLIST, Blocks: &[]Block{{}}}},
		{ResponseTypeBLOCKHEADER, `{"type":7,"data":{"blockHeight":0,"transactionHeight":0,"consensusSequenceNumber":0,"transactionsMerkleRoot":"0000000000000000000000000000000000000000000000000000000000000000","transactionsReceiptRoot":"0000000000000000000000000000000000000000000000000000000000000000","stateRoot":"0000000000000000000000000000000000000000000000000000000000000000","previousHeader":"0000000000000000000000000000000000000000000000000000000000000000","status":0}}`,
			Response{Type: ResponseTypeBLOCKHEADER, BlockHeader: &BlockHeader{}}},
		{ResponseTypeBLOCKHEADERLIST, `{"type":8,"data":[{"type":1,"data":{"signature":"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","signer":{"type":0,"data":""},"action":{"address":"0000000000000000000000000000000000000000000000000000000000000000","channelID":"0000000000000000000000000000000000000000000000000000000000000000","nonce":1,"blockExpirationNumber":0,"category":{"type":0,"data":""}}}}]}`,
			Response{Type: ResponseTypeBLOCKHEADERLIST, BlockHeaders: &[]BlockHeader{{}}}},
		{ResponseTypeCHANNEL, `{"type":9,"data":{"owner":"0000000000000000000000000000000000000000000000000000000000000000","admins":null}}`, Response{Type: ResponseTypeCHANNEL, Channel: &ChannelConfig{}}},
		{ResponseTypeCHANNELLIST, `{"type":10,"data":[{"type":1,"data":{"signature":"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","signer":{"type":0,"data":""},"action":{"address":"0000000000000000000000000000000000000000000000000000000000000000","channelID":"0000000000000000000000000000000000000000000000000000000000000000","nonce":1,"blockExpirationNumber":0,"category":{"type":0,"data":""}}}}]}`,
			Response{Type: ResponseTypeCHANNELLIST, Channels: &[]ChannelConfig{{}}}},
		{ResponseTypeACCOUNT, `{"type":11,"data":{"alias":"","transactionCount":0,"authorizedAccounts":null}}`, Response{Type: ResponseTypeACCOUNT, Account: &Account{}}},
		{ResponseTypeHEIGHT, `{"type":12,"data":{"height":0}}`, Response{Type: ResponseTypeHEIGHT, Height: &BlockHeight{}}},
		{ResponseTypeABI, `{"type":13,"data":{"version":"1.0","functions":[]}}`, Response{Type: ResponseTypeABI, Abi: &Abi{Version: "1.0", Functions: []FunctionSignature{}}}},
		{ResponseTypeABI, `{"type":13,"data":{"version":"1.0","functions":[{"functionName": "One"},{"functionName": "Two"},{"functionName": "Three"},{"functionName": "Four"},{"functionName": "Five"}]}}`, Response{Type: ResponseTypeABI, Abi: &Abi{Version: "1.0", Functions: []FunctionSignature{{FunctionName: "One"}, {FunctionName: "Two"}, {FunctionName: "Three"}, {FunctionName: "Four"}, {FunctionName: "Five"}}}}},
	}

	for i, tt := range tests {
		// Unmarshal JSON to response
		response := Response{}
		err := json.Unmarshal([]byte(tt.data), &response)
		if err != nil {
			t.Errorf("Error unmarshalling response from JSON for type %d: %v", i, err)
			continue
		}

		require.Equal(t, tt.expected, response, "Expected %v got %v", tt.expected, response)
	}
}
