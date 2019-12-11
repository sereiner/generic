package generic

import (
	"strconv"
)

type float64T struct {
	baseT
	val float64
}

func (T *float64T) ValueType() ValueType {
	return NumberValue
}

func (T *float64T) LastError() error {
	return nil
}

func (T *float64T) ToBool() bool {
	return T.ToFloat64() != 0
}

func (T *float64T) ToInt() int {
	return int(T.val)
}

func (T *float64T) ToInt32() int32 {
	return int32(T.val)
}

func (T *float64T) ToInt64() int64 {
	return int64(T.val)
}

func (T *float64T) ToUint() uint {
	if T.val > 0 {
		return uint(T.val)
	}
	return 0
}

func (T *float64T) ToUint32() uint32 {
	if T.val > 0 {
		return uint32(T.val)
	}
	return 0
}

func (T *float64T) ToUint64() uint64 {
	if T.val > 0 {
		return uint64(T.val)
	}
	return 0
}

func (T *float64T) ToFloat32() float32 {
	return float32(T.val)
}

func (T *float64T) ToFloat64() float64 {
	return T.val
}

func (T *float64T) ToString() string {
	return strconv.FormatFloat(T.val, 'E', -1, 64)
}

func (T *float64T) GetInterface() interface{} {
	return T.val
}
