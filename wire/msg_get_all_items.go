package wire

import "pubsub_poc/common"

type MsgGetAllItems struct{}

// A compile time check to ensure that MsgGetAllItems implements common.Encoder and common.Decoder.
var _ common.Encoder = (*MsgGetAllItems)(nil)
var _ common.Decoder = (*MsgGetAllItems)(nil)

// Encode encodes MsgGetAllItems according to the following format:
// +-----------+
// | type      |
// | (1 byte)  |
// +-----------+
func (m *MsgGetAllItems) Encode() []byte {
	size := 1
	b := make([]byte, size)

	b[0] = byte(MsgTypeGetAllItems)

	return b
}

func (m *MsgGetAllItems) Decode([]byte) {}

func (m *MsgGetAllItems) String() string {
	return "getAll"
}
