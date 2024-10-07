# sub_commands

This is demo how sub_commands could be implemented with golang flags

## main_logic

```golang
package main

import (
	"flag"
	"log"
)

var name string

func main() {
	flag.Parse()
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "Go", "name")
	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	phpCmd.StringVar(&name, "n", "Php", "name")

	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		_ = goCmd.Parse(args)
	case "php":
		_ = phpCmd.Parse(args)
	}

	log.Printf("name: %s\n", name)
}
```

## usage
```shell
sub_commands php -n=gson
```
expect output
```
name: gson
```
```shell
sub_commands go -n=eddie
```

expect output
```
flag provided but not defined: -n
```