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
	FunctionName string       `json:"functionName"`

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
	name := FunctionTypeMap[int32(s)]
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
	name := RequestTypeMap[int32(s)]
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
	// ResponseTypeRECEIPT enum value 3
	ResponseTypeRECEIPT ResponseType = 3
	// ResponseTypeBLOCK enum value 4
	ResponseTypeBLOCK ResponseType = 4
	// ResponseTypeBLOCKLIST enum value 5
	ResponseTypeBLOCKLIST ResponseType = 5
	// ResponseTypeBLOCKHEADER enum value 6
	ResponseTypeBLOCKHEADER ResponseType = 6
	// ResponseTypeBLOCKHEADERLIST enum value 7
	ResponseTypeBLOCKHEADERLIST ResponseType = 7
	// ResponseTypeCONTRACT enum value 8
	ResponseTypeCONTRACT ResponseType = 8
	// ResponseTypeHEIGHT enum value 9
	ResponseTypeHEIGHT ResponseType = 9
	// ResponseTypeABI enum value 10
	ResponseTypeABI ResponseType = 10
)

// ResponseTypeMap generated enum map
var ResponseTypeMap = map[int32]string{
	0:  "ResponseTypeUNKNOWN",
	1:  "ResponseTypeTRANSACTIONID",
	2:  "ResponseTypeTRANSACTION",
	3:  "ResponseTypeRECEIPT",
	4:  "ResponseTypeBLOCK",
	5:  "ResponseTypeBLOCKLIST",
	6:  "ResponseTypeBLOCKHEADER",
	7:  "ResponseTypeBLOCKHEADERLIST",
	8:  "ResponseTypeCONTRACT",
	9:  "ResponseTypeHEIGHT",
	10: "ResponseTypeABI",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ResponseType
func (s ResponseType) ValidEnum(v int32) bool {
	_, ok := ResponseTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s ResponseType) String() string {
	name := ResponseTypeMap[int32(s)]
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
	Type        RequestType
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

// NewRequest creates a new Request.
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
	Type          ResponseType
	TransactionID *ID

	Transaction *Transaction

	Receipt *Receipt

	Block *Block

	Blocks *[]Block

	BlockHeader *BlockHeader

	BlockHeaders *[]BlockHeader

	Contract *Contract

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
	case ResponseTypeRECEIPT:
		return "Receipt", true
	case ResponseTypeBLOCK:
		return "Block", true
	case ResponseTypeBLOCKLIST:
		return "Blocks", true
	case ResponseTypeBLOCKHEADER:
		return "BlockHeader", true
	case ResponseTypeBLOCKHEADERLIST:
		return "BlockHeaders", true
	case ResponseTypeCONTRACT:
		return "Contract", true
	case ResponseTypeHEIGHT:
		return "Height", true
	case ResponseTypeABI:
		return "Abi", true
	}
	return "-", false
}

// NewResponse creates a new Response.
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
	case ResponseTypeRECEIPT:
		tv, ok := value.(Receipt)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Receipt = &tv
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
	case ResponseTypeCONTRACT:
		tv, ok := value.(Contract)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Contract = &tv
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

// MustContract retrieves the Contract value from the union,
// panicing if the value is not set.
func (u Response) MustContract() Contract {

	val, ok := u.GetContract()
	if !ok {
		panic("arm Contract is not set")
	}

	return val
}

// GetContract retrieves the Contract value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Response) GetContract() (result Contract, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Contract" {
		result = *u.Contract
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
	case ResponseTypeRECEIPT:
		temp.Data = u.Receipt
	case ResponseTypeBLOCK:
		temp.Data = u.Block
	case ResponseTypeBLOCKLIST:
		temp.Data = u.Blocks
	case ResponseTypeBLOCKHEADER:
		temp.Data = u.BlockHeader
	case ResponseTypeBLOCKHEADERLIST:
		temp.Data = u.BlockHeaders
	case ResponseTypeCONTRACT:
		temp.Data = u.Contract
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
	case ResponseTypeRECEIPT:
		response := struct {
			Receipt Receipt `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Receipt = &response.Receipt
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
	case ResponseTypeCONTRACT:
		response := struct {
			Contract Contract `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Contract = &response.Contract
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
	TransactionsMerkleRoot  Hash   `json:"transactionsMerkleRoot"`
	TransactionsReceiptRoot Hash   `json:"transactionsReceiptRoot"`
	StateRoot               Hash   `json:"stateRoot"`
	PreviousHeader          Hash   `json:"previousHeader"`
	Status                  Status `json:"status"`
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

// Status generated enum
type Status int32

const (
	// StatusUNKNOWN enum value 0
	StatusUNKNOWN Status = 0
	// StatusSUCCESS enum value 1
	StatusSUCCESS Status = 1
	// StatusFAILURE enum value 2
	StatusFAILURE Status = 2
	// StatusPENDING enum value 3
	StatusPENDING Status = 3
	// StatusFINALIZED enum value 4
	StatusFINALIZED Status = 4
)

// StatusMap generated enum map
var StatusMap = map[int32]string{
	0: "StatusUNKNOWN",
	1: "StatusSUCCESS",
	2: "StatusFAILURE",
	3: "StatusPENDING",
	4: "StatusFINALIZED",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for Status
func (s Status) ValidEnum(v int32) bool {
	_, ok := StatusMap[v]
	return ok
}

// String returns the name of `e`
func (s Status) String() string {
	name := StatusMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Status) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Status) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Status)(nil)
	_ encoding.BinaryUnmarshaler = (*Status)(nil)
)

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
	TransactionID ID         `json:"transactionID"`
	Status        Status     `json:"status"`
	StateRoot     Hash       `json:"stateRoot"`
	Result        string     `json:"result"`
	StatusInfo    StatusInfo `json:"statusInfo"`
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
	Version      string `xdrmaxsize:"100" json:"version"`
	Owner        ID     `json:"owner"`
	Abi          Abi    `json:"abi"`
	ContractHash Hash   `json:"contractHash"`

	ContractBytes []byte `json:"contractBytes"`
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
	Account   ID   `json:"account"`
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

// Data generated struct
type Data struct {
	ChannelID ID `json:"channelID"`

	Nonce uint64 `json:"nonce,string"`

	BlockExpirationNumber uint64   `json:"blockExpirationNumber,string"`
	Category              Category `json:"category"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Data) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Data) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Data)(nil)
	_ encoding.BinaryUnmarshaler = (*Data)(nil)
)

// Transaction generated struct
type Transaction struct {
	Sender    ID        `json:"sender"`
	Signature Signature `json:"signature"`
	Data      Data      `json:"data"`
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

// CategoryType generated enum
type CategoryType int32

const (
	// CategoryTypeUNKNOWN enum value 0
	CategoryTypeUNKNOWN CategoryType = 0
	// CategoryTypeCALL enum value 1
	CategoryTypeCALL CategoryType = 1
	// CategoryTypeDEPLOY enum value 2
	CategoryTypeDEPLOY CategoryType = 2
	// CategoryTypePAUSE enum value 3
	CategoryTypePAUSE CategoryType = 3
	// CategoryTypeDELETE enum value 4
	CategoryTypeDELETE CategoryType = 4
)

// CategoryTypeMap generated enum map
var CategoryTypeMap = map[int32]string{
	0: "CategoryTypeUNKNOWN",
	1: "CategoryTypeCALL",
	2: "CategoryTypeDEPLOY",
	3: "CategoryTypePAUSE",
	4: "CategoryTypeDELETE",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CategoryType
func (s CategoryType) ValidEnum(v int32) bool {
	_, ok := CategoryTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s CategoryType) String() string {
	name := CategoryTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s CategoryType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *CategoryType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*CategoryType)(nil)
	_ encoding.BinaryUnmarshaler = (*CategoryType)(nil)
)

// End enum section

// Start union section

// Category generated union
type Category struct {
	Type CategoryType
	Call *Call

	Contract *Contract

	Pause *bool
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u Category) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of Category
func (u Category) ArmForSwitch(sw int32) (string, bool) {
	switch CategoryType(sw) {
	case CategoryTypeUNKNOWN:
		return "", true
	case CategoryTypeCALL:
		return "Call", true
	case CategoryTypeDEPLOY:
		return "Contract", true
	case CategoryTypePAUSE:
		return "Pause", true
	case CategoryTypeDELETE:
		return "", true
	}
	return "-", false
}

// NewCategory creates a new Category.
func NewCategory(aType CategoryType, value interface{}) (result Category, err error) {
	result.Type = aType
	switch aType {
	case CategoryTypeUNKNOWN:
	case CategoryTypeCALL:
		tv, ok := value.(Call)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Call = &tv
	case CategoryTypeDEPLOY:
		tv, ok := value.(Contract)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Contract = &tv
	case CategoryTypePAUSE:
		tv, ok := value.(bool)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Pause = &tv
	case CategoryTypeDELETE:
	}
	return
}

// MustCall retrieves the Call value from the union,
// panicing if the value is not set.
func (u Category) MustCall() Call {

	val, ok := u.GetCall()
	if !ok {
		panic("arm Call is not set")
	}

	return val
}

// GetCall retrieves the Call value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Category) GetCall() (result Call, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Call" {
		result = *u.Call
		ok = true
	}

	return
}

// MustContract retrieves the Contract value from the union,
// panicing if the value is not set.
func (u Category) MustContract() Contract {

	val, ok := u.GetContract()
	if !ok {
		panic("arm Contract is not set")
	}

	return val
}

// GetContract retrieves the Contract value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Category) GetContract() (result Contract, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Contract" {
		result = *u.Contract
		ok = true
	}

	return
}

// MustPause retrieves the Pause value from the union,
// panicing if the value is not set.
func (u Category) MustPause() bool {

	val, ok := u.GetPause()
	if !ok {
		panic("arm Pause is not set")
	}

	return val
}

// GetPause retrieves the Pause value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Category) GetPause() (result bool, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Pause" {
		result = *u.Pause
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u Category) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *Category) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Category)(nil)
	_ encoding.BinaryUnmarshaler = (*Category)(nil)
)

// MarshalJSON implements json.Marshaler.
func (u Category) MarshalJSON() ([]byte, error) {
	temp := struct {
		Type int32       `json:"type"`
		Data interface{} `json:"data"`
	}{}

	temp.Type = int32(u.Type)
	temp.Data = ""
	switch u.Type {
	case CategoryTypeUNKNOWN:
	case CategoryTypeCALL:
		temp.Data = u.Call
	case CategoryTypeDEPLOY:
		temp.Data = u.Contract
	case CategoryTypePAUSE:
		temp.Data = u.Pause
	case CategoryTypeDELETE:
	default:
		return nil, fmt.Errorf("invalid union type")
	}

	return json.Marshal(temp)
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *Category) UnmarshalJSON(data []byte) error {
	temp := struct {
		Type int32 `json:"type"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	u.Type = CategoryType(temp.Type)
	switch u.Type {
	case CategoryTypeUNKNOWN:
	case CategoryTypeCALL:
		response := struct {
			Call Call `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Call = &response.Call
	case CategoryTypeDEPLOY:
		response := struct {
			Contract Contract `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Contract = &response.Contract
	case CategoryTypePAUSE:
		response := struct {
			Pause bool `json:"data"`
		}{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}
		u.Pause = &response.Pause
	case CategoryTypeDELETE:
	default:
		return fmt.Errorf("invalid union type")
	}

	return nil
}

// End union section

// Namespace end mazzaroth
var fmtTest = fmt.Sprint("this is a dummy usage of fmt")
