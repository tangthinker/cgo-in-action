package main

// #include "hello.h"
import "C"

func main() {
	C.PrintHello(C.CString("Go"))
}
