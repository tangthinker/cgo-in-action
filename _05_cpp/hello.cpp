
#include <iostream>

extern "C" {
    #include "hello.h"
}

void printHello(const char* name) {
    std::cout << "Hello, " << name << "!" << std::endl;
}