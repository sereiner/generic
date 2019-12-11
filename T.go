package generic

import (
	"fmt"

	"github.com/modern-go/reflect2"

	"reflect"
	"strconv"
)

const ptrSize = 32 << uintptr(^uintptr(0)>>63)

type T interface {
	LastError() error
	ValueType() ValueType
	ToBool() bool
	ToInt() int
	ToInt32() int32
	ToInt64() int64
	ToUint() uint
	ToUint32() uint32
	ToUint64() uint64
	ToFloat32() float32
	ToFloat64() float64
	String() string
	Get(path ...interface{}) T
	Size() int
	Keys() []string
	GetInterface() interface{}
}

type baseT struct{}

func (T *baseT) Get(path ...interface{}) T {
	return &invalidT{baseT{}, fmt.Errorf("GetIndex %v from simple value", path)}
}

func (T *baseT) Size() int {
	return 0
}

func (T *baseT) Keys() []string {
	return []string{}
}

func Int32(val int32) T {
	return &int32T{baseT{}, val}
}

func Int64(val int64) T {
	return &int64T{baseT{}, val}
}

func Uint32(val uint32) T {
	return &uint32T{baseT{}, val}
}

func Uint64(val uint64) T {
	return &uint64T{baseT{}, val}
}

func Float64(val float64) T {
	return &float64T{baseT{}, val}
}
func Float32(val float32) T {
	return &float32T{baseT{}, val}
}

func String(val string) T {
	return &stringT{baseT{}, val}
}

func Bool(b bool) T {
	return &boolT{baseT{}, b}
}

func Wrap(val interface{}) T {
	if val == nil {
		return &nilT{}
	}
	if asT, isT := val.(T); isT {
		return asT
	}
	typ := reflect2.TypeOf(val)
	switch typ.Kind() {
	case reflect.Array:
		return Array(val)
	case reflect.Slice:
		return Slice(val)
	case reflect.Struct:
		return Struct(val)
	case reflect.Map:
		return Map(val)
	case reflect.String:
		return String(val.(string))
	case reflect.Int:
		if strconv.IntSize == 32 {
			return Int32(int32(val.(int)))
		}
		return Int64(int64(val.(int)))
	case reflect.Int8:
		return Int32(int32(val.(int8)))
	case reflect.Int16:
		return Int32(int32(val.(int16)))
	case reflect.Int32:
		return Int32(val.(int32))
	case reflect.Int64:
		return Int64(val.(int64))
	case reflect.Uint:
		if strconv.IntSize == 32 {
			return Uint32(uint32(val.(uint)))
		}
		return Uint64(uint64(val.(uint)))
	case reflect.Uintptr:
		if ptrSize == 32 {
			return Uint32(uint32(val.(uintptr)))
		}
		return Uint64(uint64(val.(uintptr)))
	case reflect.Uint8:
		return Uint32(uint32(val.(uint8)))
	case reflect.Uint16:
		return Uint32(uint32(val.(uint16)))
	case reflect.Uint32:
		return Uint32(uint32(val.(uint32)))
	case reflect.Uint64:
		return Uint64(val.(uint64))
	case reflect.Float32:
		return Float32(val.(float32))
	case reflect.Float64:
		return Float64(val.(float64))
	case reflect.Bool:
		return Bool(val.(bool))

	}
	return &invalidT{baseT{}, fmt.Errorf("unsupported type: %v", typ)}
}
