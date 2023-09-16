package main

func casting() {
	println("Casting:")
	var a int = 10
	println(a)
	var b float64 = float64(a)
	println(b)
	var c uint = uint(b)
	println(c)
	println()
}

func assertion() {
	println("Assertion:")
	var a interface{} = 10
	println(a)
	b := a.(int)
	println(b)
	println()
}

func main() {
	casting()
	assertion()
}
