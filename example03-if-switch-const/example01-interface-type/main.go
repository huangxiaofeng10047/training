package main

var (
	i interface{}
)

func convert(i interface{}) {
	switch t := i.(type) {
	case int:
		println("i is interger", t)
	case string:
		println("i is string", t)
	case float64:
		println("i is float64", t)
	case float32:
		println("i is float32", t)
	case int64:
		println("i is int64", t)
	default:
		println("type not found")
	}
}

func main() {
	i = 100
	convert(i)
	i = float64(45.55)
	convert(i)
	i = "foo"
	convert(i)
	convert(float32(10.0))
	i = 100000000
	convert(i)
	convert(int64(1000))
}
