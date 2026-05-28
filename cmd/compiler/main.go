package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/compiler"
)

var (
	input  = flag.String("i", "examples/reverse_array.dart", "source .dart file")
	output = flag.String("o", "examples/out", "executable program")
)

func main() {
	flag.Parse()

	err := compiler.Compiler(*input, *output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
