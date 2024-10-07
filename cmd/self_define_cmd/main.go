package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

type Name string

func (i *Name) String() string {
	return fmt.Sprint(*i)
}

func (i *Name) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("name flag already set")
	}
	*i = Name("special:" + value)
	return nil
}
func main() {
	var name Name
	flag.Var(&name, "name", "name")
	flag.Parse()
	log.Printf("name: %s\n", name)
}
