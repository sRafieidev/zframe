package libtest

/*
 // Test Code For Import C And Assembliy In Go Lang

#include <stdio.h>
int addInC(int a, int b) {

int src = 1;
int dst;

asm ("mov %1, %0\n\t"
     "add $1, %0"
     :"=r" (dst)
     :"r" (src));
	printf("  Siavash  Print Data In C .. :  %d\n", dst);
		return a + b;
}
*/
import "C"
import "fmt"

func Gofunctest() {
	a := 3
	b := 5
	c := C.addInC(C.int(a), C.int(b))
	fmt.Println("Add in C:", a, "+", b, "=", int(c))
}
