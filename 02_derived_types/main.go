package main

var (
	Int    int  = 0
	IntPtr *int = &Int

	IntSlice []int         = []int{1, 2, 3}
	Slice    []interface{} = []interface{}{"a", 1, true, 1.1}

	Map map[string]interface{} = map[string]interface{}{
		"string": "Hello World!",
		"int":    1,
		"bool":   true,
	}
)

type Person struct {
	Name string
	Age  int
}

type PersonInterface interface {
	GetName() string
	GetAge() int
}
