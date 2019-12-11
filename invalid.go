package generic

import "fmt"

type invalidT struct {
	baseT
	err error
}

func newInvalidT(path []interface{}) *invalidT {
	return &invalidT{baseT{}, fmt.Errorf("%v not found", path)}
}

func (T *invalidT) LastError() error {
	return T.err
}

func (T *invalidT) ValueType() ValueType {
	return InvalidValue
}

func (T *invalidT) ToBool() bool {
	return false
}

func (T *invalidT) ToInt() int {
	return 0
}

func (T *invalidT) ToInt32() int32 {
	return 0
}

func (T *invalidT) ToInt64() int64 {
	return 0
}

func (T *invalidT) ToUint() uint {
	return 0
}

func (T *invalidT) ToUint32() uint32 {
	return 0
}

func (T *invalidT) ToUint64() uint64 {
	return 0
}

func (T *invalidT) ToFloat32() float32 {
	return 0
}

func (T *invalidT) ToFloat64() float64 {
	return 0
}

func (T *invalidT) ToString() string {
	return ""
}

func (T *invalidT) Get(path ...interface{}) T {
	if T.err == nil {
		return &invalidT{baseT{}, fmt.Errorf("get %v from invalid", path)}
	}
	return &invalidT{baseT{}, fmt.Errorf("%v, get %v from invalid", T.err, path)}
}

func (T *invalidT) GetInterface() interface{} {
	return nil
}
