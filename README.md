# generic
golang 泛型对象的表示(虽然go目前没有泛型)


##### 示例
```go

func main() {
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

```