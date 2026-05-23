file ?= hello.dart

antlr4:
	antlr4 -Dlanguage=Go Dart2Lexer.g4 Dart2Parser.g4 -visitor -o internal/parser -package parser

tree:
	cat examples/${file} | antlr4-parse Dart2Parser.g4 Dart2Lexer.g4 compilationUnit -gui

run:
	go run cmd/compiler/main.go -i examples/$(file) -o examples/out