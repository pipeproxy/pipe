package codec

type Encoder interface {
	Encode(v interface{}) error
	Bytes() []byte
}

type Decoder interface {
	Decode(v interface{}) error
	Bytes() []byte
}
