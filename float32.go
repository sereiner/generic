package generic

import (
	"strconv"
)

type float32T struct {
	baseT
	val float64
}

func (T *float32T) ValueType() ValueType {
	return NumberValue
}

func (T *float32T) LastError() error {
	return nil
}

func (T *float32T) ToBool() bool {
	return T.ToFloat64() != 0
}

func (T *float32T) ToInt() int {
	return int(T.val)
}

func (T *float32T) ToInt32() int32 {
	return int32(T.val)
}

func (T *float32T) ToInt64() int64 {
	return int64(T.val)
}

func (T *float32T) ToUint() uint {
	if T.val > 0 {
		return uint(T.val)
	}
	return 0
}

func (T *float32T) ToUint32() uint32 {
	if T.val > 0 {
		return uint32(T.val)
	}
	return 0
}

func (T *float32T) ToUint64() uint64 {
	if T.val > 0 {
		return uint64(T.val)
	}
	return 0
}

func (T *float32T) ToFloat32() float32 {
	return float32(T.val)
}

func (T *float32T) ToFloat64() float64 {
	return T.val
}

func (T *float32T) ToString() string {
	return strconv.FormatFloat(T.val, 'E', -1, 64)
}

func (T *float32T) GetInterface() interface{} {
	return T.val
}
