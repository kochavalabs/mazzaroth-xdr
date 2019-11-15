// Package xdr is automatically generated
// DO NOT EDIT or your changes may be overwritten
package xdr

import (
	"bytes"
	"encoding"
	"fmt"
	"io"

	xdr "github.com/stellar/go-xdr/xdr3"
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

// Account generated struct
type Account struct {
	Name string

	Nonce uint64

	PermissionedKeys []ID
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
	Header BlockHeader

	Transactions []Transaction
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
	Timestamp string `xdrmaxsize:"256"`

	BlockHeight uint64

	TransactionHeight uint64

	ConsensusSequenceNumber uint64

	TxMerkleRoot Hash

	TxReceiptRoot Hash

	StateRoot Hash

	PreviousHeader Hash

	BlockProducerAddress ID
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

// End typedef section

// Start struct section

// ChannelConfig generated struct
type ChannelConfig struct {
	ChannelID ID

	Owner ID

	ConsensusConfig ConsensusConfig
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

// PBFTConfig generated struct
type PBFTConfig struct {
	Validators []ID

	CheckpointPeriod uint64

	WatermarkRange uint64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s PBFTConfig) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *PBFTConfig) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*PBFTConfig)(nil)
	_ encoding.BinaryUnmarshaler = (*PBFTConfig)(nil)
)

// End struct section

// Start enum section

// ConsensusConfigType generated enum
type ConsensusConfigType int32

const (

	// ConsensusConfigTypeNONE enum value 0
	ConsensusConfigTypeNONE ConsensusConfigType = 0

	// ConsensusConfigTypePBFT enum value 1
	ConsensusConfigTypePBFT ConsensusConfigType = 1
)

// ConsensusConfigTypeMap generated enum map
var ConsensusConfigTypeMap = map[int32]string{

	0: "ConsensusConfigTypeNONE",

	1: "ConsensusConfigTypePBFT",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ConsensusConfigType
func (s ConsensusConfigType) ValidEnum(v int32) bool {
	_, ok := ConsensusConfigTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s ConsensusConfigType) String() string {
	name, _ := ConsensusConfigTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ConsensusConfigType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ConsensusConfigType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ConsensusConfigType)(nil)
	_ encoding.BinaryUnmarshaler = (*ConsensusConfigType)(nil)
)

// End enum section

// Start union section

// ConsensusConfig generated union
type ConsensusConfig struct {
	Type ConsensusConfigType

	PbftConfig *PBFTConfig
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ConsensusConfig) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ConsensusConfig
func (u ConsensusConfig) ArmForSwitch(sw int32) (string, bool) {
	switch ConsensusConfigType(sw) {

	case ConsensusConfigTypeNONE:
		return "", true

	case ConsensusConfigTypePBFT:
		return "PbftConfig", true
	}
	return "-", false
}

// NewConsensusConfig creates a new  ConsensusConfig.
func NewConsensusConfig(aType ConsensusConfigType, value interface{}) (result ConsensusConfig, err error) {
	result.Type = aType
	switch aType {

	case ConsensusConfigTypeNONE:

	case ConsensusConfigTypePBFT:

		tv, ok := value.(PBFTConfig)
		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.PbftConfig = &tv

	}
	return
}

// MustPbftConfig retrieves the PbftConfig value from the union,
// panicing if the value is not set.
func (u ConsensusConfig) MustPbftConfig() PBFTConfig {
	val, ok := u.GetPbftConfig()

	if !ok {
		panic("arm PbftConfig is not set")
	}

	return val
}

// GetPbftConfig retrieves the PbftConfig value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ConsensusConfig) GetPbftConfig() (result PBFTConfig, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PbftConfig" {
		result = *u.PbftConfig
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u ConsensusConfig) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *ConsensusConfig) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ConsensusConfig)(nil)
	_ encoding.BinaryUnmarshaler = (*ConsensusConfig)(nil)
)

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

// Parameter generated typedef
type Parameter []byte

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

// ContractMetadata generated struct
type ContractMetadata struct {
	Hash Hash

	Version uint64
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ContractMetadata) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ContractMetadata) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ContractMetadata)(nil)
	_ encoding.BinaryUnmarshaler = (*ContractMetadata)(nil)
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
	DownloadRequestPayload DownloadRequestPayload
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

// DownloadResponse generated struct
type DownloadResponse struct {
	DownloadStatus DownloadStatus

	DownloadResponsePayload DownloadResponsePayload
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
)

// DownloadRequestTypeMap generated enum map
var DownloadRequestTypeMap = map[int32]string{

	0: "DownloadRequestTypeUNKNOWN",

	1: "DownloadRequestTypeHEIGHT",

	2: "DownloadRequestTypeBLOCK",
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

	Height *uint64

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
		return "Height", true

	case DownloadRequestTypeBLOCK:
		return "Block", true
	}
	return "-", false
}

// NewDownloadResponsePayload creates a new  DownloadResponsePayload.
func NewDownloadResponsePayload(aType DownloadRequestType, value interface{}) (result DownloadResponsePayload, err error) {
	result.Type = aType
	switch aType {

	case DownloadRequestTypeUNKNOWN:

	case DownloadRequestTypeHEIGHT:

		tv, ok := value.(uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Height = &tv

	case DownloadRequestTypeBLOCK:

		tv, ok := value.(Block)
		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Block = &tv

	}
	return
}

// MustHeight retrieves the Height value from the union,
// panicing if the value is not set.
func (u DownloadResponsePayload) MustHeight() uint64 {
	val, ok := u.GetHeight()

	if !ok {
		panic("arm Height is not set")
	}

	return val
}

// GetHeight retrieves the Height value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u DownloadResponsePayload) GetHeight() (result uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Height" {
		result = *u.Height
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

// Event generated struct
type Event struct {
	Key string `xdrmaxsize:"256"`

	Parameters []Parameter
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Event) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Event) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Event)(nil)
	_ encoding.BinaryUnmarshaler = (*Event)(nil)
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

// ExecutionPlan generated struct
type ExecutionPlan struct {
	Host string `xdrmaxsize:"256"`

	ChannelID ID

	Calls []Call `xdrmaxsize:"100"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ExecutionPlan) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ExecutionPlan) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ExecutionPlan)(nil)
	_ encoding.BinaryUnmarshaler = (*ExecutionPlan)(nil)
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

// Receipt generated struct
type Receipt struct {
	Status ReceiptStatus

	StateRoot Hash

	Events []Event

	Result []byte
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

// StateStatus generated struct
type StateStatus struct {
	PreviousBlock uint64

	TransactionCount uint64
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
	ID Identifier
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
	ID Identifier
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
	Block Block

	StateStatus StateStatus

	Status BlockStatus

	StatusInfo StatusInfo
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
	Header BlockHeader

	StateStatus StateStatus

	Status BlockStatus

	StatusInfo StatusInfo
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
	TransactionID ID
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
	Transaction Transaction

	StateStatus StateStatus

	Status TransactionStatus

	StatusInfo StatusInfo
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
	Transaction Transaction
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
	TransactionID ID

	Status TransactionStatus

	StatusInfo StatusInfo
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

// ReadonlyRequest generated struct
type ReadonlyRequest struct {
	Call Call
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ReadonlyRequest) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ReadonlyRequest) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ReadonlyRequest)(nil)
	_ encoding.BinaryUnmarshaler = (*ReadonlyRequest)(nil)
)

// ReadonlyResponse generated struct
type ReadonlyResponse struct {
	Result []byte

	StateStatus StateStatus

	Status ReadonlyStatus

	StatusInfo StatusInfo
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ReadonlyResponse) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ReadonlyResponse) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ReadonlyResponse)(nil)
	_ encoding.BinaryUnmarshaler = (*ReadonlyResponse)(nil)
)

// ReceiptLookupRequest generated struct
type ReceiptLookupRequest struct {
	TransactionID ID
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
	Receipt Receipt

	StateStatus StateStatus

	Status ReceiptLookupStatus

	StatusInfo StatusInfo
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

// AccountNonceLookupRequest generated struct
type AccountNonceLookupRequest struct {
	Account ID
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountNonceLookupRequest) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountNonceLookupRequest) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountNonceLookupRequest)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountNonceLookupRequest)(nil)
)

// AccountNonceLookupResponse generated struct
type AccountNonceLookupResponse struct {
	Nonce uint64

	StateStatus StateStatus

	Status NonceLookupStatus

	StatusInfo StatusInfo
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s AccountNonceLookupResponse) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *AccountNonceLookupResponse) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*AccountNonceLookupResponse)(nil)
	_ encoding.BinaryUnmarshaler = (*AccountNonceLookupResponse)(nil)
)

// AccountInfoLookupRequest generated struct
type AccountInfoLookupRequest struct {
	Account ID
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
	AccountInfo Account

	StateStatus StateStatus

	Status InfoLookupStatus

	StatusInfo StatusInfo
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

// ReadonlyStatus generated enum
type ReadonlyStatus int32

const (

	// ReadonlyStatusUNKNOWN enum value 0
	ReadonlyStatusUNKNOWN ReadonlyStatus = 0

	// ReadonlyStatusSUCCESS enum value 1
	ReadonlyStatusSUCCESS ReadonlyStatus = 1

	// ReadonlyStatusFAILURE enum value 2
	ReadonlyStatusFAILURE ReadonlyStatus = 2
)

// ReadonlyStatusMap generated enum map
var ReadonlyStatusMap = map[int32]string{

	0: "ReadonlyStatusUNKNOWN",

	1: "ReadonlyStatusSUCCESS",

	2: "ReadonlyStatusFAILURE",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ReadonlyStatus
func (s ReadonlyStatus) ValidEnum(v int32) bool {
	_, ok := ReadonlyStatusMap[v]
	return ok
}

// String returns the name of `e`
func (s ReadonlyStatus) String() string {
	name, _ := ReadonlyStatusMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ReadonlyStatus) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ReadonlyStatus) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ReadonlyStatus)(nil)
	_ encoding.BinaryUnmarshaler = (*ReadonlyStatus)(nil)
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

// NonceLookupStatus generated enum
type NonceLookupStatus int32

const (

	// NonceLookupStatusUNKNOWN enum value 0
	NonceLookupStatusUNKNOWN NonceLookupStatus = 0

	// NonceLookupStatusFOUND enum value 1
	NonceLookupStatusFOUND NonceLookupStatus = 1

	// NonceLookupStatusNOT_FOUND enum value 2
	NonceLookupStatusNOT_FOUND NonceLookupStatus = 2
)

// NonceLookupStatusMap generated enum map
var NonceLookupStatusMap = map[int32]string{

	0: "NonceLookupStatusUNKNOWN",

	1: "NonceLookupStatusFOUND",

	2: "NonceLookupStatusNOT_FOUND",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for NonceLookupStatus
func (s NonceLookupStatus) ValidEnum(v int32) bool {
	_, ok := NonceLookupStatusMap[v]
	return ok
}

// String returns the name of `e`
func (s NonceLookupStatus) String() string {
	name, _ := NonceLookupStatusMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s NonceLookupStatus) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *NonceLookupStatus) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*NonceLookupStatus)(nil)
	_ encoding.BinaryUnmarshaler = (*NonceLookupStatus)(nil)
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

// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

// Call generated struct
type Call struct {
	Function string `xdrmaxsize:"256"`

	Parameters []Parameter
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

// Update generated struct
type Update struct {
	Contract []byte
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Update) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Update) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Update)(nil)
	_ encoding.BinaryUnmarshaler = (*Update)(nil)
)

// Permission generated struct
type Permission struct {
	Key ID

	Action PermissionAction
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

// Config generated struct
type Config struct {
	Action ConfigAction
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Config) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Config) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Config)(nil)
	_ encoding.BinaryUnmarshaler = (*Config)(nil)
)

// Action generated struct
type Action struct {
	Address ID

	ChannelID ID

	Nonce uint64

	Category ActionCategory
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
	Signature Signature

	Signer Authority

	Action Action
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

// Input generated struct
type Input struct {
	InputType InputType

	Function string `xdrmaxsize:"256"`

	Parameters []Parameter
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Input) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Input) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Input)(nil)
	_ encoding.BinaryUnmarshaler = (*Input)(nil)
)

// End struct section

// Start enum section

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

// ConfigType generated enum
type ConfigType int32

const (

	// ConfigTypeNONE enum value 0
	ConfigTypeNONE ConfigType = 0

	// ConfigTypeCHANNELID enum value 1
	ConfigTypeCHANNELID ConfigType = 1

	// ConfigTypeOWNER enum value 2
	ConfigTypeOWNER ConfigType = 2

	// ConfigTypeCONSENSUS enum value 3
	ConfigTypeCONSENSUS ConfigType = 3
)

// ConfigTypeMap generated enum map
var ConfigTypeMap = map[int32]string{

	0: "ConfigTypeNONE",

	1: "ConfigTypeCHANNELID",

	2: "ConfigTypeOWNER",

	3: "ConfigTypeCONSENSUS",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ConfigType
func (s ConfigType) ValidEnum(v int32) bool {
	_, ok := ConfigTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s ConfigType) String() string {
	name, _ := ConfigTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ConfigType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ConfigType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ConfigType)(nil)
	_ encoding.BinaryUnmarshaler = (*ConfigType)(nil)
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

	// ActionCategoryTypePERMISSION enum value 3
	ActionCategoryTypePERMISSION ActionCategoryType = 3

	// ActionCategoryTypeCONFIG enum value 4
	ActionCategoryTypeCONFIG ActionCategoryType = 4
)

// ActionCategoryTypeMap generated enum map
var ActionCategoryTypeMap = map[int32]string{

	0: "ActionCategoryTypeNONE",

	1: "ActionCategoryTypeCALL",

	2: "ActionCategoryTypeUPDATE",

	3: "ActionCategoryTypePERMISSION",

	4: "ActionCategoryTypeCONFIG",
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

// InputType generated enum
type InputType int32

const (

	// InputTypeNONE enum value 0
	InputTypeNONE InputType = 0

	// InputTypeREADONLY enum value 1
	InputTypeREADONLY InputType = 1

	// InputTypeEXECUTE enum value 2
	InputTypeEXECUTE InputType = 2

	// InputTypeCONSTRUCTOR enum value 3
	InputTypeCONSTRUCTOR InputType = 3
)

// InputTypeMap generated enum map
var InputTypeMap = map[int32]string{

	0: "InputTypeNONE",

	1: "InputTypeREADONLY",

	2: "InputTypeEXECUTE",

	3: "InputTypeCONSTRUCTOR",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for InputType
func (s InputType) ValidEnum(v int32) bool {
	_, ok := InputTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s InputType) String() string {
	name, _ := InputTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s InputType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *InputType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*InputType)(nil)
	_ encoding.BinaryUnmarshaler = (*InputType)(nil)
)

// End enum section

// Start union section

// ConfigAction generated union
type ConfigAction struct {
	Type ConfigType

	ChannelID *ID

	Owner *ID

	ConsensusConfig *ConsensusConfig
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ConfigAction) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ConfigAction
func (u ConfigAction) ArmForSwitch(sw int32) (string, bool) {
	switch ConfigType(sw) {

	case ConfigTypeNONE:
		return "", true

	case ConfigTypeCHANNELID:
		return "ChannelID", true

	case ConfigTypeOWNER:
		return "Owner", true

	case ConfigTypeCONSENSUS:
		return "ConsensusConfig", true
	}
	return "-", false
}

// NewConfigAction creates a new  ConfigAction.
func NewConfigAction(aType ConfigType, value interface{}) (result ConfigAction, err error) {
	result.Type = aType
	switch aType {

	case ConfigTypeNONE:

	case ConfigTypeCHANNELID:

		tv, ok := value.(ID)
		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.ChannelID = &tv

	case ConfigTypeOWNER:

		tv, ok := value.(ID)
		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Owner = &tv

	case ConfigTypeCONSENSUS:

		tv, ok := value.(ConsensusConfig)
		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.ConsensusConfig = &tv

	}
	return
}

// MustChannelID retrieves the ChannelID value from the union,
// panicing if the value is not set.
func (u ConfigAction) MustChannelID() ID {
	val, ok := u.GetChannelID()

	if !ok {
		panic("arm ChannelID is not set")
	}

	return val
}

// GetChannelID retrieves the ChannelID value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ConfigAction) GetChannelID() (result ID, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ChannelID" {
		result = *u.ChannelID
		ok = true
	}

	return
}

// MustOwner retrieves the Owner value from the union,
// panicing if the value is not set.
func (u ConfigAction) MustOwner() ID {
	val, ok := u.GetOwner()

	if !ok {
		panic("arm Owner is not set")
	}

	return val
}

// GetOwner retrieves the Owner value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ConfigAction) GetOwner() (result ID, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Owner" {
		result = *u.Owner
		ok = true
	}

	return
}

// MustConsensusConfig retrieves the ConsensusConfig value from the union,
// panicing if the value is not set.
func (u ConfigAction) MustConsensusConfig() ConsensusConfig {
	val, ok := u.GetConsensusConfig()

	if !ok {
		panic("arm ConsensusConfig is not set")
	}

	return val
}

// GetConsensusConfig retrieves the ConsensusConfig value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ConfigAction) GetConsensusConfig() (result ConsensusConfig, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ConsensusConfig" {
		result = *u.ConsensusConfig
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u ConfigAction) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, u)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *ConfigAction) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), u)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ConfigAction)(nil)
	_ encoding.BinaryUnmarshaler = (*ConfigAction)(nil)
)

// ActionCategory generated union
type ActionCategory struct {
	Type ActionCategoryType

	Call *Call

	Update *Update

	Permission *Permission

	Config *Config
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

	case ActionCategoryTypePERMISSION:
		return "Permission", true

	case ActionCategoryTypeCONFIG:
		return "Config", true
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

	case ActionCategoryTypePERMISSION:

		tv, ok := value.(Permission)
		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Permission = &tv

	case ActionCategoryTypeCONFIG:

		tv, ok := value.(Config)
		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Config = &tv

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

// MustPermission retrieves the Permission value from the union,
// panicing if the value is not set.
func (u ActionCategory) MustPermission() Permission {
	val, ok := u.GetPermission()

	if !ok {
		panic("arm Permission is not set")
	}

	return val
}

// GetPermission retrieves the Permission value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ActionCategory) GetPermission() (result Permission, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Permission" {
		result = *u.Permission
		ok = true
	}

	return
}

// MustConfig retrieves the Config value from the union,
// panicing if the value is not set.
func (u ActionCategory) MustConfig() Config {
	val, ok := u.GetConfig()

	if !ok {
		panic("arm Config is not set")
	}

	return val
}

// GetConfig retrieves the Config value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ActionCategory) GetConfig() (result Config, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Config" {
		result = *u.Config
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
