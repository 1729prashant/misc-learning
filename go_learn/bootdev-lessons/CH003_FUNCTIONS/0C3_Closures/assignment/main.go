package main

func adder() func(int) int {
	totalTexts := 0
	
	return func(addVal int) int {
		totalTexts += addVal
		return totalTexts
	}
	
}

