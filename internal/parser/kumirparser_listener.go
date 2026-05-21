// Code generated from KumirParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // KumirParser
import "github.com/antlr4-go/antlr/v4"

// KumirParserListener is a complete listener for a parse tree produced by KumirParser.
type KumirParserListener interface {
	antlr.ParseTreeListener

	// EnterQualifiedIdentifier is called when entering the qualifiedIdentifier production.
	EnterQualifiedIdentifier(c *QualifiedIdentifierContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterIndexList is called when entering the indexList production.
	EnterIndexList(c *IndexListContext)

	// EnterPostfixExpression is called when entering the postfixExpression production.
	EnterPostfixExpression(c *PostfixExpressionContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterPowerExpression is called when entering the powerExpression production.
	EnterPowerExpression(c *PowerExpressionContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterLogicalAndExpression is called when entering the logicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterLogicalOrExpression is called when entering the logicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTypeSpecifier is called when entering the typeSpecifier production.
	EnterTypeSpecifier(c *TypeSpecifierContext)

	// EnterBasicType is called when entering the basicType production.
	EnterBasicType(c *BasicTypeContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterArrayBounds is called when entering the arrayBounds production.
	EnterArrayBounds(c *ArrayBoundsContext)

	// EnterVariableDeclarationItem is called when entering the variableDeclarationItem production.
	EnterVariableDeclarationItem(c *VariableDeclarationItemContext)

	// EnterVariableList is called when entering the variableList production.
	EnterVariableList(c *VariableListContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterGlobalDeclaration is called when entering the globalDeclaration production.
	EnterGlobalDeclaration(c *GlobalDeclarationContext)

	// EnterGlobalAssignment is called when entering the globalAssignment production.
	EnterGlobalAssignment(c *GlobalAssignmentContext)

	// EnterParameterDeclaration is called when entering the parameterDeclaration production.
	EnterParameterDeclaration(c *ParameterDeclarationContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterAlgorithmNameTokens is called when entering the algorithmNameTokens production.
	EnterAlgorithmNameTokens(c *AlgorithmNameTokensContext)

	// EnterAlgorithmName is called when entering the algorithmName production.
	EnterAlgorithmName(c *AlgorithmNameContext)

	// EnterAlgorithmHeader is called when entering the algorithmHeader production.
	EnterAlgorithmHeader(c *AlgorithmHeaderContext)

	// EnterAlgorithmBody is called when entering the algorithmBody production.
	EnterAlgorithmBody(c *AlgorithmBodyContext)

	// EnterStatementSequence is called when entering the statementSequence production.
	EnterStatementSequence(c *StatementSequenceContext)

	// EnterLvalue is called when entering the lvalue production.
	EnterLvalue(c *LvalueContext)

	// EnterAssignmentStatement is called when entering the assignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterIoArgument is called when entering the ioArgument production.
	EnterIoArgument(c *IoArgumentContext)

	// EnterIoArgumentList is called when entering the ioArgumentList production.
	EnterIoArgumentList(c *IoArgumentListContext)

	// EnterIoStatement is called when entering the ioStatement production.
	EnterIoStatement(c *IoStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterCaseBlock is called when entering the caseBlock production.
	EnterCaseBlock(c *CaseBlockContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterEndLoopCondition is called when entering the endLoopCondition production.
	EnterEndLoopCondition(c *EndLoopConditionContext)

	// EnterLoopSpecifier is called when entering the loopSpecifier production.
	EnterLoopSpecifier(c *LoopSpecifierContext)

	// EnterLoopStatement is called when entering the loopStatement production.
	EnterLoopStatement(c *LoopStatementContext)

	// EnterExitStatement is called when entering the exitStatement production.
	EnterExitStatement(c *ExitStatementContext)

	// EnterStopStatement is called when entering the stopStatement production.
	EnterStopStatement(c *StopStatementContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAlgorithmDefinition is called when entering the algorithmDefinition production.
	EnterAlgorithmDefinition(c *AlgorithmDefinitionContext)

	// EnterProgramItem is called when entering the programItem production.
	EnterProgramItem(c *ProgramItemContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// ExitQualifiedIdentifier is called when exiting the qualifiedIdentifier production.
	ExitQualifiedIdentifier(c *QualifiedIdentifierContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitIndexList is called when exiting the indexList production.
	ExitIndexList(c *IndexListContext)

	// ExitPostfixExpression is called when exiting the postfixExpression production.
	ExitPostfixExpression(c *PostfixExpressionContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitPowerExpression is called when exiting the powerExpression production.
	ExitPowerExpression(c *PowerExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitLogicalAndExpression is called when exiting the logicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitLogicalOrExpression is called when exiting the logicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTypeSpecifier is called when exiting the typeSpecifier production.
	ExitTypeSpecifier(c *TypeSpecifierContext)

	// ExitBasicType is called when exiting the basicType production.
	ExitBasicType(c *BasicTypeContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitArrayBounds is called when exiting the arrayBounds production.
	ExitArrayBounds(c *ArrayBoundsContext)

	// ExitVariableDeclarationItem is called when exiting the variableDeclarationItem production.
	ExitVariableDeclarationItem(c *VariableDeclarationItemContext)

	// ExitVariableList is called when exiting the variableList production.
	ExitVariableList(c *VariableListContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitGlobalDeclaration is called when exiting the globalDeclaration production.
	ExitGlobalDeclaration(c *GlobalDeclarationContext)

	// ExitGlobalAssignment is called when exiting the globalAssignment production.
	ExitGlobalAssignment(c *GlobalAssignmentContext)

	// ExitParameterDeclaration is called when exiting the parameterDeclaration production.
	ExitParameterDeclaration(c *ParameterDeclarationContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitAlgorithmNameTokens is called when exiting the algorithmNameTokens production.
	ExitAlgorithmNameTokens(c *AlgorithmNameTokensContext)

	// ExitAlgorithmName is called when exiting the algorithmName production.
	ExitAlgorithmName(c *AlgorithmNameContext)

	// ExitAlgorithmHeader is called when exiting the algorithmHeader production.
	ExitAlgorithmHeader(c *AlgorithmHeaderContext)

	// ExitAlgorithmBody is called when exiting the algorithmBody production.
	ExitAlgorithmBody(c *AlgorithmBodyContext)

	// ExitStatementSequence is called when exiting the statementSequence production.
	ExitStatementSequence(c *StatementSequenceContext)

	// ExitLvalue is called when exiting the lvalue production.
	ExitLvalue(c *LvalueContext)

	// ExitAssignmentStatement is called when exiting the assignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitIoArgument is called when exiting the ioArgument production.
	ExitIoArgument(c *IoArgumentContext)

	// ExitIoArgumentList is called when exiting the ioArgumentList production.
	ExitIoArgumentList(c *IoArgumentListContext)

	// ExitIoStatement is called when exiting the ioStatement production.
	ExitIoStatement(c *IoStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitCaseBlock is called when exiting the caseBlock production.
	ExitCaseBlock(c *CaseBlockContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitEndLoopCondition is called when exiting the endLoopCondition production.
	ExitEndLoopCondition(c *EndLoopConditionContext)

	// ExitLoopSpecifier is called when exiting the loopSpecifier production.
	ExitLoopSpecifier(c *LoopSpecifierContext)

	// ExitLoopStatement is called when exiting the loopStatement production.
	ExitLoopStatement(c *LoopStatementContext)

	// ExitExitStatement is called when exiting the exitStatement production.
	ExitExitStatement(c *ExitStatementContext)

	// ExitStopStatement is called when exiting the stopStatement production.
	ExitStopStatement(c *StopStatementContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAlgorithmDefinition is called when exiting the algorithmDefinition production.
	ExitAlgorithmDefinition(c *AlgorithmDefinitionContext)

	// ExitProgramItem is called when exiting the programItem production.
	ExitProgramItem(c *ProgramItemContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)
}
