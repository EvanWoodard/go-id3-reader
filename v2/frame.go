package v2

// Frame ...
type Frame interface {
	Id() string
	Size() uint
	StatusFlags() byte
	FormatFlags() byte
	String() string
	Bytes() []byte
}
