package main

// void PrintHello(_GoString_ s);
import "C"
import "fmt"

//export PrintHello
func PrintHello(str string) {
	fmt.Println("Hello, " + str + "!")
}

func main() {
	C.PrintHello("Go")
}
