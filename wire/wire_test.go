package wire

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	key = "key1"
	val = "val1"
)

func TestMsgAddItem(t *testing.T) {
	r := require.New(t)

	msg1 := &MsgAddItem{Key: key, Val: val}
	b := msg1.Encode()

	msg2 := new(MsgAddItem)
	msg2.Decode(b)

	r.Equal(msg1, msg2)
}

func TestMsgRemoveItem(t *testing.T) {
	r := require.New(t)

	msg1 := &MsgRemoveItem{Key: key}
	b := msg1.Encode()

	msg2 := new(MsgRemoveItem)
	msg2.Decode(b)

	r.Equal(msg1, msg2)
}

func TestMsgGetItem(t *testing.T) {
	r := require.New(t)

	msg1 := &MsgGetItem{Key: key}
	b := msg1.Encode()

	msg2 := new(MsgGetItem)
	msg2.Decode(b)

	r.Equal(msg1, msg2)
}

func TestMsgGetAllItems(t *testing.T) {
	r := require.New(t)

	msg1 := &MsgGetAllItems{}
	b := msg1.Encode()

	msg2 := new(MsgGetAllItems)
	msg2.Decode(b)

	r.Equal(msg1, msg2)
}
