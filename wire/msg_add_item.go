package wire

import (
	"encoding/binary"
	"fmt"
	"pubsub_poc/common"
)

type MsgAddItem struct {
	Key string
	Val string
}

// A compile time check to ensure that MsgAddItem implements common.Encoder and common.Decoder.
var _ common.Encoder = (*MsgAddItem)(nil)
var _ common.Decoder = (*MsgAddItem)(nil)

// Encode encodes MsgAddItem according to the following format:
// +-----------+-----------+------------+-----------------+
// | type      | key size  | key        | value           |
// | (1 byte)  | (4 bytes) | (key size) | remaining bytes |
// +-----------+-----------+------------+-----------------+
func (m *MsgAddItem) Encode() []byte {
	size := 1 + 4 + len(m.Key) + len(m.Val)
	b := make([]byte, size)
	offset := 0

	b[offset] = byte(MsgTypeAddItem)
	offset += 1

	binary.LittleEndian.PutUint32(b[offset:], uint32(len(m.Key)))
	offset += 4

	copy(b[offset:], m.Key)
	offset += len(m.Key)

	copy(b[offset:], m.Val)

	return b
}

func (m *MsgAddItem) Decode(data []byte) {
	msg := MsgAddItem{}
	data = data[1:] // Dispose the first msg type byte.

	keyLen := binary.LittleEndian.Uint32(data[:4])

	msg.Key = string(data[4 : 4+keyLen])
	msg.Val = string(data[4+keyLen:])

	*m = msg // Override the method pointer receiver value.
}

func (m *MsgAddItem) String() string {
	return fmt.Sprintf("add [key: %v, val: %v]", m.Key, m.Val)
}
