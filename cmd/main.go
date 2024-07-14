package main

import (
	"log"

	io "github.com/clinto-bean/go-dispatch/internal/io"
)

func main() {
	err := io.Repl()
	if err.Error() == "kill" {
		log.Fatal("Shutting down.")
	}
	log.Fatalf("Exiting: Uncaught exception: %v", err.Error())
}
