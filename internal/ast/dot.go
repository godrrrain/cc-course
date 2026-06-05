package ast

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type dotVisitor struct {
	builder strings.Builder
	nodeID  int
}

func (v *dotVisitor) nextID() int {
	v.nodeID++
	return v.nodeID
}

func (v *dotVisitor) writeNode(id int, label string) {
	label = strings.ReplaceAll(label, "\"", "\\\"")
	fmt.Fprintf(&v.builder, "  node%d [label=\"%s\"];\n", id, label)
}

func (v *dotVisitor) writeEdge(from, to int) {
	fmt.Fprintf(&v.builder, "  node%d -> node%d;\n", from, to)
}

func (v *dotVisitor) visitNode(n Node) int {
	id := v.nextID()
	v.writeNode(id, n.nodeStr())
	return id
}

func (v *dotVisitor) visitProgram(prog *Program) int {
	id := v.visitNode(prog)
	for _, d := range prog.Declarations {
		childID := v.visitDecl(d)
		v.writeEdge(id, childID)
	}
	return id
}

func (v *dotVisitor) visitDecl(d Decl) int {
	switch dd := d.(type) {
	case *FuncDecl:
		return v.visitFuncDecl(dd)
	case *VarDecl:
		return v.visitVarDecl(dd)
	}
	return v.visitNode(d.(Node))
}

func (v *dotVisitor) visitFuncDecl(fd *FuncDecl) int {
	id := v.visitNode(fd)

	for _, p := range fd.Params {
		pid := v.nextID()
		v.writeNode(pid, fmt.Sprintf("Param %s %s", p.Name, p.Type.Name))
		v.writeEdge(id, pid)
	}

	if fd.ReturnType != nil {
		rid := v.nextID()
		v.writeNode(rid, "Returns "+fd.ReturnType.Name)
		v.writeEdge(id, rid)
	}

	if fd.Body != nil {
		bid := v.visitBlockStmt(fd.Body)
		v.writeEdge(id, bid)
	}

	if fd.BodyExpr != nil {
		eid := v.visitExpr(fd.BodyExpr)
		v.writeEdge(id, eid)
	}

	return id
}

func (v *dotVisitor) visitVarDecl(vd *VarDecl) int {
	id := v.visitNode(vd)
	if vd.Init != nil {
		v.writeEdge(id, v.visitExpr(vd.Init))
	}
	return id
}

func (v *dotVisitor) visitStmt(s Stmt) int {
	switch ss := s.(type) {
	case *BlockStmt:
		return v.visitBlockStmt(ss)
	case *IfStmt:
		return v.visitIfStmt(ss)
	case *WhileStmt:
		return v.visitWhileStmt(ss)
	case *DoWhileStmt:
		return v.visitDoWhileStmt(ss)
	case *ForStmt:
		return v.visitForStmt(ss)
	case *ForInStmt:
		return v.visitForInStmt(ss)
	case *ExprStmt:
		return v.visitExprStmt(ss)
	case *ReturnStmt:
		return v.visitReturnStmt(ss)
	case *BreakStmt:
		return v.visitNode(ss)
	case *ContinueStmt:
		return v.visitNode(ss)
	case *AssertStmt:
		return v.visitAssertStmt(ss)
	case *VarDecl:
		return v.visitVarDecl(ss)
	}
	return v.visitNode(s.(Node))
}

func (v *dotVisitor) visitBlockStmt(bs *BlockStmt) int {
	id := v.visitNode(bs)
	for _, s := range bs.Statements {
		v.writeEdge(id, v.visitStmt(s))
	}
	return id
}

func (v *dotVisitor) visitIfStmt(is *IfStmt) int {
	id := v.visitNode(is)
	if is.Condition != nil {
		v.writeEdge(id, v.visitExpr(is.Condition))
	}
	if is.Then != nil {
		v.writeEdge(id, v.visitStmt(is.Then))
	}
	if is.Else != nil {
		v.writeEdge(id, v.visitStmt(is.Else))
	}
	return id
}

func (v *dotVisitor) visitWhileStmt(ws *WhileStmt) int {
	id := v.visitNode(ws)
	if ws.Condition != nil {
		v.writeEdge(id, v.visitExpr(ws.Condition))
	}
	if ws.Body != nil {
		v.writeEdge(id, v.visitStmt(ws.Body))
	}
	return id
}

func (v *dotVisitor) visitDoWhileStmt(ds *DoWhileStmt) int {
	id := v.visitNode(ds)
	if ds.Body != nil {
		v.writeEdge(id, v.visitStmt(ds.Body))
	}
	if ds.Condition != nil {
		v.writeEdge(id, v.visitExpr(ds.Condition))
	}
	return id
}

func (v *dotVisitor) visitForStmt(fs *ForStmt) int {
	id := v.visitNode(fs)
	if fs.Init != nil {
		v.writeEdge(id, v.visitStmt(fs.Init))
	}
	if fs.Condition != nil {
		v.writeEdge(id, v.visitExpr(fs.Condition))
	}
	if fs.Update != nil {
		v.writeEdge(id, v.visitExpr(fs.Update))
	}
	if fs.Body != nil {
		v.writeEdge(id, v.visitStmt(fs.Body))
	}
	return id
}

func (v *dotVisitor) visitForInStmt(fis *ForInStmt) int {
	id := v.visitNode(fis)
	if fis.Iterable != nil {
		v.writeEdge(id, v.visitExpr(fis.Iterable))
	}
	if fis.Body != nil {
		v.writeEdge(id, v.visitStmt(fis.Body))
	}
	return id
}

func (v *dotVisitor) visitExprStmt(es *ExprStmt) int {
	id := v.visitNode(es)
	if es.Expression != nil {
		v.writeEdge(id, v.visitExpr(es.Expression))
	}
	return id
}

func (v *dotVisitor) visitReturnStmt(rs *ReturnStmt) int {
	id := v.visitNode(rs)
	if rs.Value != nil {
		v.writeEdge(id, v.visitExpr(rs.Value))
	}
	return id
}

func (v *dotVisitor) visitAssertStmt(as *AssertStmt) int {
	id := v.visitNode(as)
	if as.Condition != nil {
		v.writeEdge(id, v.visitExpr(as.Condition))
	}
	return id
}

func (v *dotVisitor) visitExpr(e Expr) int {
	if e == nil {
		return v.visitNode(nilNode{})
	}
	switch ee := e.(type) {
	case *BinaryExpr:
		return v.visitBinaryExpr(ee)
	case *UnaryExpr:
		return v.visitUnaryExpr(ee)
	case *PostfixExpr:
		return v.visitPostfixExpr(ee)
	case *AssignExpr:
		return v.visitAssignExpr(ee)
	case *CallExpr:
		return v.visitCallExpr(ee)
	case *Identifier:
		return v.visitNode(ee)
	case *IntLiteral:
		return v.visitNode(ee)
	case *FloatLiteral:
		return v.visitNode(ee)
	case *BoolLiteral:
		return v.visitNode(ee)
	case *StringLiteral:
		return v.visitNode(ee)
	case *NullLiteral:
		return v.visitNode(ee)
	case *ListLiteral:
		return v.visitListLiteral(ee)
	case *IndexExpr:
		return v.visitIndexExpr(ee)
	case *TernaryExpr:
		return v.visitTernaryExpr(ee)
	case *InterpolatedStringExpr:
		return v.visitInterpolatedStringExpr(ee)
	}
	return v.visitNode(e.(Node))
}

func (v *dotVisitor) visitBinaryExpr(be *BinaryExpr) int {
	id := v.visitNode(be)
	v.writeEdge(id, v.visitExpr(be.Left))
	v.writeEdge(id, v.visitExpr(be.Right))
	return id
}

func (v *dotVisitor) visitUnaryExpr(ue *UnaryExpr) int {
	id := v.visitNode(ue)
	v.writeEdge(id, v.visitExpr(ue.Operand))
	return id
}

func (v *dotVisitor) visitPostfixExpr(pe *PostfixExpr) int {
	id := v.visitNode(pe)
	v.writeEdge(id, v.visitExpr(pe.Target))
	return id
}

func (v *dotVisitor) visitAssignExpr(ae *AssignExpr) int {
	id := v.visitNode(ae)
	v.writeEdge(id, v.visitExpr(ae.Target))
	v.writeEdge(id, v.visitExpr(ae.Value))
	return id
}

func (v *dotVisitor) visitCallExpr(ce *CallExpr) int {
	id := v.visitNode(ce)
	v.writeEdge(id, v.visitExpr(ce.Function))
	for _, a := range ce.Args {
		v.writeEdge(id, v.visitExpr(a))
	}
	return id
}

func (v *dotVisitor) visitListLiteral(ll *ListLiteral) int {
	id := v.visitNode(ll)
	for _, e := range ll.Elements {
		v.writeEdge(id, v.visitExpr(e))
	}
	return id
}

func (v *dotVisitor) visitIndexExpr(ix *IndexExpr) int {
	id := v.visitNode(ix)
	v.writeEdge(id, v.visitExpr(ix.Target))
	v.writeEdge(id, v.visitExpr(ix.Index))
	return id
}

func (v *dotVisitor) visitTernaryExpr(t *TernaryExpr) int {
	id := v.visitNode(t)
	v.writeEdge(id, v.visitExpr(t.Condition))
	v.writeEdge(id, v.visitExpr(t.Then))
	v.writeEdge(id, v.visitExpr(t.Else))
	return id
}

func (v *dotVisitor) visitInterpolatedStringExpr(is *InterpolatedStringExpr) int {
	id := v.visitNode(is)
	for _, p := range is.Parts {
		var pid int
		if p.IsExpr {
			pid = v.visitExpr(p.Expr)
		} else {
			pid = v.nextID()
			v.writeNode(pid, "Text: "+p.Text)
		}
		v.writeEdge(id, pid)
	}
	return id
}

type nilNode struct{}

func (n nilNode) nodeStr() string { return "nil" }

func generateASTDOT(prog *Program) string {
	v := &dotVisitor{}
	v.builder.WriteString("digraph AST {\n")
	v.builder.WriteString("  node [shape=box];\n")
	v.visitProgram(prog)
	v.builder.WriteString("}\n")
	return v.builder.String()
}

func SaveASTToFile(prog *Program, filename string) error {
	dot := generateASTDOT(prog)

	err := os.WriteFile(filename+".dot", []byte(dot), 0644)
	if err != nil {
		return fmt.Errorf("file %s write error: %w", filename+".dot", err)
	}

	cmd := exec.Command("dot", "-Tpng", "-o", filename+".png", filename+".dot")
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("dot -Tpng error: %w\n%s", err, string(output))
	}

	return nil
}
