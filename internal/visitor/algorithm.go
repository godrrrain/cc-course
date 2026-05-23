package visitor

import (
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
)

func (v *IRVisitor) VisitDeclaration(ctx *parser.DeclarationContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitMethodSignature(ctx *parser.MethodSignatureContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitClassBody(ctx *parser.ClassBodyContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitClassMemberDeclaration(ctx *parser.ClassMemberDeclarationContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitSuperclass(ctx *parser.SuperclassContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitMixinTypes(ctx *parser.MixinTypesContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitMixinApplication(ctx *parser.MixinApplicationContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitInterfaces(ctx *parser.InterfacesContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitEnumEntry(ctx *parser.EnumEntryContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitFunctionSignature(ctx *parser.FunctionSignatureContext) interface{} {
	return nil
}
