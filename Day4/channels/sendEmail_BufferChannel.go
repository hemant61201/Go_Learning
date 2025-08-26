package main

func addEmailsToQueue(emails []string) chan string {
	emailsChan := make(chan string, len(emails))
	for _, email := range emails {
		emailsChan <- email
	}
	return emailsChan
}
