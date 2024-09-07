package peruse

const (
	INVALID_PRECEDENCE Precedence = iota
	LOWEST
)

func NewParser[T any](g *Grammar[T], l TokenStream) *Parser[T] {
	p := new(Parser[T])
	p.gram = g
	p.lex = l
	// cycle first two tokens into place
	p.Consume()
	p.Consume()
	return p
}

type Parser[T any] struct {
	gram      *Grammar[T]
	lex       TokenStream
	Cur, Next Token
	empty     T
}

func (p *Parser[T]) Current() Token {
	return p.Cur
}

func (p *Parser[T]) Peek() Token {
	return p.Next
}

func (p *Parser[T]) HasMore() bool {
	return !(p.Next.Is(EOF) || p.Next.Is(ERR))
}

func (p *Parser[T]) Parse() (T, error) {
	return p.ParseAt(LOWEST)
}

func (p *Parser[T]) ParseAt(prd Precedence) (T, error) {
	var t T
	parser, ok := p.gram.parselets[p.Cur.Type]

	if !ok {
		return t, UnexpectedToken{p.Cur}
	}

	left, err := parser(p)

	if err != nil {
		return t, err
	}

	for thisPrd := p.NextPrecedence(); prd < thisPrd; thisPrd = p.NextPrecedence() {
		p.Consume()

		parselet, exists := p.gram.infixes[p.Cur.Type]
		if !exists {
			break
		}

		left, err = parselet(p, left, thisPrd)
		if err != nil {
			return left, err
		}
	}

	return left, nil
}

func (p *Parser[T]) NextPrecedence() Precedence {
	return p.gram.precedence[p.Next.Type]
}

func (p *Parser[T]) Consume() {
	p.Cur = p.Next
	p.Next = p.lex.NextToken()
}

func (p *Parser[T]) Expect(n TokenType) bool {
	if p.Next.Is(n) {
		p.Consume()
		return true
	}
	return false
}

func (p *Parser[T]) ExpectOrError(n TokenType) error {
	if !p.Expect(n) {
		return InvalidToken{Wanted: n, Have: p.Next}
	}

	return nil
}
