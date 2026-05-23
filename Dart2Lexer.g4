/* Generated from Dart specification
 * 
 * Dart2 Lexer Grammar for ANTLR v4
 * Adapted for compiler course
 * Simplified for educational purposes
 */

lexer grammar Dart2Lexer;

// --- Keywords ---
ABSTRACT_   : 'abstract';
AS_         : 'as';
ASSERT_     : 'assert';
ASYNC_      : 'async';
AWAIT_      : 'await';
BREAK_      : 'break';
CASE_       : 'case';
CATCH_      : 'catch';
CLASS_      : 'class';
CONST_      : 'const';
CONTINUE_   : 'continue';
COVARIANT_  : 'covariant';
DEFAULT_    : 'default';
DEFERRED_   : 'deferred';
DO_         : 'do';
DYNAMIC_    : 'dynamic';
ELSE_       : 'else';
ENUM_       : 'enum';
EXPORT_     : 'export';
EXTENDS_    : 'extends';
EXTENSION_  : 'extension';
EXTERNAL_   : 'external';
FACTORY_    : 'factory';
FALSE_      : 'false';
FINAL_      : 'final';
FINALLY_    : 'finally';
FOR_        : 'for';
FUNCTION_   : 'Function';
GET_        : 'get';
HIDE_       : 'hide';
IF_         : 'if';
IMPLEMENTS_ : 'implements';
IMPORT_     : 'import';
IN_         : 'in';
INTERFACE_  : 'interface';
IS_         : 'is';
LATE_       : 'late';
LIBRARY_    : 'library';
MIXIN_      : 'mixin';
NATIVE_     : 'native';
NEW_        : 'new';
NULL_       : 'null';
OF_         : 'of';
ON_         : 'on';
OPERATOR_   : 'operator';
PART_       : 'part';
REQUIRED_   : 'required';
RETHROW_    : 'rethrow';
RETURN_     : 'return';
SET_        : 'set';
SHOW_       : 'show';
STATIC_     : 'static';
SUPER_      : 'super';
SWITCH_     : 'switch';
SYNC_       : 'sync';
THIS_       : 'this';
THROW_      : 'throw';
TRUE_       : 'true';
TRY_        : 'try';
TYPEDEF_    : 'typedef';
VAR_        : 'var';
VOID_       : 'void';
WHILE_      : 'while';
WITH_       : 'with';
YIELD_      : 'yield';

// --- Operators ---
PLUS        : '+';
MINUS       : '-';
MUL         : '*';
DIV         : '/';
TILDE_DIV   : '~/';
MOD         : '%';
POWER       : '**';
ASSIGN      : '=';
PLUS_ASSIGN : '+=';
MINUS_ASSIGN: '-=';
MUL_ASSIGN  : '*=';
DIV_ASSIGN  : '/=';
MOD_ASSIGN  : '%=';
AND_ASSIGN  : '&=';
OR_ASSIGN   : '|=';
XOR_ASSIGN  : '^=';
TILDE_DIV_ASSIGN : '~/=';
SHL_ASSIGN  : '<<=';
SHR_ASSIGN  : '>>=';
QUES_QUES_ASSIGN : '??=';
EQ          : '==';
NE          : '!=';
LT          : '<';
GT          : '>';
LTE         : '<=';
GTE         : '>=';
AND         : '&&';
OR          : '||';
NOT         : '!';
AMPERSAND   : '&';
PIPE        : '|';
CARET       : '^';
TILDE       : '~';
SHL         : '<<';
SHR         : '>>';
PLUS_PLUS   : '++';
MINUS_MINUS : '--';
QUES        : '?';
QUES_DOT    : '?.';
QUES_QUES   : '??';
SEMICOLON   : ';';
COMMA       : ',';
DOT         : '.';
DOTDOT      : '..';
DOTDOTDOT   : '...';
DOTDOTDOT_QUES : '...?';
COLON       : ':';
DOUBLE_COLON: '::';
LPAREN      : '(';
RPAREN      : ')';
LBRACE      : '{';
RBRACE      : '}';
LBRACKET    : '[';
RBRACKET    : ']';
ARROW       : '=>';
AT          : '@';
HASH        : '#';

// --- Literals ---
NUMBER      : DIGIT+ ( '.' DIGIT+)? EXPONENT?
            | '.' DIGIT+ EXPONENT?;
HEX_NUMBER  : '0x' HEX_DIGIT+ | '0X' HEX_DIGIT+;

// String literals
SingleLineString
    : StringDQ
    | StringSQ
    | 'r\'' ~('\'' | '\n' | '\r')* '\''
    | 'r"' ~('"' | '\n' | '\r')* '"'
    ;

MultiLineString
    : '"""' StringContentTDQ*? '"""'
    | '\'\'\'' StringContentTSQ*? '\'\'\''
    | 'r"""' (~'"' | '"' ~'"' | '""' ~'"')* '"""'
    | 'r\'\'\'' (~'\'' | '\'' ~'\'' | '\'\'' ~'\'')*'\'\'\''
    ;

// --- Identifiers ---
IDENTIFIER  : IDENTIFIER_START IDENTIFIER_PART*;

// --- Comments ---
WHITESPACE  : ( '\t' | ' ' | NEWLINE)+ -> skip;
SINGLE_LINE_COMMENT : '//' ~[\r\n]* -> skip;
MULTI_LINE_COMMENT  : '/*' (MULTI_LINE_COMMENT | .)*? '*/' -> skip;

// --- Fragments ---
fragment EXPONENT : ( 'e' | 'E') ( '+' | '-')? DIGIT+;
fragment HEX_DIGIT : 'a' .. 'f' | 'A' .. 'F' | DIGIT;

fragment StringDQ
    : '"' StringContentDQ*? '"';

fragment StringContentDQ
    : ~('\\' | '"' | '\n' | '\r' | '$')
    | '\\' ~('\n' | '\r')
    | '${' StringContentDQ*? '}'
    | '$' IDENTIFIER
    ;

fragment StringSQ
    : '\'' StringContentSQ*? '\'';

fragment StringContentSQ
    : ~('\\' | '\'' | '\n' | '\r' | '$')
    | '\\' ~('\n' | '\r')
    | '${' StringContentSQ*? '}'
    | '$' IDENTIFIER
    ;

fragment StringContentTDQ : ~('\\' | '"') | '"' ~'"' | '""' ~'"';
fragment StringContentTSQ : '\'' ~'\'' | '\'\'' ~'\'' | .;

fragment ESCAPE_SEQUENCE
    : '\n'
    | '\r'
    | '\\f'
    | '\\b'
    | '\t'
    | '\\v'
    | '\\x' HEX_DIGIT HEX_DIGIT
    | '\\u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
    | '\\u{' HEX_DIGIT_SEQUENCE '}'
    ;

fragment HEX_DIGIT_SEQUENCE : HEX_DIGIT HEX_DIGIT? HEX_DIGIT? HEX_DIGIT? HEX_DIGIT? HEX_DIGIT?;

fragment NEWLINE : '\n' | '\r' | '\r\n';

fragment IDENTIFIER_START : LETTER | '_' | '$';
fragment IDENTIFIER_PART : IDENTIFIER_START | DIGIT;
fragment LETTER : 'a' .. 'z' | 'A' .. 'Z';
fragment DIGIT : '0' .. '9';
