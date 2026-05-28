package visitor

import (
	"cc-course/internal/parser"
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
