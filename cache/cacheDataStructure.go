package cache


type Frame struct {
	key string
	keySize int
	value []byte
	valueSize int
}

func (f *Frame) SetData(key string, value []byte)  {
	f.key = key
	f.keySize = len(key)
	f.value = value
	f.valueSize = len(value)
}

func (f *Frame)GetKey() string {
	return f.key
}

func (f *Frame)GetKeySize() int {
	return f.keySize
}

func (f *Frame)GetValueSize() int {
	return f.valueSize
}

func NewFrame(key string, value []byte) *Frame {
	keySize := len(key)
	valueSize := len(value)
	return &Frame{key: key, keySize: keySize, value: value, valueSize: valueSize}
}

