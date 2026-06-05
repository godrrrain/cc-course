package ast

import "fmt"

// ----- Node interfaces -----

type Node interface {
	nodeStr() string
}

type Decl interface {
	Node
	declNode()
}

type Stmt interface {
	Node
	stmtNode()
}

type Expr interface {
	Node
	exprNode()
}

// ----- Types -----

type Type struct {
	Name string
}

var (
	TypeInt     = &Type{Name: "int"}
	TypeDouble  = &Type{Name: "double"}
	TypeBool    = &Type{Name: "bool"}
	TypeString  = &Type{Name: "String"}
	TypeVoid    = &Type{Name: "void"}
	TypeDynamic = &Type{Name: "dynamic"}
)

// ----- Program -----

type Program struct {
	Declarations []Decl
}

func (p *Program) nodeStr() string { return "Program" }

// ----- Declarations -----

type FuncDecl struct {
	Name       string
	ReturnType *Type
	Params     []*Param
	Body       *BlockStmt
	BodyExpr   Expr // for arrow functions
}

func (f *FuncDecl) nodeStr() string { return "FuncDecl " + f.Name }
func (f *FuncDecl) declNode()       {}

type Param struct {
	Name string
	Type *Type
}

type VarDecl struct {
	Name string
	Type *Type
	Init Expr
}

func (v *VarDecl) nodeStr() string {
	typ := "dynamic"
	if v.Type != nil {
		typ = v.Type.Name
	}
	return fmt.Sprintf("VarDecl %s %s", v.Name, typ)
}
func (v *VarDecl) declNode()       {}
func (v *VarDecl) stmtNode()       {}

// ----- Statements -----

type BlockStmt struct {
	Statements []Stmt
}

func (b *BlockStmt) nodeStr() string { return "Block" }
func (b *BlockStmt) stmtNode()       {}

type IfStmt struct {
	Condition Expr
	Then      Stmt
	Else      Stmt
}

func (i *IfStmt) nodeStr() string { return "If" }
func (i *IfStmt) stmtNode()       {}

type WhileStmt struct {
	Condition Expr
	Body      Stmt
}

func (w *WhileStmt) nodeStr() string { return "While" }
func (w *WhileStmt) stmtNode()       {}

type DoWhileStmt struct {
	Body      Stmt
	Condition Expr
}

func (d *DoWhileStmt) nodeStr() string { return "DoWhile" }
func (d *DoWhileStmt) stmtNode()       {}

type ForStmt struct {
	Init      Stmt
	Condition Expr
	Update    Expr
	Body      Stmt
}

func (f *ForStmt) nodeStr() string { return "For" }
func (f *ForStmt) stmtNode()       {}

type ForInStmt struct {
	VarName  string
	VarType  *Type
	Iterable Expr
	Body     Stmt
}

func (f *ForInStmt) nodeStr() string {
	typ := "dynamic"
	if f.VarType != nil {
		typ = f.VarType.Name
	}
	return fmt.Sprintf("ForIn %s %s", f.VarName, typ)
}
func (f *ForInStmt) stmtNode()       {}

type ExprStmt struct {
	Expression Expr
}

func (e *ExprStmt) nodeStr() string { return "ExprStmt" }
func (e *ExprStmt) stmtNode()       {}

type ReturnStmt struct {
	Value Expr
}

func (r *ReturnStmt) nodeStr() string { return "Return" }
func (r *ReturnStmt) stmtNode()       {}

type BreakStmt struct{}

func (b *BreakStmt) nodeStr() string { return "Break" }
func (b *BreakStmt) stmtNode()       {}

type ContinueStmt struct{}

func (c *ContinueStmt) nodeStr() string { return "Continue" }
func (c *ContinueStmt) stmtNode()       {}

type AssertStmt struct {
	Condition Expr
}

func (a *AssertStmt) nodeStr() string { return "Assert" }
func (a *AssertStmt) stmtNode()       {}

// ----- Expressions -----

type BinaryExpr struct {
	Op          string
	Left, Right Expr
}

func (b *BinaryExpr) nodeStr() string { return "Binary " + b.Op }
func (b *BinaryExpr) exprNode()       {}

type UnaryExpr struct {
	Op      string // "-", "!", "~"
	Operand Expr
}

func (u *UnaryExpr) nodeStr() string { return "Unary " + u.Op }
func (u *UnaryExpr) exprNode()       {}

type PostfixExpr struct {
	Op     string // "++", "--"
	Target Expr
}

func (p *PostfixExpr) nodeStr() string { return "Postfix " + p.Op }
func (p *PostfixExpr) exprNode()       {}

type AssignExpr struct {
	Target Expr
	Value  Expr
}

func (a *AssignExpr) nodeStr() string { return "Assign" }
func (a *AssignExpr) exprNode()       {}

type CallExpr struct {
	Function Expr
	Args     []Expr
}

func (c *CallExpr) nodeStr() string { return "Call" }
func (c *CallExpr) exprNode()       {}

type Identifier struct {
	Name string
}

func (i *Identifier) nodeStr() string { return "Ident " + i.Name }
func (i *Identifier) exprNode()       {}

type IntLiteral struct {
	Value int64
}

func (i *IntLiteral) nodeStr() string { return fmt.Sprintf("Int %d", i.Value) }
func (i *IntLiteral) exprNode()       {}

type FloatLiteral struct {
	Value float64
}

func (f *FloatLiteral) nodeStr() string { return fmt.Sprintf("Float %v", f.Value) }
func (f *FloatLiteral) exprNode()       {}

type BoolLiteral struct {
	Value bool
}

func (b *BoolLiteral) nodeStr() string { return fmt.Sprintf("Bool %v", b.Value) }
func (b *BoolLiteral) exprNode()       {}

type StringLiteral struct {
	Value string
}

func (s *StringLiteral) nodeStr() string { return fmt.Sprintf("String %q", s.Value) }
func (s *StringLiteral) exprNode()       {}

type NullLiteral struct{}

func (n *NullLiteral) nodeStr() string { return "Null" }
func (n *NullLiteral) exprNode()       {}

type ListLiteral struct {
	Elements []Expr
}

func (l *ListLiteral) nodeStr() string { return "List" }
func (l *ListLiteral) exprNode()       {}

type IndexExpr struct {
	Target Expr
	Index  Expr
}

func (i *IndexExpr) nodeStr() string { return "Index" }
func (i *IndexExpr) exprNode()       {}

type TernaryExpr struct {
	Condition, Then, Else Expr
}

func (t *TernaryExpr) nodeStr() string { return "Ternary" }
func (t *TernaryExpr) exprNode()       {}

type InterpolatedStringExpr struct {
	Parts []InterpolationPart
}

func (i *InterpolatedStringExpr) nodeStr() string { return "InterpolatedString" }
func (i *InterpolatedStringExpr) exprNode()       {}

type InterpolationPart struct {
	IsExpr bool
	Text   string // for raw text
	Expr   Expr   // for interpolated expression
}
