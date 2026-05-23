// Code generated from Dart2Parser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Dart2Parser
import "github.com/antlr4-go/antlr/v4"

// Dart2ParserListener is a complete listener for a parse tree produced by Dart2Parser.
type Dart2ParserListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterTopLevelItem is called when entering the topLevelItem production.
	EnterTopLevelItem(c *TopLevelItemContext)

	// EnterImportOrExportStatement is called when entering the importOrExportStatement production.
	EnterImportOrExportStatement(c *ImportOrExportStatementContext)

	// EnterLibraryStatement is called when entering the libraryStatement production.
	EnterLibraryStatement(c *LibraryStatementContext)

	// EnterPartStatement is called when entering the partStatement production.
	EnterPartStatement(c *PartStatementContext)

	// EnterTopLevelDeclaration is called when entering the topLevelDeclaration production.
	EnterTopLevelDeclaration(c *TopLevelDeclarationContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterClassBody is called when entering the classBody production.
	EnterClassBody(c *ClassBodyContext)

	// EnterClassMemberDeclaration is called when entering the classMemberDeclaration production.
	EnterClassMemberDeclaration(c *ClassMemberDeclarationContext)

	// EnterSuperclass is called when entering the superclass production.
	EnterSuperclass(c *SuperclassContext)

	// EnterMixinTypes is called when entering the mixinTypes production.
	EnterMixinTypes(c *MixinTypesContext)

	// EnterMixinApplication is called when entering the mixinApplication production.
	EnterMixinApplication(c *MixinApplicationContext)

	// EnterInterfaces is called when entering the interfaces production.
	EnterInterfaces(c *InterfacesContext)

	// EnterEnumDeclaration is called when entering the enumDeclaration production.
	EnterEnumDeclaration(c *EnumDeclarationContext)

	// EnterEnumEntry is called when entering the enumEntry production.
	EnterEnumEntry(c *EnumEntryContext)

	// EnterTypeAliasDeclaration is called when entering the typeAliasDeclaration production.
	EnterTypeAliasDeclaration(c *TypeAliasDeclarationContext)

	// EnterFunctionSignature is called when entering the functionSignature production.
	EnterFunctionSignature(c *FunctionSignatureContext)

	// EnterMixinDeclaration is called when entering the mixinDeclaration production.
	EnterMixinDeclaration(c *MixinDeclarationContext)

	// EnterExtensionDeclaration is called when entering the extensionDeclaration production.
	EnterExtensionDeclaration(c *ExtensionDeclarationContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterGetterDeclaration is called when entering the getterDeclaration production.
	EnterGetterDeclaration(c *GetterDeclarationContext)

	// EnterSetterDeclaration is called when entering the setterDeclaration production.
	EnterSetterDeclaration(c *SetterDeclarationContext)

	// EnterMethodSignature is called when entering the methodSignature production.
	EnterMethodSignature(c *MethodSignatureContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterVariableDeclarationList is called when entering the variableDeclarationList production.
	EnterVariableDeclarationList(c *VariableDeclarationListContext)

	// EnterVariableDeclarationItem is called when entering the variableDeclarationItem production.
	EnterVariableDeclarationItem(c *VariableDeclarationItemContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterNonLabelledStatement is called when entering the nonLabelledStatement production.
	EnterNonLabelledStatement(c *NonLabelledStatementContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterForLoopParts is called when entering the forLoopParts production.
	EnterForLoopParts(c *ForLoopPartsContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterDoStatement is called when entering the doStatement production.
	EnterDoStatement(c *DoStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterSwitchCase is called when entering the switchCase production.
	EnterSwitchCase(c *SwitchCaseContext)

	// EnterDefaultCase is called when entering the defaultCase production.
	EnterDefaultCase(c *DefaultCaseContext)

	// EnterTryStatement is called when entering the tryStatement production.
	EnterTryStatement(c *TryStatementContext)

	// EnterOnPart is called when entering the onPart production.
	EnterOnPart(c *OnPartContext)

	// EnterFinallyPart is called when entering the finallyPart production.
	EnterFinallyPart(c *FinallyPartContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterYieldStatement is called when entering the yieldStatement production.
	EnterYieldStatement(c *YieldStatementContext)

	// EnterYieldEachStatement is called when entering the yieldEachStatement production.
	EnterYieldEachStatement(c *YieldEachStatementContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterAssertStatement is called when entering the assertStatement production.
	EnterAssertStatement(c *AssertStatementContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterLocalFunctionDeclaration is called when entering the localFunctionDeclaration production.
	EnterLocalFunctionDeclaration(c *LocalFunctionDeclarationContext)

	// EnterRethrowStatement is called when entering the rethrowStatement production.
	EnterRethrowStatement(c *RethrowStatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpressionWithoutCascade is called when entering the expressionWithoutCascade production.
	EnterExpressionWithoutCascade(c *ExpressionWithoutCascadeContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterAssignableExpression is called when entering the assignableExpression production.
	EnterAssignableExpression(c *AssignableExpressionContext)

	// EnterAssignableSelectorPart is called when entering the assignableSelectorPart production.
	EnterAssignableSelectorPart(c *AssignableSelectorPartContext)

	// EnterAssignableSelector is called when entering the assignableSelector production.
	EnterAssignableSelector(c *AssignableSelectorContext)

	// EnterUnconditionalAssignableSelector is called when entering the unconditionalAssignableSelector production.
	EnterUnconditionalAssignableSelector(c *UnconditionalAssignableSelectorContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterCascade is called when entering the cascade production.
	EnterCascade(c *CascadeContext)

	// EnterCascadeSection is called when entering the cascadeSection production.
	EnterCascadeSection(c *CascadeSectionContext)

	// EnterCascadeSelector is called when entering the cascadeSelector production.
	EnterCascadeSelector(c *CascadeSelectorContext)

	// EnterCascadeSectionTail is called when entering the cascadeSectionTail production.
	EnterCascadeSectionTail(c *CascadeSectionTailContext)

	// EnterCascadeAssignment is called when entering the cascadeAssignment production.
	EnterCascadeAssignment(c *CascadeAssignmentContext)

	// EnterConditionalExpression is called when entering the conditionalExpression production.
	EnterConditionalExpression(c *ConditionalExpressionContext)

	// EnterIfNullExpression is called when entering the ifNullExpression production.
	EnterIfNullExpression(c *IfNullExpressionContext)

	// EnterLogicalOrExpression is called when entering the logicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterLogicalAndExpression is called when entering the logicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterBitwiseOrExpression is called when entering the bitwiseOrExpression production.
	EnterBitwiseOrExpression(c *BitwiseOrExpressionContext)

	// EnterBitwiseXorExpression is called when entering the bitwiseXorExpression production.
	EnterBitwiseXorExpression(c *BitwiseXorExpressionContext)

	// EnterBitwiseAndExpression is called when entering the bitwiseAndExpression production.
	EnterBitwiseAndExpression(c *BitwiseAndExpressionContext)

	// EnterShiftExpression is called when entering the shiftExpression production.
	EnterShiftExpression(c *ShiftExpressionContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterAwaitExpression is called when entering the awaitExpression production.
	EnterAwaitExpression(c *AwaitExpressionContext)

	// EnterPostfixExpression is called when entering the postfixExpression production.
	EnterPostfixExpression(c *PostfixExpressionContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterThisExpression is called when entering the thisExpression production.
	EnterThisExpression(c *ThisExpressionContext)

	// EnterNewExpr is called when entering the newExpr production.
	EnterNewExpr(c *NewExprContext)

	// EnterConstExpression is called when entering the constExpression production.
	EnterConstExpression(c *ConstExpressionContext)

	// EnterFunctionExpression is called when entering the functionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterFunctionExpressionBody is called when entering the functionExpressionBody production.
	EnterFunctionExpressionBody(c *FunctionExpressionBodyContext)

	// EnterConstructorDesignation is called when entering the constructorDesignation production.
	EnterConstructorDesignation(c *ConstructorDesignationContext)

	// EnterThrowExpression is called when entering the throwExpression production.
	EnterThrowExpression(c *ThrowExpressionContext)

	// EnterThrowExpressionWithoutCascade is called when entering the throwExpressionWithoutCascade production.
	EnterThrowExpressionWithoutCascade(c *ThrowExpressionWithoutCascadeContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterTypeNotVoid is called when entering the typeNotVoid production.
	EnterTypeNotVoid(c *TypeNotVoidContext)

	// EnterTypeNotVoidNotFunction is called when entering the typeNotVoidNotFunction production.
	EnterTypeNotVoidNotFunction(c *TypeNotVoidNotFunctionContext)

	// EnterTypeNotFunction is called when entering the typeNotFunction production.
	EnterTypeNotFunction(c *TypeNotFunctionContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterTypeIdentifier is called when entering the typeIdentifier production.
	EnterTypeIdentifier(c *TypeIdentifierContext)

	// EnterTypeArguments is called when entering the typeArguments production.
	EnterTypeArguments(c *TypeArgumentsContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterReturnType is called when entering the returnType production.
	EnterReturnType(c *ReturnTypeContext)

	// EnterTypeParameters is called when entering the typeParameters production.
	EnterTypeParameters(c *TypeParametersContext)

	// EnterTypeParameterList is called when entering the typeParameterList production.
	EnterTypeParameterList(c *TypeParameterListContext)

	// EnterTypeParameter is called when entering the typeParameter production.
	EnterTypeParameter(c *TypeParameterContext)

	// EnterFunctionType is called when entering the functionType production.
	EnterFunctionType(c *FunctionTypeContext)

	// EnterFunctionTypeTails is called when entering the functionTypeTails production.
	EnterFunctionTypeTails(c *FunctionTypeTailsContext)

	// EnterFunctionTypeTail is called when entering the functionTypeTail production.
	EnterFunctionTypeTail(c *FunctionTypeTailContext)

	// EnterParameterTypeList is called when entering the parameterTypeList production.
	EnterParameterTypeList(c *ParameterTypeListContext)

	// EnterNormalParameterTypes is called when entering the normalParameterTypes production.
	EnterNormalParameterTypes(c *NormalParameterTypesContext)

	// EnterOptionalParameterTypes is called when entering the optionalParameterTypes production.
	EnterOptionalParameterTypes(c *OptionalParameterTypesContext)

	// EnterOptionalPositionalParameterTypes is called when entering the optionalPositionalParameterTypes production.
	EnterOptionalPositionalParameterTypes(c *OptionalPositionalParameterTypesContext)

	// EnterNamedParameterTypes is called when entering the namedParameterTypes production.
	EnterNamedParameterTypes(c *NamedParameterTypesContext)

	// EnterNamedParameterType is called when entering the namedParameterType production.
	EnterNamedParameterType(c *NamedParameterTypeContext)

	// EnterTypedIdentifier is called when entering the typedIdentifier production.
	EnterTypedIdentifier(c *TypedIdentifierContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterNormalFormalParameters is called when entering the normalFormalParameters production.
	EnterNormalFormalParameters(c *NormalFormalParametersContext)

	// EnterOptionalOrNamedFormalParameters is called when entering the optionalOrNamedFormalParameters production.
	EnterOptionalOrNamedFormalParameters(c *OptionalOrNamedFormalParametersContext)

	// EnterOptionalPositionalFormalParameters is called when entering the optionalPositionalFormalParameters production.
	EnterOptionalPositionalFormalParameters(c *OptionalPositionalFormalParametersContext)

	// EnterNamedFormalParameters is called when entering the namedFormalParameters production.
	EnterNamedFormalParameters(c *NamedFormalParametersContext)

	// EnterNormalFormalParameter is called when entering the normalFormalParameter production.
	EnterNormalFormalParameter(c *NormalFormalParameterContext)

	// EnterNormalFormalParameterNoMetadata is called when entering the normalFormalParameterNoMetadata production.
	EnterNormalFormalParameterNoMetadata(c *NormalFormalParameterNoMetadataContext)

	// EnterFunctionFormalParameter is called when entering the functionFormalParameter production.
	EnterFunctionFormalParameter(c *FunctionFormalParameterContext)

	// EnterFieldFormalParameter is called when entering the fieldFormalParameter production.
	EnterFieldFormalParameter(c *FieldFormalParameterContext)

	// EnterFormalParameterPart is called when entering the formalParameterPart production.
	EnterFormalParameterPart(c *FormalParameterPartContext)

	// EnterSimpleFormalParameter is called when entering the simpleFormalParameter production.
	EnterSimpleFormalParameter(c *SimpleFormalParameterContext)

	// EnterDeclaredIdentifier is called when entering the declaredIdentifier production.
	EnterDeclaredIdentifier(c *DeclaredIdentifierContext)

	// EnterFinalConstVarOrType is called when entering the finalConstVarOrType production.
	EnterFinalConstVarOrType(c *FinalConstVarOrTypeContext)

	// EnterDefaultFormalParameter is called when entering the defaultFormalParameter production.
	EnterDefaultFormalParameter(c *DefaultFormalParameterContext)

	// EnterDefaultNamedParameter is called when entering the defaultNamedParameter production.
	EnterDefaultNamedParameter(c *DefaultNamedParameterContext)

	// EnterArgumentPart is called when entering the argumentPart production.
	EnterArgumentPart(c *ArgumentPartContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterNamedArgument is called when entering the namedArgument production.
	EnterNamedArgument(c *NamedArgumentContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterNullLiteral is called when entering the nullLiteral production.
	EnterNullLiteral(c *NullLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterNumericLiteral is called when entering the numericLiteral production.
	EnterNumericLiteral(c *NumericLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterSymbolLiteral is called when entering the symbolLiteral production.
	EnterSymbolLiteral(c *SymbolLiteralContext)

	// EnterListLiteral is called when entering the listLiteral production.
	EnterListLiteral(c *ListLiteralContext)

	// EnterSetOrMapLiteral is called when entering the setOrMapLiteral production.
	EnterSetOrMapLiteral(c *SetOrMapLiteralContext)

	// EnterElements is called when entering the elements production.
	EnterElements(c *ElementsContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterExpressionElement is called when entering the expressionElement production.
	EnterExpressionElement(c *ExpressionElementContext)

	// EnterMapElement is called when entering the mapElement production.
	EnterMapElement(c *MapElementContext)

	// EnterSpreadElement is called when entering the spreadElement production.
	EnterSpreadElement(c *SpreadElementContext)

	// EnterIfElement is called when entering the ifElement production.
	EnterIfElement(c *IfElementContext)

	// EnterForElement is called when entering the forElement production.
	EnterForElement(c *ForElementContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterMetadata is called when entering the metadata production.
	EnterMetadata(c *MetadataContext)

	// EnterMetadatum is called when entering the metadatum production.
	EnterMetadatum(c *MetadatumContext)

	// EnterDottedIdentifier is called when entering the dottedIdentifier production.
	EnterDottedIdentifier(c *DottedIdentifierContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitTopLevelItem is called when exiting the topLevelItem production.
	ExitTopLevelItem(c *TopLevelItemContext)

	// ExitImportOrExportStatement is called when exiting the importOrExportStatement production.
	ExitImportOrExportStatement(c *ImportOrExportStatementContext)

	// ExitLibraryStatement is called when exiting the libraryStatement production.
	ExitLibraryStatement(c *LibraryStatementContext)

	// ExitPartStatement is called when exiting the partStatement production.
	ExitPartStatement(c *PartStatementContext)

	// ExitTopLevelDeclaration is called when exiting the topLevelDeclaration production.
	ExitTopLevelDeclaration(c *TopLevelDeclarationContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitClassBody is called when exiting the classBody production.
	ExitClassBody(c *ClassBodyContext)

	// ExitClassMemberDeclaration is called when exiting the classMemberDeclaration production.
	ExitClassMemberDeclaration(c *ClassMemberDeclarationContext)

	// ExitSuperclass is called when exiting the superclass production.
	ExitSuperclass(c *SuperclassContext)

	// ExitMixinTypes is called when exiting the mixinTypes production.
	ExitMixinTypes(c *MixinTypesContext)

	// ExitMixinApplication is called when exiting the mixinApplication production.
	ExitMixinApplication(c *MixinApplicationContext)

	// ExitInterfaces is called when exiting the interfaces production.
	ExitInterfaces(c *InterfacesContext)

	// ExitEnumDeclaration is called when exiting the enumDeclaration production.
	ExitEnumDeclaration(c *EnumDeclarationContext)

	// ExitEnumEntry is called when exiting the enumEntry production.
	ExitEnumEntry(c *EnumEntryContext)

	// ExitTypeAliasDeclaration is called when exiting the typeAliasDeclaration production.
	ExitTypeAliasDeclaration(c *TypeAliasDeclarationContext)

	// ExitFunctionSignature is called when exiting the functionSignature production.
	ExitFunctionSignature(c *FunctionSignatureContext)

	// ExitMixinDeclaration is called when exiting the mixinDeclaration production.
	ExitMixinDeclaration(c *MixinDeclarationContext)

	// ExitExtensionDeclaration is called when exiting the extensionDeclaration production.
	ExitExtensionDeclaration(c *ExtensionDeclarationContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitGetterDeclaration is called when exiting the getterDeclaration production.
	ExitGetterDeclaration(c *GetterDeclarationContext)

	// ExitSetterDeclaration is called when exiting the setterDeclaration production.
	ExitSetterDeclaration(c *SetterDeclarationContext)

	// ExitMethodSignature is called when exiting the methodSignature production.
	ExitMethodSignature(c *MethodSignatureContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitVariableDeclarationList is called when exiting the variableDeclarationList production.
	ExitVariableDeclarationList(c *VariableDeclarationListContext)

	// ExitVariableDeclarationItem is called when exiting the variableDeclarationItem production.
	ExitVariableDeclarationItem(c *VariableDeclarationItemContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitNonLabelledStatement is called when exiting the nonLabelledStatement production.
	ExitNonLabelledStatement(c *NonLabelledStatementContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitForLoopParts is called when exiting the forLoopParts production.
	ExitForLoopParts(c *ForLoopPartsContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitDoStatement is called when exiting the doStatement production.
	ExitDoStatement(c *DoStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitSwitchCase is called when exiting the switchCase production.
	ExitSwitchCase(c *SwitchCaseContext)

	// ExitDefaultCase is called when exiting the defaultCase production.
	ExitDefaultCase(c *DefaultCaseContext)

	// ExitTryStatement is called when exiting the tryStatement production.
	ExitTryStatement(c *TryStatementContext)

	// ExitOnPart is called when exiting the onPart production.
	ExitOnPart(c *OnPartContext)

	// ExitFinallyPart is called when exiting the finallyPart production.
	ExitFinallyPart(c *FinallyPartContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitYieldStatement is called when exiting the yieldStatement production.
	ExitYieldStatement(c *YieldStatementContext)

	// ExitYieldEachStatement is called when exiting the yieldEachStatement production.
	ExitYieldEachStatement(c *YieldEachStatementContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitAssertStatement is called when exiting the assertStatement production.
	ExitAssertStatement(c *AssertStatementContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitLocalFunctionDeclaration is called when exiting the localFunctionDeclaration production.
	ExitLocalFunctionDeclaration(c *LocalFunctionDeclarationContext)

	// ExitRethrowStatement is called when exiting the rethrowStatement production.
	ExitRethrowStatement(c *RethrowStatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpressionWithoutCascade is called when exiting the expressionWithoutCascade production.
	ExitExpressionWithoutCascade(c *ExpressionWithoutCascadeContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitAssignableExpression is called when exiting the assignableExpression production.
	ExitAssignableExpression(c *AssignableExpressionContext)

	// ExitAssignableSelectorPart is called when exiting the assignableSelectorPart production.
	ExitAssignableSelectorPart(c *AssignableSelectorPartContext)

	// ExitAssignableSelector is called when exiting the assignableSelector production.
	ExitAssignableSelector(c *AssignableSelectorContext)

	// ExitUnconditionalAssignableSelector is called when exiting the unconditionalAssignableSelector production.
	ExitUnconditionalAssignableSelector(c *UnconditionalAssignableSelectorContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitCascade is called when exiting the cascade production.
	ExitCascade(c *CascadeContext)

	// ExitCascadeSection is called when exiting the cascadeSection production.
	ExitCascadeSection(c *CascadeSectionContext)

	// ExitCascadeSelector is called when exiting the cascadeSelector production.
	ExitCascadeSelector(c *CascadeSelectorContext)

	// ExitCascadeSectionTail is called when exiting the cascadeSectionTail production.
	ExitCascadeSectionTail(c *CascadeSectionTailContext)

	// ExitCascadeAssignment is called when exiting the cascadeAssignment production.
	ExitCascadeAssignment(c *CascadeAssignmentContext)

	// ExitConditionalExpression is called when exiting the conditionalExpression production.
	ExitConditionalExpression(c *ConditionalExpressionContext)

	// ExitIfNullExpression is called when exiting the ifNullExpression production.
	ExitIfNullExpression(c *IfNullExpressionContext)

	// ExitLogicalOrExpression is called when exiting the logicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitLogicalAndExpression is called when exiting the logicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitBitwiseOrExpression is called when exiting the bitwiseOrExpression production.
	ExitBitwiseOrExpression(c *BitwiseOrExpressionContext)

	// ExitBitwiseXorExpression is called when exiting the bitwiseXorExpression production.
	ExitBitwiseXorExpression(c *BitwiseXorExpressionContext)

	// ExitBitwiseAndExpression is called when exiting the bitwiseAndExpression production.
	ExitBitwiseAndExpression(c *BitwiseAndExpressionContext)

	// ExitShiftExpression is called when exiting the shiftExpression production.
	ExitShiftExpression(c *ShiftExpressionContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitAwaitExpression is called when exiting the awaitExpression production.
	ExitAwaitExpression(c *AwaitExpressionContext)

	// ExitPostfixExpression is called when exiting the postfixExpression production.
	ExitPostfixExpression(c *PostfixExpressionContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitThisExpression is called when exiting the thisExpression production.
	ExitThisExpression(c *ThisExpressionContext)

	// ExitNewExpr is called when exiting the newExpr production.
	ExitNewExpr(c *NewExprContext)

	// ExitConstExpression is called when exiting the constExpression production.
	ExitConstExpression(c *ConstExpressionContext)

	// ExitFunctionExpression is called when exiting the functionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitFunctionExpressionBody is called when exiting the functionExpressionBody production.
	ExitFunctionExpressionBody(c *FunctionExpressionBodyContext)

	// ExitConstructorDesignation is called when exiting the constructorDesignation production.
	ExitConstructorDesignation(c *ConstructorDesignationContext)

	// ExitThrowExpression is called when exiting the throwExpression production.
	ExitThrowExpression(c *ThrowExpressionContext)

	// ExitThrowExpressionWithoutCascade is called when exiting the throwExpressionWithoutCascade production.
	ExitThrowExpressionWithoutCascade(c *ThrowExpressionWithoutCascadeContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitTypeNotVoid is called when exiting the typeNotVoid production.
	ExitTypeNotVoid(c *TypeNotVoidContext)

	// ExitTypeNotVoidNotFunction is called when exiting the typeNotVoidNotFunction production.
	ExitTypeNotVoidNotFunction(c *TypeNotVoidNotFunctionContext)

	// ExitTypeNotFunction is called when exiting the typeNotFunction production.
	ExitTypeNotFunction(c *TypeNotFunctionContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitTypeIdentifier is called when exiting the typeIdentifier production.
	ExitTypeIdentifier(c *TypeIdentifierContext)

	// ExitTypeArguments is called when exiting the typeArguments production.
	ExitTypeArguments(c *TypeArgumentsContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitReturnType is called when exiting the returnType production.
	ExitReturnType(c *ReturnTypeContext)

	// ExitTypeParameters is called when exiting the typeParameters production.
	ExitTypeParameters(c *TypeParametersContext)

	// ExitTypeParameterList is called when exiting the typeParameterList production.
	ExitTypeParameterList(c *TypeParameterListContext)

	// ExitTypeParameter is called when exiting the typeParameter production.
	ExitTypeParameter(c *TypeParameterContext)

	// ExitFunctionType is called when exiting the functionType production.
	ExitFunctionType(c *FunctionTypeContext)

	// ExitFunctionTypeTails is called when exiting the functionTypeTails production.
	ExitFunctionTypeTails(c *FunctionTypeTailsContext)

	// ExitFunctionTypeTail is called when exiting the functionTypeTail production.
	ExitFunctionTypeTail(c *FunctionTypeTailContext)

	// ExitParameterTypeList is called when exiting the parameterTypeList production.
	ExitParameterTypeList(c *ParameterTypeListContext)

	// ExitNormalParameterTypes is called when exiting the normalParameterTypes production.
	ExitNormalParameterTypes(c *NormalParameterTypesContext)

	// ExitOptionalParameterTypes is called when exiting the optionalParameterTypes production.
	ExitOptionalParameterTypes(c *OptionalParameterTypesContext)

	// ExitOptionalPositionalParameterTypes is called when exiting the optionalPositionalParameterTypes production.
	ExitOptionalPositionalParameterTypes(c *OptionalPositionalParameterTypesContext)

	// ExitNamedParameterTypes is called when exiting the namedParameterTypes production.
	ExitNamedParameterTypes(c *NamedParameterTypesContext)

	// ExitNamedParameterType is called when exiting the namedParameterType production.
	ExitNamedParameterType(c *NamedParameterTypeContext)

	// ExitTypedIdentifier is called when exiting the typedIdentifier production.
	ExitTypedIdentifier(c *TypedIdentifierContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitNormalFormalParameters is called when exiting the normalFormalParameters production.
	ExitNormalFormalParameters(c *NormalFormalParametersContext)

	// ExitOptionalOrNamedFormalParameters is called when exiting the optionalOrNamedFormalParameters production.
	ExitOptionalOrNamedFormalParameters(c *OptionalOrNamedFormalParametersContext)

	// ExitOptionalPositionalFormalParameters is called when exiting the optionalPositionalFormalParameters production.
	ExitOptionalPositionalFormalParameters(c *OptionalPositionalFormalParametersContext)

	// ExitNamedFormalParameters is called when exiting the namedFormalParameters production.
	ExitNamedFormalParameters(c *NamedFormalParametersContext)

	// ExitNormalFormalParameter is called when exiting the normalFormalParameter production.
	ExitNormalFormalParameter(c *NormalFormalParameterContext)

	// ExitNormalFormalParameterNoMetadata is called when exiting the normalFormalParameterNoMetadata production.
	ExitNormalFormalParameterNoMetadata(c *NormalFormalParameterNoMetadataContext)

	// ExitFunctionFormalParameter is called when exiting the functionFormalParameter production.
	ExitFunctionFormalParameter(c *FunctionFormalParameterContext)

	// ExitFieldFormalParameter is called when exiting the fieldFormalParameter production.
	ExitFieldFormalParameter(c *FieldFormalParameterContext)

	// ExitFormalParameterPart is called when exiting the formalParameterPart production.
	ExitFormalParameterPart(c *FormalParameterPartContext)

	// ExitSimpleFormalParameter is called when exiting the simpleFormalParameter production.
	ExitSimpleFormalParameter(c *SimpleFormalParameterContext)

	// ExitDeclaredIdentifier is called when exiting the declaredIdentifier production.
	ExitDeclaredIdentifier(c *DeclaredIdentifierContext)

	// ExitFinalConstVarOrType is called when exiting the finalConstVarOrType production.
	ExitFinalConstVarOrType(c *FinalConstVarOrTypeContext)

	// ExitDefaultFormalParameter is called when exiting the defaultFormalParameter production.
	ExitDefaultFormalParameter(c *DefaultFormalParameterContext)

	// ExitDefaultNamedParameter is called when exiting the defaultNamedParameter production.
	ExitDefaultNamedParameter(c *DefaultNamedParameterContext)

	// ExitArgumentPart is called when exiting the argumentPart production.
	ExitArgumentPart(c *ArgumentPartContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitNamedArgument is called when exiting the namedArgument production.
	ExitNamedArgument(c *NamedArgumentContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitNullLiteral is called when exiting the nullLiteral production.
	ExitNullLiteral(c *NullLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitNumericLiteral is called when exiting the numericLiteral production.
	ExitNumericLiteral(c *NumericLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitSymbolLiteral is called when exiting the symbolLiteral production.
	ExitSymbolLiteral(c *SymbolLiteralContext)

	// ExitListLiteral is called when exiting the listLiteral production.
	ExitListLiteral(c *ListLiteralContext)

	// ExitSetOrMapLiteral is called when exiting the setOrMapLiteral production.
	ExitSetOrMapLiteral(c *SetOrMapLiteralContext)

	// ExitElements is called when exiting the elements production.
	ExitElements(c *ElementsContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitExpressionElement is called when exiting the expressionElement production.
	ExitExpressionElement(c *ExpressionElementContext)

	// ExitMapElement is called when exiting the mapElement production.
	ExitMapElement(c *MapElementContext)

	// ExitSpreadElement is called when exiting the spreadElement production.
	ExitSpreadElement(c *SpreadElementContext)

	// ExitIfElement is called when exiting the ifElement production.
	ExitIfElement(c *IfElementContext)

	// ExitForElement is called when exiting the forElement production.
	ExitForElement(c *ForElementContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitMetadata is called when exiting the metadata production.
	ExitMetadata(c *MetadataContext)

	// ExitMetadatum is called when exiting the metadatum production.
	ExitMetadatum(c *MetadatumContext)

	// ExitDottedIdentifier is called when exiting the dottedIdentifier production.
	ExitDottedIdentifier(c *DottedIdentifierContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)
}
