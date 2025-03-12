package main

import "fmt"

// zeroval does NOT modify the original variable because it's passed by value.
func zeroval(ival int) {
	ival = 0 // This change is local to zeroval()
}

// zeroptr modifies the original variable because it's passed by reference (pointer).
func zeroptr(iptr *int) {
	*iptr = 0 // Dereferencing the pointer and modifying the value
}

// Function to modify value via pointer
func modify(n *int) {
	*n += 10
}

// Allocates memory on heap using new()
func createPointer(val int) *int {
	ptr := new(int)
	*ptr = val
	return ptr // Now returning a safely allocated pointer
}

// Function to modify a slice (modifies underlying array)
func modifySlice(s []int) {
	s[0] = 42
}

// Function to modify struct fields via pointer
func modifyStruct(p *Person) {
	p.age++ // Same as (*p).age++
}

type Person struct {
	name string
	age  int
}

func main() {
	i := 1
	fmt.Println("Value of i before zeroval:", i)

	zeroval(i)                                  // Passed by value, no effect on i
	fmt.Println("Value of i after zeroval:", i) // Still 1

	// The &i syntax gives the memory address of i
	zeroptr(&i)                                 // Passed by reference, modifies i
	fmt.Println("Value of i after zeroptr:", i) // Now 0

	fmt.Println("Address of i:", &i) // Prints memory address

	fmt.Println("====================")

	// 1️⃣ Basic pointer declaration and dereferencing
	var x int = 10
	var p *int = &x                           // p holds the address of x
	fmt.Println("Value of x:", x)             // 10
	fmt.Println("Pointer p points to x:", *p) // Dereferencing p -> 10

	// 2️⃣ Modifying a value via a pointer
	*p = 20
	fmt.Println("Modified x through p:", x) // 20

	// 3️⃣ Nil pointer and dereferencing panic
	var np *int
	fmt.Println("Nil pointer value:", np) // Prints <nil>
	// fmt.Println(*np)  // ❌ PANIC: Dereferencing nil pointer

	// 5️⃣ Returning pointers from a function (Fixed!)
	ptr := createPointer(100)
	fmt.Println("Value from returned pointer:", *ptr) // 100

	// 6️⃣ Array pointer vs slice pointer
	arr := [3]int{1, 2, 3}
	arrPtr := &arr                              // Pointer to the entire array
	fmt.Println("Array pointer:", (*arrPtr)[1]) // Access via pointer -> 2

	slice := arr[:]
	modifySlice(slice)
	fmt.Println("Modified array via slice:", arr) // [42, 2, 3] (changes reflect in array)

	// 7️⃣ Struct pointers
	person := Person{"Alice", 25}
	personPtr := &person // Pointer to struct

	fmt.Println("Person name via pointer:", personPtr.name) // Go allows shorthand without (*personPtr).name

	modifyStruct(personPtr)
	fmt.Println("Modified person age:", person.age) // 26

}
