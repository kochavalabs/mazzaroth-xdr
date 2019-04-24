          // Package xdr is generated from:
          //
          //  ./idl/block.x
//  ./idl/common.x
//  ./idl/event.x
//  ./idl/receipt.x
//  ./idl/rpc.x
//  ./idl/transaction.x
          //
          // DO NOT EDIT or your changes may be overwritten
          package xdr

          import (
            "bytes"
            "encoding"
            "io"
            "fmt"

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

// Block is an XDR Struct defines as:
//
//   struct Block
//      {
//        // Order is preserved in repeated fields
//        BlockHeader header;
//        Transaction transactions<>;
//      };
//
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

// BlockHeader is an XDR Struct defines as:
//
//   struct BlockHeader
//      {
//    
//        string timestamp<256>; 
//    
//        unsigned hyper blockHeight;
//    
//        Hash txMerkleRoot;
//    
//        Hash txReceiptRoot;
//    
//        Hash stateRoot;
//    
//        Hash previousHeader;
//    
//        ID blockProducerAddress;
//        
//      };
//
type BlockHeader struct {
  Timestamp string `xdrmaxsize:"256"`
  BlockHeight uint64 
  TxMerkleRoot Hash 
  TxReceiptRoot Hash 
  StateRoot Hash 
  PreviousHeader Hash 
  BlockProducerAddress Id 
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

// Signature is an XDR Typedef defines as:
//
//   typedef opaque Signature[64];
//
type Signature [64]byte
// XDRMaxSize implements the Sized interface for Signature
func (e Signature) XDRMaxSize() int {
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

// Id is an XDR Typedef defines as:
//
//   typedef opaque ID[32];
//
type Id [32]byte
// XDRMaxSize implements the Sized interface for Id
func (e Id) XDRMaxSize() int {
  return 32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Id) MarshalBinary() ([]byte, error) {
  b := new(bytes.Buffer)
  _, err := Marshal(b, s)
  return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Id) UnmarshalBinary(inp []byte) error {
  _, err := Unmarshal(bytes.NewReader(inp), s)
  return err
}

var (
  _ encoding.BinaryMarshaler   = (*Id)(nil)
  _ encoding.BinaryUnmarshaler = (*Id)(nil)
)

// Hash is an XDR Typedef defines as:
//
//   typedef opaque Hash[32];
//
type Hash [32]byte
// XDRMaxSize implements the Sized interface for Hash
func (e Hash) XDRMaxSize() int {
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

// Parameter is an XDR Typedef defines as:
//
//   typedef opaque Parameter<>;
//
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

// Event is an XDR Struct defines as:
//
//   struct Event {
//        // Name of Event (Function Name)
//        string key<256>;
//    
//        opaque values<>;
//      };
//
type Event struct {
  Key string `xdrmaxsize:"256"`
  Values []byte 
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

// Receipt is an XDR Struct defines as:
//
//   struct Receipt
//      {
//        // Status failure or success
//        ReceiptStatus status;
//     
//        // The state root after execution of the transaction
//        Hash stateRoot;
//     
//        // The list of events fired during execution of this transaction
//        Event events<>;
//     
//        // Return results of execution if there is one for function called
//        opaque result<>;
//      };
//
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

// ReceiptStatus is an XDR Enum defines as:
//
//   enum ReceiptStatus
//      {
//        FAILURE = 0,
//        SUCCESS = 1
//      };
//
type ReceiptStatus int32
const (
  ReceiptStatusFailure ReceiptStatus = 0
  ReceiptStatusSuccess ReceiptStatus = 1
)
var receiptStatusMap = map[int32]string{
  0: "ReceiptStatusFailure",
  1: "ReceiptStatusSuccess",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ReceiptStatus
func (e ReceiptStatus) ValidEnum(v int32) bool {
  _, ok := receiptStatusMap[v]
  return ok
}
// String returns the name of `e`
func (e ReceiptStatus) String() string {
  name, _ := receiptStatusMap[int32(e)]
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

// StatusInfo is an XDR Typedef defines as:
//
//   typedef string StatusInfo<256>;
//
type StatusInfo string
// XDRMaxSize implements the Sized interface for StatusInfo
func (e StatusInfo) XDRMaxSize() int {
  return 256
}

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

// IdentifierType is an XDR Enum defines as:
//
//   enum IdentifierType
//      {
//        NUMBER = 0,
//        HASH = 1
//      };
//
type IdentifierType int32
const (
  IdentifierTypeNumber IdentifierType = 0
  IdentifierTypeHash IdentifierType = 1
)
var identifierTypeMap = map[int32]string{
  0: "IdentifierTypeNumber",
  1: "IdentifierTypeHash",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for IdentifierType
func (e IdentifierType) ValidEnum(v int32) bool {
  _, ok := identifierTypeMap[v]
  return ok
}
// String returns the name of `e`
func (e IdentifierType) String() string {
  name, _ := identifierTypeMap[int32(e)]
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

// Identifier is an XDR Union defines as:
//
//   union Identifier switch (IdentifierType type)
//      {
//        case NUMBER:
//          unsigned hyper number;
//        case HASH:
//          Hash hash;
//      };
//
type Identifier struct{
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
    case IdentifierTypeNumber:
      return "Number", true
    case IdentifierTypeHash:
      return "Hash", true
}
return "-", false
}

// NewIdentifier creates a new  Identifier.
func NewIdentifier(aType IdentifierType, value interface{}) (result Identifier, err error) {
  result.Type = aType
switch IdentifierType(aType) {
    case IdentifierTypeNumber:
                  tv, ok := value.(uint64)
            if !ok {
              err = fmt.Errorf("invalid value, must be uint64")
              return
            }
            result.Number = &tv
    case IdentifierTypeHash:
                  tv, ok := value.(Hash)
            if !ok {
              err = fmt.Errorf("invalid value, must be Hash")
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

// BlockLookupRequest is an XDR Struct defines as:
//
//   struct BlockLookupRequest
//      {
//        Identifier id; 
//      };
//
type BlockLookupRequest struct {
  Id Identifier 
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

// BlockHeaderLookupRequest is an XDR Struct defines as:
//
//   struct BlockHeaderLookupRequest
//      {
//        Identifier id; 
//      };
//
type BlockHeaderLookupRequest struct {
  Id Identifier 
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

// BlockLookupResponse is an XDR Struct defines as:
//
//   struct BlockLookupResponse
//      {
//    
//        // Block that was requested if status is found.
//        Block block;
//    
//        // Status for the requested block.
//        BlockStatus status;
//    
//        // Human readable information to help understand the block status.
//        StatusInfo statusInfo;
//      };
//
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

// BlockHeaderLookupResponse is an XDR Struct defines as:
//
//   struct BlockHeaderLookupResponse
//      {
//    
//        // Block header that was requested if status is found.
//        BlockHeader header;
//    
//        // Status for the requested block.
//        BlockStatus status;
//    
//        // Human readable information to help understand the block status.
//        StatusInfo statusInfo;
//    
//      };
//
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

// BlockStatus is an XDR Enum defines as:
//
//   enum BlockStatus
//      {
//    
//        // Status of the block is unkown.
//        UNKNOWN = 0,
//    
//        // This block has been created.
//        CREATED = 1,
//    
//        // This block has not been created yet.
//        FUTURE = 2,
//    
//        // The block that was looked up was not found.
//        NOT_FOUND = 3
//      };
//
type BlockStatus int32
const (
  BlockStatusUnknown BlockStatus = 0
  BlockStatusCreated BlockStatus = 1
  BlockStatusFuture BlockStatus = 2
  BlockStatusNotFound BlockStatus = 3
)
var blockStatusMap = map[int32]string{
  0: "BlockStatusUnknown",
  1: "BlockStatusCreated",
  2: "BlockStatusFuture",
  3: "BlockStatusNotFound",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for BlockStatus
func (e BlockStatus) ValidEnum(v int32) bool {
  _, ok := blockStatusMap[v]
  return ok
}
// String returns the name of `e`
func (e BlockStatus) String() string {
  name, _ := blockStatusMap[int32(e)]
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

// TransactionLookupRequest is an XDR Struct defines as:
//
//   struct TransactionLookupRequest 
//      {
//        // Unique transaction identifier.
//        ID transactionId;
//      };
//
type TransactionLookupRequest struct {
  TransactionId Id 
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

// TransactionLookupResponse is an XDR Struct defines as:
//
//   struct TransactionLookupResponse
//      {
//        // Final transaction written to the blockchain.
//        Transaction transaction;
//    
//        // Current status of the transaction.
//        TransactionStatus status;
//    
//        // Human readable information to help understand the transaction status.
//        StatusInfo statusInfo;
//      };
//
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

// TransactionSubmitRequest is an XDR Struct defines as:
//
//   struct TransactionSubmitRequest
//      {
//        Transaction transaction;
//      };
//
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

// TransactionSubmitResponse is an XDR Struct defines as:
//
//   struct TransactionSubmitResponse
//      {
//        // Final transaction written to the blockchain. (if successful)
//        ID transactionId;
//    
//        // Current status of the transaction.
//        TransactionStatus status;
//    
//        // Human readable information to help understand the transaction status.
//        StatusInfo statusInfo;
//      };
//
type TransactionSubmitResponse struct {
  TransactionId Id 
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

// TransactionStatus is an XDR Enum defines as:
//
//   enum TransactionStatus
//      {
//    
//        // The transaction status is either not known or not set.
//        UNKNOWN = 0,
//    
//        // The transaction has been accepted by a node and is in the process of being
//        // submitted to the blockchain.
//        ACCEPTED = 1,
//    
//        // This transaction was not accepted by the blockchain.
//        REJECTED = 2,
//    
//        // The transaction has succesfully been added to the blockchain.
//        CONFIRMED = 3,
//    
//        // This transaction was not found.
//        NOT_FOUND = 4
//      };
//
type TransactionStatus int32
const (
  TransactionStatusUnknown TransactionStatus = 0
  TransactionStatusAccepted TransactionStatus = 1
  TransactionStatusRejected TransactionStatus = 2
  TransactionStatusConfirmed TransactionStatus = 3
  TransactionStatusNotFound TransactionStatus = 4
)
var transactionStatusMap = map[int32]string{
  0: "TransactionStatusUnknown",
  1: "TransactionStatusAccepted",
  2: "TransactionStatusRejected",
  3: "TransactionStatusConfirmed",
  4: "TransactionStatusNotFound",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for TransactionStatus
func (e TransactionStatus) ValidEnum(v int32) bool {
  _, ok := transactionStatusMap[v]
  return ok
}
// String returns the name of `e`
func (e TransactionStatus) String() string {
  name, _ := transactionStatusMap[int32(e)]
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

// ReceiptLookupRequest is an XDR Struct defines as:
//
//   struct ReceiptLookupRequest 
//      {
//        // Unique transaction identifier.
//        ID transactionId;
//      };
//
type ReceiptLookupRequest struct {
  TransactionId Id 
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

// ReceiptLookupResponse is an XDR Struct defines as:
//
//   struct ReceiptLookupResponse
//      {
//        // Final receipt written to the blockchain.
//        Receipt receipt; 
//    
//        // Current status of the receipt
//        ReceiptLookupStatus status;
//    
//        // Human readable information to help understand the receipt status.
//        StatusInfo statusInfo;
//      };
//
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

// ReceiptLookupStatus is an XDR Enum defines as:
//
//   enum ReceiptLookupStatus
//      {
//        // The receipt status is either not known or not set.
//        UNKNOWN = 0,
//    
//        // The transaction receipt was found
//        FOUND = 1,
//    
//        // This transaction receipt was not found.
//        NOT_FOUND = 2
//      };
//
type ReceiptLookupStatus int32
const (
  ReceiptLookupStatusUnknown ReceiptLookupStatus = 0
  ReceiptLookupStatusFound ReceiptLookupStatus = 1
  ReceiptLookupStatusNotFound ReceiptLookupStatus = 2
)
var receiptLookupStatusMap = map[int32]string{
  0: "ReceiptLookupStatusUnknown",
  1: "ReceiptLookupStatusFound",
  2: "ReceiptLookupStatusNotFound",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ReceiptLookupStatus
func (e ReceiptLookupStatus) ValidEnum(v int32) bool {
  _, ok := receiptLookupStatusMap[v]
  return ok
}
// String returns the name of `e`
func (e ReceiptLookupStatus) String() string {
  name, _ := receiptLookupStatusMap[int32(e)]
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

// Call is an XDR Struct defines as:
//
//   struct Call
//      {
//        // Contract function to execute.
//        string function<256>;
//    
//        // Parameters to the contract function. The serialization format is defined
//        // by the contract itself.
//        Parameter parameters<>;
//      };
//
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

// Update is an XDR Struct defines as:
//
//   struct Update
//      {
//        // Contract binary bytes.
//        opaque contract<>;
//      };
//
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

// ActionCategoryType is an XDR Enum defines as:
//
//   enum ActionCategoryType
//      {
//        CALL = 0,
//        UPDATE = 1
//      };
//
type ActionCategoryType int32
const (
  ActionCategoryTypeCall ActionCategoryType = 0
  ActionCategoryTypeUpdate ActionCategoryType = 1
)
var actionCategoryTypeMap = map[int32]string{
  0: "ActionCategoryTypeCall",
  1: "ActionCategoryTypeUpdate",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ActionCategoryType
func (e ActionCategoryType) ValidEnum(v int32) bool {
  _, ok := actionCategoryTypeMap[v]
  return ok
}
// String returns the name of `e`
func (e ActionCategoryType) String() string {
  name, _ := actionCategoryTypeMap[int32(e)]
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

// ActionCategory is an XDR Union defines as:
//
//   union ActionCategory switch (ActionCategoryType type)
//      {
//        case CALL:
//            Call call;
//        case UPDATE:
//            Update update;
//      };
//
type ActionCategory struct{
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
    case ActionCategoryTypeCall:
      return "Call", true
    case ActionCategoryTypeUpdate:
      return "Update", true
}
return "-", false
}

// NewActionCategory creates a new  ActionCategory.
func NewActionCategory(aType ActionCategoryType, value interface{}) (result ActionCategory, err error) {
  result.Type = aType
switch ActionCategoryType(aType) {
    case ActionCategoryTypeCall:
                  tv, ok := value.(Call)
            if !ok {
              err = fmt.Errorf("invalid value, must be Call")
              return
            }
            result.Call = &tv
    case ActionCategoryTypeUpdate:
                  tv, ok := value.(Update)
            if !ok {
              err = fmt.Errorf("invalid value, must be Update")
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

// Action is an XDR Struct defines as:
//
//   struct Action 
//      {
//        ID channelId;    
//    
//        unsigned hyper nonce;
//    
//        ActionCategory category;
//    
//      };
//
type Action struct {
  ChannelId Id 
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

// Transaction is an XDR Struct defines as:
//
//   struct Transaction
//      {
//        // Byte array signature of the Transaction bytes signed by the Transaction 
//        // sender's private key.
//        Signature signature;
//    
//        // Byte array representing the id of the sender, this also happens
//        // to be the sender's account public key.
//        ID address;
//    
//        // The action data for this transaction
//        Action action;
//      };
//
type Transaction struct {
  Signature Signature 
  Address Id 
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

// CommittedTransaction is an XDR Struct defines as:
//
//   struct CommittedTransaction
//      {
//        // The transaction itself
//        Transaction transaction;
//    
//        // The execution ordering of the transaction, provided by consensus
//        unsigned hyper sequenceNumber;
//    
//        // The receipt hash generated by consensus, to be compared to local execution results
//        ID receiptId;
//    
//        // The transaction merkle root after adding this transaction to the merkle tree, for validation
//        Hash currentTransactionRoot;
//    
//         // Consensus signatures
//        Signature signatures<>;
//      };
//
type CommittedTransaction struct {
  Transaction Transaction 
  SequenceNumber uint64 
  ReceiptId Id 
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

        var fmtTest = fmt.Sprint("this is a dummy usage of fmt")

