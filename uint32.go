package generic

import (
	"strconv"
)

type uint32T struct {
	baseT
	val uint32
}

func (t *uint32T) LastError() error {
	return nil
}

func (t *uint32T) ValueType() ValueType {
	return NumberValue
}

func (t *uint32T) ToBool() bool {
	return t.val != 0
}

func (t *uint32T) ToInt() int {
	return int(t.val)
}

func (t *uint32T) ToInt32() int32 {
	return int32(t.val)
}

func (t *uint32T) ToInt64() int64 {
	return int64(t.val)
}

func (t *uint32T) ToUint() uint {
	return uint(t.val)
}

func (t *uint32T) ToUint32() uint32 {
	return t.val
}

func (t *uint32T) ToUint64() uint64 {
	return uint64(t.val)
}

func (t *uint32T) ToFloat32() float32 {
	return float32(t.val)
}

func (t *uint32T) ToFloat64() float64 {
	return float64(t.val)
}

func (t *uint32T) String() string {
	return strconv.FormatInt(int64(t.val), 10)
}

func (t *uint32T) GetInterface() interface{} {
	return t.val
}
