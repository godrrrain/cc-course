// Code generated from KumirParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // KumirParser
import "github.com/antlr4-go/antlr/v4"

type BaseKumirParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseKumirParserVisitor) VisitQualifiedIdentifier(ctx *QualifiedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitIndexList(ctx *IndexListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitPowerExpression(ctx *PowerExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitTypeSpecifier(ctx *TypeSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitBasicType(ctx *BasicTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitArrayBounds(ctx *ArrayBoundsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitVariableDeclarationItem(ctx *VariableDeclarationItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitVariableList(ctx *VariableListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitGlobalDeclaration(ctx *GlobalDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitGlobalAssignment(ctx *GlobalAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitParameterDeclaration(ctx *ParameterDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitParameterList(ctx *ParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitAlgorithmNameTokens(ctx *AlgorithmNameTokensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitAlgorithmName(ctx *AlgorithmNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitAlgorithmHeader(ctx *AlgorithmHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitAlgorithmBody(ctx *AlgorithmBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitStatementSequence(ctx *StatementSequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitLvalue(ctx *LvalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitIoArgument(ctx *IoArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitIoArgumentList(ctx *IoArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitIoStatement(ctx *IoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitCaseBlock(ctx *CaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitEndLoopCondition(ctx *EndLoopConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitLoopSpecifier(ctx *LoopSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitLoopStatement(ctx *LoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitExitStatement(ctx *ExitStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitStopStatement(ctx *StopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitAlgorithmDefinition(ctx *AlgorithmDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitProgramItem(ctx *ProgramItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKumirParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}
