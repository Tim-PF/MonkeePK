import  "github.com/Tim-PF/MonkeePK/token"

package ast 
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

// Erstellt die root Node.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()

	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token .Let token
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()  {}
func (ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}

type Identifier struct {
	Token token.Token // the token.Ident token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {return i.Token.Literal }