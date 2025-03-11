package main

import "fmt"

// function that will take an arbitrary number of ints as arguments.
func sums(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	// Within the function, the type of nums is equivalent to []int.
	// We can call len(nums), iterate over it with range, etc.
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {

	// Variadic functions can be called in the usual way with individual arguments.
	sums(1, 2)
	sums(1, 2, 3)

	// If you already have multiple args in a slice,
	// apply them to a variadic function using func(slice...) like this.
	nums := []int{1, 2, 3, 4}
	sums(nums...)
}
