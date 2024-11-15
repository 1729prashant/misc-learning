package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func getMessageText(analytics *Analytics, message Message) {
	// Increment total messages
	analytics.MessagesTotal++

	// Update success or failure counts
	if message.Success {
		analytics.MessagesSucceeded++
	} else {
		analytics.MessagesFailed++
	}
}
