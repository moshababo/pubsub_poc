package wire

import (
	"fmt"
	"pubsub_poc/common"
)

type MsgRemoveItem struct {
	Key string
}

// A compile time check to ensure that MsgRemoveItem implements common.Encoder and common.Decoder.
var _ common.Encoder = (*MsgRemoveItem)(nil)
var _ common.Decoder = (*MsgRemoveItem)(nil)

// Encode encodes MsgRemoveItem according to the following format:
// +-----------+-----------------+
// | type      | key             |
// | (1 byte)  | remaining bytes |
// +-----------+-----------------+
func (m *MsgRemoveItem) Encode() []byte {
	size := 1 + len(m.Key)
	b := make([]byte, size)

	b[0] = byte(MsgTypeRemoveItem)
	copy(b[1:], m.Key)

	return b
}

func (m *MsgRemoveItem) Decode(data []byte) {
	msg := MsgRemoveItem{}
	data = data[1:] // Dispose the first msg type byte.

	msg.Key = string(data)

	*m = msg // Override the method pointer receiver value.
}

func (m *MsgRemoveItem) String() string {
	return fmt.Sprintf("remove [key: %v]", m.Key)
}
