package generic

import (
	"reflect"

	"github.com/modern-go/reflect2"

	jsoniter "github.com/json-iterator/go"
)

type arrayT struct {
	baseT
	val reflect.Value
}

func wrapArray(val interface{}) *arrayT {
	if reflect2.TypeOf(val).Kind() != reflect.Array {
		panic("wrapArray param must be array!")
	}
	return &arrayT{baseT{}, reflect.ValueOf(val)}
}

func (t *arrayT) ValueType() ValueType {
	return ArrayValue
}

func (t *arrayT) LastError() error {
	return nil
}

func (t *arrayT) ToBool() bool {
	return t.val.Len() != 0
}

func (t *arrayT) ToInt() int {
	if t.val.Len() == 0 {
		return 0
	}
	return 1
}

func (t *arrayT) ToInt32() int32 {
	if t.val.Len() == 0 {
		return 0
	}
	return 1
}

func (t *arrayT) ToInt64() int64 {
	if t.val.Len() == 0 {
		return 0
	}
	return 1
}

func (t *arrayT) ToUint() uint {
	if t.val.Len() == 0 {
		return 0
	}
	return 1
}

func (t *arrayT) ToUint32() uint32 {
	if t.val.Len() == 0 {
		return 0
	}
	return 1
}

func (t *arrayT) ToUint64() uint64 {
	if t.val.Len() == 0 {
		return 0
	}
	return 1
}

func (t *arrayT) ToFloat32() float32 {
	if t.val.Len() == 0 {
		return 0
	}
	return 1
}

func (t *arrayT) ToFloat64() float64 {
	if t.val.Len() == 0 {
		return 0
	}
	return 1
}

func (t *arrayT) ToString() string {
	str, _ := jsoniter.MarshalToString(t.val.Interface())
	return str
}

func (t *arrayT) Get(path ...interface{}) T {
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
			return wrapArray(mappedAll)
		}
		return newInvalidT(path)
	default:
		return newInvalidT(path)
	}
}

func (t *arrayT) Size() int {
	return t.val.Len()
}

func (t *arrayT) GetInterface() interface{} {
	return t.val.Interface()
}
