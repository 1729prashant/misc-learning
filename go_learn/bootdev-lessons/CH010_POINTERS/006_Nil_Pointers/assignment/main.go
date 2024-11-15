//before code correction
/*
package main

import (
	"strings"
)

func removeProfanity(message *string) {
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}
*/

//after code correction
package main

import (
	"strings"
)

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}
