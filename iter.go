package generic

type ValueType int

const (
	// InvalidValue invalid JSON element
	InvalidValue ValueType = iota
	// StringValue JSON element "string"
	StringValue
	// NumberValue JSON element 100 or 0.10
	NumberValue
	// NilValue JSON element null
	NilValue
	// BoolValue JSON element true or false
	BoolValue
	// ArrayValue JSON element []
	ArrayValue
	// ObjectValue JSON element {}
	ObjectValue
)

func (v ValueType) String() string {
	switch v {
	case InvalidValue:
		return "invalid JSON element"
	case StringValue:
		return "string"
	case NumberValue:
		return "number"
	case NilValue:
		return "nil"
	case BoolValue:
		return "bool"
	case ArrayValue:
		return "array"
	case ObjectValue:
		return "object or map"
	}
	return ""
}

var hexDigits []byte
var valueTypes []ValueType

func init() {
	hexDigits = make([]byte, 256)
	for i := 0; i < len(hexDigits); i++ {
		hexDigits[i] = 255
	}
	for i := '0'; i <= '9'; i++ {
		hexDigits[i] = byte(i - '0')
	}
	for i := 'a'; i <= 'f'; i++ {
		hexDigits[i] = byte((i - 'a') + 10)
	}
	for i := 'A'; i <= 'F'; i++ {
		hexDigits[i] = byte((i - 'A') + 10)
	}
	valueTypes = make([]ValueType, 256)
	for i := 0; i < len(valueTypes); i++ {
		valueTypes[i] = InvalidValue
	}
	valueTypes['"'] = StringValue
	valueTypes['-'] = NumberValue
	valueTypes['0'] = NumberValue
	valueTypes['1'] = NumberValue
	valueTypes['2'] = NumberValue
	valueTypes['3'] = NumberValue
	valueTypes['4'] = NumberValue
	valueTypes['5'] = NumberValue
	valueTypes['6'] = NumberValue
	valueTypes['7'] = NumberValue
	valueTypes['8'] = NumberValue
	valueTypes['9'] = NumberValue
	valueTypes['t'] = BoolValue
	valueTypes['f'] = BoolValue
	valueTypes['n'] = NilValue
	valueTypes['['] = ArrayValue
	valueTypes['{'] = ObjectValue
}
