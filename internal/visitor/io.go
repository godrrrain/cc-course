package visitor

import (
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
)

func (v *IRVisitor) VisitOnPart(ctx *parser.OnPartContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitFinallyPart(ctx *parser.FinallyPartContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitLabel(ctx *parser.LabelContext) interface{} {
	return nil
}
