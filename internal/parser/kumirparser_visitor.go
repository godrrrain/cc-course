// Code generated from KumirParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // KumirParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by KumirParser.
type KumirParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by KumirParser#qualifiedIdentifier.
	VisitQualifiedIdentifier(ctx *QualifiedIdentifierContext) interface{}

	// Visit a parse tree produced by KumirParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by KumirParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by KumirParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by KumirParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by KumirParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by KumirParser#indexList.
	VisitIndexList(ctx *IndexListContext) interface{}

	// Visit a parse tree produced by KumirParser#postfixExpression.
	VisitPostfixExpression(ctx *PostfixExpressionContext) interface{}

	// Visit a parse tree produced by KumirParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by KumirParser#powerExpression.
	VisitPowerExpression(ctx *PowerExpressionContext) interface{}

	// Visit a parse tree produced by KumirParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by KumirParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by KumirParser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by KumirParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by KumirParser#logicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by KumirParser#logicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by KumirParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by KumirParser#typeSpecifier.
	VisitTypeSpecifier(ctx *TypeSpecifierContext) interface{}

	// Visit a parse tree produced by KumirParser#basicType.
	VisitBasicType(ctx *BasicTypeContext) interface{}

	// Visit a parse tree produced by KumirParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by KumirParser#arrayBounds.
	VisitArrayBounds(ctx *ArrayBoundsContext) interface{}

	// Visit a parse tree produced by KumirParser#variableDeclarationItem.
	VisitVariableDeclarationItem(ctx *VariableDeclarationItemContext) interface{}

	// Visit a parse tree produced by KumirParser#variableList.
	VisitVariableList(ctx *VariableListContext) interface{}

	// Visit a parse tree produced by KumirParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by KumirParser#globalDeclaration.
	VisitGlobalDeclaration(ctx *GlobalDeclarationContext) interface{}

	// Visit a parse tree produced by KumirParser#globalAssignment.
	VisitGlobalAssignment(ctx *GlobalAssignmentContext) interface{}

	// Visit a parse tree produced by KumirParser#parameterDeclaration.
	VisitParameterDeclaration(ctx *ParameterDeclarationContext) interface{}

	// Visit a parse tree produced by KumirParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by KumirParser#algorithmNameTokens.
	VisitAlgorithmNameTokens(ctx *AlgorithmNameTokensContext) interface{}

	// Visit a parse tree produced by KumirParser#algorithmName.
	VisitAlgorithmName(ctx *AlgorithmNameContext) interface{}

	// Visit a parse tree produced by KumirParser#algorithmHeader.
	VisitAlgorithmHeader(ctx *AlgorithmHeaderContext) interface{}

	// Visit a parse tree produced by KumirParser#algorithmBody.
	VisitAlgorithmBody(ctx *AlgorithmBodyContext) interface{}

	// Visit a parse tree produced by KumirParser#statementSequence.
	VisitStatementSequence(ctx *StatementSequenceContext) interface{}

	// Visit a parse tree produced by KumirParser#lvalue.
	VisitLvalue(ctx *LvalueContext) interface{}

	// Visit a parse tree produced by KumirParser#assignmentStatement.
	VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{}

	// Visit a parse tree produced by KumirParser#ioArgument.
	VisitIoArgument(ctx *IoArgumentContext) interface{}

	// Visit a parse tree produced by KumirParser#ioArgumentList.
	VisitIoArgumentList(ctx *IoArgumentListContext) interface{}

	// Visit a parse tree produced by KumirParser#ioStatement.
	VisitIoStatement(ctx *IoStatementContext) interface{}

	// Visit a parse tree produced by KumirParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by KumirParser#caseBlock.
	VisitCaseBlock(ctx *CaseBlockContext) interface{}

	// Visit a parse tree produced by KumirParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by KumirParser#endLoopCondition.
	VisitEndLoopCondition(ctx *EndLoopConditionContext) interface{}

	// Visit a parse tree produced by KumirParser#loopSpecifier.
	VisitLoopSpecifier(ctx *LoopSpecifierContext) interface{}

	// Visit a parse tree produced by KumirParser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by KumirParser#exitStatement.
	VisitExitStatement(ctx *ExitStatementContext) interface{}

	// Visit a parse tree produced by KumirParser#stopStatement.
	VisitStopStatement(ctx *StopStatementContext) interface{}

	// Visit a parse tree produced by KumirParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by KumirParser#algorithmDefinition.
	VisitAlgorithmDefinition(ctx *AlgorithmDefinitionContext) interface{}

	// Visit a parse tree produced by KumirParser#programItem.
	VisitProgramItem(ctx *ProgramItemContext) interface{}

	// Visit a parse tree produced by KumirParser#program.
	VisitProgram(ctx *ProgramContext) interface{}
}
