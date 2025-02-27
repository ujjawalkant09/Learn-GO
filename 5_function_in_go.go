// Function In go

package main

import "fmt"


// Functions can take arguments and return values.
//  So this function will take two integer and will return the sum of type int 

func add(a int,b int) int{
	return a+b 
}

//  Functions can also return Multiple values

func divide(a int, b int)(int,int){
	return a/b , a%b
}

func main(){
	sum := add(5,10)
	fmt.Println(sum)

	div,mod := divide(10,5)
	fmt.Println(div)
	fmt.Println(mod)

}
