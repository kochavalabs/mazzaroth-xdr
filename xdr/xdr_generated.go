// Package xdr is automatically generated
// DO NOT EDIT or your changes may be overwritten
package xdr

import (
	"bytes"
	"encoding"
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
	FunctionType string `json:"functionType"`

	Name string `json:"name"`

	Inputs []Parameter `json:"inputs"`

	Outputs []Parameter `json:"outputs"`
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
	Name string `json:"name"`

	ParameterType string `json:"parameterType"`

	Codec string `json:"codec"`
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

// End enum section

// Start union section

// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

// Account generated struct
type Account struct {
	PermissionedKeys []ID `json:"permissionedKeys"`
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

// End struct section

// Start enum section

// End enum section

// Start union section

// End union section

// Namspace end mazzaroth
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
	Timestamp string `xdrmaxsize:"256" json:"timestamp"`

	BlockHeight uint64 `json:"blockHeight"`

	TransactionHeight uint64 `json:"transactionHeight"`

	ConsensusSequenceNumber uint64 `json:"consensusSequenceNumber"`

	TxMerkleRoot Hash `json:"txMerkleRoot"`

	TxReceiptRoot Hash `json:"txReceiptRoot"`

	StateRoot Hash `json:"stateRoot"`

	PreviousHeader Hash `json:"previousHeader"`

	BlockProducerAddress ID `json:"blockProducerAddress"`
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

// End struct section

// Start enum section

// End enum section

// Start union section

// End union section

// Namspace end mazzaroth
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

// Hash32 generated typedef
type Hash32 [32]byte

// XDRMaxSize implements the Sized interface for Hash32
func (s Hash32) XDRMaxSize() int {
	return 32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Hash32) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Hash32) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Hash32)(nil)
	_ encoding.BinaryUnmarshaler = (*Hash32)(nil)
)

// Hash64 generated typedef
type Hash64 [64]byte

// XDRMaxSize implements the Sized interface for Hash64
func (s Hash64) XDRMaxSize() int {
	return 64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Hash64) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Hash64) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Hash64)(nil)
	_ encoding.BinaryUnmarshaler = (*Hash64)(nil)
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

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

// ChannelConfig generated struct
type ChannelConfig struct {
	Owner ID `json:"owner"`

	ChannelName string `xdrmaxsize:"200" json:"channelName"`

	Admins []ID `xdrmaxsize:"200" json:"admins"`
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

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

// DownloadRequest generated struct
type DownloadRequest struct {
	DownloadRequestPayload DownloadRequestPayload `json:"downloadRequestPayload"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s DownloadRequest) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *DownloadRequest) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*DownloadRequest)(nil)
	_ encoding.BinaryUnmarshaler = (*DownloadRequest)(nil)
)

// BatchesRequest generated struct
type BatchesRequest struct {
	SeqNum uint64 `json:"seqNum"`

	Id string `json:"id"`

	Ip string `json:"ip"`

	Port uint64 `json:"port"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BatchesRequest) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BatchesRequest) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BatchesRequest)(nil)
	_ encoding.BinaryUnmarshaler = (*BatchesRequest)(nil)
)

// DownloadResponse generated struct
type DownloadResponse struct {
	DownloadStatus DownloadStatus `json:"downloadStatus"`

	DownloadResponsePayload DownloadResponsePayload `json:"downloadResponsePayload"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s DownloadResponse) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *DownloadResponse) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*DownloadResponse)(nil)
	_ encoding.BinaryUnmarshaler = (*DownloadResponse)(nil)
)

// DownloadHeight generated struct
type DownloadHeight struct {
	Height uint64 `json:"height"`

	SeqNum uint64 `json:"seqNum"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s DownloadHeight) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *DownloadHeight) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*DownloadHeight)(nil)
	_ encoding.BinaryUnmarshaler = (*DownloadHeight)(nil)
)

// End struct section

// Start enum section

// DownloadRequestType generated enum
type DownloadRequestType int32

const (

	// DownloadRequestTypeUNKNOWN enum value 0
	DownloadRequestTypeUNKNOWN DownloadRequestType = 0

	// DownloadRequestTypeHEIGHT enum value 1
	DownloadRequestTypeHEIGHT DownloadRequestType = 1

	// DownloadRequestTypeBLOCK enum value 2
	DownloadRequestTypeBLOCK DownloadRequestType = 2

	// DownloadRequestTypeBATCHES enum value 3
	DownloadRequestTypeBATCHES DownloadRequestType = 3
)

// DownloadRequestTypeMap generated enum map
var DownloadRequestTypeMap = map[int32]string{

	0: "DownloadRequestTypeUNKNOWN",

	1: "DownloadRequestTypeHEIGHT",

	2: "DownloadRequestTypeBLOCK",

	3: "DownloadRequestTypeBATCHES",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for DownloadRequestType
func (s DownloadRequestType) ValidEnum(v int32) bool {
	_, ok := DownloadRequestTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s DownloadRequestType) String() string {
	name, _ := DownloadRequestTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s DownloadRequestType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *DownloadRequestType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*DownloadRequestType)(nil)
	_ encoding.BinaryUnmarshaler = (*DownloadRequestType)(nil)
)

// DownloadStatus generated enum
type DownloadStatus int32

const (

	// DownloadStatusUNKNOWN enum value 0
	DownloadStatusUNKNOWN DownloadStatus = 0

	// DownloadStatusSUCCESS enum value 1
	DownloadStatusSUCCESS DownloadStatus = 1

	// DownloadStatusFAILURE enum value 2
	DownloadStatusFAILURE DownloadStatus = 2
)

// DownloadStatusMap generated enum map
var DownloadStatusMap = map[int32]string{

	0: "DownloadStatusUNKNOWN",

	1: "DownloadStatusSUCCESS",

	2: "DownloadStatusFAILURE",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for DownloadStatus
func (s DownloadStatus) ValidEnum(v int32) bool {
	_, ok := DownloadStatusMap[v]
	return ok
}

// String returns the name of `e`
func (s DownloadStatus) String() string {
	name, _ := DownloadStatusMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s DownloadStatus) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *DownloadStatus) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*DownloadStatus)(nil)
	_ encoding.BinaryUnmarshaler = (*DownloadStatus)(nil)
)

// End enum section

// Start union section

// DownloadRequestPayload generated union
type DownloadRequestPayload struct {
	Type DownloadRequestType

	BlockNumber *uint64

	BatchesRequest *BatchesRequest
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u DownloadRequestPayload) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of DownloadRequestPayload
func (u DownloadRequestPayload) ArmForSwitch(sw int32) (string, bool) {
	switch DownloadRequestType(sw) {

	case DownloadRequestTypeUNKNOWN:
		return "", true

	case DownloadRequestTypeHEIGHT:
		return "", true

	case DownloadRequestTypeBLOCK:
		return "BlockNumber", true

	case DownloadRequestTypeBATCHES:
		return "BatchesRequest", true
	}
	return "-", false
}

// NewDownloadRequestPayload creates a new  DownloadRequestPayload.
func NewDownloadRequestPayload(aType DownloadRequestType, value interface{}) (result DownloadRequestPayload, err error) {
	result.Type = aType
	switch aType {

	case DownloadRequestTypeUNKNOWN:

	case DownloadRequestTypeHEIGHT:

	case DownloadRequestTypeBLOCK:

		tv, ok := value.(uint64)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.BlockNumber = &tv

	case DownloadRequestTypeBATCHES:

		tv, ok := value.(BatchesRequest)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.BatchesRequest = &tv

	}
	return
}

// MustBlockNumber retrieves the BlockNumber value from the union,
// panicing if the value is not set.
func (u DownloadRequestPayload) MustBlockNumber() uint64 {

	val, ok := u.GetBlockNumber()
	if !ok {
		panic("arm BlockNumber is not set")
	}

	return val
}

// GetBlockNumber retrieves the BlockNumber value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u DownloadRequestPayload) GetBlockNumber() (result uint64, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "BlockNumber" {
		result = *u.BlockNumber
		ok = true
	}

	return
}

// MustBatchesRequest retrieves the BatchesRequest value from the union,
// panicing if the value is not set.
func (u DownloadRequestPayload) MustBatchesRequest() BatchesRequest {

	val, ok := u.GetBatchesRequest()
	if !ok {
		panic("arm BatchesRequest is not set")
	}

	return val
}

// GetBatchesRequest retrieves the BatchesRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u DownloadRequestPayload) GetBatchesRequest() (result BatchesRequest, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "BatchesRequest" {
		result = *u.BatchesRequest
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u DownloadRequestPayload) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *DownloadRequestPayload) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*DownloadRequestPayload)(nil)
	_ encoding.BinaryUnmarshaler = (*DownloadRequestPayload)(nil)
)

// DownloadResponsePayload generated union
type DownloadResponsePayload struct {
	Type DownloadRequestType

	DownloadHeight *DownloadHeight

	Block *Block
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u DownloadResponsePayload) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of DownloadResponsePayload
func (u DownloadResponsePayload) ArmForSwitch(sw int32) (string, bool) {
	switch DownloadRequestType(sw) {

	case DownloadRequestTypeUNKNOWN:
		return "", true

	case DownloadRequestTypeHEIGHT:
		return "DownloadHeight", true

	case DownloadRequestTypeBLOCK:
		return "Block", true

	case DownloadRequestTypeBATCHES:
		return "", true
	}
	return "-", false
}

// NewDownloadResponsePayload creates a new  DownloadResponsePayload.
func NewDownloadResponsePayload(aType DownloadRequestType, value interface{}) (result DownloadResponsePayload, err error) {
	result.Type = aType
	switch aType {

	case DownloadRequestTypeUNKNOWN:

	case DownloadRequestTypeHEIGHT:

		tv, ok := value.(DownloadHeight)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.DownloadHeight = &tv

	case DownloadRequestTypeBLOCK:

		tv, ok := value.(Block)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Block = &tv

	case DownloadRequestTypeBATCHES:

	}
	return
}

// MustDownloadHeight retrieves the DownloadHeight value from the union,
// panicing if the value is not set.
func (u DownloadResponsePayload) MustDownloadHeight() DownloadHeight {

	val, ok := u.GetDownloadHeight()
	if !ok {
		panic("arm DownloadHeight is not set")
	}

	return val
}

// GetDownloadHeight retrieves the DownloadHeight value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u DownloadResponsePayload) GetDownloadHeight() (result DownloadHeight, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "DownloadHeight" {
		result = *u.DownloadHeight
		ok = true
	}

	return
}

// MustBlock retrieves the Block value from the union,
// panicing if the value is not set.
func (u DownloadResponsePayload) MustBlock() Block {

	val, ok := u.GetBlock()
	if !ok {
		panic("arm Block is not set")
	}

	return val
}

// GetBlock retrieves the Block value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u DownloadResponsePayload) GetBlock() (result Block, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Block" {
		result = *u.Block
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u DownloadResponsePayload) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *DownloadResponsePayload) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*DownloadResponsePayload)(nil)
	_ encoding.BinaryUnmarshaler = (*DownloadResponsePayload)(nil)
)

// End union section

// Namspace end mazzaroth
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

	// ReceiptStatusFAILURE enum value 0
	ReceiptStatusFAILURE ReceiptStatus = 0

	// ReceiptStatusSUCCESS enum value 1
	ReceiptStatusSUCCESS ReceiptStatus = 1
)

// ReceiptStatusMap generated enum map
var ReceiptStatusMap = map[int32]string{

	0: "ReceiptStatusFAILURE",

	1: "ReceiptStatusSUCCESS",
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

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

// StateStatus generated struct
type StateStatus struct {
	PreviousBlock uint64 `json:"previousBlock"`

	TransactionCount uint64 `json:"transactionCount"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s StateStatus) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *StateStatus) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*StateStatus)(nil)
	_ encoding.BinaryUnmarshaler = (*StateStatus)(nil)
)

// BlockLookupRequest generated struct
type BlockLookupRequest struct {
	Identifier Identifier `json:"identifier"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BlockLookupRequest) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BlockLookupRequest) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BlockLookupRequest)(nil)
	_ encoding.BinaryUnmarshaler = (*BlockLookupRequest)(nil)
)

// BlockHeaderLookupRequest generated struct
type BlockHeaderLookupRequest struct {
	Identifier Identifier `json:"identifier"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BlockHeaderLookupRequest) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BlockHeaderLookupRequest) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BlockHeaderLookupRequest)(nil)
	_ encoding.BinaryUnmarshaler = (*BlockHeaderLookupRequest)(nil)
)

// BlockLookupResponse generated struct
type BlockLookupResponse struct {
	Block Block `json:"block"`

	StateStatus StateStatus `json:"stateStatus"`

	Status BlockStatus `json:"status"`

	StatusInfo StatusInfo `json:"statusInfo"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BlockLookupResponse) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BlockLookupResponse) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BlockLookupResponse)(nil)
	_ encoding.BinaryUnmarshaler = (*BlockLookupResponse)(nil)
)

// BlockHeaderLookupResponse generated struct
type BlockHeaderLookupResponse struct {
	Header BlockHeader `json:"header"`

	StateStatus StateStatus `json:"stateStatus"`

	Status BlockStatus `json:"status"`

	StatusInfo StatusInfo `json:"statusInfo"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BlockHeaderLookupResponse) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BlockHeaderLookupResponse) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BlockHeaderLookupResponse)(nil)
	_ encoding.BinaryUnmarshaler = (*BlockHeaderLookupResponse)(nil)
)

// TransactionLookupRequest generated struct
type TransactionLookupRequest struct {
	TransactionID ID `json:"transactionID"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionLookupRequest) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionLookupRequest) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionLookupRequest)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionLookupRequest)(nil)
)

// TransactionLookupResponse generated struct
type TransactionLookupResponse struct {
	Transaction Transaction `json:"transaction"`

	StateStatus StateStatus `json:"stateStatus"`

	Status TransactionStatus `json:"status"`

	StatusInfo StatusInfo `json:"statusInfo"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionLookupResponse) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionLookupResponse) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionLookupResponse)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionLookupResponse)(nil)
)

// TransactionSubmitRequest generated struct
type TransactionSubmitRequest struct {
	Transaction Transaction `json:"transaction"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionSubmitRequest) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionSubmitRequest) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionSubmitRequest)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionSubmitRequest)(nil)
)

// TransactionSubmitResponse generated struct
type TransactionSubmitResponse struct {
	TransactionInfo TransactionInfo `json:"transactionInfo"`

	Status TransactionStatus `json:"status"`

	StatusInfo StatusInfo `json:"statusInfo"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionSubmitResponse) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionSubmitResponse) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionSubmitResponse)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionSubmitResponse)(nil)
)

// ReceiptLookupRequest generated struct
type ReceiptLookupRequest struct {
	TransactionID ID `json:"transactionID"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ReceiptLookupRequest) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ReceiptLookupRequest) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ReceiptLookupRequest)(nil)
	_ encoding.BinaryUnmarshaler = (*ReceiptLookupRequest)(nil)
)

// ReceiptLookupResponse generated struct
type ReceiptLookupResponse struct {
	Receipt Receipt `json:"receipt"`

	StateStatus StateStatus `json:"stateStatus"`

	Status ReceiptLookupStatus `json:"status"`

	StatusInfo StatusInfo `json:"statusInfo"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ReceiptLookupResponse) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ReceiptLookupResponse) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ReceiptLookupResponse)(nil)
	_ encoding.BinaryUnmarshaler = (*ReceiptLookupResponse)(nil)
)

// AccountInfoLookupRequest generated struct
type AccountInfoLookupRequest struct {
	Account ID `json:"account"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountInfoLookupRequest) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountInfoLookupRequest) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountInfoLookupRequest)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountInfoLookupRequest)(nil)
)

// AccountInfoLookupResponse generated struct
type AccountInfoLookupResponse struct {
	AccountInfo Account `json:"accountInfo"`

	StateStatus StateStatus `json:"stateStatus"`

	Status InfoLookupStatus `json:"status"`

	StatusInfo StatusInfo `json:"statusInfo"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountInfoLookupResponse) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountInfoLookupResponse) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountInfoLookupResponse)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountInfoLookupResponse)(nil)
)

// ChannelInfoLookupRequest generated struct
type ChannelInfoLookupRequest struct {
	InfoType ChannelInfoType `json:"infoType"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ChannelInfoLookupRequest) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ChannelInfoLookupRequest) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ChannelInfoLookupRequest)(nil)
	_ encoding.BinaryUnmarshaler = (*ChannelInfoLookupRequest)(nil)
)

// ChannelInfoLookupResponse generated struct
type ChannelInfoLookupResponse struct {
	ChannelInfo ChannelInfo `json:"channelInfo"`

	StateStatus StateStatus `json:"stateStatus"`

	Status InfoLookupStatus `json:"status"`

	StatusInfo StatusInfo `json:"statusInfo"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ChannelInfoLookupResponse) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ChannelInfoLookupResponse) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ChannelInfoLookupResponse)(nil)
	_ encoding.BinaryUnmarshaler = (*ChannelInfoLookupResponse)(nil)
)

// End struct section

// Start enum section

// IdentifierType generated enum
type IdentifierType int32

const (

	// IdentifierTypeNONE enum value 0
	IdentifierTypeNONE IdentifierType = 0

	// IdentifierTypeNUMBER enum value 1
	IdentifierTypeNUMBER IdentifierType = 1

	// IdentifierTypeHASH enum value 2
	IdentifierTypeHASH IdentifierType = 2
)

// IdentifierTypeMap generated enum map
var IdentifierTypeMap = map[int32]string{

	0: "IdentifierTypeNONE",

	1: "IdentifierTypeNUMBER",

	2: "IdentifierTypeHASH",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for IdentifierType
func (s IdentifierType) ValidEnum(v int32) bool {
	_, ok := IdentifierTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s IdentifierType) String() string {
	name, _ := IdentifierTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s IdentifierType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *IdentifierType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*IdentifierType)(nil)
	_ encoding.BinaryUnmarshaler = (*IdentifierType)(nil)
)

// BlockStatus generated enum
type BlockStatus int32

const (

	// BlockStatusUNKNOWN enum value 0
	BlockStatusUNKNOWN BlockStatus = 0

	// BlockStatusCREATED enum value 1
	BlockStatusCREATED BlockStatus = 1

	// BlockStatusFUTURE enum value 2
	BlockStatusFUTURE BlockStatus = 2

	// BlockStatusNOT_FOUND enum value 3
	BlockStatusNOT_FOUND BlockStatus = 3
)

// BlockStatusMap generated enum map
var BlockStatusMap = map[int32]string{

	0: "BlockStatusUNKNOWN",

	1: "BlockStatusCREATED",

	2: "BlockStatusFUTURE",

	3: "BlockStatusNOT_FOUND",
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

// TransactionType generated enum
type TransactionType int32

const (

	// TransactionTypeNONE enum value 0
	TransactionTypeNONE TransactionType = 0

	// TransactionTypeREAD enum value 1
	TransactionTypeREAD TransactionType = 1

	// TransactionTypeWRITE enum value 2
	TransactionTypeWRITE TransactionType = 2
)

// TransactionTypeMap generated enum map
var TransactionTypeMap = map[int32]string{

	0: "TransactionTypeNONE",

	1: "TransactionTypeREAD",

	2: "TransactionTypeWRITE",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for TransactionType
func (s TransactionType) ValidEnum(v int32) bool {
	_, ok := TransactionTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s TransactionType) String() string {
	name, _ := TransactionTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionType)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionType)(nil)
)

// TransactionStatus generated enum
type TransactionStatus int32

const (

	// TransactionStatusUNKNOWN enum value 0
	TransactionStatusUNKNOWN TransactionStatus = 0

	// TransactionStatusACCEPTED enum value 1
	TransactionStatusACCEPTED TransactionStatus = 1

	// TransactionStatusREJECTED enum value 2
	TransactionStatusREJECTED TransactionStatus = 2

	// TransactionStatusCONFIRMED enum value 3
	TransactionStatusCONFIRMED TransactionStatus = 3

	// TransactionStatusNOT_FOUND enum value 4
	TransactionStatusNOT_FOUND TransactionStatus = 4
)

// TransactionStatusMap generated enum map
var TransactionStatusMap = map[int32]string{

	0: "TransactionStatusUNKNOWN",

	1: "TransactionStatusACCEPTED",

	2: "TransactionStatusREJECTED",

	3: "TransactionStatusCONFIRMED",

	4: "TransactionStatusNOT_FOUND",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for TransactionStatus
func (s TransactionStatus) ValidEnum(v int32) bool {
	_, ok := TransactionStatusMap[v]
	return ok
}

// String returns the name of `e`
func (s TransactionStatus) String() string {
	name, _ := TransactionStatusMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionStatus) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionStatus) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionStatus)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionStatus)(nil)
)

// ReceiptLookupStatus generated enum
type ReceiptLookupStatus int32

const (

	// ReceiptLookupStatusUNKNOWN enum value 0
	ReceiptLookupStatusUNKNOWN ReceiptLookupStatus = 0

	// ReceiptLookupStatusFOUND enum value 1
	ReceiptLookupStatusFOUND ReceiptLookupStatus = 1

	// ReceiptLookupStatusNOT_FOUND enum value 2
	ReceiptLookupStatusNOT_FOUND ReceiptLookupStatus = 2
)

// ReceiptLookupStatusMap generated enum map
var ReceiptLookupStatusMap = map[int32]string{

	0: "ReceiptLookupStatusUNKNOWN",

	1: "ReceiptLookupStatusFOUND",

	2: "ReceiptLookupStatusNOT_FOUND",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ReceiptLookupStatus
func (s ReceiptLookupStatus) ValidEnum(v int32) bool {
	_, ok := ReceiptLookupStatusMap[v]
	return ok
}

// String returns the name of `e`
func (s ReceiptLookupStatus) String() string {
	name, _ := ReceiptLookupStatusMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ReceiptLookupStatus) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ReceiptLookupStatus) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ReceiptLookupStatus)(nil)
	_ encoding.BinaryUnmarshaler = (*ReceiptLookupStatus)(nil)
)

// ChannelInfoType generated enum
type ChannelInfoType int32

const (

	// ChannelInfoTypeNONE enum value 0
	ChannelInfoTypeNONE ChannelInfoType = 0

	// ChannelInfoTypeCONTRACT enum value 1
	ChannelInfoTypeCONTRACT ChannelInfoType = 1

	// ChannelInfoTypeCONFIG enum value 2
	ChannelInfoTypeCONFIG ChannelInfoType = 2
)

// ChannelInfoTypeMap generated enum map
var ChannelInfoTypeMap = map[int32]string{

	0: "ChannelInfoTypeNONE",

	1: "ChannelInfoTypeCONTRACT",

	2: "ChannelInfoTypeCONFIG",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ChannelInfoType
func (s ChannelInfoType) ValidEnum(v int32) bool {
	_, ok := ChannelInfoTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s ChannelInfoType) String() string {
	name, _ := ChannelInfoTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ChannelInfoType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ChannelInfoType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ChannelInfoType)(nil)
	_ encoding.BinaryUnmarshaler = (*ChannelInfoType)(nil)
)

// InfoLookupStatus generated enum
type InfoLookupStatus int32

const (

	// InfoLookupStatusUNKNOWN enum value 0
	InfoLookupStatusUNKNOWN InfoLookupStatus = 0

	// InfoLookupStatusFOUND enum value 1
	InfoLookupStatusFOUND InfoLookupStatus = 1

	// InfoLookupStatusNOT_FOUND enum value 2
	InfoLookupStatusNOT_FOUND InfoLookupStatus = 2
)

// InfoLookupStatusMap generated enum map
var InfoLookupStatusMap = map[int32]string{

	0: "InfoLookupStatusUNKNOWN",

	1: "InfoLookupStatusFOUND",

	2: "InfoLookupStatusNOT_FOUND",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for InfoLookupStatus
func (s InfoLookupStatus) ValidEnum(v int32) bool {
	_, ok := InfoLookupStatusMap[v]
	return ok
}

// String returns the name of `e`
func (s InfoLookupStatus) String() string {
	name, _ := InfoLookupStatusMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s InfoLookupStatus) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *InfoLookupStatus) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*InfoLookupStatus)(nil)
	_ encoding.BinaryUnmarshaler = (*InfoLookupStatus)(nil)
)

// End enum section

// Start union section

// Identifier generated union
type Identifier struct {
	Type IdentifierType

	Number *uint64

	Hash *Hash
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u Identifier) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of Identifier
func (u Identifier) ArmForSwitch(sw int32) (string, bool) {
	switch IdentifierType(sw) {

	case IdentifierTypeNONE:
		return "", true

	case IdentifierTypeNUMBER:
		return "Number", true

	case IdentifierTypeHASH:
		return "Hash", true
	}
	return "-", false
}

// NewIdentifier creates a new  Identifier.
func NewIdentifier(aType IdentifierType, value interface{}) (result Identifier, err error) {
	result.Type = aType
	switch aType {

	case IdentifierTypeNONE:

	case IdentifierTypeNUMBER:

		tv, ok := value.(uint64)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Number = &tv

	case IdentifierTypeHASH:

		tv, ok := value.(Hash)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Hash = &tv

	}
	return
}

// MustNumber retrieves the Number value from the union,
// panicing if the value is not set.
func (u Identifier) MustNumber() uint64 {

	val, ok := u.GetNumber()
	if !ok {
		panic("arm Number is not set")
	}

	return val
}

// GetNumber retrieves the Number value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Identifier) GetNumber() (result uint64, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Number" {
		result = *u.Number
		ok = true
	}

	return
}

// MustHash retrieves the Hash value from the union,
// panicing if the value is not set.
func (u Identifier) MustHash() Hash {

	val, ok := u.GetHash()
	if !ok {
		panic("arm Hash is not set")
	}

	return val
}

// GetHash retrieves the Hash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Identifier) GetHash() (result Hash, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Hash" {
		result = *u.Hash
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u Identifier) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *Identifier) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Identifier)(nil)
	_ encoding.BinaryUnmarshaler = (*Identifier)(nil)
)

// TransactionInfo generated union
type TransactionInfo struct {
	Type TransactionType

	Receipt *Receipt

	TransactionID *ID
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionInfo) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionInfo
func (u TransactionInfo) ArmForSwitch(sw int32) (string, bool) {
	switch TransactionType(sw) {

	case TransactionTypeNONE:
		return "", true

	case TransactionTypeREAD:
		return "Receipt", true

	case TransactionTypeWRITE:
		return "TransactionID", true
	}
	return "-", false
}

// NewTransactionInfo creates a new  TransactionInfo.
func NewTransactionInfo(aType TransactionType, value interface{}) (result TransactionInfo, err error) {
	result.Type = aType
	switch aType {

	case TransactionTypeNONE:

	case TransactionTypeREAD:

		tv, ok := value.(Receipt)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Receipt = &tv

	case TransactionTypeWRITE:

		tv, ok := value.(ID)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.TransactionID = &tv

	}
	return
}

// MustReceipt retrieves the Receipt value from the union,
// panicing if the value is not set.
func (u TransactionInfo) MustReceipt() Receipt {

	val, ok := u.GetReceipt()
	if !ok {
		panic("arm Receipt is not set")
	}

	return val
}

// GetReceipt retrieves the Receipt value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionInfo) GetReceipt() (result Receipt, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Receipt" {
		result = *u.Receipt
		ok = true
	}

	return
}

// MustTransactionID retrieves the TransactionID value from the union,
// panicing if the value is not set.
func (u TransactionInfo) MustTransactionID() ID {

	val, ok := u.GetTransactionID()
	if !ok {
		panic("arm TransactionID is not set")
	}

	return val
}

// GetTransactionID retrieves the TransactionID value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionInfo) GetTransactionID() (result ID, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "TransactionID" {
		result = *u.TransactionID
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u TransactionInfo) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *TransactionInfo) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionInfo)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionInfo)(nil)
)

// ChannelInfo generated union
type ChannelInfo struct {
	Type ChannelInfoType

	Contract *Contract

	ChannelConfig *ChannelConfig
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ChannelInfo) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ChannelInfo
func (u ChannelInfo) ArmForSwitch(sw int32) (string, bool) {
	switch ChannelInfoType(sw) {

	case ChannelInfoTypeNONE:
		return "", true

	case ChannelInfoTypeCONTRACT:
		return "Contract", true

	case ChannelInfoTypeCONFIG:
		return "ChannelConfig", true
	}
	return "-", false
}

// NewChannelInfo creates a new  ChannelInfo.
func NewChannelInfo(aType ChannelInfoType, value interface{}) (result ChannelInfo, err error) {
	result.Type = aType
	switch aType {

	case ChannelInfoTypeNONE:

	case ChannelInfoTypeCONTRACT:

		tv, ok := value.(Contract)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Contract = &tv

	case ChannelInfoTypeCONFIG:

		tv, ok := value.(ChannelConfig)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.ChannelConfig = &tv

	}
	return
}

// MustContract retrieves the Contract value from the union,
// panicing if the value is not set.
func (u ChannelInfo) MustContract() Contract {

	val, ok := u.GetContract()
	if !ok {
		panic("arm Contract is not set")
	}

	return val
}

// GetContract retrieves the Contract value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ChannelInfo) GetContract() (result Contract, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Contract" {
		result = *u.Contract
		ok = true
	}

	return
}

// MustChannelConfig retrieves the ChannelConfig value from the union,
// panicing if the value is not set.
func (u ChannelInfo) MustChannelConfig() ChannelConfig {

	val, ok := u.GetChannelConfig()
	if !ok {
		panic("arm ChannelConfig is not set")
	}

	return val
}

// GetChannelConfig retrieves the ChannelConfig value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ChannelInfo) GetChannelConfig() (result ChannelConfig, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ChannelConfig" {
		result = *u.ChannelConfig
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u ChannelInfo) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *ChannelInfo) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ChannelInfo)(nil)
	_ encoding.BinaryUnmarshaler = (*ChannelInfo)(nil)
)

// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

// ReceiptSubscription generated struct
type ReceiptSubscription struct {
	TransactionFilter TransactionFilter `json:"transactionFilter"`

	ReceiptFilter ReceiptFilter `json:"receiptFilter"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ReceiptSubscription) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ReceiptSubscription) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ReceiptSubscription)(nil)
	_ encoding.BinaryUnmarshaler = (*ReceiptSubscription)(nil)
)

// ReceiptSubscriptionResult generated struct
type ReceiptSubscriptionResult struct {
	Receipt Receipt `json:"receipt"`

	TransactionID ID `json:"transactionID"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ReceiptSubscriptionResult) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ReceiptSubscriptionResult) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ReceiptSubscriptionResult)(nil)
	_ encoding.BinaryUnmarshaler = (*ReceiptSubscriptionResult)(nil)
)

// ReceiptValueFilter generated struct
type ReceiptValueFilter struct {
	Status ValueFilter `json:"status"`

	StateRoot ValueFilter `json:"stateRoot"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ReceiptValueFilter) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ReceiptValueFilter) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ReceiptValueFilter)(nil)
	_ encoding.BinaryUnmarshaler = (*ReceiptValueFilter)(nil)
)

// ActionFilter generated struct
type ActionFilter struct {
	Signature ValueFilter `json:"signature"`

	Signer ValueFilter `json:"signer"`

	Address ValueFilter `json:"address"`

	ChannelID ValueFilter `json:"channelID"`

	Nonce ValueFilter `json:"nonce"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ActionFilter) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ActionFilter) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ActionFilter)(nil)
	_ encoding.BinaryUnmarshaler = (*ActionFilter)(nil)
)

// ContractFilter generated struct
type ContractFilter struct {
	ActionFilter ActionFilter `json:"actionFilter"`

	Version ValueFilter `json:"version"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ContractFilter) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ContractFilter) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ContractFilter)(nil)
	_ encoding.BinaryUnmarshaler = (*ContractFilter)(nil)
)

// ConfigFilter generated struct
type ConfigFilter struct {
	ActionFilter ActionFilter `json:"actionFilter"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ConfigFilter) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ConfigFilter) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ConfigFilter)(nil)
	_ encoding.BinaryUnmarshaler = (*ConfigFilter)(nil)
)

// PermissionFilter generated struct
type PermissionFilter struct {
	ActionFilter ActionFilter `json:"actionFilter"`

	Key ValueFilter `json:"key"`

	Action ValueFilter `json:"action"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PermissionFilter) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PermissionFilter) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PermissionFilter)(nil)
	_ encoding.BinaryUnmarshaler = (*PermissionFilter)(nil)
)

// CallFilter generated struct
type CallFilter struct {
	ActionFilter ActionFilter `json:"actionFilter"`

	Function ValueFilter `json:"function"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s CallFilter) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *CallFilter) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*CallFilter)(nil)
	_ encoding.BinaryUnmarshaler = (*CallFilter)(nil)
)

// End struct section

// Start enum section

// ValueFilterType generated enum
type ValueFilterType int32

const (

	// ValueFilterTypeNONE enum value 0
	ValueFilterTypeNONE ValueFilterType = 0

	// ValueFilterTypeSTRING enum value 1
	ValueFilterTypeSTRING ValueFilterType = 1

	// ValueFilterTypeHASH32 enum value 2
	ValueFilterTypeHASH32 ValueFilterType = 2

	// ValueFilterTypeHASH64 enum value 3
	ValueFilterTypeHASH64 ValueFilterType = 3

	// ValueFilterTypeUHYPER enum value 4
	ValueFilterTypeUHYPER ValueFilterType = 4

	// ValueFilterTypeINT enum value 5
	ValueFilterTypeINT ValueFilterType = 5
)

// ValueFilterTypeMap generated enum map
var ValueFilterTypeMap = map[int32]string{

	0: "ValueFilterTypeNONE",

	1: "ValueFilterTypeSTRING",

	2: "ValueFilterTypeHASH32",

	3: "ValueFilterTypeHASH64",

	4: "ValueFilterTypeUHYPER",

	5: "ValueFilterTypeINT",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ValueFilterType
func (s ValueFilterType) ValidEnum(v int32) bool {
	_, ok := ValueFilterTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s ValueFilterType) String() string {
	name, _ := ValueFilterTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ValueFilterType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ValueFilterType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ValueFilterType)(nil)
	_ encoding.BinaryUnmarshaler = (*ValueFilterType)(nil)
)

// TransactionFilterType generated enum
type TransactionFilterType int32

const (

	// TransactionFilterTypeNONE enum value 0
	TransactionFilterTypeNONE TransactionFilterType = 0

	// TransactionFilterTypeGENERIC enum value 1
	TransactionFilterTypeGENERIC TransactionFilterType = 1

	// TransactionFilterTypeCONTRACT enum value 2
	TransactionFilterTypeCONTRACT TransactionFilterType = 2

	// TransactionFilterTypeCONFIG enum value 3
	TransactionFilterTypeCONFIG TransactionFilterType = 3

	// TransactionFilterTypePERMISSION enum value 4
	TransactionFilterTypePERMISSION TransactionFilterType = 4

	// TransactionFilterTypeCALL enum value 5
	TransactionFilterTypeCALL TransactionFilterType = 5
)

// TransactionFilterTypeMap generated enum map
var TransactionFilterTypeMap = map[int32]string{

	0: "TransactionFilterTypeNONE",

	1: "TransactionFilterTypeGENERIC",

	2: "TransactionFilterTypeCONTRACT",

	3: "TransactionFilterTypeCONFIG",

	4: "TransactionFilterTypePERMISSION",

	5: "TransactionFilterTypeCALL",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for TransactionFilterType
func (s TransactionFilterType) ValidEnum(v int32) bool {
	_, ok := TransactionFilterTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s TransactionFilterType) String() string {
	name, _ := TransactionFilterTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TransactionFilterType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TransactionFilterType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionFilterType)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionFilterType)(nil)
)

// ReceiptFilterType generated enum
type ReceiptFilterType int32

const (

	// ReceiptFilterTypeNONE enum value 0
	ReceiptFilterTypeNONE ReceiptFilterType = 0

	// ReceiptFilterTypeRECEIPT enum value 1
	ReceiptFilterTypeRECEIPT ReceiptFilterType = 1
)

// ReceiptFilterTypeMap generated enum map
var ReceiptFilterTypeMap = map[int32]string{

	0: "ReceiptFilterTypeNONE",

	1: "ReceiptFilterTypeRECEIPT",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ReceiptFilterType
func (s ReceiptFilterType) ValidEnum(v int32) bool {
	_, ok := ReceiptFilterTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s ReceiptFilterType) String() string {
	name, _ := ReceiptFilterTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ReceiptFilterType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ReceiptFilterType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ReceiptFilterType)(nil)
	_ encoding.BinaryUnmarshaler = (*ReceiptFilterType)(nil)
)

// End enum section

// Start union section

// ValueFilter generated union
type ValueFilter struct {
	Type ValueFilterType

	Regex *string

	Hash32Value *Hash32

	Hash64Value *Hash64

	UhyperValue *uint64

	IntValue *int32
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ValueFilter) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ValueFilter
func (u ValueFilter) ArmForSwitch(sw int32) (string, bool) {
	switch ValueFilterType(sw) {

	case ValueFilterTypeNONE:
		return "", true

	case ValueFilterTypeSTRING:
		return "Regex", true

	case ValueFilterTypeHASH32:
		return "Hash32Value", true

	case ValueFilterTypeHASH64:
		return "Hash64Value", true

	case ValueFilterTypeUHYPER:
		return "UhyperValue", true

	case ValueFilterTypeINT:
		return "IntValue", true
	}
	return "-", false
}

// NewValueFilter creates a new  ValueFilter.
func NewValueFilter(aType ValueFilterType, value interface{}) (result ValueFilter, err error) {
	result.Type = aType
	switch aType {

	case ValueFilterTypeNONE:

	case ValueFilterTypeSTRING:

		tv, ok := value.(string)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Regex = &tv

	case ValueFilterTypeHASH32:

		tv, ok := value.(Hash32)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Hash32Value = &tv

	case ValueFilterTypeHASH64:

		tv, ok := value.(Hash64)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Hash64Value = &tv

	case ValueFilterTypeUHYPER:

		tv, ok := value.(uint64)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.UhyperValue = &tv

	case ValueFilterTypeINT:

		tv, ok := value.(int32)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.IntValue = &tv

	}
	return
}

// MustRegex retrieves the Regex value from the union,
// panicing if the value is not set.
func (u ValueFilter) MustRegex() string {

	val, ok := u.GetRegex()
	if !ok {
		panic("arm Regex is not set")
	}

	return val
}

// GetRegex retrieves the Regex value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ValueFilter) GetRegex() (result string, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Regex" {
		result = *u.Regex
		ok = true
	}

	return
}

// MustHash32Value retrieves the Hash32Value value from the union,
// panicing if the value is not set.
func (u ValueFilter) MustHash32Value() Hash32 {

	val, ok := u.GetHash32Value()
	if !ok {
		panic("arm Hash32Value is not set")
	}

	return val
}

// GetHash32Value retrieves the Hash32Value value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ValueFilter) GetHash32Value() (result Hash32, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Hash32Value" {
		result = *u.Hash32Value
		ok = true
	}

	return
}

// MustHash64Value retrieves the Hash64Value value from the union,
// panicing if the value is not set.
func (u ValueFilter) MustHash64Value() Hash64 {

	val, ok := u.GetHash64Value()
	if !ok {
		panic("arm Hash64Value is not set")
	}

	return val
}

// GetHash64Value retrieves the Hash64Value value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ValueFilter) GetHash64Value() (result Hash64, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Hash64Value" {
		result = *u.Hash64Value
		ok = true
	}

	return
}

// MustUhyperValue retrieves the UhyperValue value from the union,
// panicing if the value is not set.
func (u ValueFilter) MustUhyperValue() uint64 {

	val, ok := u.GetUhyperValue()
	if !ok {
		panic("arm UhyperValue is not set")
	}

	return val
}

// GetUhyperValue retrieves the UhyperValue value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ValueFilter) GetUhyperValue() (result uint64, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UhyperValue" {
		result = *u.UhyperValue
		ok = true
	}

	return
}

// MustIntValue retrieves the IntValue value from the union,
// panicing if the value is not set.
func (u ValueFilter) MustIntValue() int32 {

	val, ok := u.GetIntValue()
	if !ok {
		panic("arm IntValue is not set")
	}

	return val
}

// GetIntValue retrieves the IntValue value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ValueFilter) GetIntValue() (result int32, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "IntValue" {
		result = *u.IntValue
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u ValueFilter) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *ValueFilter) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ValueFilter)(nil)
	_ encoding.BinaryUnmarshaler = (*ValueFilter)(nil)
)

// TransactionFilter generated union
type TransactionFilter struct {
	Type TransactionFilterType

	GenericFilter *ActionFilter

	ContractFilter *ContractFilter

	ConfigFilter *ConfigFilter

	PermissionFilter *PermissionFilter

	CallFilter *CallFilter
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionFilter) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionFilter
func (u TransactionFilter) ArmForSwitch(sw int32) (string, bool) {
	switch TransactionFilterType(sw) {

	case TransactionFilterTypeNONE:
		return "", true

	case TransactionFilterTypeGENERIC:
		return "GenericFilter", true

	case TransactionFilterTypeCONTRACT:
		return "ContractFilter", true

	case TransactionFilterTypeCONFIG:
		return "ConfigFilter", true

	case TransactionFilterTypePERMISSION:
		return "PermissionFilter", true

	case TransactionFilterTypeCALL:
		return "CallFilter", true
	}
	return "-", false
}

// NewTransactionFilter creates a new  TransactionFilter.
func NewTransactionFilter(aType TransactionFilterType, value interface{}) (result TransactionFilter, err error) {
	result.Type = aType
	switch aType {

	case TransactionFilterTypeNONE:

	case TransactionFilterTypeGENERIC:

		tv, ok := value.(ActionFilter)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.GenericFilter = &tv

	case TransactionFilterTypeCONTRACT:

		tv, ok := value.(ContractFilter)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.ContractFilter = &tv

	case TransactionFilterTypeCONFIG:

		tv, ok := value.(ConfigFilter)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.ConfigFilter = &tv

	case TransactionFilterTypePERMISSION:

		tv, ok := value.(PermissionFilter)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.PermissionFilter = &tv

	case TransactionFilterTypeCALL:

		tv, ok := value.(CallFilter)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.CallFilter = &tv

	}
	return
}

// MustGenericFilter retrieves the GenericFilter value from the union,
// panicing if the value is not set.
func (u TransactionFilter) MustGenericFilter() ActionFilter {

	val, ok := u.GetGenericFilter()
	if !ok {
		panic("arm GenericFilter is not set")
	}

	return val
}

// GetGenericFilter retrieves the GenericFilter value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionFilter) GetGenericFilter() (result ActionFilter, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "GenericFilter" {
		result = *u.GenericFilter
		ok = true
	}

	return
}

// MustContractFilter retrieves the ContractFilter value from the union,
// panicing if the value is not set.
func (u TransactionFilter) MustContractFilter() ContractFilter {

	val, ok := u.GetContractFilter()
	if !ok {
		panic("arm ContractFilter is not set")
	}

	return val
}

// GetContractFilter retrieves the ContractFilter value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionFilter) GetContractFilter() (result ContractFilter, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ContractFilter" {
		result = *u.ContractFilter
		ok = true
	}

	return
}

// MustConfigFilter retrieves the ConfigFilter value from the union,
// panicing if the value is not set.
func (u TransactionFilter) MustConfigFilter() ConfigFilter {

	val, ok := u.GetConfigFilter()
	if !ok {
		panic("arm ConfigFilter is not set")
	}

	return val
}

// GetConfigFilter retrieves the ConfigFilter value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionFilter) GetConfigFilter() (result ConfigFilter, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ConfigFilter" {
		result = *u.ConfigFilter
		ok = true
	}

	return
}

// MustPermissionFilter retrieves the PermissionFilter value from the union,
// panicing if the value is not set.
func (u TransactionFilter) MustPermissionFilter() PermissionFilter {

	val, ok := u.GetPermissionFilter()
	if !ok {
		panic("arm PermissionFilter is not set")
	}

	return val
}

// GetPermissionFilter retrieves the PermissionFilter value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionFilter) GetPermissionFilter() (result PermissionFilter, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PermissionFilter" {
		result = *u.PermissionFilter
		ok = true
	}

	return
}

// MustCallFilter retrieves the CallFilter value from the union,
// panicing if the value is not set.
func (u TransactionFilter) MustCallFilter() CallFilter {

	val, ok := u.GetCallFilter()
	if !ok {
		panic("arm CallFilter is not set")
	}

	return val
}

// GetCallFilter retrieves the CallFilter value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionFilter) GetCallFilter() (result CallFilter, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CallFilter" {
		result = *u.CallFilter
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u TransactionFilter) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *TransactionFilter) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TransactionFilter)(nil)
	_ encoding.BinaryUnmarshaler = (*TransactionFilter)(nil)
)

// ReceiptFilter generated union
type ReceiptFilter struct {
	Type ReceiptFilterType

	ValueFilter *ReceiptValueFilter
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReceiptFilter) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReceiptFilter
func (u ReceiptFilter) ArmForSwitch(sw int32) (string, bool) {
	switch ReceiptFilterType(sw) {

	case ReceiptFilterTypeNONE:
		return "", true

	case ReceiptFilterTypeRECEIPT:
		return "ValueFilter", true
	}
	return "-", false
}

// NewReceiptFilter creates a new  ReceiptFilter.
func NewReceiptFilter(aType ReceiptFilterType, value interface{}) (result ReceiptFilter, err error) {
	result.Type = aType
	switch aType {

	case ReceiptFilterTypeNONE:

	case ReceiptFilterTypeRECEIPT:

		tv, ok := value.(ReceiptValueFilter)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.ValueFilter = &tv

	}
	return
}

// MustValueFilter retrieves the ValueFilter value from the union,
// panicing if the value is not set.
func (u ReceiptFilter) MustValueFilter() ReceiptValueFilter {

	val, ok := u.GetValueFilter()
	if !ok {
		panic("arm ValueFilter is not set")
	}

	return val
}

// GetValueFilter retrieves the ValueFilter value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReceiptFilter) GetValueFilter() (result ReceiptValueFilter, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ValueFilter" {
		result = *u.ValueFilter
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u ReceiptFilter) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *ReceiptFilter) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ReceiptFilter)(nil)
	_ encoding.BinaryUnmarshaler = (*ReceiptFilter)(nil)
)

// End union section

// Namspace end mazzaroth
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

// Permission generated struct
type Permission struct {
	Key ID `json:"key"`

	Action PermissionAction `json:"action"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Permission) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Permission) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Permission)(nil)
	_ encoding.BinaryUnmarshaler = (*Permission)(nil)
)

// Action generated struct
type Action struct {
	Address ID `json:"address"`

	ChannelID ID `json:"channelID"`

	Nonce uint64 `json:"nonce"`

	BlockExpirationNumber uint64 `json:"blockExpirationNumber"`

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

	// UpdateTypeNONE enum value 0
	UpdateTypeNONE UpdateType = 0

	// UpdateTypeCONTRACT enum value 1
	UpdateTypeCONTRACT UpdateType = 1

	// UpdateTypeCONFIG enum value 2
	UpdateTypeCONFIG UpdateType = 2

	// UpdateTypePERMISSION enum value 3
	UpdateTypePERMISSION UpdateType = 3
)

// UpdateTypeMap generated enum map
var UpdateTypeMap = map[int32]string{

	0: "UpdateTypeNONE",

	1: "UpdateTypeCONTRACT",

	2: "UpdateTypeCONFIG",

	3: "UpdateTypePERMISSION",
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

// PermissionAction generated enum
type PermissionAction int32

const (

	// PermissionActionREVOKE enum value 0
	PermissionActionREVOKE PermissionAction = 0

	// PermissionActionGRANT enum value 1
	PermissionActionGRANT PermissionAction = 1
)

// PermissionActionMap generated enum map
var PermissionActionMap = map[int32]string{

	0: "PermissionActionREVOKE",

	1: "PermissionActionGRANT",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PermissionAction
func (s PermissionAction) ValidEnum(v int32) bool {
	_, ok := PermissionActionMap[v]
	return ok
}

// String returns the name of `e`
func (s PermissionAction) String() string {
	name, _ := PermissionActionMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PermissionAction) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PermissionAction) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PermissionAction)(nil)
	_ encoding.BinaryUnmarshaler = (*PermissionAction)(nil)
)

// ActionCategoryType generated enum
type ActionCategoryType int32

const (

	// ActionCategoryTypeNONE enum value 0
	ActionCategoryTypeNONE ActionCategoryType = 0

	// ActionCategoryTypeCALL enum value 1
	ActionCategoryTypeCALL ActionCategoryType = 1

	// ActionCategoryTypeUPDATE enum value 2
	ActionCategoryTypeUPDATE ActionCategoryType = 2
)

// ActionCategoryTypeMap generated enum map
var ActionCategoryTypeMap = map[int32]string{

	0: "ActionCategoryTypeNONE",

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

	// AuthorityTypeNONE enum value 0
	AuthorityTypeNONE AuthorityType = 0

	// AuthorityTypePERMISSIONED enum value 1
	AuthorityTypePERMISSIONED AuthorityType = 1
)

// AuthorityTypeMap generated enum map
var AuthorityTypeMap = map[int32]string{

	0: "AuthorityTypeNONE",

	1: "AuthorityTypePERMISSIONED",
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

	Permission *Permission
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

	case UpdateTypeNONE:
		return "", true

	case UpdateTypeCONTRACT:
		return "Contract", true

	case UpdateTypeCONFIG:
		return "ChannelConfig", true

	case UpdateTypePERMISSION:
		return "Permission", true
	}
	return "-", false
}

// NewUpdate creates a new  Update.
func NewUpdate(aType UpdateType, value interface{}) (result Update, err error) {
	result.Type = aType
	switch aType {

	case UpdateTypeNONE:

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

	case UpdateTypePERMISSION:

		tv, ok := value.(Permission)

		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Permission = &tv

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

// MustPermission retrieves the Permission value from the union,
// panicing if the value is not set.
func (u Update) MustPermission() Permission {

	val, ok := u.GetPermission()
	if !ok {
		panic("arm Permission is not set")
	}

	return val
}

// GetPermission retrieves the Permission value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Update) GetPermission() (result Permission, ok bool) {

	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Permission" {
		result = *u.Permission
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

	case ActionCategoryTypeNONE:
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

	case ActionCategoryTypeNONE:

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

	case AuthorityTypeNONE:
		return "", true

	case AuthorityTypePERMISSIONED:
		return "Origin", true
	}
	return "-", false
}

// NewAuthority creates a new  Authority.
func NewAuthority(aType AuthorityType, value interface{}) (result Authority, err error) {
	result.Type = aType
	switch aType {

	case AuthorityTypeNONE:

	case AuthorityTypePERMISSIONED:

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

// End union section

// Namspace end mazzaroth
var fmtTest = fmt.Sprint("this is a dummy usage of fmt")
