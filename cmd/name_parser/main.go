package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "Go Cli Example", "Name")
	flag.StringVar(&name, "n", "Go Cli Example", "Name")
	// parse input
	flag.Parse()
	log.Printf("name: %s\n", name)
}
