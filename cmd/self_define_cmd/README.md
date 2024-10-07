# self_define_cmd

This is demo how to define self struct cmd with golang flag package

## main logic

```golang
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
```

## usage

```shell
self_define_cmd -name=gson
```

expect output
```
name: special:gson
```