# name_parser

usage use -flag to access input name

## main logic

```golang
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
```

## usage

```shell
./name_parser -name=test -n=eddie
```
expect output

```
name: eddie
```