package sse

import (
	"fmt"
)

func ExampleData() {
	fmt.Println(Data("message"))
	// Output: data: message
}

func ExampleData_multiline() {
	fmt.Println(Data("message\nmessage"))
	// Output:
	// data: message
	// data: message
}

func ExampleId() {
	fmt.Println(Id("identifier"))
	// Output: id: identifier
}

func ExampleEvent() {
	fmt.Println(Event("eventname"))
	// Output: event: eventname
}

func ExampleRetry() {
	fmt.Println(Retry(5))
	// Output: retry: 5
}

