// Code generated from Dart2Parser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Dart2Parser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by Dart2Parser.
type Dart2ParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Dart2Parser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by Dart2Parser#topLevelItem.
	VisitTopLevelItem(ctx *TopLevelItemContext) interface{}

	// Visit a parse tree produced by Dart2Parser#importOrExportStatement.
	VisitImportOrExportStatement(ctx *ImportOrExportStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#libraryStatement.
	VisitLibraryStatement(ctx *LibraryStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#partStatement.
	VisitPartStatement(ctx *PartStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#topLevelDeclaration.
	VisitTopLevelDeclaration(ctx *TopLevelDeclarationContext) interface{}

	// Visit a parse tree produced by Dart2Parser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by Dart2Parser#classBody.
	VisitClassBody(ctx *ClassBodyContext) interface{}

	// Visit a parse tree produced by Dart2Parser#classMemberDeclaration.
	VisitClassMemberDeclaration(ctx *ClassMemberDeclarationContext) interface{}

	// Visit a parse tree produced by Dart2Parser#superclass.
	VisitSuperclass(ctx *SuperclassContext) interface{}

	// Visit a parse tree produced by Dart2Parser#mixinTypes.
	VisitMixinTypes(ctx *MixinTypesContext) interface{}

	// Visit a parse tree produced by Dart2Parser#mixinApplication.
	VisitMixinApplication(ctx *MixinApplicationContext) interface{}

	// Visit a parse tree produced by Dart2Parser#interfaces.
	VisitInterfaces(ctx *InterfacesContext) interface{}

	// Visit a parse tree produced by Dart2Parser#enumDeclaration.
	VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{}

	// Visit a parse tree produced by Dart2Parser#enumEntry.
	VisitEnumEntry(ctx *EnumEntryContext) interface{}

	// Visit a parse tree produced by Dart2Parser#typeAliasDeclaration.
	VisitTypeAliasDeclaration(ctx *TypeAliasDeclarationContext) interface{}

	// Visit a parse tree produced by Dart2Parser#functionSignature.
	VisitFunctionSignature(ctx *FunctionSignatureContext) interface{}

	// Visit a parse tree produced by Dart2Parser#mixinDeclaration.
	VisitMixinDeclaration(ctx *MixinDeclarationContext) interface{}

	// Visit a parse tree produced by Dart2Parser#extensionDeclaration.
	VisitExtensionDeclaration(ctx *ExtensionDeclarationContext) interface{}

	// Visit a parse tree produced by Dart2Parser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by Dart2Parser#getterDeclaration.
	VisitGetterDeclaration(ctx *GetterDeclarationContext) interface{}

	// Visit a parse tree produced by Dart2Parser#setterDeclaration.
	VisitSetterDeclaration(ctx *SetterDeclarationContext) interface{}

	// Visit a parse tree produced by Dart2Parser#methodSignature.
	VisitMethodSignature(ctx *MethodSignatureContext) interface{}

	// Visit a parse tree produced by Dart2Parser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by Dart2Parser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by Dart2Parser#variableDeclarationList.
	VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{}

	// Visit a parse tree produced by Dart2Parser#variableDeclarationItem.
	VisitVariableDeclarationItem(ctx *VariableDeclarationItemContext) interface{}

	// Visit a parse tree produced by Dart2Parser#functionBody.
	VisitFunctionBody(ctx *FunctionBodyContext) interface{}

	// Visit a parse tree produced by Dart2Parser#statements.
	VisitStatements(ctx *StatementsContext) interface{}

	// Visit a parse tree produced by Dart2Parser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#nonLabelledStatement.
	VisitNonLabelledStatement(ctx *NonLabelledStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by Dart2Parser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#forLoopParts.
	VisitForLoopParts(ctx *ForLoopPartsContext) interface{}

	// Visit a parse tree produced by Dart2Parser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#doStatement.
	VisitDoStatement(ctx *DoStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#switchCase.
	VisitSwitchCase(ctx *SwitchCaseContext) interface{}

	// Visit a parse tree produced by Dart2Parser#defaultCase.
	VisitDefaultCase(ctx *DefaultCaseContext) interface{}

	// Visit a parse tree produced by Dart2Parser#tryStatement.
	VisitTryStatement(ctx *TryStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#onPart.
	VisitOnPart(ctx *OnPartContext) interface{}

	// Visit a parse tree produced by Dart2Parser#finallyPart.
	VisitFinallyPart(ctx *FinallyPartContext) interface{}

	// Visit a parse tree produced by Dart2Parser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#yieldStatement.
	VisitYieldStatement(ctx *YieldStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#yieldEachStatement.
	VisitYieldEachStatement(ctx *YieldEachStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#assertStatement.
	VisitAssertStatement(ctx *AssertStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#label.
	VisitLabel(ctx *LabelContext) interface{}

	// Visit a parse tree produced by Dart2Parser#localFunctionDeclaration.
	VisitLocalFunctionDeclaration(ctx *LocalFunctionDeclarationContext) interface{}

	// Visit a parse tree produced by Dart2Parser#rethrowStatement.
	VisitRethrowStatement(ctx *RethrowStatementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#expressionWithoutCascade.
	VisitExpressionWithoutCascade(ctx *ExpressionWithoutCascadeContext) interface{}

	// Visit a parse tree produced by Dart2Parser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by Dart2Parser#assignableExpression.
	VisitAssignableExpression(ctx *AssignableExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#assignableSelectorPart.
	VisitAssignableSelectorPart(ctx *AssignableSelectorPartContext) interface{}

	// Visit a parse tree produced by Dart2Parser#assignableSelector.
	VisitAssignableSelector(ctx *AssignableSelectorContext) interface{}

	// Visit a parse tree produced by Dart2Parser#unconditionalAssignableSelector.
	VisitUnconditionalAssignableSelector(ctx *UnconditionalAssignableSelectorContext) interface{}

	// Visit a parse tree produced by Dart2Parser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by Dart2Parser#cascade.
	VisitCascade(ctx *CascadeContext) interface{}

	// Visit a parse tree produced by Dart2Parser#cascadeSection.
	VisitCascadeSection(ctx *CascadeSectionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#cascadeSelector.
	VisitCascadeSelector(ctx *CascadeSelectorContext) interface{}

	// Visit a parse tree produced by Dart2Parser#cascadeSectionTail.
	VisitCascadeSectionTail(ctx *CascadeSectionTailContext) interface{}

	// Visit a parse tree produced by Dart2Parser#cascadeAssignment.
	VisitCascadeAssignment(ctx *CascadeAssignmentContext) interface{}

	// Visit a parse tree produced by Dart2Parser#conditionalExpression.
	VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#ifNullExpression.
	VisitIfNullExpression(ctx *IfNullExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#logicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#logicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#bitwiseOrExpression.
	VisitBitwiseOrExpression(ctx *BitwiseOrExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#bitwiseXorExpression.
	VisitBitwiseXorExpression(ctx *BitwiseXorExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#bitwiseAndExpression.
	VisitBitwiseAndExpression(ctx *BitwiseAndExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#shiftExpression.
	VisitShiftExpression(ctx *ShiftExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#awaitExpression.
	VisitAwaitExpression(ctx *AwaitExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#postfixExpression.
	VisitPostfixExpression(ctx *PostfixExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by Dart2Parser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by Dart2Parser#thisExpression.
	VisitThisExpression(ctx *ThisExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#newExpr.
	VisitNewExpr(ctx *NewExprContext) interface{}

	// Visit a parse tree produced by Dart2Parser#constExpression.
	VisitConstExpression(ctx *ConstExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#functionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#functionExpressionBody.
	VisitFunctionExpressionBody(ctx *FunctionExpressionBodyContext) interface{}

	// Visit a parse tree produced by Dart2Parser#constructorDesignation.
	VisitConstructorDesignation(ctx *ConstructorDesignationContext) interface{}

	// Visit a parse tree produced by Dart2Parser#throwExpression.
	VisitThrowExpression(ctx *ThrowExpressionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#throwExpressionWithoutCascade.
	VisitThrowExpressionWithoutCascade(ctx *ThrowExpressionWithoutCascadeContext) interface{}

	// Visit a parse tree produced by Dart2Parser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by Dart2Parser#typeNotVoid.
	VisitTypeNotVoid(ctx *TypeNotVoidContext) interface{}

	// Visit a parse tree produced by Dart2Parser#typeNotVoidNotFunction.
	VisitTypeNotVoidNotFunction(ctx *TypeNotVoidNotFunctionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#typeNotFunction.
	VisitTypeNotFunction(ctx *TypeNotFunctionContext) interface{}

	// Visit a parse tree produced by Dart2Parser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by Dart2Parser#typeIdentifier.
	VisitTypeIdentifier(ctx *TypeIdentifierContext) interface{}

	// Visit a parse tree produced by Dart2Parser#typeArguments.
	VisitTypeArguments(ctx *TypeArgumentsContext) interface{}

	// Visit a parse tree produced by Dart2Parser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by Dart2Parser#returnType.
	VisitReturnType(ctx *ReturnTypeContext) interface{}

	// Visit a parse tree produced by Dart2Parser#typeParameters.
	VisitTypeParameters(ctx *TypeParametersContext) interface{}

	// Visit a parse tree produced by Dart2Parser#typeParameterList.
	VisitTypeParameterList(ctx *TypeParameterListContext) interface{}

	// Visit a parse tree produced by Dart2Parser#typeParameter.
	VisitTypeParameter(ctx *TypeParameterContext) interface{}

	// Visit a parse tree produced by Dart2Parser#functionType.
	VisitFunctionType(ctx *FunctionTypeContext) interface{}

	// Visit a parse tree produced by Dart2Parser#functionTypeTails.
	VisitFunctionTypeTails(ctx *FunctionTypeTailsContext) interface{}

	// Visit a parse tree produced by Dart2Parser#functionTypeTail.
	VisitFunctionTypeTail(ctx *FunctionTypeTailContext) interface{}

	// Visit a parse tree produced by Dart2Parser#parameterTypeList.
	VisitParameterTypeList(ctx *ParameterTypeListContext) interface{}

	// Visit a parse tree produced by Dart2Parser#normalParameterTypes.
	VisitNormalParameterTypes(ctx *NormalParameterTypesContext) interface{}

	// Visit a parse tree produced by Dart2Parser#optionalParameterTypes.
	VisitOptionalParameterTypes(ctx *OptionalParameterTypesContext) interface{}

	// Visit a parse tree produced by Dart2Parser#optionalPositionalParameterTypes.
	VisitOptionalPositionalParameterTypes(ctx *OptionalPositionalParameterTypesContext) interface{}

	// Visit a parse tree produced by Dart2Parser#namedParameterTypes.
	VisitNamedParameterTypes(ctx *NamedParameterTypesContext) interface{}

	// Visit a parse tree produced by Dart2Parser#namedParameterType.
	VisitNamedParameterType(ctx *NamedParameterTypeContext) interface{}

	// Visit a parse tree produced by Dart2Parser#typedIdentifier.
	VisitTypedIdentifier(ctx *TypedIdentifierContext) interface{}

	// Visit a parse tree produced by Dart2Parser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by Dart2Parser#normalFormalParameters.
	VisitNormalFormalParameters(ctx *NormalFormalParametersContext) interface{}

	// Visit a parse tree produced by Dart2Parser#optionalOrNamedFormalParameters.
	VisitOptionalOrNamedFormalParameters(ctx *OptionalOrNamedFormalParametersContext) interface{}

	// Visit a parse tree produced by Dart2Parser#optionalPositionalFormalParameters.
	VisitOptionalPositionalFormalParameters(ctx *OptionalPositionalFormalParametersContext) interface{}

	// Visit a parse tree produced by Dart2Parser#namedFormalParameters.
	VisitNamedFormalParameters(ctx *NamedFormalParametersContext) interface{}

	// Visit a parse tree produced by Dart2Parser#normalFormalParameter.
	VisitNormalFormalParameter(ctx *NormalFormalParameterContext) interface{}

	// Visit a parse tree produced by Dart2Parser#normalFormalParameterNoMetadata.
	VisitNormalFormalParameterNoMetadata(ctx *NormalFormalParameterNoMetadataContext) interface{}

	// Visit a parse tree produced by Dart2Parser#functionFormalParameter.
	VisitFunctionFormalParameter(ctx *FunctionFormalParameterContext) interface{}

	// Visit a parse tree produced by Dart2Parser#fieldFormalParameter.
	VisitFieldFormalParameter(ctx *FieldFormalParameterContext) interface{}

	// Visit a parse tree produced by Dart2Parser#formalParameterPart.
	VisitFormalParameterPart(ctx *FormalParameterPartContext) interface{}

	// Visit a parse tree produced by Dart2Parser#simpleFormalParameter.
	VisitSimpleFormalParameter(ctx *SimpleFormalParameterContext) interface{}

	// Visit a parse tree produced by Dart2Parser#declaredIdentifier.
	VisitDeclaredIdentifier(ctx *DeclaredIdentifierContext) interface{}

	// Visit a parse tree produced by Dart2Parser#finalConstVarOrType.
	VisitFinalConstVarOrType(ctx *FinalConstVarOrTypeContext) interface{}

	// Visit a parse tree produced by Dart2Parser#defaultFormalParameter.
	VisitDefaultFormalParameter(ctx *DefaultFormalParameterContext) interface{}

	// Visit a parse tree produced by Dart2Parser#defaultNamedParameter.
	VisitDefaultNamedParameter(ctx *DefaultNamedParameterContext) interface{}

	// Visit a parse tree produced by Dart2Parser#argumentPart.
	VisitArgumentPart(ctx *ArgumentPartContext) interface{}

	// Visit a parse tree produced by Dart2Parser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by Dart2Parser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by Dart2Parser#namedArgument.
	VisitNamedArgument(ctx *NamedArgumentContext) interface{}

	// Visit a parse tree produced by Dart2Parser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by Dart2Parser#nullLiteral.
	VisitNullLiteral(ctx *NullLiteralContext) interface{}

	// Visit a parse tree produced by Dart2Parser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by Dart2Parser#numericLiteral.
	VisitNumericLiteral(ctx *NumericLiteralContext) interface{}

	// Visit a parse tree produced by Dart2Parser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by Dart2Parser#symbolLiteral.
	VisitSymbolLiteral(ctx *SymbolLiteralContext) interface{}

	// Visit a parse tree produced by Dart2Parser#listLiteral.
	VisitListLiteral(ctx *ListLiteralContext) interface{}

	// Visit a parse tree produced by Dart2Parser#setOrMapLiteral.
	VisitSetOrMapLiteral(ctx *SetOrMapLiteralContext) interface{}

	// Visit a parse tree produced by Dart2Parser#elements.
	VisitElements(ctx *ElementsContext) interface{}

	// Visit a parse tree produced by Dart2Parser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#expressionElement.
	VisitExpressionElement(ctx *ExpressionElementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#mapElement.
	VisitMapElement(ctx *MapElementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#spreadElement.
	VisitSpreadElement(ctx *SpreadElementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#ifElement.
	VisitIfElement(ctx *IfElementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#forElement.
	VisitForElement(ctx *ForElementContext) interface{}

	// Visit a parse tree produced by Dart2Parser#operator.
	VisitOperator(ctx *OperatorContext) interface{}

	// Visit a parse tree produced by Dart2Parser#metadata.
	VisitMetadata(ctx *MetadataContext) interface{}

	// Visit a parse tree produced by Dart2Parser#metadatum.
	VisitMetadatum(ctx *MetadatumContext) interface{}

	// Visit a parse tree produced by Dart2Parser#dottedIdentifier.
	VisitDottedIdentifier(ctx *DottedIdentifierContext) interface{}

	// Visit a parse tree produced by Dart2Parser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}
}
