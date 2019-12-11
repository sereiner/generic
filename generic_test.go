package generic

import (
	"fmt"
	"testing"
)

type Foo struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func Test_Generic(t *testing.T) {
	f := Struct(Foo{"jack", 22})
	fmt.Println(genericGet(f, "Age"))

	a := String("12345")
	fmt.Println(genericGet(a, 3))

	c := []T{String("123"), Int64(45), Float32(12.3)}

	fmt.Println(c)
	ReverseT(c)

	fmt.Println(c)
}

func ReverseT(s []T) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

func genericGet(arg T, key interface{}) int {
	return arg.Get(key).ToInt()
}
