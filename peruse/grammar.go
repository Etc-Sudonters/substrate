package peruse

type Parselet[T any] func(*Parser[T]) (T, error)
type InflixParselet[T any] func(*Parser[T], T, Precedence) (T, error)
type Precedence uint

type Grammar[T any] struct {
	parselets  map[TokenType]Parselet[T]
	infixes    map[TokenType]InflixParselet[T]
	precedence map[TokenType]Precedence
}

func NewGrammar[T any]() Grammar[T] {
	var g Grammar[T]
	g.parselets = make(map[TokenType]Parselet[T])
	g.infixes = make(map[TokenType]InflixParselet[T])
	g.precedence = make(map[TokenType]Precedence)
	return g
}

func (g Grammar[T]) Parse(t TokenType, p Parselet[T]) {
	g.parselets[t] = p
}

func (g Grammar[T]) Infix(p Precedence, i InflixParselet[T], ts ...TokenType) {
	for _, t := range ts {
		g.infixes[t] = i
		g.precedence[t] = p
	}
}

func (g Grammar[T]) Precedence(t TokenType) Precedence {
	return g.precedence[t]
}
