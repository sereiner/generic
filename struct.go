package generic

import (
	"reflect"

	"github.com/modern-go/reflect2"

	jsoniter "github.com/json-iterator/go"
)

type structT struct {
	baseT
	err error
	val reflect.Value
}

func wrapStruct(val interface{}) *structT {
	if reflect2.TypeOf(val).Kind() != reflect.Struct {
		panic("wrapArray param must be struct!")
	}

	return &structT{baseT{}, nil, reflect.ValueOf(val)}
}

func (t *structT) ValueType() ValueType {
	return ObjectValue
}

func (t *structT) LastError() error {
	return t.err
}

func (t *structT) ToBool() bool {
	return t.val.NumField() != 0
}

func (t *structT) ToInt() int {
	return 0
}

func (t *structT) ToInt32() int32 {
	return 0
}

func (t *structT) ToInt64() int64 {
	return 0
}

func (t *structT) ToUint() uint {
	return 0
}

func (t *structT) ToUint32() uint32 {
	return 0
}

func (t *structT) ToUint64() uint64 {
	return 0
}

func (t *structT) ToFloat32() float32 {
	return 0
}

func (t *structT) ToFloat64() float64 {
	return 0
}

func (t *structT) ToString() string {
	str, err := jsoniter.MarshalToString(t.val.Interface())
	t.err = err
	return str
}

func (t *structT) Get(path ...interface{}) T {
	if len(path) == 0 {
		return t
	}
	switch firstPath := path[0].(type) {
	case string:
		field := t.val.FieldByName(firstPath)
		if !field.IsValid() {
			return newInvalidT(path)
		}
		return Wrap(field.Interface())
	case int32:
		if '*' == firstPath {
			mappedAll := map[string]T{}
			for i := 0; i < t.val.NumField(); i++ {
				field := t.val.Field(i)
				if field.CanInterface() {
					mapped := Wrap(field.Interface()).Get(path[1:]...)
					if mapped.ValueType() != InvalidValue {
						mappedAll[t.val.Type().Field(i).Name] = mapped
					}
				}
			}
			return wrapMap(mappedAll)
		}
		return newInvalidT(path)
	default:
		return newInvalidT(path)
	}
}

func (t *structT) Keys() []string {
	keys := make([]string, 0, t.val.NumField())
	for i := 0; i < t.val.NumField(); i++ {
		keys = append(keys, t.val.Type().Field(i).Name)
	}
	return keys
}

func (t *structT) Size() int {
	return t.val.NumField()
}

func (t *structT) GetInterface() interface{} {
	return t.val.Interface()
}
