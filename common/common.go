package common

type Encoder interface {
	Encode() []byte
}

type Decoder interface {
	Decode([]byte)
}
