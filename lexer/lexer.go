package lexer

import (
	"fmt"
	"io"
)

type location struct {
	line uint
	col  uint
}

type keyword string

const (
	selectKeyword keyword = "select"
	fromKeyword   keyword = "from"
	asKeyword     keyword = "as"
	tableKeyword  keyword = "table"
	createKeyword keyword = "create"
	insertKeyword keyword = "insert"
	intoKeyword   keyword = "into"
	valuesKeyword keyword = "values"
	intKeyword    keyword = "int"
	textKeyword   keyword = "text"
)

type symbol string

const (
	semicolonSymbol  symbol = ";"
	asteriskSymbol   symbol = "*"
	commaSymbol      symbol = ","
	leftparenSymbol  symbol = "("
	rightparenSymbol symbol = ")"
)

type tokenKind uint

const (
	keywordKind tokenKind = iota
	symbolKind
	identifierKind
	stringKind
	numbericKind
)

type token struct {
	value string
	kind  tokenKind
	loc   location
}

func (t *token) equals(other *token) bool {
	return t.value == other.value && t.kind == other.kind
}

func (t *token) finalize() bool {
	return true
}

func lex(source io.Reader) ([]*token, error) {
	buf := make([]byte, 1)
	tokens := []*token{}
	current := token{}
	var (
		line uint = 0
		col  uint = 0
	)
	for {
		// 从源文件读取字节
		_, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		var c byte = ';'
		if err == nil {
			c = buf[0]
		}
		switch c {
		case '\n':
			line++
			col = 0
			continue
		case ' ', ',', '(', ')', ';':
			if !current.finalize() {
				return nil, fmt.Errorf("未知的 token ‘%s’ at %d:%d", current.value, current.loc.line, current.loc.col)
			}
			if current.value != "" {
				copy1 := current
				tokens = append(tokens, &copy1)
			}
			if c == ';' || c == ',' || c == '(' || c == ')' {
				tokens = append(tokens, &token{
					value: string(c),
					kind:  symbolKind,
					loc: location{
						col:  col,
						line: line,
					},
				})
			}
			current = token{}
			current.loc.col = col
			current.loc.line = line
		default:
			current.value += string(c)
		}
		if err == io.EOF {
			break
		}
		col++
	}
	return tokens, nil
}
