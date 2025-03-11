package main

import (
	"fmt"
	"slices"
)

func main() {

	// Unlike arrays, slices are typed only by the
	// elements they contain (not the number of elements).
	// An uninitialized slice equals to nil and has length 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// To create an empty slice with non-zero length,
	// use the builtin make. Here we make a slice of
	// strings of length 3 (initially zero-valued).
	// By default a new slice’s capacity is equal to
	// its length; if we know the slice is going to
	// grow ahead of time, it’s possible to pass a
	// capacity explicitly as an additional parameter
	// to make.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// append returns a slice containing one
	// or more new values. Note that we need to
	// accept a return value from append as we may
	// get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy slices
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// can use slicing slice[low:high].
	// this gets a slice of the elements s[2], s[3], and s[4]
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("sl2:", l)

	// And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("sl3:", l)

	// can declare and initialize a variable for slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// use slices package to compare slices
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println(" t == t2")
	}

	// Slices can be composed into multi-dimensional
	// data structures. The length of the inner slices
	// can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
