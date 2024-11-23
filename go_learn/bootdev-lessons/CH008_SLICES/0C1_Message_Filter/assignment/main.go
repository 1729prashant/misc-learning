
package main

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
		var filteredMessages []Message

	// Loop through each message in the slice
	for _, message := range messages {
		// Check if the message type matches the filterType
		if message.Type() == filterType {
			// Append the message to the filteredMessages slice if it matches
			filteredMessages = append(filteredMessages, message)
		}
	}

	// Return the filtered slice of messages
	return filteredMessages
}



/*
provided solution

package main

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	var filtered []Message
	for _, message := range messages {
		if message.Type() == filterType {
			filtered = append(filtered, message)
		}
	}
	return filtered
}

*/
