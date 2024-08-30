package main

// void PrintHello(char *s);
import "C"
import "fmt"

//export PrintHello
func PrintHello(s *C.char) {
	fmt.Println("Hello from Go, C says:", C.GoString(s))
}

func main() {
	C.PrintHello(C.CString("World"))
}
