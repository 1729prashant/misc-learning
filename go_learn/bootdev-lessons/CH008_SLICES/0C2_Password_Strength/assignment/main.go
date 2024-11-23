package main

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	
	AZ := "QWERTYUIOPASDFGHJKLZXCVBNM"
	NUM := "1234567890"
	foundUpperCase := false
	foundNumber := false
	
	for _, ele := range password {

		for _, AZ_ele := range AZ {
			if ele == AZ_ele {
				foundUpperCase = true

			}
		}

		
		for _, NUM_ele := range NUM {
			
			if ele == NUM_ele {
				foundNumber = true
			}
		}

	}
	return foundUpperCase && foundNumber
}




/*
provided solution


package main

func isValidPassword(password string) bool {
	length := len(password)
	hasUpper := false
	hasNumber := false

	for i := 0; i < length; i++ {
		char := password[i]
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		}
		if char >= '0' && char <= '9' {
			hasNumber = true
		}
	}

	return length >= 5 && length <= 12 && hasUpper && hasNumber
}

*/
