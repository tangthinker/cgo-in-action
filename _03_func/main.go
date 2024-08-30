package main

/*
#include <stdio.h>

static void printHello(const char* message) {
    printf("Hello, world!\n");
	printf("%s\n", message);
}
*/
import "C"

func main() {
	C.printHello(C.CString("Hello, Go!"))
}
