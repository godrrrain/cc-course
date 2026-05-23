package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/compiler"
)

var (
	input  = flag.String("i", "examples/for_loop.dart", "source .dart file")
	output = flag.String("o", "examples/out", "executable program")
	isDOT  = flag.Bool("d", true, "output as DOT")
)

func main() {
	flag.Parse()

	err := compiler.Compiler(*input, *output, *isDOT)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
