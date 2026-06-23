package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		flagString string
		flagInt    int
		flagBool   bool
	)

	flag.StringVar(&flagString, "stringvar", "default", "an example string var")
	flag.IntVar(&flagInt, "intval", 42, "an example int var")
	flag.BoolVar(&flagBool, "boolval", false, "an example bool var")

	flag.Parse()

	fmt.Println("String var:", flagString)
	fmt.Println("Int var:", flagInt)
	fmt.Println("Boolean var:", flagBool)
}
