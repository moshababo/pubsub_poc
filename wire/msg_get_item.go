package wire

import (
	"fmt"
	"pubsub_poc/common"
)

type MsgGetItem struct {
	Key string
}

// A compile time check to ensure that MsgGetItem implements common.Encoder and common.Decoder.
var _ common.Encoder = (*MsgGetItem)(nil)
var _ common.Decoder = (*MsgGetItem)(nil)

// Encode encodes MsgGetItem according to the following format:
// +-----------+-----------------+
// | type      | key             |
// | (1 byte)  | remaining bytes |
// +-----------+-----------------+
func (m *MsgGetItem) Encode() []byte {
	size := 1 + len(m.Key)
	b := make([]byte, size)

	b[0] = byte(MsgTypeGetItem)
	copy(b[1:], m.Key)

	return b
}

func (m *MsgGetItem) Decode(data []byte) {
	msg := MsgGetItem{}
	data = data[1:] // Dispose the first msg type byte.

	msg.Key = string(data)

	*m = msg // Override the method pointer receiver value.
}

func (m *MsgGetItem) String() string {
	return fmt.Sprintf("get [key: %v]", m.Key)
}
