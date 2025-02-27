"
2. Variables and Constants
Variable Declaration
"

package main

import "fmt"

func main() {
    var a int = 10  // Explicit type
    var b = "Go"    // Type inferred
    c := 3.14       // Short declaration (type inferred)

    fmt.Println(a, b, c)
}


"
var name type = value (Explicit type)
var name = value (Type inferred)  Since 'Go' is a string, Go automatically infers b as a string (string type) without needing an explicit type declaration.
name := value (Short declaration, only inside functions)
"


// Constants
// const Pi = 3.1415
// const Name string = "GoLang"
// -> const values cannot be changed after declaration.


// --------------------------------------------
//     Data Types
// Type	  		Example
// int   		var x int = 10
// float64 		var pi float64 = 3.14
// string		var name string = "Go"
// bool			var flag bool = true
// ---------------------------------------------------------
