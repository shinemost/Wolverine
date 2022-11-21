package test

//#include <stdio.h>
//void say(){
// printf("Hello world\n");
//}
import "C"

func Xx() {
	C.say()
}
