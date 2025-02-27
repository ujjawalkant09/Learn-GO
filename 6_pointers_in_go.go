// Pointers
// & gets the memory address.
// * dereferences the pointer.

package main

import "fmt"

func main(){
	x := 10
	ptr := &x //Pointer to X
	fmt.Println("Value : ",*ptr)
	fmt.Println("Address : ",ptr)
}