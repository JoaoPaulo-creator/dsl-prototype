package ast

type ExpressionStmt struct {
	ExpressionStmt Expr
}

func (node ExpressionStmt) stmt() {}

type SetDeclStmt struct {
	Body []Stmt
}

func (node SetDeclStmt) stmt() {}

type BlockStmt struct {
	Body []Stmt
}

func (node BlockStmt) stmt() {}

type StructProperty struct {
	IsStatic      bool // is property static?
	AssignedValue Expr
}

type StructDeclStmt struct {
	StructName string
	Properties map[string]StructProperty
}

func (node StructDeclStmt) stmt() {}
