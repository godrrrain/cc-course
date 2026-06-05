package visitor

import (
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"cc-course/internal/ast"
	"cc-course/internal/parser"
)

type AstBuilder struct {
	*parser.BaseDart2ParserVisitor
	Errors []error
}

func NewAstBuilder() *AstBuilder {
	return &AstBuilder{
		BaseDart2ParserVisitor: &parser.BaseDart2ParserVisitor{},
		Errors:                 make([]error, 0),
	}
}

func (b *AstBuilder) Build(tree antlr.ParseTree) *ast.Program {
	if cu, ok := tree.(*parser.CompilationUnitContext); ok {
		return b.visitCompilationUnit(cu)
	}
	return &ast.Program{}
}

func (b *AstBuilder) visitCompilationUnit(ctx *parser.CompilationUnitContext) *ast.Program {
	prog := &ast.Program{}
	for _, item := range ctx.AllTopLevelItem() {
		if d := b.visitTopLevelItem(item); d != nil {
			prog.Declarations = append(prog.Declarations, d...)
		}
	}
	return prog
}

func (b *AstBuilder) visitTopLevelItem(ctx parser.ITopLevelItemContext) []ast.Decl {
	if tc, ok := ctx.(*parser.TopLevelItemContext); ok && tc.TopLevelDeclaration() != nil {
		return b.visitTopLevelDeclaration(tc.TopLevelDeclaration())
	}
	return nil
}

func (b *AstBuilder) visitTopLevelDeclaration(ctx parser.ITopLevelDeclarationContext) []ast.Decl {
	if dc, ok := ctx.(*parser.TopLevelDeclarationContext); ok {
		if dc.FunctionDeclaration() != nil {
			if fd := b.visitFunctionDeclaration(dc.FunctionDeclaration()); fd != nil {
				return []ast.Decl{fd}
			}
		}
		if dc.VariableDeclaration() != nil {
			return b.visitVariableDeclaration(dc.VariableDeclaration())
		}
	}
	return nil
}

func (b *AstBuilder) visitFunctionDeclaration(ctx parser.IFunctionDeclarationContext) *ast.FuncDecl {
	fc, ok := ctx.(*parser.FunctionDeclarationContext)
	if !ok || fc.IDENTIFIER() == nil {
		return nil
	}
	fd := &ast.FuncDecl{
		Name: fc.IDENTIFIER().GetText(),
	}
	switch {
	case fc.ReturnType() != nil:
		fd.ReturnType = b.visitReturnType(fc.ReturnType())
	default:
		fd.ReturnType = ast.TypeVoid
	}

	if fpl := fc.FormalParameterList(); fpl != nil {
		fd.Params = b.visitFormalParameterList(fpl)
	}

	if fc.FunctionBody() != nil {
		b.visitFunctionBody(fc.FunctionBody(), fd)
	}

	return fd
}

func (b *AstBuilder) visitFunctionBody(ctx parser.IFunctionBodyContext, fd *ast.FuncDecl) {
	fbc, ok := ctx.(*parser.FunctionBodyContext)
	if !ok {
		return
	}
	if fbc.Statements() != nil {
		fd.Body = &ast.BlockStmt{}
		fd.Body.Statements = b.visitStatements(fbc.Statements())
	} else if fbc.ARROW() != nil && fbc.Expression() != nil {
		fd.BodyExpr = b.visitExpression(fbc.Expression())
	}
}

func (b *AstBuilder) visitReturnType(ctx parser.IReturnTypeContext) *ast.Type {
	if rc, ok := ctx.(*parser.ReturnTypeContext); ok && rc.Type_() != nil {
		return b.visitType(rc.Type_())
	}
	return ast.TypeVoid
}

func (b *AstBuilder) visitFormalParameterList(ctx parser.IFormalParameterListContext) []*ast.Param {
	fpc, ok := ctx.(*parser.FormalParameterListContext)
	if !ok {
		return nil
	}
	var params []*ast.Param
	if nfps := fpc.NormalFormalParameters(); nfps != nil {
		for _, p := range nfps.AllNormalFormalParameter() {
			params = append(params, b.visitNormalFormalParameter(p)...)
		}
	}
	return params
}

func (b *AstBuilder) visitNormalFormalParameter(ctx parser.INormalFormalParameterContext) []*ast.Param {
	if nc, ok := ctx.(*parser.NormalFormalParameterContext); ok && nc.NormalFormalParameterNoMetadata() != nil {
		return b.visitNormalFormalParameterNoMetadata(nc.NormalFormalParameterNoMetadata())
	}
	return nil
}

func (b *AstBuilder) visitNormalFormalParameterNoMetadata(ctx parser.INormalFormalParameterNoMetadataContext) []*ast.Param {
	if nc, ok := ctx.(*parser.NormalFormalParameterNoMetadataContext); ok && nc.SimpleFormalParameter() != nil {
		return b.visitSimpleFormalParameter(nc.SimpleFormalParameter())
	}
	return nil
}

func (b *AstBuilder) visitSimpleFormalParameter(ctx parser.ISimpleFormalParameterContext) []*ast.Param {
	if sc, ok := ctx.(*parser.SimpleFormalParameterContext); ok {
		if sc.DeclaredIdentifier() != nil {
			return b.visitDeclaredIdentifier(sc.DeclaredIdentifier())
		}
		if sc.IDENTIFIER() != nil {
			return []*ast.Param{{Name: sc.IDENTIFIER().GetText(), Type: ast.TypeDynamic}}
		}
	}
	return nil
}

func (b *AstBuilder) visitDeclaredIdentifier(ctx parser.IDeclaredIdentifierContext) []*ast.Param {
	dc, ok := ctx.(*parser.DeclaredIdentifierContext)
	if !ok {
		return nil
	}
	typ := ast.TypeDynamic
	if dc.FinalConstVarOrType() != nil {
		if t := b.visitFinalConstVarOrType(dc.FinalConstVarOrType()); t != nil {
			typ = t
		}
	}
	if dc.IDENTIFIER() != nil {
		return []*ast.Param{{Name: dc.IDENTIFIER().GetText(), Type: typ}}
	}
	return nil
}

func (b *AstBuilder) visitFinalConstVarOrType(ctx parser.IFinalConstVarOrTypeContext) *ast.Type {
	if fc, ok := ctx.(*parser.FinalConstVarOrTypeContext); ok && fc.Type_() != nil {
		return b.visitType(fc.Type_())
	}
	return ast.TypeDynamic
}

func (b *AstBuilder) visitType(ctx parser.ITypeContext) *ast.Type {
	if tc, ok := ctx.(*parser.TypeContext); ok && tc.TypeNotFunction() != nil {
		return b.visitTypeNotFunction(tc.TypeNotFunction())
	}
	return ast.TypeDynamic
}

func (b *AstBuilder) visitTypeNotFunction(ctx parser.ITypeNotFunctionContext) *ast.Type {
	if tc, ok := ctx.(*parser.TypeNotFunctionContext); ok {
		if tc.VOID_() != nil {
			return ast.TypeVoid
		}
		if tc.TypeNotVoidNotFunction() != nil {
			return b.visitTypeNotVoidNotFunction(tc.TypeNotVoidNotFunction())
		}
	}
	return ast.TypeDynamic
}

func (b *AstBuilder) visitTypeNotVoidNotFunction(ctx parser.ITypeNotVoidNotFunctionContext) *ast.Type {
	if tc, ok := ctx.(*parser.TypeNotVoidNotFunctionContext); ok && tc.TypeName() != nil {
		name := tc.TypeName().GetText()
		switch name {
		case "int":
			return ast.TypeInt
		case "double":
			return ast.TypeDouble
		case "bool":
			return ast.TypeBool
		case "String":
			return ast.TypeString
		case "num":
			return ast.TypeDouble
		case "dynamic", "var":
			return ast.TypeDynamic
		}
	}
	return ast.TypeDynamic
}

// ----- Variable declaration -----

func (b *AstBuilder) visitVariableDeclaration(ctx parser.IVariableDeclarationContext) []ast.Decl {
	vc, ok := ctx.(*parser.VariableDeclarationContext)
	if !ok || vc.VariableDeclarationList() == nil {
		return nil
	}
	declaredType := ast.TypeDynamic
	if vc.VAR_() == nil {
		if vc.Type_() != nil {
			declaredType = b.visitType(vc.Type_())
		} else if vc.FINAL_() != nil && vc.Type_() != nil {
			declaredType = b.visitType(vc.Type_())
		}
	}
	return b.visitVariableDeclarationList(vc.VariableDeclarationList(), declaredType)
}

func (b *AstBuilder) visitVariableDeclarationList(ctx parser.IVariableDeclarationListContext, declaredType *ast.Type) []ast.Decl {
	vc, ok := ctx.(*parser.VariableDeclarationListContext)
	if !ok {
		return nil
	}
	var decls []ast.Decl
	for _, item := range vc.AllVariableDeclarationItem() {
		if d := b.visitVariableDeclarationItem(item, declaredType); d != nil {
			decls = append(decls, d)
		}
	}
	return decls
}

func (b *AstBuilder) visitVariableDeclarationItem(ctx parser.IVariableDeclarationItemContext, declaredType *ast.Type) *ast.VarDecl {
	vc, ok := ctx.(*parser.VariableDeclarationItemContext)
	if !ok || vc.IDENTIFIER() == nil {
		return nil
	}
	vd := &ast.VarDecl{
		Name: vc.IDENTIFIER().GetText(),
		Type: declaredType,
	}
	if vc.Expression() != nil {
		vd.Init = b.visitExpression(vc.Expression())
	}
	return vd
}

// ----- Statements -----

func (b *AstBuilder) visitStatements(ctx parser.IStatementsContext) []ast.Stmt {
	sc, ok := ctx.(*parser.StatementsContext)
	if !ok {
		return nil
	}
	var stmts []ast.Stmt
	for _, stmt := range sc.AllStatement() {
		if s := b.visitStatement(stmt); s != nil {
			stmts = append(stmts, s)
		}
	}
	return stmts
}

func (b *AstBuilder) visitStatement(ctx parser.IStatementContext) ast.Stmt {
	if sc, ok := ctx.(*parser.StatementContext); ok && sc.NonLabelledStatement() != nil {
		return b.visitNonLabelledStatement(sc.NonLabelledStatement())
	}
	return nil
}

func (b *AstBuilder) visitNonLabelledStatement(ctx parser.INonLabelledStatementContext) ast.Stmt {
	nc, ok := ctx.(*parser.NonLabelledStatementContext)
	if !ok {
		return nil
	}
	switch {
	case nc.Block() != nil:
		return b.visitBlock(nc.Block())
	case nc.VariableDeclaration() != nil:
		decls := b.visitVariableDeclaration(nc.VariableDeclaration())
		if len(decls) > 0 {
			if vd, ok := decls[0].(*ast.VarDecl); ok {
				return vd
			}
		}
		return nil
	case nc.IfStatement() != nil:
		return b.visitIfStatement(nc.IfStatement())
	case nc.ForStatement() != nil:
		return b.visitForStatement(nc.ForStatement())
	case nc.WhileStatement() != nil:
		return b.visitWhileStatement(nc.WhileStatement())
	case nc.DoStatement() != nil:
		return b.visitDoStatement(nc.DoStatement())
	case nc.BreakStatement() != nil:
		return &ast.BreakStmt{}
	case nc.ContinueStatement() != nil:
		return &ast.ContinueStmt{}
	case nc.ReturnStatement() != nil:
		return b.visitReturnStatement(nc.ReturnStatement())
	case nc.ExpressionStatement() != nil:
		return b.visitExpressionStatement(nc.ExpressionStatement())
	case nc.AssertStatement() != nil:
		return b.visitAssertStatement(nc.AssertStatement())
	}
	return nil
}

func (b *AstBuilder) visitBlock(ctx parser.IBlockContext) *ast.BlockStmt {
	bc, ok := ctx.(*parser.BlockContext)
	if !ok {
		return &ast.BlockStmt{}
	}
	blk := &ast.BlockStmt{}
	if bc.Statements() != nil {
		blk.Statements = b.visitStatements(bc.Statements())
	}
	return blk
}

func (b *AstBuilder) visitIfStatement(ctx parser.IIfStatementContext) *ast.IfStmt {
	ic, ok := ctx.(*parser.IfStatementContext)
	if !ok {
		return nil
	}
	is := &ast.IfStmt{}
	if ic.Expression() != nil {
		is.Condition = b.visitExpression(ic.Expression())
	}
	if len(ic.AllStatement()) > 0 {
		is.Then = b.visitStatement(ic.Statement(0))
	}
	if ic.ELSE_() != nil && len(ic.AllStatement()) > 1 {
		is.Else = b.visitStatement(ic.Statement(1))
	}
	return is
}

func (b *AstBuilder) visitWhileStatement(ctx parser.IWhileStatementContext) *ast.WhileStmt {
	wc, ok := ctx.(*parser.WhileStatementContext)
	if !ok {
		return nil
	}
	ws := &ast.WhileStmt{}
	if wc.Expression() != nil {
		ws.Condition = b.visitExpression(wc.Expression())
	}
	if wc.Statement() != nil {
		ws.Body = b.visitStatement(wc.Statement())
	}
	return ws
}

func (b *AstBuilder) visitDoStatement(ctx parser.IDoStatementContext) *ast.DoWhileStmt {
	dc, ok := ctx.(*parser.DoStatementContext)
	if !ok {
		return nil
	}
	ds := &ast.DoWhileStmt{}
	if dc.Statement() != nil {
		ds.Body = b.visitStatement(dc.Statement())
	}
	if dc.Expression() != nil {
		ds.Condition = b.visitExpression(dc.Expression())
	}
	return ds
}

func (b *AstBuilder) visitForStatement(ctx parser.IForStatementContext) ast.Stmt {
	fc, ok := ctx.(*parser.ForStatementContext)
	if !ok || fc.ForLoopParts() == nil {
		return nil
	}
	parts := fc.ForLoopParts()

	if parts.IN_() != nil {
		return b.visitForInLoop(parts, fc.Statement())
	}

	fs := &ast.ForStmt{}
	if parts.VariableDeclaration() != nil {
		decls := b.visitVariableDeclaration(parts.VariableDeclaration())
		if len(decls) > 0 {
			if vd, ok := decls[0].(*ast.VarDecl); ok {
				fs.Init = vd
			}
		}
	}
	if parts.Expression() != nil {
		fs.Condition = b.visitExpression(parts.Expression())
	}
	if parts.ExpressionList() != nil {
		fs.Update = b.visitExpressionListAsExpr(parts.ExpressionList())
	}
	if fc.Statement() != nil {
		fs.Body = b.visitStatement(fc.Statement())
	}
	return fs
}

func (b *AstBuilder) visitForInLoop(parts parser.IForLoopPartsContext, bodyCtx parser.IStatementContext) *ast.ForInStmt {
	fis := &ast.ForInStmt{
		VarName: parts.IDENTIFIER().GetText(),
	}
	if parts.Type_() != nil {
		fis.VarType = b.visitType(parts.Type_())
	} else {
		fis.VarType = ast.TypeDynamic
	}
	if parts.Expression() != nil {
		fis.Iterable = b.visitExpression(parts.Expression())
	}
	if bodyCtx != nil {
		fis.Body = b.visitStatement(bodyCtx)
	}
	return fis
}

func (b *AstBuilder) visitReturnStatement(ctx parser.IReturnStatementContext) *ast.ReturnStmt {
	rc, ok := ctx.(*parser.ReturnStatementContext)
	if !ok {
		return nil
	}
	rs := &ast.ReturnStmt{}
	if rc.Expression() != nil {
		rs.Value = b.visitExpression(rc.Expression())
	}
	return rs
}

func (b *AstBuilder) visitAssertStatement(ctx parser.IAssertStatementContext) *ast.AssertStmt {
	if ac, ok := ctx.(*parser.AssertStatementContext); ok && len(ac.AllExpression()) > 0 {
		return &ast.AssertStmt{Condition: b.visitExpression(ac.Expression(0))}
	}
	return nil
}

func (b *AstBuilder) visitExpressionStatement(ctx parser.IExpressionStatementContext) *ast.ExprStmt {
	if ec, ok := ctx.(*parser.ExpressionStatementContext); ok && ec.Expression() != nil {
		return &ast.ExprStmt{Expression: b.visitExpression(ec.Expression())}
	}
	return nil
}

// ----- Expressions -----

func (b *AstBuilder) visitExpression(ctx parser.IExpressionContext) ast.Expr {
	ec, ok := ctx.(*parser.ExpressionContext)
	if !ok {
		return nil
	}
	switch {
	case ec.AssignableExpression() != nil && ec.AssignmentOperator() != nil:
		target := b.visitAssignableExpression(ec.AssignableExpression())
		value := b.visitExpression(ec.Expression())
		return &ast.AssignExpr{Target: target, Value: value}
	case ec.ConditionalExpression() != nil:
		return b.visitConditionalExpression(ec.ConditionalExpression())
	}
	return nil
}

func (b *AstBuilder) visitConditionalExpression(ctx parser.IConditionalExpressionContext) ast.Expr {
	cc, ok := ctx.(*parser.ConditionalExpressionContext)
	if !ok || cc.IfNullExpression() == nil {
		return nil
	}
	expr := b.visitIfNullExpression(cc.IfNullExpression())
	if cc.QUES() != nil && len(cc.AllExpressionWithoutCascade()) == 2 {
		return &ast.TernaryExpr{
			Condition: expr,
			Then:      b.visitExpressionWithoutCascade(cc.ExpressionWithoutCascade(0)),
			Else:      b.visitExpressionWithoutCascade(cc.ExpressionWithoutCascade(1)),
		}
	}
	return expr
}

func (b *AstBuilder) visitIfNullExpression(ctx parser.IIfNullExpressionContext) ast.Expr {
	if ic, ok := ctx.(*parser.IfNullExpressionContext); ok && len(ic.AllLogicalOrExpression()) > 0 {
		return b.visitLogicalOrExpression(ic.LogicalOrExpression(0))
	}
	return nil
}

func (b *AstBuilder) visitLogicalOrExpression(ctx parser.ILogicalOrExpressionContext) ast.Expr {
	lc, ok := ctx.(*parser.LogicalOrExpressionContext)
	if !ok || len(lc.AllLogicalAndExpression()) == 0 {
		return nil
	}
	expr := b.visitLogicalAndExpression(lc.LogicalAndExpression(0))
	for i := 1; i < len(lc.AllLogicalAndExpression()); i++ {
		rhs := b.visitLogicalAndExpression(lc.LogicalAndExpression(i))
		expr = &ast.BinaryExpr{Op: "||", Left: expr, Right: rhs}
	}
	return expr
}

func (b *AstBuilder) visitLogicalAndExpression(ctx parser.ILogicalAndExpressionContext) ast.Expr {
	lc, ok := ctx.(*parser.LogicalAndExpressionContext)
	if !ok || len(lc.AllEqualityExpression()) == 0 {
		return nil
	}
	expr := b.visitEqualityExpression(lc.EqualityExpression(0))
	for i := 1; i < len(lc.AllEqualityExpression()); i++ {
		rhs := b.visitEqualityExpression(lc.EqualityExpression(i))
		expr = &ast.BinaryExpr{Op: "&&", Left: expr, Right: rhs}
	}
	return expr
}

func (b *AstBuilder) visitEqualityExpression(ctx parser.IEqualityExpressionContext) ast.Expr {
	ec, ok := ctx.(*parser.EqualityExpressionContext)
	if !ok || len(ec.AllRelationalExpression()) == 0 {
		return nil
	}
	expr := b.visitRelationalExpression(ec.RelationalExpression(0))
	for i := 1; i < ec.GetChildCount(); i += 2 {
		op := ec.GetChild(i).(antlr.TerminalNode).GetText()
		rhs := b.visitRelationalExpression(ec.RelationalExpression((i + 1) / 2))
		expr = &ast.BinaryExpr{Op: op, Left: expr, Right: rhs}
	}
	return expr
}

func (b *AstBuilder) visitRelationalExpression(ctx parser.IRelationalExpressionContext) ast.Expr {
	rc, ok := ctx.(*parser.RelationalExpressionContext)
	if !ok || len(rc.AllBitwiseOrExpression()) == 0 {
		return nil
	}
	expr := b.visitBitwiseOrExpression(rc.BitwiseOrExpression(0))
	for i := 1; i < rc.GetChildCount(); i += 2 {
		op := rc.GetChild(i).(antlr.TerminalNode).GetText()
		rhs := b.visitBitwiseOrExpression(rc.BitwiseOrExpression((i + 1) / 2))
		expr = &ast.BinaryExpr{Op: op, Left: expr, Right: rhs}
	}
	return expr
}

func (b *AstBuilder) visitBitwiseOrExpression(ctx parser.IBitwiseOrExpressionContext) ast.Expr {
	bc, ok := ctx.(*parser.BitwiseOrExpressionContext)
	if !ok || len(bc.AllBitwiseXorExpression()) == 0 {
		return nil
	}
	expr := b.visitBitwiseXorExpression(bc.BitwiseXorExpression(0))
	for i := 1; i < len(bc.AllBitwiseXorExpression()); i++ {
		rhs := b.visitBitwiseXorExpression(bc.BitwiseXorExpression(i))
		expr = &ast.BinaryExpr{Op: "|", Left: expr, Right: rhs}
	}
	return expr
}

func (b *AstBuilder) visitBitwiseXorExpression(ctx parser.IBitwiseXorExpressionContext) ast.Expr {
	bc, ok := ctx.(*parser.BitwiseXorExpressionContext)
	if !ok || len(bc.AllBitwiseAndExpression()) == 0 {
		return nil
	}
	expr := b.visitBitwiseAndExpression(bc.BitwiseAndExpression(0))
	for i := 1; i < len(bc.AllBitwiseAndExpression()); i++ {
		rhs := b.visitBitwiseAndExpression(bc.BitwiseAndExpression(i))
		expr = &ast.BinaryExpr{Op: "^", Left: expr, Right: rhs}
	}
	return expr
}

func (b *AstBuilder) visitBitwiseAndExpression(ctx parser.IBitwiseAndExpressionContext) ast.Expr {
	bc, ok := ctx.(*parser.BitwiseAndExpressionContext)
	if !ok || len(bc.AllShiftExpression()) == 0 {
		return nil
	}
	expr := b.visitShiftExpression(bc.ShiftExpression(0))
	for i := 1; i < len(bc.AllShiftExpression()); i++ {
		rhs := b.visitShiftExpression(bc.ShiftExpression(i))
		expr = &ast.BinaryExpr{Op: "&", Left: expr, Right: rhs}
	}
	return expr
}

func (b *AstBuilder) visitShiftExpression(ctx parser.IShiftExpressionContext) ast.Expr {
	sc, ok := ctx.(*parser.ShiftExpressionContext)
	if !ok || len(sc.AllAdditiveExpression()) == 0 {
		return nil
	}
	expr := b.visitAdditiveExpression(sc.AdditiveExpression(0))
	for i := 1; i < sc.GetChildCount(); i += 2 {
		op := sc.GetChild(i).(antlr.TerminalNode).GetText()
		rhs := b.visitAdditiveExpression(sc.AdditiveExpression((i + 1) / 2))
		expr = &ast.BinaryExpr{Op: op, Left: expr, Right: rhs}
	}
	return expr
}

func (b *AstBuilder) visitAdditiveExpression(ctx parser.IAdditiveExpressionContext) ast.Expr {
	ac, ok := ctx.(*parser.AdditiveExpressionContext)
	if !ok || len(ac.AllMultiplicativeExpression()) == 0 {
		return nil
	}
	expr := b.visitMultiplicativeExpression(ac.MultiplicativeExpression(0))
	for i := 1; i < ac.GetChildCount(); i += 2 {
		op := ac.GetChild(i).(antlr.TerminalNode).GetText()
		rhs := b.visitMultiplicativeExpression(ac.MultiplicativeExpression((i + 1) / 2))
		expr = &ast.BinaryExpr{Op: op, Left: expr, Right: rhs}
	}
	return expr
}

func (b *AstBuilder) visitMultiplicativeExpression(ctx parser.IMultiplicativeExpressionContext) ast.Expr {
	mc, ok := ctx.(*parser.MultiplicativeExpressionContext)
	if !ok || len(mc.AllUnaryExpression()) == 0 {
		return nil
	}
	expr := b.visitUnaryExpression(mc.UnaryExpression(0))
	for i := 1; i < mc.GetChildCount(); i += 2 {
		op := mc.GetChild(i).(antlr.TerminalNode).GetText()
		rhs := b.visitUnaryExpression(mc.UnaryExpression((i + 1) / 2))
		expr = &ast.BinaryExpr{Op: op, Left: expr, Right: rhs}
	}
	return expr
}

func (b *AstBuilder) visitUnaryExpression(ctx parser.IUnaryExpressionContext) ast.Expr {
	uc, ok := ctx.(*parser.UnaryExpressionContext)
	if !ok {
		return nil
	}
	switch {
	case uc.PLUS() != nil:
		return b.visitUnaryExpression(uc.UnaryExpression())
	case uc.MINUS() != nil && uc.UnaryExpression() != nil:
		return &ast.UnaryExpr{Op: "-", Operand: b.visitUnaryExpression(uc.UnaryExpression())}
	case uc.NOT() != nil:
		return &ast.UnaryExpr{Op: "!", Operand: b.visitUnaryExpression(uc.UnaryExpression())}
	case uc.TILDE() != nil:
		return &ast.UnaryExpr{Op: "~", Operand: b.visitUnaryExpression(uc.UnaryExpression())}
	case uc.PostfixExpression() != nil:
		return b.visitPostfixExpression(uc.PostfixExpression())
	}
	return nil
}

func (b *AstBuilder) visitPostfixExpression(ctx parser.IPostfixExpressionContext) ast.Expr {
	pc, ok := ctx.(*parser.PostfixExpressionContext)
	if !ok {
		return nil
	}
	if pc.AssignableExpression() != nil && (pc.PLUS_PLUS() != nil || pc.MINUS_MINUS() != nil) {
		op := "++"
		if pc.MINUS_MINUS() != nil {
			op = "--"
		}
		target := b.visitAssignableExpression(pc.AssignableExpression())
		return &ast.PostfixExpr{Op: op, Target: target}
	}
	if pc.Primary() != nil {
		expr := b.visitPrimary(pc.Primary())
		for _, sel := range pc.AllSelector() {
			if ap := sel.ArgumentPart(); ap != nil {
				expr = &ast.CallExpr{
					Function: expr,
					Args:     b.visitArgumentPart(ap),
				}
			} else if sel.LBRACKET() != nil && sel.Expression() != nil {
				idx := b.visitExpression(sel.Expression())
				expr = &ast.IndexExpr{Target: expr, Index: idx}
			}
		}
		return expr
	}
	return nil
}

func (b *AstBuilder) visitPrimary(ctx parser.IPrimaryContext) ast.Expr {
	pc, ok := ctx.(*parser.PrimaryContext)
	if !ok {
		return nil
	}
	switch {
	case pc.IDENTIFIER() != nil:
		return &ast.Identifier{Name: pc.IDENTIFIER().GetText()}
	case pc.Literal() != nil:
		return b.visitLiteral(pc.Literal())
	case pc.Expression() != nil:
		return b.visitExpression(pc.Expression())
	}
	return nil
}

func (b *AstBuilder) visitLiteral(ctx parser.ILiteralContext) ast.Expr {
	lc, ok := ctx.(*parser.LiteralContext)
	if !ok {
		return nil
	}
	switch {
	case lc.NullLiteral() != nil:
		return &ast.NullLiteral{}
	case lc.BooleanLiteral() != nil:
		return b.visitBooleanLiteral(lc.BooleanLiteral())
	case lc.NumericLiteral() != nil:
		return b.visitNumericLiteral(lc.NumericLiteral())
	case lc.StringLiteral() != nil:
		return b.visitStringLiteral(lc.StringLiteral())
	case lc.ListLiteral() != nil:
		return b.visitListLiteral(lc.ListLiteral())
	}
	return nil
}

func (b *AstBuilder) visitBooleanLiteral(ctx parser.IBooleanLiteralContext) *ast.BoolLiteral {
	if bc, ok := ctx.(*parser.BooleanLiteralContext); ok {
		return &ast.BoolLiteral{Value: bc.TRUE_() != nil}
	}
	return &ast.BoolLiteral{Value: false}
}

func (b *AstBuilder) visitNumericLiteral(ctx parser.INumericLiteralContext) ast.Expr {
	nc, ok := ctx.(*parser.NumericLiteralContext)
	if !ok {
		return &ast.IntLiteral{Value: 0}
	}
	if nc.NUMBER() != nil {
		text := nc.NUMBER().GetText()
		if val, err := strconv.ParseFloat(text, 64); err == nil {
			if val == float64(int64(val)) {
				return &ast.IntLiteral{Value: int64(val)}
			}
			return &ast.FloatLiteral{Value: val}
		}
	}
	if nc.HEX_NUMBER() != nil {
		text := nc.HEX_NUMBER().GetText()
		if val, err := strconv.ParseInt(text, 0, 64); err == nil {
			return &ast.IntLiteral{Value: val}
		}
	}
	return &ast.IntLiteral{Value: 0}
}

func (b *AstBuilder) visitStringLiteral(ctx parser.IStringLiteralContext) ast.Expr {
	sc, ok := ctx.(*parser.StringLiteralContext)
	if !ok {
		return &ast.StringLiteral{}
	}
	text := sc.GetText()
	unquoted := stripQuotes(text)
	if strings.Contains(unquoted, "$") {
		return b.buildInterpolatedString(unquoted)
	}
	return &ast.StringLiteral{Value: unquoted}
}

func (b *AstBuilder) buildInterpolatedString(text string) *ast.InterpolatedStringExpr {
	is := &ast.InterpolatedStringExpr{}
	for i := 0; i < len(text); i++ {
		if text[i] == '$' && i+1 < len(text) && isIdentStart(text[i+1]) {
			j := i + 1
			for j < len(text) && isIdentPart(text[j]) {
				j++
			}
			name := text[i+1 : j]
			is.Parts = append(is.Parts, ast.InterpolationPart{
				IsExpr: true,
				Expr:   &ast.Identifier{Name: name},
			})
			i = j - 1
		} else {
			j := i
			for j < len(text) && text[j] != '$' {
				j++
			}
			if j > i {
				is.Parts = append(is.Parts, ast.InterpolationPart{
					Text: text[i:j],
				})
			}
			i = j - 1
		}
	}
	return is
}

func (b *AstBuilder) visitListLiteral(ctx parser.IListLiteralContext) *ast.ListLiteral {
	lc, ok := ctx.(*parser.ListLiteralContext)
	if !ok {
		return &ast.ListLiteral{}
	}
	ll := &ast.ListLiteral{}
	if lc.Elements() != nil {
		for _, el := range lc.Elements().AllElement() {
			if ec, ok := el.(*parser.ElementContext); ok && ec.ExpressionElement() != nil {
				if eec, ok := ec.ExpressionElement().(*parser.ExpressionElementContext); ok && eec.Expression() != nil {
					ll.Elements = append(ll.Elements, b.visitExpression(eec.Expression()))
				}
			}
		}
	}
	return ll
}

func (b *AstBuilder) visitAssignableExpression(ctx parser.IAssignableExpressionContext) ast.Expr {
	ac, ok := ctx.(*parser.AssignableExpressionContext)
	if !ok {
		return nil
	}
	if ac.IDENTIFIER() != nil {
		return &ast.Identifier{Name: ac.IDENTIFIER().GetText()}
	}
	if ac.Primary() != nil {
		expr := b.visitPrimary(ac.Primary())
		if asp := ac.AssignableSelectorPart(); asp != nil {
			if asc, ok := asp.(*parser.AssignableSelectorPartContext); ok {
				if assignSel := asc.AssignableSelector(); assignSel != nil {
					if usc, ok := assignSel.(*parser.AssignableSelectorContext); ok {
						if uas := usc.UnconditionalAssignableSelector(); uas != nil {
							if uc, ok := uas.(*parser.UnconditionalAssignableSelectorContext); ok && uc.LBRACKET() != nil && uc.Expression() != nil {
								idx := b.visitExpression(uc.Expression())
								expr = &ast.IndexExpr{Target: expr, Index: idx}
							}
						}
					}
				}
			}
		}
		return expr
	}
	return nil
}

func (b *AstBuilder) visitArgumentPart(ctx parser.IArgumentPartContext) []ast.Expr {
	if ac, ok := ctx.(*parser.ArgumentPartContext); ok && ac.Arguments() != nil {
		return b.visitArguments(ac.Arguments())
	}
	return nil
}

func (b *AstBuilder) visitArguments(ctx parser.IArgumentsContext) []ast.Expr {
	if ac, ok := ctx.(*parser.ArgumentsContext); ok && ac.ArgumentList() != nil {
		return b.visitArgumentList(ac.ArgumentList())
	}
	return nil
}

func (b *AstBuilder) visitArgumentList(ctx parser.IArgumentListContext) []ast.Expr {
	ac, ok := ctx.(*parser.ArgumentListContext)
	if !ok || ac.ExpressionList() == nil {
		return nil
	}
	var args []ast.Expr
	for _, expr := range ac.ExpressionList().AllExpression() {
		args = append(args, b.visitExpression(expr))
	}
	return args
}

func (b *AstBuilder) visitExpressionListAsExpr(ctx parser.IExpressionListContext) ast.Expr {
	if ctx != nil && len(ctx.AllExpression()) > 0 {
		return b.visitExpression(ctx.Expression(0))
	}
	return nil
}

func (b *AstBuilder) visitExpressionWithoutCascade(ctx parser.IExpressionWithoutCascadeContext) ast.Expr {
	for _, e := range ctx.GetChildren() {
		if exprCtx, ok := e.(*parser.ExpressionContext); ok {
			return b.visitExpression(exprCtx)
		}
	}
	return nil
}
