package ast

import "github.com/keisuke123/monkey/token"

type Node interface {
	TokenLiteral() string
}

// ASTのすべてのノードはNodeインタフェースを実装しなければならない
type Statement interface {
	Node // embedding
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type Identifier struct {
	Token token.Token
	Value string
}

// Expression interfaceとNode interfaceを満たす
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier // 束縛の識別子
	Value Expression  // 右辺の式
}

// Statement interfaceとNode interfaceを満たす
func (l *LetStatement) statementNode()       {}
func (l *LetStatement) TokenLiteral() string { return l.Token.Literal }

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}
