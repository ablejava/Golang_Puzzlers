package main

import (
	"flag"
	"fmt"
)

var name string

// flag必须在flag.Parse()之前即可
var age *string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
	age = flag.String("age", "10", "age usage")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, name=%s, age=%s!\n", name, *age)
}
