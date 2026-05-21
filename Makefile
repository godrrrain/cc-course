file ?= ifs.kum

antlr4:
	antlr4 -Dlanguage=Go KumirLexer.g4 KumirParser.g4 -visitor -o internal/parser -package parser

tree:
	cat examples/${file} | antlr4-parse KumirParser.g4 KumirLexer.g4 program -gui