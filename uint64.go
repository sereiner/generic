package generic

import (
	"strconv"
)

type uint64T struct {
	baseT
	val uint64
}

func (t *uint64T) LastError() error {
	return nil
}

func (t *uint64T) ValueType() ValueType {
	return NumberValue
}

func (t *uint64T) ToBool() bool {
	return t.val != 0
}

func (t *uint64T) ToInt() int {
	return int(t.val)
}

func (t *uint64T) ToInt32() int32 {
	return int32(t.val)
}

func (t *uint64T) ToInt64() int64 {
	return int64(t.val)
}

func (t *uint64T) ToUint() uint {
	return uint(t.val)
}

func (t *uint64T) ToUint32() uint32 {
	return uint32(t.val)
}

func (t *uint64T) ToUint64() uint64 {
	return t.val
}

func (t *uint64T) ToFloat32() float32 {
	return float32(t.val)
}

func (t *uint64T) ToFloat64() float64 {
	return float64(t.val)
}

func (t *uint64T) String() string {
	return strconv.FormatUint(t.val, 10)
}

func (t *uint64T) GetInterface() interface{} {
	return t.val
}
