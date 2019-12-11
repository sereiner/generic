package generic

import (
	"fmt"
	"strconv"
)

type stringT struct {
	baseT
	val string
}

func (t *stringT) Get(path ...interface{}) T {
	if len(path) == 0 {
		return t
	}
	switch i := path[0].(type) {
	case int:
		return String(string(t.val[i]))
	}

	return &invalidT{baseT{}, fmt.Errorf("GetIndex %v from simple value", path)}
}

func (t *stringT) ValueType() ValueType {
	return StringValue
}

func (t *stringT) LastError() error {
	return nil
}

func (t *stringT) ToBool() bool {
	str := t.String()
	if str == "0" {
		return false
	}
	for _, c := range str {
		switch c {
		case ' ', '\n', '\r', '\t':
		default:
			return true
		}
	}
	return false
}

func (t *stringT) ToInt() int {
	return int(t.ToInt64())

}

func (t *stringT) ToInt32() int32 {
	return int32(t.ToInt64())
}

func (t *stringT) ToInt64() int64 {
	if t.val == "" {
		return 0
	}

	flag := 1
	startPos := 0
	endPos := 0
	if t.val[0] == '+' || t.val[0] == '-' {
		startPos = 1
	}

	if t.val[0] == '-' {
		flag = -1
	}

	for i := startPos; i < len(t.val); i++ {
		if t.val[i] >= '0' && t.val[i] <= '9' {
			endPos = i + 1
		} else {
			break
		}
	}
	parsed, _ := strconv.ParseInt(t.val[startPos:endPos], 10, 64)
	return int64(flag) * parsed
}

func (t *stringT) ToUint() uint {
	return uint(t.ToUint64())
}

func (t *stringT) ToUint32() uint32 {
	return uint32(t.ToUint64())
}

func (t *stringT) ToUint64() uint64 {
	if t.val == "" {
		return 0
	}

	startPos := 0
	endPos := 0

	if t.val[0] == '-' {
		return 0
	}
	if t.val[0] == '+' {
		startPos = 1
	}

	for i := startPos; i < len(t.val); i++ {
		if t.val[i] >= '0' && t.val[i] <= '9' {
			endPos = i + 1
		} else {
			break
		}
	}
	parsed, _ := strconv.ParseUint(t.val[startPos:endPos], 10, 64)
	return parsed
}

func (t *stringT) ToFloat32() float32 {
	return float32(t.ToFloat64())
}

func (t *stringT) ToFloat64() float64 {
	if len(t.val) == 0 {
		return 0
	}

	// first char invalid
	if t.val[0] != '+' && t.val[0] != '-' && (t.val[0] > '9' || t.val[0] < '0') {
		return 0
	}

	// extract valid num expression from string
	// eg 123true => 123, -12.12xxa => -12.12
	endPos := 1
	for i := 1; i < len(t.val); i++ {
		if t.val[i] == '.' || t.val[i] == 'e' || t.val[i] == 'E' || t.val[i] == '+' || t.val[i] == '-' {
			endPos = i + 1
			continue
		}

		// end position is the first char which is not digit
		if t.val[i] >= '0' && t.val[i] <= '9' {
			endPos = i + 1
		} else {
			endPos = i
			break
		}
	}
	parsed, _ := strconv.ParseFloat(t.val[:endPos], 64)
	return parsed
}

func (t *stringT) String() string {
	return t.val
}

func (t *stringT) Size() int {
	return len(t.val)
}

func (t *stringT) GetInterface() interface{} {
	return t.val
}
