package tree

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"os/exec"
	"strings"
)

type dotTreeVisitor struct {
	builder strings.Builder
	nodeID  int
}

func (v *dotTreeVisitor) nextID() int {
	v.nodeID++
	return v.nodeID
}

func (v *dotTreeVisitor) writeNode(id int, label string) {
	label = strings.ReplaceAll(label, "\"", "\\\"")
	fmt.Fprintf(&v.builder, "  node%d [label=\"%s\"];\n", id, label)
}

func (v *dotTreeVisitor) writeEdge(from, to int) {
	fmt.Fprintf(&v.builder, "  node%d -> node%d;\n", from, to)
}

func (v *dotTreeVisitor) visit(tree antlr.Tree, recog antlr.Recognizer) int {
	id := v.nextID()

	var label string
	switch t := tree.(type) {
	case antlr.TerminalNode:
		label = fmt.Sprintf("%s", t.GetText())
	case antlr.ParserRuleContext:
		ruleName := recog.GetRuleNames()[t.GetRuleIndex()]
		label = fmt.Sprintf("%s", ruleName)
		label = label[strings.LastIndex(label, ".")+1:]
	default:
		label = "?"
	}

	v.writeNode(id, label)

	for i := 0; i < tree.GetChildCount(); i++ {
		child := tree.GetChild(i)
		childID := v.visit(child, recog)
		v.writeEdge(id, childID)
	}

	return id
}

func generateDOT(tree antlr.Tree, recog antlr.Recognizer) string {
	v := &dotTreeVisitor{}
	v.builder.WriteString("digraph ParseTree {\n")
	v.builder.WriteString("  node [shape=box];\n")
	v.visit(tree, recog)
	v.builder.WriteString("}\n")
	return v.builder.String()
}

func SaveTreeToFile(tree antlr.Tree, recog antlr.Recognizer, filename string) error {
	dot := generateDOT(tree, recog)

	err := os.WriteFile(filename+".dot", []byte(dot), 0644)
	if err != nil {
		return fmt.Errorf("file %s write error: %w", filename+".dot", err)
	}

	cmd := exec.Command("dot", "-Tpng", "-o", filename+".png", filename+".dot")
	if _, err := cmd.CombinedOutput(); err != nil {

		return fmt.Errorf("file %s write error: %w", filename+".png", err)
	}

	return nil
}
