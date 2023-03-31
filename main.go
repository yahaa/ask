package main

import (
	"log"
	"os"

	"github.com/yahaa/ask/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Printf("run err: %v", err)
		os.Exit(1)
	}
}
