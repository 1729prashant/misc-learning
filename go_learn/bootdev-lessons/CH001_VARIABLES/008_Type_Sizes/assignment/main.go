package main

import "fmt"

func main() {
	accountAgeFloat := 2.6

	// create a new "accountAgeInt" here
	// it should be the result of casting "accountAgeFloat" to an integer
	accountAgeInt := int(accountAgeFloat)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}

