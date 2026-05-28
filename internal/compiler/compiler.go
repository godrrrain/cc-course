package compiler

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	ast "github.com/AskaryanKarine/BMSTU-CC/cource/internal/tree"
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/visitor"
	"github.com/antlr4-go/antlr/v4"
)

func Compiler(input, output string) error {
	fileStream, err := antlr.NewFileStream(input)
	if err != nil {
		return fmt.Errorf("compiler input error: %w", err)
	}

	lexer := parser.NewDart2Lexer(fileStream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewDart2Parser(tokens)

	errListener := newDartErrorListener()
	p.AddErrorListener(errListener)
	if len(errListener.errs) > 0 {
		return fmt.Errorf("compiler parser error: %w", errListener.errs[0])
	}

	tree := p.CompilationUnit()

	err = ast.SaveTreeToFile(tree, p, output)
	if err != nil {
		return fmt.Errorf("ast build error: %w", err)
	}

	v := visitor.NewIRVisitor()
	v.Visit(tree)

	if len(v.Errors) > 0 {
		return fmt.Errorf("compiler errors: %v", v.Errors)
	}

	llFilename := output + ".ll"

	err = v.WriteToFile(llFilename)
	if err != nil {
		return fmt.Errorf("save llvm output: %w", err)
	}

	cmd := exec.Command("clang", llFilename, "-o", output+".exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("compile llvm to program: %w", err)
	}

	return nil
}

type dartErrorListener struct {
	*antlr.DefaultErrorListener
	errs []error
}

func newDartErrorListener() *dartErrorListener {
	return &dartErrorListener{
		DefaultErrorListener: new(antlr.DefaultErrorListener),
		errs:                 make([]error, 0),
	}
}

func (l *dartErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	l.errs = append(l.errs, fmt.Errorf("syntax error at %d:%d – %s", line, column, msg))
}
