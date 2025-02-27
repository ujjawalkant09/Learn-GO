"
Data Structures in Go
-> I will implement this in detail later; this is just a skeleton. 
If you are familiar with any other programming language, you can easily relate to it and won't waste too much time.

1️⃣ Arrays -> Fixed Size

var arr [3]int = [3]int{1, 2, 3}

----------------------------------------------------------------------

2️⃣ Slices (Dynamic Arrays) {Refer to Extra_Pie Article 1.1}
nums := []int{1, 2, 3, 4}
nums = append(nums, 5) // Append new elements

----------------------------------------------------------------------

3️⃣ Maps (Key-Value Store)

person := map[string]int{"Alice": 25, "Bob": 30}
fmt.Println(person["Alice"])

----------------------------------------------------------------------

4️⃣ Structs (Custom Data Types)

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{"John", 30}
    fmt.Println(p.Name, p.Age)
}

----------------------------------------------------------------------

5️⃣ Interfaces (Polymorphism)  ->  {Refer to Extra_Pie Article 1.2}

type Shape interface {
    Area() float64
}

---------------------------------------------------------------------------
"