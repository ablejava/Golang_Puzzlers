package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 方式1。
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	// flag.Usage必须在Parse之前
	flag.Parse()
}
