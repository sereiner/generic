package generic

import (
	"strconv"
)

type int64T struct {
	baseT
	val int64
}

func (T *int64T) LastError() error {
	return nil
}

func (T *int64T) ValueType() ValueType {
	return NumberValue
}

func (T *int64T) ToBool() bool {
	return T.val != 0
}

func (T *int64T) ToInt() int {
	return int(T.val)
}

func (T *int64T) ToInt32() int32 {
	return int32(T.val)
}

func (T *int64T) ToInt64() int64 {
	return T.val
}

func (T *int64T) ToUint() uint {
	return uint(T.val)
}

func (T *int64T) ToUint32() uint32 {
	return uint32(T.val)
}

func (T *int64T) ToUint64() uint64 {
	return uint64(T.val)
}

func (T *int64T) ToFloat32() float32 {
	return float32(T.val)
}

func (T *int64T) ToFloat64() float64 {
	return float64(T.val)
}

func (T *int64T) String() string {
	return strconv.FormatInt(T.val, 10)
}

func (T *int64T) GetInterface() interface{} {
	return T.val
}
