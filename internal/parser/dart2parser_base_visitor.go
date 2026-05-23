// Code generated from Dart2Parser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Dart2Parser
import "github.com/antlr4-go/antlr/v4"

type BaseDart2ParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDart2ParserVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitTopLevelItem(ctx *TopLevelItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitImportOrExportStatement(ctx *ImportOrExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitLibraryStatement(ctx *LibraryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitPartStatement(ctx *PartStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitTopLevelDeclaration(ctx *TopLevelDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitClassBody(ctx *ClassBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitClassMemberDeclaration(ctx *ClassMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitSuperclass(ctx *SuperclassContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitMixinTypes(ctx *MixinTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitMixinApplication(ctx *MixinApplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitInterfaces(ctx *InterfacesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitEnumEntry(ctx *EnumEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitTypeAliasDeclaration(ctx *TypeAliasDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitFunctionSignature(ctx *FunctionSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitMixinDeclaration(ctx *MixinDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitExtensionDeclaration(ctx *ExtensionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitGetterDeclaration(ctx *GetterDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitSetterDeclaration(ctx *SetterDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitMethodSignature(ctx *MethodSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitVariableDeclarationItem(ctx *VariableDeclarationItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitStatements(ctx *StatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitNonLabelledStatement(ctx *NonLabelledStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitForLoopParts(ctx *ForLoopPartsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitDoStatement(ctx *DoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitSwitchCase(ctx *SwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitDefaultCase(ctx *DefaultCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitTryStatement(ctx *TryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitOnPart(ctx *OnPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitFinallyPart(ctx *FinallyPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitYieldStatement(ctx *YieldStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitYieldEachStatement(ctx *YieldEachStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitAssertStatement(ctx *AssertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitLabel(ctx *LabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitLocalFunctionDeclaration(ctx *LocalFunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitRethrowStatement(ctx *RethrowStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitExpressionWithoutCascade(ctx *ExpressionWithoutCascadeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitAssignableExpression(ctx *AssignableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitAssignableSelectorPart(ctx *AssignableSelectorPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitAssignableSelector(ctx *AssignableSelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitUnconditionalAssignableSelector(ctx *UnconditionalAssignableSelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitCascade(ctx *CascadeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitCascadeSection(ctx *CascadeSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitCascadeSelector(ctx *CascadeSelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitCascadeSectionTail(ctx *CascadeSectionTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitCascadeAssignment(ctx *CascadeAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitIfNullExpression(ctx *IfNullExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitBitwiseOrExpression(ctx *BitwiseOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitBitwiseXorExpression(ctx *BitwiseXorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitBitwiseAndExpression(ctx *BitwiseAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitShiftExpression(ctx *ShiftExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitAwaitExpression(ctx *AwaitExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitThisExpression(ctx *ThisExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitNewExpr(ctx *NewExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitConstExpression(ctx *ConstExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitFunctionExpression(ctx *FunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitFunctionExpressionBody(ctx *FunctionExpressionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitConstructorDesignation(ctx *ConstructorDesignationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitThrowExpression(ctx *ThrowExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitThrowExpressionWithoutCascade(ctx *ThrowExpressionWithoutCascadeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitTypeNotVoid(ctx *TypeNotVoidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitTypeNotVoidNotFunction(ctx *TypeNotVoidNotFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitTypeNotFunction(ctx *TypeNotFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitTypeIdentifier(ctx *TypeIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitTypeArguments(ctx *TypeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitReturnType(ctx *ReturnTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitTypeParameters(ctx *TypeParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitTypeParameterList(ctx *TypeParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitFunctionTypeTails(ctx *FunctionTypeTailsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitFunctionTypeTail(ctx *FunctionTypeTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitParameterTypeList(ctx *ParameterTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitNormalParameterTypes(ctx *NormalParameterTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitOptionalParameterTypes(ctx *OptionalParameterTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitOptionalPositionalParameterTypes(ctx *OptionalPositionalParameterTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitNamedParameterTypes(ctx *NamedParameterTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitNamedParameterType(ctx *NamedParameterTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitTypedIdentifier(ctx *TypedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitFormalParameterList(ctx *FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitNormalFormalParameters(ctx *NormalFormalParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitOptionalOrNamedFormalParameters(ctx *OptionalOrNamedFormalParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitOptionalPositionalFormalParameters(ctx *OptionalPositionalFormalParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitNamedFormalParameters(ctx *NamedFormalParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitNormalFormalParameter(ctx *NormalFormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitNormalFormalParameterNoMetadata(ctx *NormalFormalParameterNoMetadataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitFunctionFormalParameter(ctx *FunctionFormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitFieldFormalParameter(ctx *FieldFormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitFormalParameterPart(ctx *FormalParameterPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitSimpleFormalParameter(ctx *SimpleFormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitDeclaredIdentifier(ctx *DeclaredIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitFinalConstVarOrType(ctx *FinalConstVarOrTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitDefaultFormalParameter(ctx *DefaultFormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitDefaultNamedParameter(ctx *DefaultNamedParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitArgumentPart(ctx *ArgumentPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitNamedArgument(ctx *NamedArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitSymbolLiteral(ctx *SymbolLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitListLiteral(ctx *ListLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitSetOrMapLiteral(ctx *SetOrMapLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitElements(ctx *ElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitExpressionElement(ctx *ExpressionElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitMapElement(ctx *MapElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitSpreadElement(ctx *SpreadElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitIfElement(ctx *IfElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitForElement(ctx *ForElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitOperator(ctx *OperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitMetadata(ctx *MetadataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitMetadatum(ctx *MetadatumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitDottedIdentifier(ctx *DottedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDart2ParserVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}
