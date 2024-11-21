package main

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))
	for i := 0; i < len(emails); i++ {
		ch <- emails[i]
	}
	
	return ch
}



/*
using range
package main

func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}


*/
