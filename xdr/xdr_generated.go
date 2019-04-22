// Package xdr is generated from:
//
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

// Call is an XDR Struct defines as:
//
//   struct Call
//      {
//        string function<256>;
//    
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
//        ID channelID;    
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
//        Signature signature;
//    
//        ID address;
//    
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

        var fmtTest = fmt.Sprint("this is a dummy usage of fmt")

