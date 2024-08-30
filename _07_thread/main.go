package main

/*
#include <pthread.h>
#include <stdio.h>

static void printThreadId() {
	pthread_t threadId = pthread_self();
	printf("thread id: %lu\n", threadId);
}
*/
import "C"

func main() {
	C.printThreadId()

}
