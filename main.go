package main

import (
	"flag"
	"Calculator/Parser"
)

var P = Parser.Parser{}

func main() {
	math := flag.String("math", "", "")

	flag.Parse()
	P.Parse(*math)
	P.Status()
}
