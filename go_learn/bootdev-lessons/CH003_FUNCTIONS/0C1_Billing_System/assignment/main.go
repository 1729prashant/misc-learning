package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	baseBill := calculateBaseBill(costPerMessage, numMessages)
	discountPercentage := calculateDiscountRate(numMessages)
	discountAmount := baseBill * discountPercentage
	finalBill := baseBill - discountAmount
	return finalBill
}

func calculateDiscountRate(messagesSent int) float64 {
	if messagesSent > 5000 {
		return 0.20
	}
	if messagesSent > 1000 {
		return 0.10
	}
	return 0.0
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}

