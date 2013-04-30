package sse

import (
	"fmt"
	"strings"
)

// Send a message to the client
func Data(input string) string {
	lines := []string{}
	for _, line := range strings.Split(input, "\n") {
		lines = append(lines, fmt.Sprintf("data: %s", line))
	}
	return fmt.Sprintf("%s\n\n", strings.Join(lines, "\n"))
}

// Set the last event ID.
// If a connection is dropped, the header `Last-Event-ID`
// is sent with the new request. As well, the client `message`
// event contains a `lastEventId` property
func Id(id string) string {
	return fmt.Sprintf("id: %s\n", id)
}

// Set the event name for the next data message
func Event(name string) string {
	return fmt.Sprintf("event: %s\n", name)
}

// Set the event stream's reconnection time
func Retry(timeout int) string {
	return fmt.Sprintf("retry: %d\n", timeout)
}
