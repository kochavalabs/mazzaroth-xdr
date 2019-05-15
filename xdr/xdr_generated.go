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

	Value uint64
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

type Action struct {
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

	Address ID

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

// End struct section

// Start enum section

type ActionCategoryType int32

const (
	ActionCategoryTypeNONE ActionCategoryType = 0

	ActionCategoryTypeCALL ActionCategoryType = 1

	ActionCategoryTypeUPDATE ActionCategoryType = 2
)

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

// End enum section
//
// Start union section

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

// End union section

// Namspace end mazzaroth
