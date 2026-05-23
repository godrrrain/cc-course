// Code generated from Dart2Parser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Dart2Parser
import "github.com/antlr4-go/antlr/v4"

// BaseDart2ParserListener is a complete listener for a parse tree produced by Dart2Parser.
type BaseDart2ParserListener struct{}

var _ Dart2ParserListener = &BaseDart2ParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDart2ParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDart2ParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDart2ParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDart2ParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseDart2ParserListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseDart2ParserListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterTopLevelItem is called when production topLevelItem is entered.
func (s *BaseDart2ParserListener) EnterTopLevelItem(ctx *TopLevelItemContext) {}

// ExitTopLevelItem is called when production topLevelItem is exited.
func (s *BaseDart2ParserListener) ExitTopLevelItem(ctx *TopLevelItemContext) {}

// EnterImportOrExportStatement is called when production importOrExportStatement is entered.
func (s *BaseDart2ParserListener) EnterImportOrExportStatement(ctx *ImportOrExportStatementContext) {}

// ExitImportOrExportStatement is called when production importOrExportStatement is exited.
func (s *BaseDart2ParserListener) ExitImportOrExportStatement(ctx *ImportOrExportStatementContext) {}

// EnterLibraryStatement is called when production libraryStatement is entered.
func (s *BaseDart2ParserListener) EnterLibraryStatement(ctx *LibraryStatementContext) {}

// ExitLibraryStatement is called when production libraryStatement is exited.
func (s *BaseDart2ParserListener) ExitLibraryStatement(ctx *LibraryStatementContext) {}

// EnterPartStatement is called when production partStatement is entered.
func (s *BaseDart2ParserListener) EnterPartStatement(ctx *PartStatementContext) {}

// ExitPartStatement is called when production partStatement is exited.
func (s *BaseDart2ParserListener) ExitPartStatement(ctx *PartStatementContext) {}

// EnterTopLevelDeclaration is called when production topLevelDeclaration is entered.
func (s *BaseDart2ParserListener) EnterTopLevelDeclaration(ctx *TopLevelDeclarationContext) {}

// ExitTopLevelDeclaration is called when production topLevelDeclaration is exited.
func (s *BaseDart2ParserListener) ExitTopLevelDeclaration(ctx *TopLevelDeclarationContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseDart2ParserListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseDart2ParserListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BaseDart2ParserListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BaseDart2ParserListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterClassMemberDeclaration is called when production classMemberDeclaration is entered.
func (s *BaseDart2ParserListener) EnterClassMemberDeclaration(ctx *ClassMemberDeclarationContext) {}

// ExitClassMemberDeclaration is called when production classMemberDeclaration is exited.
func (s *BaseDart2ParserListener) ExitClassMemberDeclaration(ctx *ClassMemberDeclarationContext) {}

// EnterSuperclass is called when production superclass is entered.
func (s *BaseDart2ParserListener) EnterSuperclass(ctx *SuperclassContext) {}

// ExitSuperclass is called when production superclass is exited.
func (s *BaseDart2ParserListener) ExitSuperclass(ctx *SuperclassContext) {}

// EnterMixinTypes is called when production mixinTypes is entered.
func (s *BaseDart2ParserListener) EnterMixinTypes(ctx *MixinTypesContext) {}

// ExitMixinTypes is called when production mixinTypes is exited.
func (s *BaseDart2ParserListener) ExitMixinTypes(ctx *MixinTypesContext) {}

// EnterMixinApplication is called when production mixinApplication is entered.
func (s *BaseDart2ParserListener) EnterMixinApplication(ctx *MixinApplicationContext) {}

// ExitMixinApplication is called when production mixinApplication is exited.
func (s *BaseDart2ParserListener) ExitMixinApplication(ctx *MixinApplicationContext) {}

// EnterInterfaces is called when production interfaces is entered.
func (s *BaseDart2ParserListener) EnterInterfaces(ctx *InterfacesContext) {}

// ExitInterfaces is called when production interfaces is exited.
func (s *BaseDart2ParserListener) ExitInterfaces(ctx *InterfacesContext) {}

// EnterEnumDeclaration is called when production enumDeclaration is entered.
func (s *BaseDart2ParserListener) EnterEnumDeclaration(ctx *EnumDeclarationContext) {}

// ExitEnumDeclaration is called when production enumDeclaration is exited.
func (s *BaseDart2ParserListener) ExitEnumDeclaration(ctx *EnumDeclarationContext) {}

// EnterEnumEntry is called when production enumEntry is entered.
func (s *BaseDart2ParserListener) EnterEnumEntry(ctx *EnumEntryContext) {}

// ExitEnumEntry is called when production enumEntry is exited.
func (s *BaseDart2ParserListener) ExitEnumEntry(ctx *EnumEntryContext) {}

// EnterTypeAliasDeclaration is called when production typeAliasDeclaration is entered.
func (s *BaseDart2ParserListener) EnterTypeAliasDeclaration(ctx *TypeAliasDeclarationContext) {}

// ExitTypeAliasDeclaration is called when production typeAliasDeclaration is exited.
func (s *BaseDart2ParserListener) ExitTypeAliasDeclaration(ctx *TypeAliasDeclarationContext) {}

// EnterFunctionSignature is called when production functionSignature is entered.
func (s *BaseDart2ParserListener) EnterFunctionSignature(ctx *FunctionSignatureContext) {}

// ExitFunctionSignature is called when production functionSignature is exited.
func (s *BaseDart2ParserListener) ExitFunctionSignature(ctx *FunctionSignatureContext) {}

// EnterMixinDeclaration is called when production mixinDeclaration is entered.
func (s *BaseDart2ParserListener) EnterMixinDeclaration(ctx *MixinDeclarationContext) {}

// ExitMixinDeclaration is called when production mixinDeclaration is exited.
func (s *BaseDart2ParserListener) ExitMixinDeclaration(ctx *MixinDeclarationContext) {}

// EnterExtensionDeclaration is called when production extensionDeclaration is entered.
func (s *BaseDart2ParserListener) EnterExtensionDeclaration(ctx *ExtensionDeclarationContext) {}

// ExitExtensionDeclaration is called when production extensionDeclaration is exited.
func (s *BaseDart2ParserListener) ExitExtensionDeclaration(ctx *ExtensionDeclarationContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseDart2ParserListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseDart2ParserListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterGetterDeclaration is called when production getterDeclaration is entered.
func (s *BaseDart2ParserListener) EnterGetterDeclaration(ctx *GetterDeclarationContext) {}

// ExitGetterDeclaration is called when production getterDeclaration is exited.
func (s *BaseDart2ParserListener) ExitGetterDeclaration(ctx *GetterDeclarationContext) {}

// EnterSetterDeclaration is called when production setterDeclaration is entered.
func (s *BaseDart2ParserListener) EnterSetterDeclaration(ctx *SetterDeclarationContext) {}

// ExitSetterDeclaration is called when production setterDeclaration is exited.
func (s *BaseDart2ParserListener) ExitSetterDeclaration(ctx *SetterDeclarationContext) {}

// EnterMethodSignature is called when production methodSignature is entered.
func (s *BaseDart2ParserListener) EnterMethodSignature(ctx *MethodSignatureContext) {}

// ExitMethodSignature is called when production methodSignature is exited.
func (s *BaseDart2ParserListener) ExitMethodSignature(ctx *MethodSignatureContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseDart2ParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseDart2ParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseDart2ParserListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseDart2ParserListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterVariableDeclarationList is called when production variableDeclarationList is entered.
func (s *BaseDart2ParserListener) EnterVariableDeclarationList(ctx *VariableDeclarationListContext) {}

// ExitVariableDeclarationList is called when production variableDeclarationList is exited.
func (s *BaseDart2ParserListener) ExitVariableDeclarationList(ctx *VariableDeclarationListContext) {}

// EnterVariableDeclarationItem is called when production variableDeclarationItem is entered.
func (s *BaseDart2ParserListener) EnterVariableDeclarationItem(ctx *VariableDeclarationItemContext) {}

// ExitVariableDeclarationItem is called when production variableDeclarationItem is exited.
func (s *BaseDart2ParserListener) ExitVariableDeclarationItem(ctx *VariableDeclarationItemContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseDart2ParserListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseDart2ParserListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseDart2ParserListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseDart2ParserListener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseDart2ParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseDart2ParserListener) ExitStatement(ctx *StatementContext) {}

// EnterNonLabelledStatement is called when production nonLabelledStatement is entered.
func (s *BaseDart2ParserListener) EnterNonLabelledStatement(ctx *NonLabelledStatementContext) {}

// ExitNonLabelledStatement is called when production nonLabelledStatement is exited.
func (s *BaseDart2ParserListener) ExitNonLabelledStatement(ctx *NonLabelledStatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseDart2ParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseDart2ParserListener) ExitBlock(ctx *BlockContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseDart2ParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseDart2ParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseDart2ParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseDart2ParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForLoopParts is called when production forLoopParts is entered.
func (s *BaseDart2ParserListener) EnterForLoopParts(ctx *ForLoopPartsContext) {}

// ExitForLoopParts is called when production forLoopParts is exited.
func (s *BaseDart2ParserListener) ExitForLoopParts(ctx *ForLoopPartsContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseDart2ParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseDart2ParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterDoStatement is called when production doStatement is entered.
func (s *BaseDart2ParserListener) EnterDoStatement(ctx *DoStatementContext) {}

// ExitDoStatement is called when production doStatement is exited.
func (s *BaseDart2ParserListener) ExitDoStatement(ctx *DoStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseDart2ParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseDart2ParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterSwitchCase is called when production switchCase is entered.
func (s *BaseDart2ParserListener) EnterSwitchCase(ctx *SwitchCaseContext) {}

// ExitSwitchCase is called when production switchCase is exited.
func (s *BaseDart2ParserListener) ExitSwitchCase(ctx *SwitchCaseContext) {}

// EnterDefaultCase is called when production defaultCase is entered.
func (s *BaseDart2ParserListener) EnterDefaultCase(ctx *DefaultCaseContext) {}

// ExitDefaultCase is called when production defaultCase is exited.
func (s *BaseDart2ParserListener) ExitDefaultCase(ctx *DefaultCaseContext) {}

// EnterTryStatement is called when production tryStatement is entered.
func (s *BaseDart2ParserListener) EnterTryStatement(ctx *TryStatementContext) {}

// ExitTryStatement is called when production tryStatement is exited.
func (s *BaseDart2ParserListener) ExitTryStatement(ctx *TryStatementContext) {}

// EnterOnPart is called when production onPart is entered.
func (s *BaseDart2ParserListener) EnterOnPart(ctx *OnPartContext) {}

// ExitOnPart is called when production onPart is exited.
func (s *BaseDart2ParserListener) ExitOnPart(ctx *OnPartContext) {}

// EnterFinallyPart is called when production finallyPart is entered.
func (s *BaseDart2ParserListener) EnterFinallyPart(ctx *FinallyPartContext) {}

// ExitFinallyPart is called when production finallyPart is exited.
func (s *BaseDart2ParserListener) ExitFinallyPart(ctx *FinallyPartContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseDart2ParserListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseDart2ParserListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseDart2ParserListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseDart2ParserListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseDart2ParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseDart2ParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterYieldStatement is called when production yieldStatement is entered.
func (s *BaseDart2ParserListener) EnterYieldStatement(ctx *YieldStatementContext) {}

// ExitYieldStatement is called when production yieldStatement is exited.
func (s *BaseDart2ParserListener) ExitYieldStatement(ctx *YieldStatementContext) {}

// EnterYieldEachStatement is called when production yieldEachStatement is entered.
func (s *BaseDart2ParserListener) EnterYieldEachStatement(ctx *YieldEachStatementContext) {}

// ExitYieldEachStatement is called when production yieldEachStatement is exited.
func (s *BaseDart2ParserListener) ExitYieldEachStatement(ctx *YieldEachStatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseDart2ParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseDart2ParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterAssertStatement is called when production assertStatement is entered.
func (s *BaseDart2ParserListener) EnterAssertStatement(ctx *AssertStatementContext) {}

// ExitAssertStatement is called when production assertStatement is exited.
func (s *BaseDart2ParserListener) ExitAssertStatement(ctx *AssertStatementContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseDart2ParserListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseDart2ParserListener) ExitLabel(ctx *LabelContext) {}

// EnterLocalFunctionDeclaration is called when production localFunctionDeclaration is entered.
func (s *BaseDart2ParserListener) EnterLocalFunctionDeclaration(ctx *LocalFunctionDeclarationContext) {
}

// ExitLocalFunctionDeclaration is called when production localFunctionDeclaration is exited.
func (s *BaseDart2ParserListener) ExitLocalFunctionDeclaration(ctx *LocalFunctionDeclarationContext) {
}

// EnterRethrowStatement is called when production rethrowStatement is entered.
func (s *BaseDart2ParserListener) EnterRethrowStatement(ctx *RethrowStatementContext) {}

// ExitRethrowStatement is called when production rethrowStatement is exited.
func (s *BaseDart2ParserListener) ExitRethrowStatement(ctx *RethrowStatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDart2ParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDart2ParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpressionWithoutCascade is called when production expressionWithoutCascade is entered.
func (s *BaseDart2ParserListener) EnterExpressionWithoutCascade(ctx *ExpressionWithoutCascadeContext) {
}

// ExitExpressionWithoutCascade is called when production expressionWithoutCascade is exited.
func (s *BaseDart2ParserListener) ExitExpressionWithoutCascade(ctx *ExpressionWithoutCascadeContext) {
}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseDart2ParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseDart2ParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterAssignableExpression is called when production assignableExpression is entered.
func (s *BaseDart2ParserListener) EnterAssignableExpression(ctx *AssignableExpressionContext) {}

// ExitAssignableExpression is called when production assignableExpression is exited.
func (s *BaseDart2ParserListener) ExitAssignableExpression(ctx *AssignableExpressionContext) {}

// EnterAssignableSelectorPart is called when production assignableSelectorPart is entered.
func (s *BaseDart2ParserListener) EnterAssignableSelectorPart(ctx *AssignableSelectorPartContext) {}

// ExitAssignableSelectorPart is called when production assignableSelectorPart is exited.
func (s *BaseDart2ParserListener) ExitAssignableSelectorPart(ctx *AssignableSelectorPartContext) {}

// EnterAssignableSelector is called when production assignableSelector is entered.
func (s *BaseDart2ParserListener) EnterAssignableSelector(ctx *AssignableSelectorContext) {}

// ExitAssignableSelector is called when production assignableSelector is exited.
func (s *BaseDart2ParserListener) ExitAssignableSelector(ctx *AssignableSelectorContext) {}

// EnterUnconditionalAssignableSelector is called when production unconditionalAssignableSelector is entered.
func (s *BaseDart2ParserListener) EnterUnconditionalAssignableSelector(ctx *UnconditionalAssignableSelectorContext) {
}

// ExitUnconditionalAssignableSelector is called when production unconditionalAssignableSelector is exited.
func (s *BaseDart2ParserListener) ExitUnconditionalAssignableSelector(ctx *UnconditionalAssignableSelectorContext) {
}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseDart2ParserListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseDart2ParserListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterCascade is called when production cascade is entered.
func (s *BaseDart2ParserListener) EnterCascade(ctx *CascadeContext) {}

// ExitCascade is called when production cascade is exited.
func (s *BaseDart2ParserListener) ExitCascade(ctx *CascadeContext) {}

// EnterCascadeSection is called when production cascadeSection is entered.
func (s *BaseDart2ParserListener) EnterCascadeSection(ctx *CascadeSectionContext) {}

// ExitCascadeSection is called when production cascadeSection is exited.
func (s *BaseDart2ParserListener) ExitCascadeSection(ctx *CascadeSectionContext) {}

// EnterCascadeSelector is called when production cascadeSelector is entered.
func (s *BaseDart2ParserListener) EnterCascadeSelector(ctx *CascadeSelectorContext) {}

// ExitCascadeSelector is called when production cascadeSelector is exited.
func (s *BaseDart2ParserListener) ExitCascadeSelector(ctx *CascadeSelectorContext) {}

// EnterCascadeSectionTail is called when production cascadeSectionTail is entered.
func (s *BaseDart2ParserListener) EnterCascadeSectionTail(ctx *CascadeSectionTailContext) {}

// ExitCascadeSectionTail is called when production cascadeSectionTail is exited.
func (s *BaseDart2ParserListener) ExitCascadeSectionTail(ctx *CascadeSectionTailContext) {}

// EnterCascadeAssignment is called when production cascadeAssignment is entered.
func (s *BaseDart2ParserListener) EnterCascadeAssignment(ctx *CascadeAssignmentContext) {}

// ExitCascadeAssignment is called when production cascadeAssignment is exited.
func (s *BaseDart2ParserListener) ExitCascadeAssignment(ctx *CascadeAssignmentContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *BaseDart2ParserListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *BaseDart2ParserListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterIfNullExpression is called when production ifNullExpression is entered.
func (s *BaseDart2ParserListener) EnterIfNullExpression(ctx *IfNullExpressionContext) {}

// ExitIfNullExpression is called when production ifNullExpression is exited.
func (s *BaseDart2ParserListener) ExitIfNullExpression(ctx *IfNullExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *BaseDart2ParserListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *BaseDart2ParserListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *BaseDart2ParserListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *BaseDart2ParserListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterBitwiseOrExpression is called when production bitwiseOrExpression is entered.
func (s *BaseDart2ParserListener) EnterBitwiseOrExpression(ctx *BitwiseOrExpressionContext) {}

// ExitBitwiseOrExpression is called when production bitwiseOrExpression is exited.
func (s *BaseDart2ParserListener) ExitBitwiseOrExpression(ctx *BitwiseOrExpressionContext) {}

// EnterBitwiseXorExpression is called when production bitwiseXorExpression is entered.
func (s *BaseDart2ParserListener) EnterBitwiseXorExpression(ctx *BitwiseXorExpressionContext) {}

// ExitBitwiseXorExpression is called when production bitwiseXorExpression is exited.
func (s *BaseDart2ParserListener) ExitBitwiseXorExpression(ctx *BitwiseXorExpressionContext) {}

// EnterBitwiseAndExpression is called when production bitwiseAndExpression is entered.
func (s *BaseDart2ParserListener) EnterBitwiseAndExpression(ctx *BitwiseAndExpressionContext) {}

// ExitBitwiseAndExpression is called when production bitwiseAndExpression is exited.
func (s *BaseDart2ParserListener) ExitBitwiseAndExpression(ctx *BitwiseAndExpressionContext) {}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *BaseDart2ParserListener) EnterShiftExpression(ctx *ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *BaseDart2ParserListener) ExitShiftExpression(ctx *ShiftExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseDart2ParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseDart2ParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseDart2ParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseDart2ParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseDart2ParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseDart2ParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterAwaitExpression is called when production awaitExpression is entered.
func (s *BaseDart2ParserListener) EnterAwaitExpression(ctx *AwaitExpressionContext) {}

// ExitAwaitExpression is called when production awaitExpression is exited.
func (s *BaseDart2ParserListener) ExitAwaitExpression(ctx *AwaitExpressionContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseDart2ParserListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseDart2ParserListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseDart2ParserListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseDart2ParserListener) ExitSelector(ctx *SelectorContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseDart2ParserListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseDart2ParserListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterThisExpression is called when production thisExpression is entered.
func (s *BaseDart2ParserListener) EnterThisExpression(ctx *ThisExpressionContext) {}

// ExitThisExpression is called when production thisExpression is exited.
func (s *BaseDart2ParserListener) ExitThisExpression(ctx *ThisExpressionContext) {}

// EnterNewExpr is called when production newExpr is entered.
func (s *BaseDart2ParserListener) EnterNewExpr(ctx *NewExprContext) {}

// ExitNewExpr is called when production newExpr is exited.
func (s *BaseDart2ParserListener) ExitNewExpr(ctx *NewExprContext) {}

// EnterConstExpression is called when production constExpression is entered.
func (s *BaseDart2ParserListener) EnterConstExpression(ctx *ConstExpressionContext) {}

// ExitConstExpression is called when production constExpression is exited.
func (s *BaseDart2ParserListener) ExitConstExpression(ctx *ConstExpressionContext) {}

// EnterFunctionExpression is called when production functionExpression is entered.
func (s *BaseDart2ParserListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production functionExpression is exited.
func (s *BaseDart2ParserListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterFunctionExpressionBody is called when production functionExpressionBody is entered.
func (s *BaseDart2ParserListener) EnterFunctionExpressionBody(ctx *FunctionExpressionBodyContext) {}

// ExitFunctionExpressionBody is called when production functionExpressionBody is exited.
func (s *BaseDart2ParserListener) ExitFunctionExpressionBody(ctx *FunctionExpressionBodyContext) {}

// EnterConstructorDesignation is called when production constructorDesignation is entered.
func (s *BaseDart2ParserListener) EnterConstructorDesignation(ctx *ConstructorDesignationContext) {}

// ExitConstructorDesignation is called when production constructorDesignation is exited.
func (s *BaseDart2ParserListener) ExitConstructorDesignation(ctx *ConstructorDesignationContext) {}

// EnterThrowExpression is called when production throwExpression is entered.
func (s *BaseDart2ParserListener) EnterThrowExpression(ctx *ThrowExpressionContext) {}

// ExitThrowExpression is called when production throwExpression is exited.
func (s *BaseDart2ParserListener) ExitThrowExpression(ctx *ThrowExpressionContext) {}

// EnterThrowExpressionWithoutCascade is called when production throwExpressionWithoutCascade is entered.
func (s *BaseDart2ParserListener) EnterThrowExpressionWithoutCascade(ctx *ThrowExpressionWithoutCascadeContext) {
}

// ExitThrowExpressionWithoutCascade is called when production throwExpressionWithoutCascade is exited.
func (s *BaseDart2ParserListener) ExitThrowExpressionWithoutCascade(ctx *ThrowExpressionWithoutCascadeContext) {
}

// EnterType is called when production type is entered.
func (s *BaseDart2ParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseDart2ParserListener) ExitType(ctx *TypeContext) {}

// EnterTypeNotVoid is called when production typeNotVoid is entered.
func (s *BaseDart2ParserListener) EnterTypeNotVoid(ctx *TypeNotVoidContext) {}

// ExitTypeNotVoid is called when production typeNotVoid is exited.
func (s *BaseDart2ParserListener) ExitTypeNotVoid(ctx *TypeNotVoidContext) {}

// EnterTypeNotVoidNotFunction is called when production typeNotVoidNotFunction is entered.
func (s *BaseDart2ParserListener) EnterTypeNotVoidNotFunction(ctx *TypeNotVoidNotFunctionContext) {}

// ExitTypeNotVoidNotFunction is called when production typeNotVoidNotFunction is exited.
func (s *BaseDart2ParserListener) ExitTypeNotVoidNotFunction(ctx *TypeNotVoidNotFunctionContext) {}

// EnterTypeNotFunction is called when production typeNotFunction is entered.
func (s *BaseDart2ParserListener) EnterTypeNotFunction(ctx *TypeNotFunctionContext) {}

// ExitTypeNotFunction is called when production typeNotFunction is exited.
func (s *BaseDart2ParserListener) ExitTypeNotFunction(ctx *TypeNotFunctionContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseDart2ParserListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseDart2ParserListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterTypeIdentifier is called when production typeIdentifier is entered.
func (s *BaseDart2ParserListener) EnterTypeIdentifier(ctx *TypeIdentifierContext) {}

// ExitTypeIdentifier is called when production typeIdentifier is exited.
func (s *BaseDart2ParserListener) ExitTypeIdentifier(ctx *TypeIdentifierContext) {}

// EnterTypeArguments is called when production typeArguments is entered.
func (s *BaseDart2ParserListener) EnterTypeArguments(ctx *TypeArgumentsContext) {}

// ExitTypeArguments is called when production typeArguments is exited.
func (s *BaseDart2ParserListener) ExitTypeArguments(ctx *TypeArgumentsContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseDart2ParserListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseDart2ParserListener) ExitTypeList(ctx *TypeListContext) {}

// EnterReturnType is called when production returnType is entered.
func (s *BaseDart2ParserListener) EnterReturnType(ctx *ReturnTypeContext) {}

// ExitReturnType is called when production returnType is exited.
func (s *BaseDart2ParserListener) ExitReturnType(ctx *ReturnTypeContext) {}

// EnterTypeParameters is called when production typeParameters is entered.
func (s *BaseDart2ParserListener) EnterTypeParameters(ctx *TypeParametersContext) {}

// ExitTypeParameters is called when production typeParameters is exited.
func (s *BaseDart2ParserListener) ExitTypeParameters(ctx *TypeParametersContext) {}

// EnterTypeParameterList is called when production typeParameterList is entered.
func (s *BaseDart2ParserListener) EnterTypeParameterList(ctx *TypeParameterListContext) {}

// ExitTypeParameterList is called when production typeParameterList is exited.
func (s *BaseDart2ParserListener) ExitTypeParameterList(ctx *TypeParameterListContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseDart2ParserListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseDart2ParserListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *BaseDart2ParserListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *BaseDart2ParserListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterFunctionTypeTails is called when production functionTypeTails is entered.
func (s *BaseDart2ParserListener) EnterFunctionTypeTails(ctx *FunctionTypeTailsContext) {}

// ExitFunctionTypeTails is called when production functionTypeTails is exited.
func (s *BaseDart2ParserListener) ExitFunctionTypeTails(ctx *FunctionTypeTailsContext) {}

// EnterFunctionTypeTail is called when production functionTypeTail is entered.
func (s *BaseDart2ParserListener) EnterFunctionTypeTail(ctx *FunctionTypeTailContext) {}

// ExitFunctionTypeTail is called when production functionTypeTail is exited.
func (s *BaseDart2ParserListener) ExitFunctionTypeTail(ctx *FunctionTypeTailContext) {}

// EnterParameterTypeList is called when production parameterTypeList is entered.
func (s *BaseDart2ParserListener) EnterParameterTypeList(ctx *ParameterTypeListContext) {}

// ExitParameterTypeList is called when production parameterTypeList is exited.
func (s *BaseDart2ParserListener) ExitParameterTypeList(ctx *ParameterTypeListContext) {}

// EnterNormalParameterTypes is called when production normalParameterTypes is entered.
func (s *BaseDart2ParserListener) EnterNormalParameterTypes(ctx *NormalParameterTypesContext) {}

// ExitNormalParameterTypes is called when production normalParameterTypes is exited.
func (s *BaseDart2ParserListener) ExitNormalParameterTypes(ctx *NormalParameterTypesContext) {}

// EnterOptionalParameterTypes is called when production optionalParameterTypes is entered.
func (s *BaseDart2ParserListener) EnterOptionalParameterTypes(ctx *OptionalParameterTypesContext) {}

// ExitOptionalParameterTypes is called when production optionalParameterTypes is exited.
func (s *BaseDart2ParserListener) ExitOptionalParameterTypes(ctx *OptionalParameterTypesContext) {}

// EnterOptionalPositionalParameterTypes is called when production optionalPositionalParameterTypes is entered.
func (s *BaseDart2ParserListener) EnterOptionalPositionalParameterTypes(ctx *OptionalPositionalParameterTypesContext) {
}

// ExitOptionalPositionalParameterTypes is called when production optionalPositionalParameterTypes is exited.
func (s *BaseDart2ParserListener) ExitOptionalPositionalParameterTypes(ctx *OptionalPositionalParameterTypesContext) {
}

// EnterNamedParameterTypes is called when production namedParameterTypes is entered.
func (s *BaseDart2ParserListener) EnterNamedParameterTypes(ctx *NamedParameterTypesContext) {}

// ExitNamedParameterTypes is called when production namedParameterTypes is exited.
func (s *BaseDart2ParserListener) ExitNamedParameterTypes(ctx *NamedParameterTypesContext) {}

// EnterNamedParameterType is called when production namedParameterType is entered.
func (s *BaseDart2ParserListener) EnterNamedParameterType(ctx *NamedParameterTypeContext) {}

// ExitNamedParameterType is called when production namedParameterType is exited.
func (s *BaseDart2ParserListener) ExitNamedParameterType(ctx *NamedParameterTypeContext) {}

// EnterTypedIdentifier is called when production typedIdentifier is entered.
func (s *BaseDart2ParserListener) EnterTypedIdentifier(ctx *TypedIdentifierContext) {}

// ExitTypedIdentifier is called when production typedIdentifier is exited.
func (s *BaseDart2ParserListener) ExitTypedIdentifier(ctx *TypedIdentifierContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseDart2ParserListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseDart2ParserListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterNormalFormalParameters is called when production normalFormalParameters is entered.
func (s *BaseDart2ParserListener) EnterNormalFormalParameters(ctx *NormalFormalParametersContext) {}

// ExitNormalFormalParameters is called when production normalFormalParameters is exited.
func (s *BaseDart2ParserListener) ExitNormalFormalParameters(ctx *NormalFormalParametersContext) {}

// EnterOptionalOrNamedFormalParameters is called when production optionalOrNamedFormalParameters is entered.
func (s *BaseDart2ParserListener) EnterOptionalOrNamedFormalParameters(ctx *OptionalOrNamedFormalParametersContext) {
}

// ExitOptionalOrNamedFormalParameters is called when production optionalOrNamedFormalParameters is exited.
func (s *BaseDart2ParserListener) ExitOptionalOrNamedFormalParameters(ctx *OptionalOrNamedFormalParametersContext) {
}

// EnterOptionalPositionalFormalParameters is called when production optionalPositionalFormalParameters is entered.
func (s *BaseDart2ParserListener) EnterOptionalPositionalFormalParameters(ctx *OptionalPositionalFormalParametersContext) {
}

// ExitOptionalPositionalFormalParameters is called when production optionalPositionalFormalParameters is exited.
func (s *BaseDart2ParserListener) ExitOptionalPositionalFormalParameters(ctx *OptionalPositionalFormalParametersContext) {
}

// EnterNamedFormalParameters is called when production namedFormalParameters is entered.
func (s *BaseDart2ParserListener) EnterNamedFormalParameters(ctx *NamedFormalParametersContext) {}

// ExitNamedFormalParameters is called when production namedFormalParameters is exited.
func (s *BaseDart2ParserListener) ExitNamedFormalParameters(ctx *NamedFormalParametersContext) {}

// EnterNormalFormalParameter is called when production normalFormalParameter is entered.
func (s *BaseDart2ParserListener) EnterNormalFormalParameter(ctx *NormalFormalParameterContext) {}

// ExitNormalFormalParameter is called when production normalFormalParameter is exited.
func (s *BaseDart2ParserListener) ExitNormalFormalParameter(ctx *NormalFormalParameterContext) {}

// EnterNormalFormalParameterNoMetadata is called when production normalFormalParameterNoMetadata is entered.
func (s *BaseDart2ParserListener) EnterNormalFormalParameterNoMetadata(ctx *NormalFormalParameterNoMetadataContext) {
}

// ExitNormalFormalParameterNoMetadata is called when production normalFormalParameterNoMetadata is exited.
func (s *BaseDart2ParserListener) ExitNormalFormalParameterNoMetadata(ctx *NormalFormalParameterNoMetadataContext) {
}

// EnterFunctionFormalParameter is called when production functionFormalParameter is entered.
func (s *BaseDart2ParserListener) EnterFunctionFormalParameter(ctx *FunctionFormalParameterContext) {}

// ExitFunctionFormalParameter is called when production functionFormalParameter is exited.
func (s *BaseDart2ParserListener) ExitFunctionFormalParameter(ctx *FunctionFormalParameterContext) {}

// EnterFieldFormalParameter is called when production fieldFormalParameter is entered.
func (s *BaseDart2ParserListener) EnterFieldFormalParameter(ctx *FieldFormalParameterContext) {}

// ExitFieldFormalParameter is called when production fieldFormalParameter is exited.
func (s *BaseDart2ParserListener) ExitFieldFormalParameter(ctx *FieldFormalParameterContext) {}

// EnterFormalParameterPart is called when production formalParameterPart is entered.
func (s *BaseDart2ParserListener) EnterFormalParameterPart(ctx *FormalParameterPartContext) {}

// ExitFormalParameterPart is called when production formalParameterPart is exited.
func (s *BaseDart2ParserListener) ExitFormalParameterPart(ctx *FormalParameterPartContext) {}

// EnterSimpleFormalParameter is called when production simpleFormalParameter is entered.
func (s *BaseDart2ParserListener) EnterSimpleFormalParameter(ctx *SimpleFormalParameterContext) {}

// ExitSimpleFormalParameter is called when production simpleFormalParameter is exited.
func (s *BaseDart2ParserListener) ExitSimpleFormalParameter(ctx *SimpleFormalParameterContext) {}

// EnterDeclaredIdentifier is called when production declaredIdentifier is entered.
func (s *BaseDart2ParserListener) EnterDeclaredIdentifier(ctx *DeclaredIdentifierContext) {}

// ExitDeclaredIdentifier is called when production declaredIdentifier is exited.
func (s *BaseDart2ParserListener) ExitDeclaredIdentifier(ctx *DeclaredIdentifierContext) {}

// EnterFinalConstVarOrType is called when production finalConstVarOrType is entered.
func (s *BaseDart2ParserListener) EnterFinalConstVarOrType(ctx *FinalConstVarOrTypeContext) {}

// ExitFinalConstVarOrType is called when production finalConstVarOrType is exited.
func (s *BaseDart2ParserListener) ExitFinalConstVarOrType(ctx *FinalConstVarOrTypeContext) {}

// EnterDefaultFormalParameter is called when production defaultFormalParameter is entered.
func (s *BaseDart2ParserListener) EnterDefaultFormalParameter(ctx *DefaultFormalParameterContext) {}

// ExitDefaultFormalParameter is called when production defaultFormalParameter is exited.
func (s *BaseDart2ParserListener) ExitDefaultFormalParameter(ctx *DefaultFormalParameterContext) {}

// EnterDefaultNamedParameter is called when production defaultNamedParameter is entered.
func (s *BaseDart2ParserListener) EnterDefaultNamedParameter(ctx *DefaultNamedParameterContext) {}

// ExitDefaultNamedParameter is called when production defaultNamedParameter is exited.
func (s *BaseDart2ParserListener) ExitDefaultNamedParameter(ctx *DefaultNamedParameterContext) {}

// EnterArgumentPart is called when production argumentPart is entered.
func (s *BaseDart2ParserListener) EnterArgumentPart(ctx *ArgumentPartContext) {}

// ExitArgumentPart is called when production argumentPart is exited.
func (s *BaseDart2ParserListener) ExitArgumentPart(ctx *ArgumentPartContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseDart2ParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseDart2ParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseDart2ParserListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseDart2ParserListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterNamedArgument is called when production namedArgument is entered.
func (s *BaseDart2ParserListener) EnterNamedArgument(ctx *NamedArgumentContext) {}

// ExitNamedArgument is called when production namedArgument is exited.
func (s *BaseDart2ParserListener) ExitNamedArgument(ctx *NamedArgumentContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseDart2ParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseDart2ParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseDart2ParserListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseDart2ParserListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseDart2ParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseDart2ParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseDart2ParserListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseDart2ParserListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseDart2ParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseDart2ParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterSymbolLiteral is called when production symbolLiteral is entered.
func (s *BaseDart2ParserListener) EnterSymbolLiteral(ctx *SymbolLiteralContext) {}

// ExitSymbolLiteral is called when production symbolLiteral is exited.
func (s *BaseDart2ParserListener) ExitSymbolLiteral(ctx *SymbolLiteralContext) {}

// EnterListLiteral is called when production listLiteral is entered.
func (s *BaseDart2ParserListener) EnterListLiteral(ctx *ListLiteralContext) {}

// ExitListLiteral is called when production listLiteral is exited.
func (s *BaseDart2ParserListener) ExitListLiteral(ctx *ListLiteralContext) {}

// EnterSetOrMapLiteral is called when production setOrMapLiteral is entered.
func (s *BaseDart2ParserListener) EnterSetOrMapLiteral(ctx *SetOrMapLiteralContext) {}

// ExitSetOrMapLiteral is called when production setOrMapLiteral is exited.
func (s *BaseDart2ParserListener) ExitSetOrMapLiteral(ctx *SetOrMapLiteralContext) {}

// EnterElements is called when production elements is entered.
func (s *BaseDart2ParserListener) EnterElements(ctx *ElementsContext) {}

// ExitElements is called when production elements is exited.
func (s *BaseDart2ParserListener) ExitElements(ctx *ElementsContext) {}

// EnterElement is called when production element is entered.
func (s *BaseDart2ParserListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseDart2ParserListener) ExitElement(ctx *ElementContext) {}

// EnterExpressionElement is called when production expressionElement is entered.
func (s *BaseDart2ParserListener) EnterExpressionElement(ctx *ExpressionElementContext) {}

// ExitExpressionElement is called when production expressionElement is exited.
func (s *BaseDart2ParserListener) ExitExpressionElement(ctx *ExpressionElementContext) {}

// EnterMapElement is called when production mapElement is entered.
func (s *BaseDart2ParserListener) EnterMapElement(ctx *MapElementContext) {}

// ExitMapElement is called when production mapElement is exited.
func (s *BaseDart2ParserListener) ExitMapElement(ctx *MapElementContext) {}

// EnterSpreadElement is called when production spreadElement is entered.
func (s *BaseDart2ParserListener) EnterSpreadElement(ctx *SpreadElementContext) {}

// ExitSpreadElement is called when production spreadElement is exited.
func (s *BaseDart2ParserListener) ExitSpreadElement(ctx *SpreadElementContext) {}

// EnterIfElement is called when production ifElement is entered.
func (s *BaseDart2ParserListener) EnterIfElement(ctx *IfElementContext) {}

// ExitIfElement is called when production ifElement is exited.
func (s *BaseDart2ParserListener) ExitIfElement(ctx *IfElementContext) {}

// EnterForElement is called when production forElement is entered.
func (s *BaseDart2ParserListener) EnterForElement(ctx *ForElementContext) {}

// ExitForElement is called when production forElement is exited.
func (s *BaseDart2ParserListener) ExitForElement(ctx *ForElementContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseDart2ParserListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseDart2ParserListener) ExitOperator(ctx *OperatorContext) {}

// EnterMetadata is called when production metadata is entered.
func (s *BaseDart2ParserListener) EnterMetadata(ctx *MetadataContext) {}

// ExitMetadata is called when production metadata is exited.
func (s *BaseDart2ParserListener) ExitMetadata(ctx *MetadataContext) {}

// EnterMetadatum is called when production metadatum is entered.
func (s *BaseDart2ParserListener) EnterMetadatum(ctx *MetadatumContext) {}

// ExitMetadatum is called when production metadatum is exited.
func (s *BaseDart2ParserListener) ExitMetadatum(ctx *MetadatumContext) {}

// EnterDottedIdentifier is called when production dottedIdentifier is entered.
func (s *BaseDart2ParserListener) EnterDottedIdentifier(ctx *DottedIdentifierContext) {}

// ExitDottedIdentifier is called when production dottedIdentifier is exited.
func (s *BaseDart2ParserListener) ExitDottedIdentifier(ctx *DottedIdentifierContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseDart2ParserListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseDart2ParserListener) ExitIdentifierList(ctx *IdentifierListContext) {}
