package main

/*
#include "hello.c"
*/
import "C"
import "fmt"

func main() {
	Cstr, err := C.HelloWorld()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[%T] Cstr: \"%v\"; *Cstr: \"%v\"\n", Cstr, Cstr, *Cstr)
	fmt.Printf("[%T] %v\n", C.GoString(Cstr), C.GoString(Cstr))
}
