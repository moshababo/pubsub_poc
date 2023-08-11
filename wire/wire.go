package wire

import "reflect"

type Msg any

// MsgType represents the type of message.
type MsgType byte

const (
	MsgTypeAddItem     MsgType = 1
	MsgTypeRemoveItem  MsgType = 2
	MsgTypeGetItem     MsgType = 3
	MsgTypeGetAllItems MsgType = 4
)

var TypeToStructType = map[MsgType]reflect.Type{
	MsgTypeAddItem:     reflect.TypeOf(MsgAddItem{}),
	MsgTypeRemoveItem:  reflect.TypeOf(MsgRemoveItem{}),
	MsgTypeGetItem:     reflect.TypeOf(MsgGetItem{}),
	MsgTypeGetAllItems: reflect.TypeOf(MsgGetAllItems{}),
}
