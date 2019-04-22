          // Package xdr is generated from:
          //
          //  ./idl/block.x
//  ./idl/common.x
//  ./idl/event.x
//  ./idl/receipt.x
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

// BlockHeader is an XDR Struct defines as:
//
//   struct BlockHeader
//      {
//    
//        string timestamp<256>; 
//    
//        uint64 blockHeight;
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
  BlockHeight Uint64 
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

// Uint64 is an XDR Typedef defines as:
//
//   typedef unsigned hyper uint64;
//
type Uint64 uint64

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Uint64) MarshalBinary() ([]byte, error) {
  b := new(bytes.Buffer)
  _, err := Marshal(b, s)
  return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Uint64) UnmarshalBinary(inp []byte) error {
  _, err := Unmarshal(bytes.NewReader(inp), s)
  return err
}

var (
  _ encoding.BinaryMarshaler   = (*Uint64)(nil)
  _ encoding.BinaryUnmarshaler = (*Uint64)(nil)
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

// Call is an XDR Struct defines as:
//
//   struct Call
//      {
//        // Contract function to execute.
//        string function<256>;
//    
//        // Parameters to the contract function. The serialization format is defined
//        // by the contract itself.
//        opaque parameters<>;
//      };
//
type Call struct {
  Function string `xdrmaxsize:"256"`
  Parameters []byte 
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
//        uint64 nonce;
//    
//        ActionCategory category;
//    
//      };
//
type Action struct {
  ChannelId Id 
  Nonce Uint64 
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
//        uint64 sequenceNumber;
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
  SequenceNumber Uint64 
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

