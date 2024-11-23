package main

func getLast[T any](s []T) T {
	if len(s) == 0 {
		var myZero T
		return myZero
	}

	return s[len(s)-1]
}

