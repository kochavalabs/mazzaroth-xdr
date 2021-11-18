// Package xdr is automatically generated
// DO NOT EDIT or your changes may be overwritten
package xdr

import (
	"bytes"
	"encoding"
	"encoding/json"
	"fmt"
	"io"

	"github.com/stellar/go-xdr/xdr3"
)

// Unmarshal reads an xdr element from `r` into `v`.
func Unmarshal(r io.Reader, v interface{}) (int, error) {
	// delegate to xdr package's Unmarshal
	return xdr.Unmarshal(r, v)
}

// Marshal writes an xdr element `v` into `w`.
func Marshal(w io.Writer, v interface{}) (int, error) {
	// delegate to xdr package's Marshal
	return xdr.Marshal(w, v)
}

// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

// Abi generated struct
type Abi struct {
	Version string `json:"version"`

	Functions []FunctionSignature `json:"functions"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Abi) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Abi) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Abi)(nil)
	_ encoding.BinaryUnmarshaler = (*Abi)(nil)
)

// FunctionSignature generated struct
type FunctionSignature struct {
	FunctionType FunctionType `json:"functionType"`

	FunctionName string `json:"functionName"`

	Parameters []Parameter `json:"parameters"`

	Returns []Parameter `json:"returns"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s FunctionSignature) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *FunctionSignature) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*FunctionSignature)(nil)
	_ encoding.BinaryUnmarshaler = (*FunctionSignature)(nil)
)

// Parameter generated struct
type Parameter struct {
	ParameterName string `json:"parameterName"`

	ParameterType string `json:"parameterType"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Parameter) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Parameter) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Parameter)(nil)
	_ encoding.BinaryUnmarshaler = (*Parameter)(nil)
)

// End struct section

// Start enum section

// FunctionType generated enum
type FunctionType int32

const (

	// FunctionTypeUNKNOWN enum value 0
	FunctionTypeUNKNOWN FunctionType = 0

	// FunctionTypeREAD enum value 1
	FunctionTypeREAD FunctionType = 1

	// FunctionTypeWRITE enum value 2
	FunctionTypeWRITE FunctionType = 2
)

// FunctionTypeMap generated enum map
var FunctionTypeMap = map[int32]string{

	0: "FunctionTypeUNKNOWN",

	1: "FunctionTypeREAD",

	2: "FunctionTypeWRITE",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for FunctionType
func (s FunctionType) ValidEnum(v int32) bool {
	_, ok := FunctionTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s FunctionType) String() string {
	name, _ := FunctionTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s FunctionType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *FunctionType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*FunctionType)(nil)
	_ encoding.BinaryUnmarshaler = (*FunctionType)(nil)
)

// End enum section

// Start union section

// End union section

// Namespace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

// Account generated struct
type Account struct {
	Alias string `xdrmaxsize:"32" json:"alias"`

	TransactionCount uint64 `json:"transactionCount,string"`

	AuthorizedAccounts []AuthorizedAccount `xdrmaxsize:"32" json:"authorizedAccounts"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Account) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Account) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Account)(nil)
	_ encoding.BinaryUnmarshaler = (*Account)(nil)
)

// AuthorizedAccount generated struct
type AuthorizedAccount struct {
	Key ID `json:"key"`

	Alias string `xdrmaxsize:"32" json:"alias"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AuthorizedAccount) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AuthorizedAccount) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AuthorizedAccount)(nil)
	_ encoding.BinaryUnmarshaler = (*AuthorizedAccount)(nil)
)

// End struct section

// Start enum section

// End enum section

// Start union section

// End union section

// Namespace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

// End struct section

// Start enum section

// RequestType generated enum
type RequestType int32

const (

	// RequestTypeUNKNOWN enum value 0
	RequestTypeUNKNOWN RequestType = 0

	// RequestTypeTRANSACTION enum value 1
	RequestTypeTRANSACTION RequestType = 1
)

// RequestTypeMap generated enum map
var RequestTypeMap = map[int32]string{

	0: "RequestTypeUNKNOWN",

	1: "RequestTypeTRANSACTION",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for RequestType
func (s RequestType) ValidEnum(v int32) bool {
	_, ok := RequestTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s RequestType) String() string {
	name, _ := RequestTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s RequestType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *RequestType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*RequestType)(nil)
	_ encoding.BinaryUnmarshaler = (*RequestType)(nil)
)

// ResponseType generated enum
type ResponseType int32

const (

	// ResponseTypeUNKNOWN enum value 0
	ResponseTypeUNKNOWN ResponseType = 0

	// ResponseTypeTRANSACTIONID enum value 1
	ResponseTypeTRANSACTIONID ResponseType = 1

	// ResponseTypeTRANSACTION enum value 2
	ResponseTypeTRANSACTION ResponseType = 2

	// ResponseTypeTRANSACTIONLIST enum value 3
	ResponseTypeTRANSACTIONLIST ResponseType = 3

	// ResponseTypeRECEIPT enum value 4
	ResponseTypeRECEIPT ResponseType = 4

	// ResponseTypeRECEIPTLIST enum value 5
	ResponseTypeRECEIPTLIST ResponseType = 5

	// ResponseTypeBLOCK enum value 6
	ResponseTypeBLOCK ResponseType = 6

	// ResponseTypeBLOCKLIST enum value 7
	ResponseTypeBLOCKLIST ResponseType = 7

	// ResponseTypeBLOCKHEADER enum value 8
	ResponseTypeBLOCKHEADER ResponseType = 8

	// ResponseTypeBLOCKHEADERLIST enum value 9
	ResponseTypeBLOCKHEADERLIST ResponseType = 9

	// ResponseTypeCHANNEL enum value 10
	ResponseTypeCHANNEL ResponseType = 10

	// ResponseTypeCHANNELLIST enum value 11
	ResponseTypeCHANNELLIST ResponseType = 11

	// ResponseTypeACCOUNT enum value 12
	ResponseTypeACCOUNT ResponseType = 12

	// ResponseTypeHEIGHT enum value 13
	ResponseTypeHEIGHT ResponseType = 13

	// ResponseTypeABI enum value 14
	ResponseTypeABI ResponseType = 14
)

// ResponseTypeMap generated enum map
var ResponseTypeMap = map[int32]string{

	0: "ResponseTypeUNKNOWN",

	1: "ResponseTypeTRANSACTIONID",

	2: "ResponseTypeTRANSACTION",

	3: "ResponseTypeTRANSACTIONLIST",

	4: "ResponseTypeRECEIPT",

	5: "ResponseTypeRECEIPTLIST",

	6: "ResponseTypeBLOCK",

	7: "ResponseTypeBLOCKLIST",

	8: "ResponseTypeBLOCKHEADER",

	9: "ResponseTypeBLOCKHEADERLIST",

	10: "ResponseTypeCHANNEL",

	11: "ResponseTypeCHANNELLIST",

	12: "ResponseTypeACCOUNT",

	13: "ResponseTypeHEIGHT",

	14: "ResponseTypeABI",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ResponseType
func (s ResponseType) ValidEnum(v int32) bool {
	_, ok := ResponseTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s ResponseType) String() string {
	name, _ := ResponseTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ResponseType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ResponseType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ResponseType)(nil)
	_ encoding.BinaryUnmarshaler = (*ResponseType)(nil)
)

// End enum section

// Start union section

// Request generated union
type Request struct {
	Type RequestType

	Transaction *Transaction
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u Request) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of Request
func (u Request) ArmForSwitch(sw int32) (string, bool) {
	switch RequestType(sw) {

	case RequestTypeUNKNOWN:
		return "", true

	case RequestTypeTRANSACTION:
		return "Transaction", true
	}
	return "-", false
}

// NewRequest creates a new  Request.
func NewRequest(aType RequestType, value interface{}) (result Request, err error) {
	result.Type = aType
	switch aType {

	case RequestTypeUNKNOWN:

	case RequestTypeTRANSACTION:

		tv, ok := value.(Transaction)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Transaction = &tv

	}
	return
}

// MustTransaction retrieves the Transaction value from the union,
// panicing if the value is not set.
func (u Request) MustTransaction() Transaction {

	val, ok := u.GetTransaction()
	if !ok {
		panic("arm Transaction is not set")
	}

	return val
}

// GetTransaction retrieves the Transaction value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Request) GetTransaction() (result Transaction, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Transaction" {
		result = *u.Transaction
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u Request) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *Request) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Request)(nil)
	_ encoding.BinaryUnmarshaler = (*Request)(nil)
)

// MarshalJSON implements json.Marshaler.
func (u Request) MarshalJSON() ([]byte, error) {
	temp := struct {
		Type int32       `json:"type"`
		Data interface{} `json:"data"`
	}{}

	temp.Type = int32(u.Type)
	temp.Data = ""
	switch u.Type {
	case RequestTypeUNKNOWN:
	case RequestTypeTRANSACTION:
		temp.Data = u.Transaction
	default:
		return nil, fmt.Errorf("invalid union type")
	}

	return json.Marshal(temp)
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *Request) UnmarshalJSON(data []byte) error {
	temp := struct {
		Type int32 `json:"type"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	u.Type = RequestType(temp.Type)
	switch u.Type {
	case RequestTypeUNKNOWN:

	case RequestTypeTRANSACTION:
		response := struct {
			Transaction Transaction `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Transaction = &response.Transaction

	default:
		return fmt.Errorf("invalid union type")
	}

	return nil
}

// Response generated union
type Response struct {
	Type ResponseType

	TransactionID *ID

	Transaction *Transaction

	Transactions *[]Transaction

	Receipt *Receipt

	Receipts *[]Receipt

	Block *Block

	Blocks *[]Block

	BlockHeader *BlockHeader

	BlockHeaders *[]BlockHeader

	Channel *ChannelConfig

	Channels *[]ChannelConfig

	Account *Account

	Height *BlockHeight

	Abi *Abi
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u Response) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of Response
func (u Response) ArmForSwitch(sw int32) (string, bool) {
	switch ResponseType(sw) {

	case ResponseTypeUNKNOWN:
		return "", true

	case ResponseTypeTRANSACTIONID:
		return "TransactionID", true

	case ResponseTypeTRANSACTION:
		return "Transaction", true

	case ResponseTypeTRANSACTIONLIST:
		return "Transactions", true

	case ResponseTypeRECEIPT:
		return "Receipt", true

	case ResponseTypeRECEIPTLIST:
		return "Receipts", true

	case ResponseTypeBLOCK:
		return "Block", true

	case ResponseTypeBLOCKLIST:
		return "Blocks", true

	case ResponseTypeBLOCKHEADER:
		return "BlockHeader", true

	case ResponseTypeBLOCKHEADERLIST:
		return "BlockHeaders", true

	case ResponseTypeCHANNEL:
		return "Channel", true

	case ResponseTypeCHANNELLIST:
		return "Channels", true

	case ResponseTypeACCOUNT:
		return "Account", true

	case ResponseTypeHEIGHT:
		return "Height", true

	case ResponseTypeABI:
		return "Abi", true
	}
	return "-", false
}

// NewResponse creates a new  Response.
func NewResponse(aType ResponseType, value interface{}) (result Response, err error) {
	result.Type = aType
	switch aType {

	case ResponseTypeUNKNOWN:

	case ResponseTypeTRANSACTIONID:

		tv, ok := value.(ID)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.TransactionID = &tv

	case ResponseTypeTRANSACTION:

		tv, ok := value.(Transaction)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Transaction = &tv

	case ResponseTypeTRANSACTIONLIST:

		tv, ok := value.([]Transaction)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Transactions = &tv

	case ResponseTypeRECEIPT:

		tv, ok := value.(Receipt)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Receipt = &tv

	case ResponseTypeRECEIPTLIST:

		tv, ok := value.([]Receipt)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Receipts = &tv

	case ResponseTypeBLOCK:

		tv, ok := value.(Block)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Block = &tv

	case ResponseTypeBLOCKLIST:

		tv, ok := value.([]Block)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Blocks = &tv

	case ResponseTypeBLOCKHEADER:

		tv, ok := value.(BlockHeader)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.BlockHeader = &tv

	case ResponseTypeBLOCKHEADERLIST:

		tv, ok := value.([]BlockHeader)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.BlockHeaders = &tv

	case ResponseTypeCHANNEL:

		tv, ok := value.(ChannelConfig)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Channel = &tv

	case ResponseTypeCHANNELLIST:

		tv, ok := value.([]ChannelConfig)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Channels = &tv

	case ResponseTypeACCOUNT:

		tv, ok := value.(Account)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Account = &tv

	case ResponseTypeHEIGHT:

		tv, ok := value.(BlockHeight)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Height = &tv

	case ResponseTypeABI:

		tv, ok := value.(Abi)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Abi = &tv

	}
	return
}

// MustTransactionID retrieves the TransactionID value from the union,
// panicing if the value is not set.
func (u Response) MustTransactionID() ID {

	val, ok := u.GetTransactionID()
	if !ok {
		panic("arm TransactionID is not set")
	}

	return val
}

// GetTransactionID retrieves the TransactionID value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Response) GetTransactionID() (result ID, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "TransactionID" {
		result = *u.TransactionID
		ok = true
	}

	return
}

// MustTransaction retrieves the Transaction value from the union,
// panicing if the value is not set.
func (u Response) MustTransaction() Transaction {

	val, ok := u.GetTransaction()
	if !ok {
		panic("arm Transaction is not set")
	}

	return val
}

// GetTransaction retrieves the Transaction value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Response) GetTransaction() (result Transaction, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Transaction" {
		result = *u.Transaction
		ok = true
	}

	return
}

// MustTransactions retrieves the Transactions value from the union,
// panicing if the value is not set.
func (u Response) MustTransactions() []Transaction {

	val, ok := u.GetTransactions()
	if !ok {
		panic("arm Transactions is not set")
	}

	return val
}

// GetTransactions retrieves the Transactions value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Response) GetTransactions() (result []Transaction, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Transactions" {
		result = *u.Transactions
		ok = true
	}

	return
}

// MustReceipt retrieves the Receipt value from the union,
// panicing if the value is not set.
func (u Response) MustReceipt() Receipt {

	val, ok := u.GetReceipt()
	if !ok {
		panic("arm Receipt is not set")
	}

	return val
}

// GetReceipt retrieves the Receipt value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Response) GetReceipt() (result Receipt, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Receipt" {
		result = *u.Receipt
		ok = true
	}

	return
}

// MustReceipts retrieves the Receipts value from the union,
// panicing if the value is not set.
func (u Response) MustReceipts() []Receipt {

	val, ok := u.GetReceipts()
	if !ok {
		panic("arm Receipts is not set")
	}

	return val
}

// GetReceipts retrieves the Receipts value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Response) GetReceipts() (result []Receipt, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Receipts" {
		result = *u.Receipts
		ok = true
	}

	return
}

// MustBlock retrieves the Block value from the union,
// panicing if the value is not set.
func (u Response) MustBlock() Block {

	val, ok := u.GetBlock()
	if !ok {
		panic("arm Block is not set")
	}

	return val
}

// GetBlock retrieves the Block value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Response) GetBlock() (result Block, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Block" {
		result = *u.Block
		ok = true
	}

	return
}

// MustBlocks retrieves the Blocks value from the union,
// panicing if the value is not set.
func (u Response) MustBlocks() []Block {

	val, ok := u.GetBlocks()
	if !ok {
		panic("arm Blocks is not set")
	}

	return val
}

// GetBlocks retrieves the Blocks value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Response) GetBlocks() (result []Block, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Blocks" {
		result = *u.Blocks
		ok = true
	}

	return
}

// MustBlockHeader retrieves the BlockHeader value from the union,
// panicing if the value is not set.
func (u Response) MustBlockHeader() BlockHeader {

	val, ok := u.GetBlockHeader()
	if !ok {
		panic("arm BlockHeader is not set")
	}

	return val
}

// GetBlockHeader retrieves the BlockHeader value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Response) GetBlockHeader() (result BlockHeader, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "BlockHeader" {
		result = *u.BlockHeader
		ok = true
	}

	return
}

// MustBlockHeaders retrieves the BlockHeaders value from the union,
// panicing if the value is not set.
func (u Response) MustBlockHeaders() []BlockHeader {

	val, ok := u.GetBlockHeaders()
	if !ok {
		panic("arm BlockHeaders is not set")
	}

	return val
}

// GetBlockHeaders retrieves the BlockHeaders value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Response) GetBlockHeaders() (result []BlockHeader, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "BlockHeaders" {
		result = *u.BlockHeaders
		ok = true
	}

	return
}

// MustChannel retrieves the Channel value from the union,
// panicing if the value is not set.
func (u Response) MustChannel() ChannelConfig {

	val, ok := u.GetChannel()
	if !ok {
		panic("arm Channel is not set")
	}

	return val
}

// GetChannel retrieves the Channel value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Response) GetChannel() (result ChannelConfig, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Channel" {
		result = *u.Channel
		ok = true
	}

	return
}

// MustChannels retrieves the Channels value from the union,
// panicing if the value is not set.
func (u Response) MustChannels() []ChannelConfig {

	val, ok := u.GetChannels()
	if !ok {
		panic("arm Channels is not set")
	}

	return val
}

// GetChannels retrieves the Channels value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Response) GetChannels() (result []ChannelConfig, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Channels" {
		result = *u.Channels
		ok = true
	}

	return
}

// MustAccount retrieves the Account value from the union,
// panicing if the value is not set.
func (u Response) MustAccount() Account {

	val, ok := u.GetAccount()
	if !ok {
		panic("arm Account is not set")
	}

	return val
}

// GetAccount retrieves the Account value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Response) GetAccount() (result Account, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Account" {
		result = *u.Account
		ok = true
	}

	return
}

// MustHeight retrieves the Height value from the union,
// panicing if the value is not set.
func (u Response) MustHeight() BlockHeight {

	val, ok := u.GetHeight()
	if !ok {
		panic("arm Height is not set")
	}

	return val
}

// GetHeight retrieves the Height value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Response) GetHeight() (result BlockHeight, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Height" {
		result = *u.Height
		ok = true
	}

	return
}

// MustAbi retrieves the Abi value from the union,
// panicing if the value is not set.
func (u Response) MustAbi() Abi {

	val, ok := u.GetAbi()
	if !ok {
		panic("arm Abi is not set")
	}

	return val
}

// GetAbi retrieves the Abi value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Response) GetAbi() (result Abi, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Abi" {
		result = *u.Abi
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u Response) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *Response) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Response)(nil)
	_ encoding.BinaryUnmarshaler = (*Response)(nil)
)

// MarshalJSON implements json.Marshaler.
func (u Response) MarshalJSON() ([]byte, error) {
	temp := struct {
		Type int32       `json:"type"`
		Data interface{} `json:"data"`
	}{}

	temp.Type = int32(u.Type)
	temp.Data = ""
	switch u.Type {
	case ResponseTypeUNKNOWN:
	case ResponseTypeTRANSACTIONID:
		temp.Data = u.TransactionID
	case ResponseTypeTRANSACTION:
		temp.Data = u.Transaction
	case ResponseTypeTRANSACTIONLIST:
		temp.Data = u.Transactions
	case ResponseTypeRECEIPT:
		temp.Data = u.Receipt
	case ResponseTypeRECEIPTLIST:
		temp.Data = u.Receipts
	case ResponseTypeBLOCK:
		temp.Data = u.Block
	case ResponseTypeBLOCKLIST:
		temp.Data = u.Blocks
	case ResponseTypeBLOCKHEADER:
		temp.Data = u.BlockHeader
	case ResponseTypeBLOCKHEADERLIST:
		temp.Data = u.BlockHeaders
	case ResponseTypeCHANNEL:
		temp.Data = u.Channel
	case ResponseTypeCHANNELLIST:
		temp.Data = u.Channels
	case ResponseTypeACCOUNT:
		temp.Data = u.Account
	case ResponseTypeHEIGHT:
		temp.Data = u.Height
	case ResponseTypeABI:
		temp.Data = u.Abi
	default:
		return nil, fmt.Errorf("invalid union type")
	}

	return json.Marshal(temp)
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *Response) UnmarshalJSON(data []byte) error {
	temp := struct {
		Type int32 `json:"type"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	u.Type = ResponseType(temp.Type)
	switch u.Type {
	case ResponseTypeUNKNOWN:

	case ResponseTypeTRANSACTIONID:
		response := struct {
			TransactionID ID `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.TransactionID = &response.TransactionID

	case ResponseTypeTRANSACTION:
		response := struct {
			Transaction Transaction `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Transaction = &response.Transaction

	case ResponseTypeTRANSACTIONLIST:
		response := struct {
			Transactions []Transaction `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Transactions = &response.Transactions

	case ResponseTypeRECEIPT:
		response := struct {
			Receipt Receipt `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Receipt = &response.Receipt

	case ResponseTypeRECEIPTLIST:
		response := struct {
			Receipts []Receipt `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Receipts = &response.Receipts

	case ResponseTypeBLOCK:
		response := struct {
			Block Block `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Block = &response.Block

	case ResponseTypeBLOCKLIST:
		response := struct {
			Blocks []Block `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Blocks = &response.Blocks

	case ResponseTypeBLOCKHEADER:
		response := struct {
			BlockHeader BlockHeader `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.BlockHeader = &response.BlockHeader

	case ResponseTypeBLOCKHEADERLIST:
		response := struct {
			BlockHeaders []BlockHeader `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.BlockHeaders = &response.BlockHeaders

	case ResponseTypeCHANNEL:
		response := struct {
			Channel ChannelConfig `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Channel = &response.Channel

	case ResponseTypeCHANNELLIST:
		response := struct {
			Channels []ChannelConfig `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Channels = &response.Channels

	case ResponseTypeACCOUNT:
		response := struct {
			Account Account `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Account = &response.Account

	case ResponseTypeHEIGHT:
		response := struct {
			Height BlockHeight `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Height = &response.Height

	case ResponseTypeABI:
		response := struct {
			Abi Abi `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Abi = &response.Abi

	default:
		return fmt.Errorf("invalid union type")
	}

	return nil
}

// End union section

// Namespace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

// Block generated struct
type Block struct {
	Header BlockHeader `json:"header"`

	Transactions []Transaction `json:"transactions"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Block) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Block) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Block)(nil)
	_ encoding.BinaryUnmarshaler = (*Block)(nil)
)

// BlockHeader generated struct
type BlockHeader struct {
	BlockHeight uint64 `json:"blockHeight,string"`

	TransactionHeight uint64 `json:"transactionHeight,string"`

	ConsensusSequenceNumber uint64 `json:"consensusSequenceNumber,string"`

	TransactionsMerkleRoot Hash `json:"transactionsMerkleRoot"`

	TransactionsReceiptRoot Hash `json:"transactionsReceiptRoot"`

	StateRoot Hash `json:"stateRoot"`

	PreviousHeader Hash `json:"previousHeader"`

	Status BlockStatus `json:"status"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BlockHeader) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BlockHeader) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BlockHeader)(nil)
	_ encoding.BinaryUnmarshaler = (*BlockHeader)(nil)
)

// BlockHeight generated struct
type BlockHeight struct {
	Height uint64 `json:"height,string"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BlockHeight) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BlockHeight) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BlockHeight)(nil)
	_ encoding.BinaryUnmarshaler = (*BlockHeight)(nil)
)

// End struct section

// Start enum section

// BlockStatus generated enum
type BlockStatus int32

const (

	// BlockStatusUNKNOWN enum value 0
	BlockStatusUNKNOWN BlockStatus = 0

	// BlockStatusPENDING enum value 1
	BlockStatusPENDING BlockStatus = 1

	// BlockStatusFINALIZED enum value 2
	BlockStatusFINALIZED BlockStatus = 2
)

// BlockStatusMap generated enum map
var BlockStatusMap = map[int32]string{

	0: "BlockStatusUNKNOWN",

	1: "BlockStatusPENDING",

	2: "BlockStatusFINALIZED",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for BlockStatus
func (s BlockStatus) ValidEnum(v int32) bool {
	_, ok := BlockStatusMap[v]
	return ok
}

// String returns the name of `e`
func (s BlockStatus) String() string {
	name, _ := BlockStatusMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BlockStatus) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BlockStatus) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BlockStatus)(nil)
	_ encoding.BinaryUnmarshaler = (*BlockStatus)(nil)
)

// End enum section

// Start union section

// End union section

// Namespace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// Signature generated typedef
type Signature [64]byte

// XDRMaxSize implements the Sized interface for Signature
func (s Signature) XDRMaxSize() int {
	return 64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Signature) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Signature) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Signature)(nil)
	_ encoding.BinaryUnmarshaler = (*Signature)(nil)
)

// ID generated typedef
type ID [32]byte

// XDRMaxSize implements the Sized interface for ID
func (s ID) XDRMaxSize() int {
	return 32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ID) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ID) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ID)(nil)
	_ encoding.BinaryUnmarshaler = (*ID)(nil)
)

// Hash generated typedef
type Hash [32]byte

// XDRMaxSize implements the Sized interface for Hash
func (s Hash) XDRMaxSize() int {
	return 32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Hash) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Hash) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Hash)(nil)
	_ encoding.BinaryUnmarshaler = (*Hash)(nil)
)

// Argument generated typedef
type Argument string

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Argument) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Argument) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Argument)(nil)
	_ encoding.BinaryUnmarshaler = (*Argument)(nil)
)

// StatusInfo generated typedef
type StatusInfo string

// MarshalBinary implements encoding.BinaryMarshaler.
func (s StatusInfo) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *StatusInfo) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*StatusInfo)(nil)
	_ encoding.BinaryUnmarshaler = (*StatusInfo)(nil)
)

// End typedef section

// Start struct section

// End struct section

// Start enum section

// End enum section

// Start union section

// End union section

// Namespace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

// ChannelConfig generated struct
type ChannelConfig struct {
	Owner ID `json:"owner"`

	Admins []ID `xdrmaxsize:"32" json:"admins"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ChannelConfig) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ChannelConfig) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ChannelConfig)(nil)
	_ encoding.BinaryUnmarshaler = (*ChannelConfig)(nil)
)

// End struct section

// Start enum section

// End enum section

// Start union section

// End union section

// Namespace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

// Receipt generated struct
type Receipt struct {
	Status ReceiptStatus `json:"status"`

	StateRoot Hash `json:"stateRoot"`

	Result string `json:"result"`

	StatusInfo StatusInfo `json:"statusInfo"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Receipt) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Receipt) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Receipt)(nil)
	_ encoding.BinaryUnmarshaler = (*Receipt)(nil)
)

// End struct section

// Start enum section

// ReceiptStatus generated enum
type ReceiptStatus int32

const (

	// ReceiptStatusUNKNOWN enum value 0
	ReceiptStatusUNKNOWN ReceiptStatus = 0

	// ReceiptStatusFAILURE enum value 1
	ReceiptStatusFAILURE ReceiptStatus = 1

	// ReceiptStatusSUCCESS enum value 2
	ReceiptStatusSUCCESS ReceiptStatus = 2
)

// ReceiptStatusMap generated enum map
var ReceiptStatusMap = map[int32]string{

	0: "ReceiptStatusUNKNOWN",

	1: "ReceiptStatusFAILURE",

	2: "ReceiptStatusSUCCESS",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ReceiptStatus
func (s ReceiptStatus) ValidEnum(v int32) bool {
	_, ok := ReceiptStatusMap[v]
	return ok
}

// String returns the name of `e`
func (s ReceiptStatus) String() string {
	name, _ := ReceiptStatusMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ReceiptStatus) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ReceiptStatus) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ReceiptStatus)(nil)
	_ encoding.BinaryUnmarshaler = (*ReceiptStatus)(nil)
)

// End enum section

// Start union section

// End union section

// Namespace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

// Call generated struct
type Call struct {
	Function string `xdrmaxsize:"256" json:"function"`

	Arguments []Argument `json:"arguments"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Call) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Call) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Call)(nil)
	_ encoding.BinaryUnmarshaler = (*Call)(nil)
)

// Contract generated struct
type Contract struct {
	ContractBytes []byte `json:"contractBytes"`

	Abi Abi `json:"abi"`

	ContractHash Hash `json:"contractHash"`

	Version string `xdrmaxsize:"100" json:"version"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Contract) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Contract) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Contract)(nil)
	_ encoding.BinaryUnmarshaler = (*Contract)(nil)
)

// Authorization generated struct
type Authorization struct {
	Account AuthorizedAccount `json:"account"`

	Authorize bool `json:"authorize"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Authorization) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Authorization) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Authorization)(nil)
	_ encoding.BinaryUnmarshaler = (*Authorization)(nil)
)

// Action generated struct
type Action struct {
	Address ID `json:"address"`

	ChannelID ID `json:"channelID"`

	Nonce uint64 `json:"nonce,string"`

	BlockExpirationNumber uint64 `json:"blockExpirationNumber,string"`

	Category ActionCategory `json:"category"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Action) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Action) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Action)(nil)
	_ encoding.BinaryUnmarshaler = (*Action)(nil)
)

// Transaction generated struct
type Transaction struct {
	Signature Signature `json:"signature"`

	Signer Authority `json:"signer"`

	Action Action `json:"action"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Transaction) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Transaction) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Transaction)(nil)
	_ encoding.BinaryUnmarshaler = (*Transaction)(nil)
)

// End struct section

// Start enum section

// UpdateType generated enum
type UpdateType int32

const (

	// UpdateTypeUNKNOWN enum value 0
	UpdateTypeUNKNOWN UpdateType = 0

	// UpdateTypeCONTRACT enum value 1
	UpdateTypeCONTRACT UpdateType = 1

	// UpdateTypeCONFIG enum value 2
	UpdateTypeCONFIG UpdateType = 2

	// UpdateTypeACCOUNT enum value 3
	UpdateTypeACCOUNT UpdateType = 3
)

// UpdateTypeMap generated enum map
var UpdateTypeMap = map[int32]string{

	0: "UpdateTypeUNKNOWN",

	1: "UpdateTypeCONTRACT",

	2: "UpdateTypeCONFIG",

	3: "UpdateTypeACCOUNT",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for UpdateType
func (s UpdateType) ValidEnum(v int32) bool {
	_, ok := UpdateTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s UpdateType) String() string {
	name, _ := UpdateTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s UpdateType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *UpdateType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*UpdateType)(nil)
	_ encoding.BinaryUnmarshaler = (*UpdateType)(nil)
)

// AccountUpdateType generated enum
type AccountUpdateType int32

const (

	// AccountUpdateTypeUNKNOWN enum value 0
	AccountUpdateTypeUNKNOWN AccountUpdateType = 0

	// AccountUpdateTypeALIAS enum value 1
	AccountUpdateTypeALIAS AccountUpdateType = 1

	// AccountUpdateTypeAUTHORIZATION enum value 2
	AccountUpdateTypeAUTHORIZATION AccountUpdateType = 2
)

// AccountUpdateTypeMap generated enum map
var AccountUpdateTypeMap = map[int32]string{

	0: "AccountUpdateTypeUNKNOWN",

	1: "AccountUpdateTypeALIAS",

	2: "AccountUpdateTypeAUTHORIZATION",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for AccountUpdateType
func (s AccountUpdateType) ValidEnum(v int32) bool {
	_, ok := AccountUpdateTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s AccountUpdateType) String() string {
	name, _ := AccountUpdateTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountUpdateType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountUpdateType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountUpdateType)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountUpdateType)(nil)
)

// ActionCategoryType generated enum
type ActionCategoryType int32

const (

	// ActionCategoryTypeUNKNOWN enum value 0
	ActionCategoryTypeUNKNOWN ActionCategoryType = 0

	// ActionCategoryTypeCALL enum value 1
	ActionCategoryTypeCALL ActionCategoryType = 1

	// ActionCategoryTypeUPDATE enum value 2
	ActionCategoryTypeUPDATE ActionCategoryType = 2
)

// ActionCategoryTypeMap generated enum map
var ActionCategoryTypeMap = map[int32]string{

	0: "ActionCategoryTypeUNKNOWN",

	1: "ActionCategoryTypeCALL",

	2: "ActionCategoryTypeUPDATE",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ActionCategoryType
func (s ActionCategoryType) ValidEnum(v int32) bool {
	_, ok := ActionCategoryTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s ActionCategoryType) String() string {
	name, _ := ActionCategoryTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ActionCategoryType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ActionCategoryType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ActionCategoryType)(nil)
	_ encoding.BinaryUnmarshaler = (*ActionCategoryType)(nil)
)

// AuthorityType generated enum
type AuthorityType int32

const (

	// AuthorityTypeUNKNOWN enum value 0
	AuthorityTypeUNKNOWN AuthorityType = 0

	// AuthorityTypeSELF enum value 1
	AuthorityTypeSELF AuthorityType = 1

	// AuthorityTypeAUTHORIZED enum value 2
	AuthorityTypeAUTHORIZED AuthorityType = 2
)

// AuthorityTypeMap generated enum map
var AuthorityTypeMap = map[int32]string{

	0: "AuthorityTypeUNKNOWN",

	1: "AuthorityTypeSELF",

	2: "AuthorityTypeAUTHORIZED",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for AuthorityType
func (s AuthorityType) ValidEnum(v int32) bool {
	_, ok := AuthorityTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s AuthorityType) String() string {
	name, _ := AuthorityTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AuthorityType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AuthorityType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AuthorityType)(nil)
	_ encoding.BinaryUnmarshaler = (*AuthorityType)(nil)
)

// End enum section

// Start union section

// Update generated union
type Update struct {
	Type UpdateType

	Contract *Contract

	ChannelConfig *ChannelConfig

	Account *AccountUpdate
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u Update) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of Update
func (u Update) ArmForSwitch(sw int32) (string, bool) {
	switch UpdateType(sw) {

	case UpdateTypeUNKNOWN:
		return "", true

	case UpdateTypeCONTRACT:
		return "Contract", true

	case UpdateTypeCONFIG:
		return "ChannelConfig", true

	case UpdateTypeACCOUNT:
		return "Account", true
	}
	return "-", false
}

// NewUpdate creates a new  Update.
func NewUpdate(aType UpdateType, value interface{}) (result Update, err error) {
	result.Type = aType
	switch aType {

	case UpdateTypeUNKNOWN:

	case UpdateTypeCONTRACT:

		tv, ok := value.(Contract)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Contract = &tv

	case UpdateTypeCONFIG:

		tv, ok := value.(ChannelConfig)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.ChannelConfig = &tv

	case UpdateTypeACCOUNT:

		tv, ok := value.(AccountUpdate)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Account = &tv

	}
	return
}

// MustContract retrieves the Contract value from the union,
// panicing if the value is not set.
func (u Update) MustContract() Contract {

	val, ok := u.GetContract()
	if !ok {
		panic("arm Contract is not set")
	}

	return val
}

// GetContract retrieves the Contract value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Update) GetContract() (result Contract, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Contract" {
		result = *u.Contract
		ok = true
	}

	return
}

// MustChannelConfig retrieves the ChannelConfig value from the union,
// panicing if the value is not set.
func (u Update) MustChannelConfig() ChannelConfig {

	val, ok := u.GetChannelConfig()
	if !ok {
		panic("arm ChannelConfig is not set")
	}

	return val
}

// GetChannelConfig retrieves the ChannelConfig value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Update) GetChannelConfig() (result ChannelConfig, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ChannelConfig" {
		result = *u.ChannelConfig
		ok = true
	}

	return
}

// MustAccount retrieves the Account value from the union,
// panicing if the value is not set.
func (u Update) MustAccount() AccountUpdate {

	val, ok := u.GetAccount()
	if !ok {
		panic("arm Account is not set")
	}

	return val
}

// GetAccount retrieves the Account value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Update) GetAccount() (result AccountUpdate, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Account" {
		result = *u.Account
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u Update) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *Update) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Update)(nil)
	_ encoding.BinaryUnmarshaler = (*Update)(nil)
)

// MarshalJSON implements json.Marshaler.
func (u Update) MarshalJSON() ([]byte, error) {
	temp := struct {
		Type int32       `json:"type"`
		Data interface{} `json:"data"`
	}{}

	temp.Type = int32(u.Type)
	temp.Data = ""
	switch u.Type {
	case UpdateTypeUNKNOWN:
	case UpdateTypeCONTRACT:
		temp.Data = u.Contract
	case UpdateTypeCONFIG:
		temp.Data = u.ChannelConfig
	case UpdateTypeACCOUNT:
		temp.Data = u.Account
	default:
		return nil, fmt.Errorf("invalid union type")
	}

	return json.Marshal(temp)
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *Update) UnmarshalJSON(data []byte) error {
	temp := struct {
		Type int32 `json:"type"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	u.Type = UpdateType(temp.Type)
	switch u.Type {
	case UpdateTypeUNKNOWN:

	case UpdateTypeCONTRACT:
		response := struct {
			Contract Contract `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Contract = &response.Contract

	case UpdateTypeCONFIG:
		response := struct {
			ChannelConfig ChannelConfig `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.ChannelConfig = &response.ChannelConfig

	case UpdateTypeACCOUNT:
		response := struct {
			Account AccountUpdate `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Account = &response.Account

	default:
		return fmt.Errorf("invalid union type")
	}

	return nil
}

// AccountUpdate generated union
type AccountUpdate struct {
	Type AccountUpdateType

	Alias *string

	Authorization *Authorization
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountUpdate) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AccountUpdate
func (u AccountUpdate) ArmForSwitch(sw int32) (string, bool) {
	switch AccountUpdateType(sw) {

	case AccountUpdateTypeUNKNOWN:
		return "", true

	case AccountUpdateTypeALIAS:
		return "Alias", true

	case AccountUpdateTypeAUTHORIZATION:
		return "Authorization", true
	}
	return "-", false
}

// NewAccountUpdate creates a new  AccountUpdate.
func NewAccountUpdate(aType AccountUpdateType, value interface{}) (result AccountUpdate, err error) {
	result.Type = aType
	switch aType {

	case AccountUpdateTypeUNKNOWN:

	case AccountUpdateTypeALIAS:

		tv, ok := value.(string)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Alias = &tv

	case AccountUpdateTypeAUTHORIZATION:

		tv, ok := value.(Authorization)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Authorization = &tv

	}
	return
}

// MustAlias retrieves the Alias value from the union,
// panicing if the value is not set.
func (u AccountUpdate) MustAlias() string {

	val, ok := u.GetAlias()
	if !ok {
		panic("arm Alias is not set")
	}

	return val
}

// GetAlias retrieves the Alias value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountUpdate) GetAlias() (result string, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Alias" {
		result = *u.Alias
		ok = true
	}

	return
}

// MustAuthorization retrieves the Authorization value from the union,
// panicing if the value is not set.
func (u AccountUpdate) MustAuthorization() Authorization {

	val, ok := u.GetAuthorization()
	if !ok {
		panic("arm Authorization is not set")
	}

	return val
}

// GetAuthorization retrieves the Authorization value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountUpdate) GetAuthorization() (result Authorization, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Authorization" {
		result = *u.Authorization
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u AccountUpdate) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *AccountUpdate) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountUpdate)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountUpdate)(nil)
)

// MarshalJSON implements json.Marshaler.
func (u AccountUpdate) MarshalJSON() ([]byte, error) {
	temp := struct {
		Type int32       `json:"type"`
		Data interface{} `json:"data"`
	}{}

	temp.Type = int32(u.Type)
	temp.Data = ""
	switch u.Type {
	case AccountUpdateTypeUNKNOWN:
	case AccountUpdateTypeALIAS:
		temp.Data = u.Alias
	case AccountUpdateTypeAUTHORIZATION:
		temp.Data = u.Authorization
	default:
		return nil, fmt.Errorf("invalid union type")
	}

	return json.Marshal(temp)
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *AccountUpdate) UnmarshalJSON(data []byte) error {
	temp := struct {
		Type int32 `json:"type"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	u.Type = AccountUpdateType(temp.Type)
	switch u.Type {
	case AccountUpdateTypeUNKNOWN:

	case AccountUpdateTypeALIAS:
		response := struct {
			Alias string `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Alias = &response.Alias

	case AccountUpdateTypeAUTHORIZATION:
		response := struct {
			Authorization Authorization `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Authorization = &response.Authorization

	default:
		return fmt.Errorf("invalid union type")
	}

	return nil
}

// ActionCategory generated union
type ActionCategory struct {
	Type ActionCategoryType

	Call *Call

	Update *Update
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ActionCategory) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ActionCategory
func (u ActionCategory) ArmForSwitch(sw int32) (string, bool) {
	switch ActionCategoryType(sw) {

	case ActionCategoryTypeUNKNOWN:
		return "", true

	case ActionCategoryTypeCALL:
		return "Call", true

	case ActionCategoryTypeUPDATE:
		return "Update", true
	}
	return "-", false
}

// NewActionCategory creates a new  ActionCategory.
func NewActionCategory(aType ActionCategoryType, value interface{}) (result ActionCategory, err error) {
	result.Type = aType
	switch aType {

	case ActionCategoryTypeUNKNOWN:

	case ActionCategoryTypeCALL:

		tv, ok := value.(Call)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Call = &tv

	case ActionCategoryTypeUPDATE:

		tv, ok := value.(Update)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Update = &tv

	}
	return
}

// MustCall retrieves the Call value from the union,
// panicing if the value is not set.
func (u ActionCategory) MustCall() Call {

	val, ok := u.GetCall()
	if !ok {
		panic("arm Call is not set")
	}

	return val
}

// GetCall retrieves the Call value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ActionCategory) GetCall() (result Call, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Call" {
		result = *u.Call
		ok = true
	}

	return
}

// MustUpdate retrieves the Update value from the union,
// panicing if the value is not set.
func (u ActionCategory) MustUpdate() Update {

	val, ok := u.GetUpdate()
	if !ok {
		panic("arm Update is not set")
	}

	return val
}

// GetUpdate retrieves the Update value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ActionCategory) GetUpdate() (result Update, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Update" {
		result = *u.Update
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u ActionCategory) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *ActionCategory) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ActionCategory)(nil)
	_ encoding.BinaryUnmarshaler = (*ActionCategory)(nil)
)

// MarshalJSON implements json.Marshaler.
func (u ActionCategory) MarshalJSON() ([]byte, error) {
	temp := struct {
		Type int32       `json:"type"`
		Data interface{} `json:"data"`
	}{}

	temp.Type = int32(u.Type)
	temp.Data = ""
	switch u.Type {
	case ActionCategoryTypeUNKNOWN:
	case ActionCategoryTypeCALL:
		temp.Data = u.Call
	case ActionCategoryTypeUPDATE:
		temp.Data = u.Update
	default:
		return nil, fmt.Errorf("invalid union type")
	}

	return json.Marshal(temp)
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *ActionCategory) UnmarshalJSON(data []byte) error {
	temp := struct {
		Type int32 `json:"type"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	u.Type = ActionCategoryType(temp.Type)
	switch u.Type {
	case ActionCategoryTypeUNKNOWN:

	case ActionCategoryTypeCALL:
		response := struct {
			Call Call `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Call = &response.Call

	case ActionCategoryTypeUPDATE:
		response := struct {
			Update Update `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Update = &response.Update

	default:
		return fmt.Errorf("invalid union type")
	}

	return nil
}

// Authority generated union
type Authority struct {
	Type AuthorityType

	Origin *ID
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u Authority) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of Authority
func (u Authority) ArmForSwitch(sw int32) (string, bool) {
	switch AuthorityType(sw) {

	case AuthorityTypeUNKNOWN:
		return "", true

	case AuthorityTypeSELF:
		return "", true

	case AuthorityTypeAUTHORIZED:
		return "Origin", true
	}
	return "-", false
}

// NewAuthority creates a new  Authority.
func NewAuthority(aType AuthorityType, value interface{}) (result Authority, err error) {
	result.Type = aType
	switch aType {

	case AuthorityTypeUNKNOWN:

	case AuthorityTypeSELF:

	case AuthorityTypeAUTHORIZED:

		tv, ok := value.(ID)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Origin = &tv

	}
	return
}

// MustOrigin retrieves the Origin value from the union,
// panicing if the value is not set.
func (u Authority) MustOrigin() ID {

	val, ok := u.GetOrigin()
	if !ok {
		panic("arm Origin is not set")
	}

	return val
}

// GetOrigin retrieves the Origin value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Authority) GetOrigin() (result ID, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Origin" {
		result = *u.Origin
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u Authority) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *Authority) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Authority)(nil)
	_ encoding.BinaryUnmarshaler = (*Authority)(nil)
)

// MarshalJSON implements json.Marshaler.
func (u Authority) MarshalJSON() ([]byte, error) {
	temp := struct {
		Type int32       `json:"type"`
		Data interface{} `json:"data"`
	}{}

	temp.Type = int32(u.Type)
	temp.Data = ""
	switch u.Type {
	case AuthorityTypeUNKNOWN:
	case AuthorityTypeSELF:
	case AuthorityTypeAUTHORIZED:
		temp.Data = u.Origin
	default:
		return nil, fmt.Errorf("invalid union type")
	}

	return json.Marshal(temp)
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *Authority) UnmarshalJSON(data []byte) error {
	temp := struct {
		Type int32 `json:"type"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	u.Type = AuthorityType(temp.Type)
	switch u.Type {
	case AuthorityTypeUNKNOWN:

	case AuthorityTypeSELF:

	case AuthorityTypeAUTHORIZED:
		response := struct {
			Origin ID `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Origin = &response.Origin

	default:
		return fmt.Errorf("invalid union type")
	}

	return nil
}

// End union section

// Namespace end mazzaroth
var fmtTest = fmt.Sprint("this is a dummy usage of fmt")
