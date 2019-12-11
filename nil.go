package generic

type nilT struct {
	baseT
}

func (T *nilT) LastError() error {
	return nil
}

func (T *nilT) ValueType() ValueType {
	return NilValue
}

func (T *nilT) ToBool() bool {
	return false
}

func (T *nilT) ToInt() int {
	return 0
}

func (T *nilT) ToInt32() int32 {
	return 0
}

func (T *nilT) ToInt64() int64 {
	return 0
}

func (T *nilT) ToUint() uint {
	return 0
}

func (T *nilT) ToUint32() uint32 {
	return 0
}

func (T *nilT) ToUint64() uint64 {
	return 0
}

func (T *nilT) ToFloat32() float32 {
	return 0
}

func (T *nilT) ToFloat64() float64 {
	return 0
}

func (T *nilT) ToString() string {
	return ""
}

func (T *nilT) GetInterface() interface{} {
	return nil
}
