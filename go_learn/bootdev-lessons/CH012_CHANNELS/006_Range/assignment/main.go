package main

func concurrentFib(n int) []int {
	ch := make(chan int) //channel
	var s []int //slice

	
	go fibonacci(n, ch)

	for i := range ch {
		s = append(s,i)
	}

	return s
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}



/*
provided solution

package main

func concurrentFib(n int) []int {
	ch := make(chan int)
	go fibonacci(n, ch)

	final := []int{}
	for v := range ch {
		final = append(final, v)
	}
	return final
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}


*/
