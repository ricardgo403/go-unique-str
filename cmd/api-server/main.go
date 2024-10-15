package main

import (
	"github.com/ricardgo403/go-unique-str"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("api-server: ")
	log.SetFlags(0)

	// Request a greeting message.
	err := uniqueStr.Run()
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	log.Println("Hello, World!")
}
