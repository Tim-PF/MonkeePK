package lexer

import "github.com/Tim-PF/MonkeePK/token"


type Lexer struct {
	input string
	position int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch byte // current char under examination. Have to be byte for readchar. Uni could be more than one byte! Maybe I will add it later with emojis.
 }
// In lexer_test erklärt!
 func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
 }

// Funktion lädt einen Token "tok" skippt leerzeichen und je nachdem welcher char ist wird der char an newToken weitergeleitet als Parameter !

func (l *Lexer) NextToken() token.Token {
	var tok token.Token  
	l.skipWhitespace()

	switch l.ch { // Erst besodnere Zeichen dann fürs erste werden Zeichen und Zahlen gesucht! Wenn nichts gefunden wird = Illegal. 
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case 0: // EOF
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok

		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// Skippt whitespace
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
	l.readChar()
	} 
}

//  Updated die Position bis EOF erreicht ist im Input. Position = aktuelle Position readPosition = Position + 1
func (l *Lexer) readChar() {
		if l.readPosition >= len(l.input) { // Wenn das nächste Zeichen leer ist also der Input vorbei => EOF
			l.ch = 0 //ASCII for NUL 
		} else {
			l.ch = l.input[l.readPosition] // Liest das nächste Zeichen ein
		}
		l.position = l.readPosition
		l.readPosition += 1
	 }

// Liest bis kein Identifier mehr also bis kein Letter mehr 
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}


// Liest Zahl bis sie vorbei ist.

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
	l.readChar()
	}
	return l.input[position:l.position]
	}

// Checkt ob input ein Buchstabe ist
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
	}
// Checkt ob Input eine Zahl ist
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
	}
	
// Erstellt aus dem switch case einen Token als return wert
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}





