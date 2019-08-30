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
//
// Start union section
// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
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

type BlockHeader struct {
	Timestamp string `xdrmaxsize:"256"`

	BlockHeight uint64

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
//
// Start union section
// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
type ChannelConfig struct {
	Owner ID

	Validators []ID

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

type PBFTConfig struct {
	CheckpointPeriod uint64
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

type ConsensusConfigType int32

const (
	ConsensusConfigTypeNONE ConsensusConfigType = 0

	ConsensusConfigTypePBFT ConsensusConfigType = 1
)

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
//
// Start union section

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
func (s ConsensusConfig) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ConsensusConfig) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
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
//
// Start union section
// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
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
//
// Start union section
// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
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
//
// Start union section
// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
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
//
// Start union section
// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
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

type ReceiptStatus int32

const (
	ReceiptStatusFAILURE ReceiptStatus = 0

	ReceiptStatusSUCCESS ReceiptStatus = 1
)

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
//
// Start union section
// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

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

type IdentifierType int32

const (
	IdentifierTypeNONE IdentifierType = 0

	IdentifierTypeNUMBER IdentifierType = 1

	IdentifierTypeHASH IdentifierType = 2
)

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

type BlockStatus int32

const (
	BlockStatusUNKNOWN BlockStatus = 0

	BlockStatusCREATED BlockStatus = 1

	BlockStatusFUTURE BlockStatus = 2

	BlockStatusNOT_FOUND BlockStatus = 3
)

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

type TransactionStatus int32

const (
	TransactionStatusUNKNOWN TransactionStatus = 0

	TransactionStatusACCEPTED TransactionStatus = 1

	TransactionStatusREJECTED TransactionStatus = 2

	TransactionStatusCONFIRMED TransactionStatus = 3

	TransactionStatusNOT_FOUND TransactionStatus = 4
)

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

type ReadonlyStatus int32

const (
	ReadonlyStatusUNKNOWN ReadonlyStatus = 0

	ReadonlyStatusSUCCESS ReadonlyStatus = 1

	ReadonlyStatusFAILURE ReadonlyStatus = 2
)

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

type ReceiptLookupStatus int32

const (
	ReceiptLookupStatusUNKNOWN ReceiptLookupStatus = 0

	ReceiptLookupStatusFOUND ReceiptLookupStatus = 1

	ReceiptLookupStatusNOT_FOUND ReceiptLookupStatus = 2
)

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

type NonceLookupStatus int32

const (
	NonceLookupStatusUNKNOWN NonceLookupStatus = 0

	NonceLookupStatusFOUND NonceLookupStatus = 1

	NonceLookupStatusNOT_FOUND NonceLookupStatus = 2
)

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

type InfoLookupStatus int32

const (
	InfoLookupStatusUNKNOWN InfoLookupStatus = 0

	InfoLookupStatusFOUND InfoLookupStatus = 1

	InfoLookupStatusNOT_FOUND InfoLookupStatus = 2
)

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
//
// Start union section

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
func (s Identifier) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Identifier) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
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
type BasicColumn struct {
	Name string `xdrmaxsize:"40"`

	Typ BasicType
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BasicColumn) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BasicColumn) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BasicColumn)(nil)
	_ encoding.BinaryUnmarshaler = (*BasicColumn)(nil)
)

type TypedefColumn struct {
	Name string `xdrmaxsize:"40"`

	Child [1]Column
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s TypedefColumn) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *TypedefColumn) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*TypedefColumn)(nil)
	_ encoding.BinaryUnmarshaler = (*TypedefColumn)(nil)
)

type StructColumn struct {
	Name string `xdrmaxsize:"40"`

	Columns []Column `xdrmaxsize:"40"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s StructColumn) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *StructColumn) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*StructColumn)(nil)
	_ encoding.BinaryUnmarshaler = (*StructColumn)(nil)
)

type ArrayColumn struct {
	Name string `xdrmaxsize:"40"`

	Fixed bool

	Length uint32

	Column [1]Column
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ArrayColumn) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ArrayColumn) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ArrayColumn)(nil)
	_ encoding.BinaryUnmarshaler = (*ArrayColumn)(nil)
)

type Table struct {
	Name string `xdrmaxsize:"40"`

	Columns []Column `xdrmaxsize:"40"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Table) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Table) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Table)(nil)
	_ encoding.BinaryUnmarshaler = (*Table)(nil)
)

type Schema struct {
	Tables []Table `xdrmaxsize:"40"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Schema) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Schema) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Schema)(nil)
	_ encoding.BinaryUnmarshaler = (*Schema)(nil)
)

// End struct section

// Start enum section

type BasicType int32

const (
	BasicTypeBOOLEAN BasicType = 0

	BasicTypeSTRING BasicType = 1

	BasicTypeOPAQUE BasicType = 2

	BasicTypeINT BasicType = 3

	BasicTypeUNSIGNED_INT BasicType = 4

	BasicTypeHYPER BasicType = 5

	BasicTypeUNSIGNED_HYPER BasicType = 6

	BasicTypeFLOAT BasicType = 7

	BasicTypeDOUBLE BasicType = 8
)

var BasicTypeMap = map[int32]string{

	0: "BasicTypeBOOLEAN",

	1: "BasicTypeSTRING",

	2: "BasicTypeOPAQUE",

	3: "BasicTypeINT",

	4: "BasicTypeUNSIGNED_INT",

	5: "BasicTypeHYPER",

	6: "BasicTypeUNSIGNED_HYPER",

	7: "BasicTypeFLOAT",

	8: "BasicTypeDOUBLE",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for BasicType
func (s BasicType) ValidEnum(v int32) bool {
	_, ok := BasicTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s BasicType) String() string {
	name, _ := BasicTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s BasicType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *BasicType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*BasicType)(nil)
	_ encoding.BinaryUnmarshaler = (*BasicType)(nil)
)

type ColumnType int32

const (
	ColumnTypeBASIC ColumnType = 0

	ColumnTypeSTRUCT ColumnType = 1

	ColumnTypeARRAY ColumnType = 2

	ColumnTypeTYPEDEF ColumnType = 3
)

var ColumnTypeMap = map[int32]string{

	0: "ColumnTypeBASIC",

	1: "ColumnTypeSTRUCT",

	2: "ColumnTypeARRAY",

	3: "ColumnTypeTYPEDEF",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ColumnType
func (s ColumnType) ValidEnum(v int32) bool {
	_, ok := ColumnTypeMap[v]
	return ok
}

// String returns the name of `e`
func (s ColumnType) String() string {
	name, _ := ColumnTypeMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ColumnType) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ColumnType) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ColumnType)(nil)
	_ encoding.BinaryUnmarshaler = (*ColumnType)(nil)
)

// End enum section
//
// Start union section

type Column struct {
	Type ColumnType

	Basic *BasicColumn

	Array *ArrayColumn

	Struc *StructColumn

	Typ *TypedefColumn
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u Column) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of Column
func (u Column) ArmForSwitch(sw int32) (string, bool) {
	switch ColumnType(sw) {

	case ColumnTypeBASIC:
		return "Basic", true

	case ColumnTypeARRAY:
		return "Array", true

	case ColumnTypeSTRUCT:
		return "Struc", true

	case ColumnTypeTYPEDEF:
		return "Typ", true
	}
	return "-", false
}

// NewColumn creates a new  Column.
func NewColumn(aType ColumnType, value interface{}) (result Column, err error) {
	result.Type = aType
	switch aType {

	case ColumnTypeBASIC:

		tv, ok := value.(BasicColumn)
		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Basic = &tv

	case ColumnTypeARRAY:

		tv, ok := value.(ArrayColumn)
		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Array = &tv

	case ColumnTypeSTRUCT:

		tv, ok := value.(StructColumn)
		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Struc = &tv

	case ColumnTypeTYPEDEF:

		tv, ok := value.(TypedefColumn)
		if !ok {
			err = fmt.Errorf("invalid value, must be [object]")
			return
		}
		result.Typ = &tv

	}
	return
}

// MustBasic retrieves the Basic value from the union,
// panicing if the value is not set.
func (u Column) MustBasic() BasicColumn {
	val, ok := u.GetBasic()

	if !ok {
		panic("arm Basic is not set")
	}

	return val
}

// GetBasic retrieves the Basic value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Column) GetBasic() (result BasicColumn, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Basic" {
		result = *u.Basic
		ok = true
	}

	return
}

// MustArray retrieves the Array value from the union,
// panicing if the value is not set.
func (u Column) MustArray() ArrayColumn {
	val, ok := u.GetArray()

	if !ok {
		panic("arm Array is not set")
	}

	return val
}

// GetArray retrieves the Array value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Column) GetArray() (result ArrayColumn, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Array" {
		result = *u.Array
		ok = true
	}

	return
}

// MustStruc retrieves the Struc value from the union,
// panicing if the value is not set.
func (u Column) MustStruc() StructColumn {
	val, ok := u.GetStruc()

	if !ok {
		panic("arm Struc is not set")
	}

	return val
}

// GetStruc retrieves the Struc value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Column) GetStruc() (result StructColumn, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Struc" {
		result = *u.Struc
		ok = true
	}

	return
}

// MustTyp retrieves the Typ value from the union,
// panicing if the value is not set.
func (u Column) MustTyp() TypedefColumn {
	val, ok := u.GetTyp()

	if !ok {
		panic("arm Typ is not set")
	}

	return val
}

// GetTyp retrieves the Typ value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Column) GetTyp() (result TypedefColumn, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Typ" {
		result = *u.Typ
		ok = true
	}

	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Column) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Column) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Column)(nil)
	_ encoding.BinaryUnmarshaler = (*Column)(nil)
)

// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
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

type CommittedTransaction struct {
	Transaction Transaction

	SequenceNumber uint64

	ReceiptID []ID `xdrmaxsize:"25"`

	CurrentTransactionRoot Hash

	Signatures []Signature
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s CommittedTransaction) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *CommittedTransaction) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*CommittedTransaction)(nil)
	_ encoding.BinaryUnmarshaler = (*CommittedTransaction)(nil)
)

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

type PermissionAction int32

const (
	PermissionActionREVOKE PermissionAction = 0

	PermissionActionGRANT PermissionAction = 1
)

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

type ActionCategoryType int32

const (
	ActionCategoryTypeNONE ActionCategoryType = 0

	ActionCategoryTypeCALL ActionCategoryType = 1

	ActionCategoryTypeUPDATE ActionCategoryType = 2

	ActionCategoryTypePERMISSION ActionCategoryType = 3
)

var ActionCategoryTypeMap = map[int32]string{

	0: "ActionCategoryTypeNONE",

	1: "ActionCategoryTypeCALL",

	2: "ActionCategoryTypeUPDATE",

	3: "ActionCategoryTypePERMISSION",
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

type AuthorityType int32

const (
	AuthorityTypeNONE AuthorityType = 0

	AuthorityTypePERMISSIONED AuthorityType = 1
)

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

type InputType int32

const (
	InputTypeNONE InputType = 0

	InputTypeREADONLY InputType = 1

	InputTypeEXECUTE InputType = 2

	InputTypeCONSTRUCTOR InputType = 3
)

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
//
// Start union section

type ActionCategory struct {
	Type ActionCategoryType

	Call *Call

	Update *Update

	Permission *Permission
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

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ActionCategory) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ActionCategory) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ActionCategory)(nil)
	_ encoding.BinaryUnmarshaler = (*ActionCategory)(nil)
)

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
func (s Authority) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Authority) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Authority)(nil)
	_ encoding.BinaryUnmarshaler = (*Authority)(nil)
)

// End union section

// Namspace end mazzaroth
var fmtTest = fmt.Sprint("this is a dummy usage of fmt")
