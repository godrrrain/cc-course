// Code generated from KumirParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // KumirParser
import "github.com/antlr4-go/antlr/v4"

// BaseKumirParserListener is a complete listener for a parse tree produced by KumirParser.
type BaseKumirParserListener struct{}

var _ KumirParserListener = &BaseKumirParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseKumirParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseKumirParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseKumirParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseKumirParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQualifiedIdentifier is called when production qualifiedIdentifier is entered.
func (s *BaseKumirParserListener) EnterQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// ExitQualifiedIdentifier is called when production qualifiedIdentifier is exited.
func (s *BaseKumirParserListener) ExitQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseKumirParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseKumirParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseKumirParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseKumirParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseKumirParserListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseKumirParserListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseKumirParserListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseKumirParserListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseKumirParserListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseKumirParserListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterIndexList is called when production indexList is entered.
func (s *BaseKumirParserListener) EnterIndexList(ctx *IndexListContext) {}

// ExitIndexList is called when production indexList is exited.
func (s *BaseKumirParserListener) ExitIndexList(ctx *IndexListContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseKumirParserListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseKumirParserListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseKumirParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseKumirParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterPowerExpression is called when production powerExpression is entered.
func (s *BaseKumirParserListener) EnterPowerExpression(ctx *PowerExpressionContext) {}

// ExitPowerExpression is called when production powerExpression is exited.
func (s *BaseKumirParserListener) ExitPowerExpression(ctx *PowerExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseKumirParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseKumirParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseKumirParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseKumirParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseKumirParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseKumirParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseKumirParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseKumirParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *BaseKumirParserListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *BaseKumirParserListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *BaseKumirParserListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *BaseKumirParserListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseKumirParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseKumirParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *BaseKumirParserListener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *BaseKumirParserListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {}

// EnterBasicType is called when production basicType is entered.
func (s *BaseKumirParserListener) EnterBasicType(ctx *BasicTypeContext) {}

// ExitBasicType is called when production basicType is exited.
func (s *BaseKumirParserListener) ExitBasicType(ctx *BasicTypeContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseKumirParserListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseKumirParserListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterArrayBounds is called when production arrayBounds is entered.
func (s *BaseKumirParserListener) EnterArrayBounds(ctx *ArrayBoundsContext) {}

// ExitArrayBounds is called when production arrayBounds is exited.
func (s *BaseKumirParserListener) ExitArrayBounds(ctx *ArrayBoundsContext) {}

// EnterVariableDeclarationItem is called when production variableDeclarationItem is entered.
func (s *BaseKumirParserListener) EnterVariableDeclarationItem(ctx *VariableDeclarationItemContext) {}

// ExitVariableDeclarationItem is called when production variableDeclarationItem is exited.
func (s *BaseKumirParserListener) ExitVariableDeclarationItem(ctx *VariableDeclarationItemContext) {}

// EnterVariableList is called when production variableList is entered.
func (s *BaseKumirParserListener) EnterVariableList(ctx *VariableListContext) {}

// ExitVariableList is called when production variableList is exited.
func (s *BaseKumirParserListener) ExitVariableList(ctx *VariableListContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseKumirParserListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseKumirParserListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterGlobalDeclaration is called when production globalDeclaration is entered.
func (s *BaseKumirParserListener) EnterGlobalDeclaration(ctx *GlobalDeclarationContext) {}

// ExitGlobalDeclaration is called when production globalDeclaration is exited.
func (s *BaseKumirParserListener) ExitGlobalDeclaration(ctx *GlobalDeclarationContext) {}

// EnterGlobalAssignment is called when production globalAssignment is entered.
func (s *BaseKumirParserListener) EnterGlobalAssignment(ctx *GlobalAssignmentContext) {}

// ExitGlobalAssignment is called when production globalAssignment is exited.
func (s *BaseKumirParserListener) ExitGlobalAssignment(ctx *GlobalAssignmentContext) {}

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *BaseKumirParserListener) EnterParameterDeclaration(ctx *ParameterDeclarationContext) {}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *BaseKumirParserListener) ExitParameterDeclaration(ctx *ParameterDeclarationContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseKumirParserListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseKumirParserListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterAlgorithmNameTokens is called when production algorithmNameTokens is entered.
func (s *BaseKumirParserListener) EnterAlgorithmNameTokens(ctx *AlgorithmNameTokensContext) {}

// ExitAlgorithmNameTokens is called when production algorithmNameTokens is exited.
func (s *BaseKumirParserListener) ExitAlgorithmNameTokens(ctx *AlgorithmNameTokensContext) {}

// EnterAlgorithmName is called when production algorithmName is entered.
func (s *BaseKumirParserListener) EnterAlgorithmName(ctx *AlgorithmNameContext) {}

// ExitAlgorithmName is called when production algorithmName is exited.
func (s *BaseKumirParserListener) ExitAlgorithmName(ctx *AlgorithmNameContext) {}

// EnterAlgorithmHeader is called when production algorithmHeader is entered.
func (s *BaseKumirParserListener) EnterAlgorithmHeader(ctx *AlgorithmHeaderContext) {}

// ExitAlgorithmHeader is called when production algorithmHeader is exited.
func (s *BaseKumirParserListener) ExitAlgorithmHeader(ctx *AlgorithmHeaderContext) {}

// EnterAlgorithmBody is called when production algorithmBody is entered.
func (s *BaseKumirParserListener) EnterAlgorithmBody(ctx *AlgorithmBodyContext) {}

// ExitAlgorithmBody is called when production algorithmBody is exited.
func (s *BaseKumirParserListener) ExitAlgorithmBody(ctx *AlgorithmBodyContext) {}

// EnterStatementSequence is called when production statementSequence is entered.
func (s *BaseKumirParserListener) EnterStatementSequence(ctx *StatementSequenceContext) {}

// ExitStatementSequence is called when production statementSequence is exited.
func (s *BaseKumirParserListener) ExitStatementSequence(ctx *StatementSequenceContext) {}

// EnterLvalue is called when production lvalue is entered.
func (s *BaseKumirParserListener) EnterLvalue(ctx *LvalueContext) {}

// ExitLvalue is called when production lvalue is exited.
func (s *BaseKumirParserListener) ExitLvalue(ctx *LvalueContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BaseKumirParserListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BaseKumirParserListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterIoArgument is called when production ioArgument is entered.
func (s *BaseKumirParserListener) EnterIoArgument(ctx *IoArgumentContext) {}

// ExitIoArgument is called when production ioArgument is exited.
func (s *BaseKumirParserListener) ExitIoArgument(ctx *IoArgumentContext) {}

// EnterIoArgumentList is called when production ioArgumentList is entered.
func (s *BaseKumirParserListener) EnterIoArgumentList(ctx *IoArgumentListContext) {}

// ExitIoArgumentList is called when production ioArgumentList is exited.
func (s *BaseKumirParserListener) ExitIoArgumentList(ctx *IoArgumentListContext) {}

// EnterIoStatement is called when production ioStatement is entered.
func (s *BaseKumirParserListener) EnterIoStatement(ctx *IoStatementContext) {}

// ExitIoStatement is called when production ioStatement is exited.
func (s *BaseKumirParserListener) ExitIoStatement(ctx *IoStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseKumirParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseKumirParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterCaseBlock is called when production caseBlock is entered.
func (s *BaseKumirParserListener) EnterCaseBlock(ctx *CaseBlockContext) {}

// ExitCaseBlock is called when production caseBlock is exited.
func (s *BaseKumirParserListener) ExitCaseBlock(ctx *CaseBlockContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseKumirParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseKumirParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterEndLoopCondition is called when production endLoopCondition is entered.
func (s *BaseKumirParserListener) EnterEndLoopCondition(ctx *EndLoopConditionContext) {}

// ExitEndLoopCondition is called when production endLoopCondition is exited.
func (s *BaseKumirParserListener) ExitEndLoopCondition(ctx *EndLoopConditionContext) {}

// EnterLoopSpecifier is called when production loopSpecifier is entered.
func (s *BaseKumirParserListener) EnterLoopSpecifier(ctx *LoopSpecifierContext) {}

// ExitLoopSpecifier is called when production loopSpecifier is exited.
func (s *BaseKumirParserListener) ExitLoopSpecifier(ctx *LoopSpecifierContext) {}

// EnterLoopStatement is called when production loopStatement is entered.
func (s *BaseKumirParserListener) EnterLoopStatement(ctx *LoopStatementContext) {}

// ExitLoopStatement is called when production loopStatement is exited.
func (s *BaseKumirParserListener) ExitLoopStatement(ctx *LoopStatementContext) {}

// EnterExitStatement is called when production exitStatement is entered.
func (s *BaseKumirParserListener) EnterExitStatement(ctx *ExitStatementContext) {}

// ExitExitStatement is called when production exitStatement is exited.
func (s *BaseKumirParserListener) ExitExitStatement(ctx *ExitStatementContext) {}

// EnterStopStatement is called when production stopStatement is entered.
func (s *BaseKumirParserListener) EnterStopStatement(ctx *StopStatementContext) {}

// ExitStopStatement is called when production stopStatement is exited.
func (s *BaseKumirParserListener) ExitStopStatement(ctx *StopStatementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseKumirParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseKumirParserListener) ExitStatement(ctx *StatementContext) {}

// EnterAlgorithmDefinition is called when production algorithmDefinition is entered.
func (s *BaseKumirParserListener) EnterAlgorithmDefinition(ctx *AlgorithmDefinitionContext) {}

// ExitAlgorithmDefinition is called when production algorithmDefinition is exited.
func (s *BaseKumirParserListener) ExitAlgorithmDefinition(ctx *AlgorithmDefinitionContext) {}

// EnterProgramItem is called when production programItem is entered.
func (s *BaseKumirParserListener) EnterProgramItem(ctx *ProgramItemContext) {}

// ExitProgramItem is called when production programItem is exited.
func (s *BaseKumirParserListener) ExitProgramItem(ctx *ProgramItemContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseKumirParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseKumirParserListener) ExitProgram(ctx *ProgramContext) {}
