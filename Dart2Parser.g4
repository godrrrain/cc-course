/* Generated from Dart specification
 * 
 * Dart2 Parser Grammar for ANTLR v4
 * Adapted for compiler course
 * Simplified for educational purposes
 */

parser grammar Dart2Parser;

options {
    tokenVocab = Dart2Lexer;
}

// ============================================================================
// Top-level rules
// ============================================================================

compilationUnit
    : (topLevelItem)* EOF
    ;

topLevelItem
    : importOrExportStatement
    | libraryStatement
    | partStatement
    | topLevelDeclaration
    | statement
    ;

importOrExportStatement
    : IMPORT_ stringLiteral (DEFERRED_? AS_ IDENTIFIER)? (SHOW_ identifierList | HIDE_ identifierList)* SEMICOLON
    | EXPORT_ stringLiteral (SHOW_ identifierList | HIDE_ identifierList)* SEMICOLON
    ;

libraryStatement
    : LIBRARY_ dottedIdentifier SEMICOLON
    ;

partStatement
    : PART_ stringLiteral SEMICOLON
    | PART_ OF_ dottedIdentifier SEMICOLON
    ;

topLevelDeclaration
    : classDeclaration
    | enumDeclaration
    | typeAliasDeclaration
    | mixinDeclaration
    | extensionDeclaration
    | functionDeclaration
    | getterDeclaration
    | setterDeclaration
    | variableDeclaration
    ;

// ============================================================================
// Class Declaration
// ============================================================================

classDeclaration
    : ABSTRACT_? CLASS_ typeIdentifier typeParameters? superclass? interfaces? LBRACE classMemberDeclaration* RBRACE
    | ABSTRACT_? CLASS_ typeIdentifier typeParameters? ASSIGN mixinApplication SEMICOLON
    ;

classBody
    : LBRACE classMemberDeclaration* RBRACE
    ;

classMemberDeclaration
    : declaration
    | methodSignature functionBody
    ;

superclass
    : EXTENDS_ typeNotVoid
    | EXTENDS_ typeNotVoid WITH_ mixinTypes
    | WITH_ mixinTypes
    ;

mixinTypes
    : typeNotVoid (COMMA typeNotVoid)*
    ;

mixinApplication
    : typeNotVoid WITH_ mixinTypes
    ;

interfaces
    : IMPLEMENTS_ typeList
    ;

// ============================================================================
// Enum Declaration
// ============================================================================

enumDeclaration
    : ENUM_ typeIdentifier LBRACE enumEntry (COMMA enumEntry)* COMMA? RBRACE
    | ENUM_ typeIdentifier LT typeParameterList GT LBRACE enumEntry (COMMA enumEntry)* COMMA? RBRACE
    ;

enumEntry
    : IDENTIFIER
    | IDENTIFIER LPAREN argumentList RPAREN
    ;

// ============================================================================
// Type Alias Declaration
// ============================================================================

typeAliasDeclaration
    : TYPEDEF_ typeIdentifier typeParameters? ASSIGN functionType SEMICOLON
    | TYPEDEF_ functionSignature SEMICOLON
    ;

functionSignature
    : returnType? IDENTIFIER typeParameters? formalParameterList
    ;

// ============================================================================
// Mixin Declaration
// ============================================================================

mixinDeclaration
    : MIXIN_ typeIdentifier typeParameters? (ON_ typeList)? interfaces? LBRACE classMemberDeclaration* RBRACE
    ;

// ============================================================================
// Extension Declaration
// ============================================================================

extensionDeclaration
    : EXTENSION_ IDENTIFIER? typeParameters? ON_ type LBRACE classMemberDeclaration* RBRACE
    ;

// ============================================================================
// Function and Member Declarations
// ============================================================================

functionDeclaration
    : returnType? IDENTIFIER typeParameters? formalParameterList functionBody
    | EXTERNAL_ returnType? IDENTIFIER typeParameters? formalParameterList SEMICOLON
    ;

getterDeclaration
    : returnType? GET_ IDENTIFIER functionBody
    | EXTERNAL_ returnType? GET_ IDENTIFIER SEMICOLON
    ;

setterDeclaration
    : VOID_? SET_ IDENTIFIER formalParameterList functionBody
    | EXTERNAL_ VOID_? SET_ IDENTIFIER formalParameterList SEMICOLON
    ;

methodSignature
    : STATIC_? EXTERNAL_? returnType? IDENTIFIER typeParameters? formalParameterList
    | STATIC_? OPERATOR_ operator typeParameters? formalParameterList
    | FACTORY_ IDENTIFIER typeParameters? formalParameterList
    | CONST_ FACTORY_ IDENTIFIER typeParameters? formalParameterList
    ;

declaration
    : CONST_ type? variableDeclarationList SEMICOLON
    | FINAL_ type? variableDeclarationList SEMICOLON
    | LATE_ FINAL_ type? variableDeclarationList SEMICOLON
    | STATIC_ CONST_ type? variableDeclarationList SEMICOLON
    | STATIC_ FINAL_ type? variableDeclarationList SEMICOLON
    | STATIC_ LATE_ FINAL_ type? variableDeclarationList SEMICOLON
    | VAR_ variableDeclarationList SEMICOLON
    | type variableDeclarationList SEMICOLON
    | LATE_ VAR_ variableDeclarationList SEMICOLON
    | LATE_ type variableDeclarationList SEMICOLON
    | COVARIANT_ VAR_ variableDeclarationList SEMICOLON
    | COVARIANT_ type variableDeclarationList SEMICOLON
    | STATIC_ VAR_ variableDeclarationList SEMICOLON
    | STATIC_ type variableDeclarationList SEMICOLON
    | STATIC_ LATE_ VAR_ variableDeclarationList SEMICOLON
    | STATIC_ LATE_ type variableDeclarationList SEMICOLON
    ;

variableDeclaration
    : (VAR_ | type | FINAL_ type?)? variableDeclarationList SEMICOLON
    ;

variableDeclarationList
    : variableDeclarationItem (COMMA variableDeclarationItem)*
    ;

variableDeclarationItem
    : IDENTIFIER (ASSIGN expression)?
    ;

functionBody
    : ASYNC_? ARROW expression SEMICOLON
    | ASYNC_? LBRACE statements RBRACE
    | ASYNC_ MUL LBRACE statements RBRACE
    | SYNC_ MUL LBRACE statements RBRACE
    | NATIVE_ stringLiteral? SEMICOLON
    ;

// ============================================================================
// Statements
// ============================================================================

statements
    : statement*
    ;

statement
    : label* nonLabelledStatement
    ;

nonLabelledStatement
    : block
    | expressionStatement
    | variableDeclaration
    | ifStatement
    | forStatement
    | whileStatement
    | doStatement
    | switchStatement
    | tryStatement
    | breakStatement
    | continueStatement
    | returnStatement
    | yieldStatement
    | yieldEachStatement
    | assertStatement
    | localFunctionDeclaration
    | rethrowStatement
    ;

block
    : LBRACE statements RBRACE
    ;

ifStatement
    : IF_ LPAREN expression RPAREN statement (ELSE_ statement)?
    ;

forStatement
    : AWAIT_? FOR_ LPAREN forLoopParts RPAREN statement
    ;

forLoopParts
    : variableDeclaration? expression? SEMICOLON expressionList?
    | metadata? (VAR_ | type)? IDENTIFIER IN_ expression
    | metadata? (VAR_ | type)? IDENTIFIER IN_ expression
    ;

whileStatement
    : WHILE_ LPAREN expression RPAREN statement
    ;

doStatement
    : DO_ statement WHILE_ LPAREN expression RPAREN SEMICOLON
    ;

switchStatement
    : SWITCH_ LPAREN expression RPAREN LBRACE (switchCase | defaultCase)* RBRACE
    ;

switchCase
    : label* CASE_ expression COLON statements
    ;

defaultCase
    : label* DEFAULT_ COLON statements
    ;

tryStatement
    : TRY_ block (onPart+ finallyPart? | finallyPart)
    ;

onPart
    : CATCH_ LPAREN IDENTIFIER (COMMA IDENTIFIER)? RPAREN block
    | ON_ typeNotVoid CATCH_? LPAREN IDENTIFIER (COMMA IDENTIFIER)? RPAREN? block
    ;

finallyPart
    : FINALLY_ block
    ;

breakStatement
    : BREAK_ IDENTIFIER? SEMICOLON
    ;

continueStatement
    : CONTINUE_ IDENTIFIER? SEMICOLON
    ;

returnStatement
    : RETURN_ expression? SEMICOLON
    ;

yieldStatement
    : YIELD_ expression SEMICOLON
    ;

yieldEachStatement
    : YIELD_ MUL expression SEMICOLON
    ;

expressionStatement
    : expression? SEMICOLON
    ;

assertStatement
    : ASSERT_ LPAREN expression (COMMA expression)? RPAREN SEMICOLON
    ;

label
    : IDENTIFIER COLON
    ;

localFunctionDeclaration
    : returnType? IDENTIFIER typeParameters? formalParameterList functionBody
    ;

rethrowStatement
    : RETHROW_ SEMICOLON
    ;

// ============================================================================
// Expressions
// ============================================================================

expression
    : assignableExpression assignmentOperator expression
    | conditionalExpression
    | cascade
    | throwExpression
    ;

expressionWithoutCascade
    : assignableExpression assignmentOperator expressionWithoutCascade
    | conditionalExpression
    | throwExpressionWithoutCascade
    ;

expressionList
    : expression (COMMA expression)*
    ;

assignableExpression
    : primary assignableSelectorPart
    | SUPER_ unconditionalAssignableSelector
    | IDENTIFIER
    ;

assignableSelectorPart
    : selector* assignableSelector
    ;

assignableSelector
    : unconditionalAssignableSelector
    | QUES_DOT IDENTIFIER
    | QUES LBRACKET expression RBRACKET
    ;

unconditionalAssignableSelector
    : LBRACKET expression RBRACKET
    | DOT IDENTIFIER
    ;

assignmentOperator
    : ASSIGN
    | MUL_ASSIGN
    | DIV_ASSIGN
    | TILDE_DIV_ASSIGN
    | MOD_ASSIGN
    | PLUS_ASSIGN
    | MINUS_ASSIGN
    | SHL_ASSIGN
    | SHR_ASSIGN
    | AND_ASSIGN
    | XOR_ASSIGN
    | OR_ASSIGN
    | QUES_QUES_ASSIGN
    ;

cascade
    : conditionalExpression (DOTDOT cascadeSection)+
    ;

cascadeSection
    : cascadeSelector cascadeSectionTail
    ;

cascadeSelector
    : LBRACKET expression RBRACKET
    | IDENTIFIER
    ;

cascadeSectionTail
    : cascadeAssignment
    | selector* (assignableSelector cascadeAssignment)?
    ;

cascadeAssignment
    : assignmentOperator expressionWithoutCascade
    ;

conditionalExpression
    : ifNullExpression (QUES expressionWithoutCascade COLON expressionWithoutCascade)?
    ;

ifNullExpression
    : logicalOrExpression (QUES_QUES logicalOrExpression)*
    ;

logicalOrExpression
    : logicalAndExpression (OR logicalAndExpression)*
    ;

logicalAndExpression
    : equalityExpression (AND equalityExpression)*
    ;

equalityExpression
    : relationalExpression ((EQ | NE) relationalExpression)*
    ;

relationalExpression
    : bitwiseOrExpression ((LT | GT | LTE | GTE) bitwiseOrExpression)*
    ;

bitwiseOrExpression
    : bitwiseXorExpression (PIPE bitwiseXorExpression)*
    ;

bitwiseXorExpression
    : bitwiseAndExpression (CARET bitwiseAndExpression)*
    ;

bitwiseAndExpression
    : shiftExpression (AMPERSAND shiftExpression)*
    ;

shiftExpression
    : additiveExpression ((SHL | SHR) additiveExpression)*
    ;

additiveExpression
    : multiplicativeExpression ((PLUS | MINUS) multiplicativeExpression)*
    ;

multiplicativeExpression
    : unaryExpression ((MUL | DIV | TILDE_DIV | MOD) unaryExpression)*
    ;

unaryExpression
    : (PLUS | MINUS | NOT | TILDE) unaryExpression
    | awaitExpression
    | postfixExpression
    | (MINUS | TILDE) SUPER_
    | (PLUS_PLUS | MINUS_MINUS) assignableExpression
    ;

awaitExpression
    : AWAIT_ unaryExpression
    ;

postfixExpression
    : assignableExpression (PLUS_PLUS | MINUS_MINUS)
    | primary selector*
    ;

selector
    : LBRACKET expression RBRACKET
    | DOT IDENTIFIER
    | QUES_DOT IDENTIFIER
    | argumentPart
    ;

primary
    : thisExpression
    | SUPER_ unconditionalAssignableSelector
    | SUPER_ argumentPart
    | IDENTIFIER
    | literal
    | newExpr
    | constExpression
    | functionExpression
    | LPAREN expression RPAREN
    ;

thisExpression
    : THIS_
    ;

newExpr
    : NEW_ constructorDesignation arguments
    ;

constExpression
    : CONST_ constructorDesignation arguments
    ;

functionExpression
    : formalParameterList functionExpressionBody
    ;

functionExpressionBody
    : ARROW expression
    | (ASYNC_? ARROW | ASYNC_? MUL | SYNC_ MUL) expression
    | ASYNC_? block
    | ASYNC_ MUL block
    | SYNC_ MUL block
    ;

constructorDesignation
    : typeIdentifier
    | IDENTIFIER DOT IDENTIFIER
    | typeName typeArguments? (DOT IDENTIFIER)?
    ;

throwExpression
    : THROW_ expression
    ;

throwExpressionWithoutCascade
    : THROW_ expressionWithoutCascade
    ;

// ============================================================================
// Types
// ============================================================================

type
    : functionType QUES?
    | typeNotFunction
    ;

typeNotVoid
    : functionType QUES?
    | typeNotVoidNotFunction
    ;

typeNotVoidNotFunction
    : typeName typeArguments? QUES?
    | FUNCTION_ QUES?
    ;

typeNotFunction
    : VOID_
    | typeNotVoidNotFunction
    ;

typeName
    : IDENTIFIER (DOT IDENTIFIER)?
    ;

typeIdentifier
    : IDENTIFIER
    | ASYNC_
    | HIDE_
    | OF_
    | ON_
    | SHOW_
    | SYNC_
    | AWAIT_
    | YIELD_
    | DYNAMIC_
    | NATIVE_
    ;

typeArguments
    : LT typeList GT
    ;

typeList
    : type (COMMA type)*
    ;

returnType
    : VOID_
    | type
    ;

typeParameters
    : LT typeParameterList GT
    ;

typeParameterList
    : typeParameter (COMMA typeParameter)*
    ;

typeParameter
    : metadata? IDENTIFIER (EXTENDS_ typeNotVoid)?
    ;

functionType
    : functionTypeTails
    | typeNotFunction functionTypeTails
    ;

functionTypeTails
    : functionTypeTail QUES? functionTypeTails
    | functionTypeTail
    ;

functionTypeTail
    : FUNCTION_ typeParameters? parameterTypeList
    ;

parameterTypeList
    : LPAREN RPAREN
    | LPAREN normalParameterTypes (COMMA optionalParameterTypes)? RPAREN
    | LPAREN optionalParameterTypes RPAREN
    ;

normalParameterTypes
    : typedIdentifier (COMMA typedIdentifier)*
    ;

optionalParameterTypes
    : optionalPositionalParameterTypes
    | namedParameterTypes
    ;

optionalPositionalParameterTypes
    : LBRACKET typedIdentifier (COMMA typedIdentifier)* RBRACKET
    ;

namedParameterTypes
    : LBRACE namedParameterType (COMMA namedParameterType)* RBRACE
    ;

namedParameterType
    : REQUIRED_? typedIdentifier
    ;

typedIdentifier
    : type IDENTIFIER
    ;

// ============================================================================
// Parameters and Arguments
// ============================================================================

formalParameterList
    : LPAREN RPAREN
    | LPAREN normalFormalParameters (COMMA optionalOrNamedFormalParameters)? RPAREN
    | LPAREN optionalOrNamedFormalParameters RPAREN
    ;

normalFormalParameters
    : normalFormalParameter (COMMA normalFormalParameter)*
    ;

optionalOrNamedFormalParameters
    : optionalPositionalFormalParameters
    | namedFormalParameters
    ;

optionalPositionalFormalParameters
    : LBRACKET defaultFormalParameter (COMMA defaultFormalParameter)* RBRACKET
    ;

namedFormalParameters
    : LBRACE defaultNamedParameter (COMMA defaultNamedParameter)* RBRACE
    ;

normalFormalParameter
    : metadata? normalFormalParameterNoMetadata
    ;

normalFormalParameterNoMetadata
    : functionFormalParameter
    | fieldFormalParameter
    | simpleFormalParameter
    ;

functionFormalParameter
    : COVARIANT_? type? IDENTIFIER typeParameters? formalParameterList QUES?
    ;

fieldFormalParameter
    : COVARIANT_? finalConstVarOrType? THIS_ DOT IDENTIFIER formalParameterPart QUES?
    ;

formalParameterPart
    : typeParameters? formalParameterList
    ;

simpleFormalParameter
    : declaredIdentifier
    | COVARIANT_? IDENTIFIER
    ;

declaredIdentifier
    : COVARIANT_? finalConstVarOrType IDENTIFIER
    ;

finalConstVarOrType
    : LATE_? FINAL_ type?
    | CONST_ type?
    | VAR_
    | type
    ;

defaultFormalParameter
    : normalFormalParameter (ASSIGN expression)?
    ;

defaultNamedParameter
    : metadata? REQUIRED_? normalFormalParameterNoMetadata (ASSIGN expression)?
    ;

argumentPart
    : typeArguments? arguments
    ;

arguments
    : LPAREN (argumentList COMMA?)? RPAREN
    ;

argumentList
    : namedArgument (COMMA namedArgument)*
    | expressionList (COMMA namedArgument)*
    ;

namedArgument
    : IDENTIFIER COLON expression
    ;

// ============================================================================
// Literals
// ============================================================================

literal
    : nullLiteral
    | booleanLiteral
    | numericLiteral
    | stringLiteral
    | symbolLiteral
    | listLiteral
    | setOrMapLiteral
    ;

nullLiteral
    : NULL_
    ;

booleanLiteral
    : TRUE_
    | FALSE_
    ;

numericLiteral
    : NUMBER
    | HEX_NUMBER
    ;

stringLiteral
    : (SingleLineString | MultiLineString)+
    ;

symbolLiteral
    : HASH (IDENTIFIER (DOT IDENTIFIER)* | operator | VOID_)
    ;

listLiteral
    : CONST_? typeArguments? LBRACKET elements? RBRACKET
    ;

setOrMapLiteral
    : CONST_? typeArguments? LBRACE elements? RBRACE
    ;

elements
    : element (COMMA element)* COMMA?
    ;

element
    : expressionElement
    | mapElement
    | spreadElement
    | ifElement
    | forElement
    ;

expressionElement
    : expression
    ;

mapElement
    : expression COLON expression
    ;

spreadElement
    : (DOTDOTDOT | DOTDOTDOT_QUES) expression
    ;

ifElement
    : IF_ LPAREN expression RPAREN element (ELSE_ element)?
    ;

forElement
    : AWAIT_? FOR_ LPAREN forLoopParts RPAREN element
    ;

// ============================================================================
// Miscellaneous
// ============================================================================

operator
    : TILDE
    | (PLUS | MINUS | MUL | DIV | TILDE_DIV | MOD | POWER)
    | (SHL | SHR)
    | (LT | GT | LTE | GTE)
    | EQ
    | NE
    | AMPERSAND
    | CARET
    | PIPE
    | LBRACKET RBRACKET
    | LBRACKET RBRACKET ASSIGN
    ;

metadata
    : (AT metadatum)*
    ;

metadatum
    : IDENTIFIER
    | IDENTIFIER DOT IDENTIFIER
    | constructorDesignation arguments
    ;

dottedIdentifier
    : IDENTIFIER (DOT IDENTIFIER)*
    ;

identifierList
    : IDENTIFIER (COMMA IDENTIFIER)*
    ;
