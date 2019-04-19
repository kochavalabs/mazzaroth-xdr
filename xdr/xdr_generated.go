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

// ActionType is an XDR Enum defines as:
//
//   enum ActionType
//      {
//        ACTION_TYPE_CALL = 0,
//        ACTION_TYPE_UPDATE = 1
//      };
//
type ActionType int32
const (
  ActionTypeActionTypeCall ActionType = 0
  ActionTypeActionTypeUpdate ActionType = 1
)
var actionTypeMap = map[int32]string{
  0: "ActionTypeActionTypeCall",
  1: "ActionTypeActionTypeUpdate",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ActionType
func (e ActionType) ValidEnum(v int32) bool {
  _, ok := actionTypeMap[v]
  return ok
}
// String returns the name of `e`
func (e ActionType) String() string {
  name, _ := actionTypeMap[int32(e)]
  return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ActionType) MarshalBinary() ([]byte, error) {
  b := new(bytes.Buffer)
  _, err := Marshal(b, s)
  return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ActionType) UnmarshalBinary(inp []byte) error {
  _, err := Unmarshal(bytes.NewReader(inp), s)
  return err
}

var (
  _ encoding.BinaryMarshaler   = (*ActionType)(nil)
  _ encoding.BinaryUnmarshaler = (*ActionType)(nil)
)

// Action is an XDR Union defines as:
//
//   union Action switch (ActionType type)
//      {
//        case ACTION_TYPE_CALL:
//            Call call;
//        case ACTION_TYPE_UPDATE:
//            Update update;
//      };
//
type Action struct{
  Type ActionType
  Call *Call 
  Update *Update 
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u Action) SwitchFieldName() string {
  return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of Action
func (u Action) ArmForSwitch(sw int32) (string, bool) {
switch ActionType(sw) {
    case ActionTypeActionTypeCall:
      return "Call", true
    case ActionTypeActionTypeUpdate:
      return "Update", true
}
return "-", false
}

// NewAction creates a new  Action.
func NewAction(aType ActionType, value interface{}) (result Action, err error) {
  result.Type = aType
switch ActionType(aType) {
    case ActionTypeActionTypeCall:
                  tv, ok := value.(Call)
            if !ok {
              err = fmt.Errorf("invalid value, must be Call")
              return
            }
            result.Call = &tv
    case ActionTypeActionTypeUpdate:
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
func (u Action) MustCall() Call {
  val, ok := u.GetCall()

  if !ok {
    panic("arm Call is not set")
  }

  return val
}

// GetCall retrieves the Call value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Action) GetCall() (result Call, ok bool) {
  armName, _ := u.ArmForSwitch(int32(u.Type))

  if armName == "Call" {
    result = *u.Call
    ok = true
  }

  return
}
// MustUpdate retrieves the Update value from the union,
// panicing if the value is not set.
func (u Action) MustUpdate() Update {
  val, ok := u.GetUpdate()

  if !ok {
    panic("arm Update is not set")
  }

  return val
}

// GetUpdate retrieves the Update value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Action) GetUpdate() (result Update, ok bool) {
  armName, _ := u.ArmForSwitch(int32(u.Type))

  if armName == "Update" {
    result = *u.Update
    ok = true
  }

  return
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

