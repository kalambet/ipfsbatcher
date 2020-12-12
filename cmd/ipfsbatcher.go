package main

import (
	"ipfsbatcher"
	"log"
)

func main() {
	err := ipfsbatcher.Do()
	if err != nil {
		log.Fatalf("error batching and calculating batch size: %s", err)
	}
}
