package generic

import (
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

type mapT struct {
	baseT
	err error
	val reflect.Value
}

func wrapMap(val interface{}) *mapT {
	return &mapT{baseT{}, nil, reflect.ValueOf(val)}
}

func (t *mapT) ValueType() ValueType {
	return ObjectValue
}

func (t *mapT) LastError() error {
	return t.err
}

func (t *mapT) ToBool() bool {
	return true
}

func (t *mapT) ToInt() int {
	return 0
}

func (t *mapT) ToInt32() int32 {
	return 0
}

func (t *mapT) ToInt64() int64 {
	return 0
}

func (t *mapT) ToUint() uint {
	return 0
}

func (t *mapT) ToUint32() uint32 {
	return 0
}

func (t *mapT) ToUint64() uint64 {
	return 0
}

func (t *mapT) ToFloat32() float32 {
	return 0
}

func (t *mapT) ToFloat64() float64 {
	return 0
}

func (t *mapT) ToString() string {
	str, err := jsoniter.MarshalToString(t.val.Interface())
	t.err = err
	return str
}

func (t *mapT) Get(path ...interface{}) T {
	if len(path) == 0 {
		return t
	}
	switch firstPath := path[0].(type) {
	case int32:
		if '*' == firstPath {
			mappedAll := map[string]T{}
			for _, key := range t.val.MapKeys() {
				keyAsStr := key.String()
				element := Wrap(t.val.MapIndex(key).Interface())
				mapped := element.Get(path[1:]...)
				if mapped.ValueType() != InvalidValue {
					mappedAll[keyAsStr] = mapped
				}
			}
			return wrapMap(mappedAll)
		}
		return newInvalidT(path)
	default:
		value := t.val.MapIndex(reflect.ValueOf(firstPath))
		if !value.IsValid() {
			return newInvalidT(path)
		}
		return Wrap(value.Interface())
	}
}

func (t *mapT) Keys() []string {
	keys := make([]string, 0, t.val.Len())
	for _, key := range t.val.MapKeys() {
		keys = append(keys, key.String())
	}
	return keys
}

func (t *mapT) Size() int {
	return t.val.Len()
}

func (t *mapT) GetInterface() interface{} {
	return t.val.Interface()
}
