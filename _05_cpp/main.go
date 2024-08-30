package main

// void printHello(const char* name);
import "C"

func main() {
	C.printHello(C.CString("Go"))
}
