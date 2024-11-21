package main

import "time"

func processMessages(messages []string) []string {
	if len(messages) == 0 {
		return []string{}
	}

	ch := make(chan string, len(messages))

	for _, msg := range messages {
		go func(m string) {
			ch <- process(m)
		}(msg)
	}

	processedMessages := make([]string, len(messages))
	for i := 0; i < len(messages); i++ {
		processedMessages[i] = <-ch
	}

	return processedMessages
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}

