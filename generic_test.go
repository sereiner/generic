package generic

import (
	"fmt"
	"testing"
)

type Foo struct {
	Name string `json:"name"`
}

func Test_Generic(t *testing.T) {
	f := wrapStruct(Foo{"jack"})
	fmt.Println(genericGet(f, "Name"))

	a := WrapString("hello")
	fmt.Println(genericGet(a, 0))
}

func genericGet(arg T, key interface{}) int {
	return arg.Get(key).Size()
}
