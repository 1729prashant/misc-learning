package main

func main() {
	var x int = 50
	var y *int = &x
	*y = 100
}



// after execution value of 
// y is 100
// x ix 100
