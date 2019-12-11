package generic

import (
	"strconv"
)

type int32T struct {
	baseT
	val int32
}

func (T *int32T) LastError() error {
	return nil
}

func (T *int32T) ValueType() ValueType {
	return NumberValue
}

func (T *int32T) ToBool() bool {
	return T.val != 0
}

func (T *int32T) ToInt() int {
	return int(T.val)
}

func (T *int32T) ToInt32() int32 {
	return T.val
}

func (T *int32T) ToInt64() int64 {
	return int64(T.val)
}

func (T *int32T) ToUint() uint {
	return uint(T.val)
}

func (T *int32T) ToUint32() uint32 {
	return uint32(T.val)
}

func (T *int32T) ToUint64() uint64 {
	return uint64(T.val)
}

func (T *int32T) ToFloat32() float32 {
	return float32(T.val)
}

func (T *int32T) ToFloat64() float64 {
	return float64(T.val)
}

func (T *int32T) String() string {
	return strconv.FormatInt(int64(T.val), 10)
}

func (T *int32T) GetInterface() interface{} {
	return T.val
}
