package main

func getMonthlyPrice(tier string) int {
	if tier == "basic" {
		return 10000
	} else if tier == "premium" {
		return 15000
	} else if tier == "enterprise" {
		return 50000
	} else {
		return 0
	}
}

