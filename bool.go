package generic

import "strconv"

type boolT struct {
	baseT
	val bool
}

func (b *boolT) LastError() error {
	return nil
}

func (b *boolT) ToBool() bool {
	return b.val
}

func (b *boolT) ToInt() int {
	if b.val {
		return 1
	}
	return 0
}

func (b *boolT) ToInt32() int32 {
	if b.val {
		return 1
	}
	return 0
}

func (b *boolT) ToInt64() int64 {
	if b.val {
		return 1
	}
	return 0
}

func (b *boolT) ToUint() uint {
	if b.val {
		return 1
	}
	return 0
}

func (b *boolT) ToUint32() uint32 {
	if b.val {
		return 1
	}
	return 0
}

func (b *boolT) ToUint64() uint64 {
	if b.val {
		return 1
	}
	return 0
}

func (b *boolT) ToFloat32() float32 {
	if b.val {
		return 1
	}
	return 0
}

func (b *boolT) ToFloat64() float64 {
	if b.val {
		return 1
	}
	return 0
}

func (b *boolT) ToString() string {
	return strconv.FormatBool(b.val)
}

func (b *boolT) GetInterface() interface{} {
	return b.val
}

func (b *boolT) ValueType() ValueType {
	return BoolValue
}
