package generic

import (
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

type sliceT struct {
	baseT
	val reflect.Value
}

func Slice(val interface{}) *sliceT {
	return &sliceT{baseT{}, reflect.ValueOf(val)}
}

func (t *sliceT) ValueType() ValueType {
	return ArrayValue
}

func (t *sliceT) LastError() error {
	return nil
}

func (t *sliceT) ToBool() bool {
	return t.val.Len() != 0
}

func (t *sliceT) ToInt() int {
	if t.val.Len() == 0 {
		return 0
	}
	return 1
}

func (t *sliceT) ToInt32() int32 {
	if t.val.Len() == 0 {
		return 0
	}
	return 1
}

func (t *sliceT) ToInt64() int64 {
	if t.val.Len() == 0 {
		return 0
	}
	return 1
}

func (t *sliceT) ToUint() uint {
	if t.val.Len() == 0 {
		return 0
	}
	return 1
}

func (t *sliceT) ToUint32() uint32 {
	if t.val.Len() == 0 {
		return 0
	}
	return 1
}

func (t *sliceT) ToUint64() uint64 {
	if t.val.Len() == 0 {
		return 0
	}
	return 1
}

func (t *sliceT) ToFloat32() float32 {
	if t.val.Len() == 0 {
		return 0
	}
	return 1
}

func (t *sliceT) ToFloat64() float64 {
	if t.val.Len() == 0 {
		return 0
	}
	return 1
}

func (t *sliceT) String() string {
	str, _ := jsoniter.MarshalToString(t.val.Interface())
	return str
}

func (t *sliceT) Get(path ...interface{}) T {
	if len(path) == 0 {
		return t
	}
	switch firstPath := path[0].(type) {
	case int:
		if firstPath < 0 || firstPath >= t.val.Len() {
			return newInvalidT(path)
		}
		return Wrap(t.val.Index(firstPath).Interface())
	case int32:
		if '*' == firstPath {
			mappedAll := make([]T, 0)
			for i := 0; i < t.val.Len(); i++ {
				mapped := Wrap(t.val.Index(i).Interface()).Get(path[1:]...)
				if mapped.ValueType() != InvalidValue {
					mappedAll = append(mappedAll, mapped)
				}
			}
			return Slice(mappedAll)
		}
		return newInvalidT(path)
	default:
		return newInvalidT(path)
	}
}

func (t *sliceT) Size() int {
	return t.val.Len()
}

func (t *sliceT) GetInterface() interface{} {
	return t.val.Interface()
}
