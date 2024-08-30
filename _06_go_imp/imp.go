package main

import "C"
import "fmt"

//export PrintHello
func PrintHello(s *C.char) {
	fmt.Println("Hello from Go:", C.GoString(s))
}
