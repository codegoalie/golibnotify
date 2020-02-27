package main

import (
	"fmt"
	"log"
	"time"

	"github.com/codegoalie/golibnotify"
)

func main() {

	// Get an instance of a SimpleNotifier
	notifier := golibnotify.NewSimpleNotifier("my-cool-app-name")

	// Show a new notification
	err := notifier.Show("A summary", "Some body text", "a/path/to/an/icon.png")
	if err != nil {
		err = fmt.Errorf("failed to send a notification: %w", err)
		log.Fatal(err)
	}

	time.Sleep(3 * time.Second)

	// Update an existing notification (or send a new one if one hasn't been sent)
	err = notifier.Update("A new summary", "Some different body text", "another/path/to/icon.png")
	if err != nil {
		err = fmt.Errorf("failed to update a notification: %w", err)
		log.Fatal(err)
	}

	time.Sleep(3 * time.Second)

	// Remove an existing notification
	err = notifier.Close()
	if err != nil {
		err = fmt.Errorf("failed to close a notification: %w", err)
		log.Fatal(err)
	}

}
