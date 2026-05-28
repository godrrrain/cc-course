package visitor

import (
	"cc-course/internal/parser"
)

func (v *IRVisitor) VisitForLoopParts(ctx *parser.ForLoopPartsContext) interface{} {
	return nil
}
