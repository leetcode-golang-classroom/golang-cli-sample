package main

import (
	"log"

	"github.com/leetcode-golang-classroom/golang-cli-sample/internal/commands"
)

func main() {
	err := commands.Execute()
	if err != nil {
		log.Fatalf("commands.Execute error: %v", err)
	}
}
