package main

import "fmt"

func main() {
	var age int = 25
	name := "Alice"

	fmt.Println("My name is ", name, " and I am ", age, " years old")
	fmt.Printf("%s is %d years old\n", name, age)


	// TYPES
	var isStudent bool = true
	var height float64 = 1.75
	
	fmt.Printf("Type %t is %T\n", isStudent, isStudent)
	fmt.Printf("Type %2.2f is %T\n", height, height)
	
	
	const pi = 3.14159
	fmt.Printf("The value of pi is %.5f\n", pi)
}
