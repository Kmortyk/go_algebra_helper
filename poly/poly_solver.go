package poly

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

const (
	NUM = iota
	SIGN
	X
	UNK
	END
)

type Term struct {
	sign  float64
	coeff float64
	pow   float64
}

func (t *Term) Eval(v float64) float64 {
	return t.sign * t.coeff * math.Pow(v, t.pow)
}

type Token struct {
	val string
	typ uint8
}

type Expr struct {
	fx  string
	cur int
}

func NewExpr(fx string) Expr {
	return Expr{fx, 0}
}

func (exp *Expr) nextChar() {
	exp.cur++
}

func (exp *Expr) curChar() rune {
	return rune(exp.fx[exp.cur])
}

func (exp *Expr) isEnd() bool {
	return exp.cur >= len(exp.fx)
}

func (exp *Expr) GetToken() Token {
	if exp.isEnd() {
		return Token{"\000", END}
	}
	char := exp.curChar()
	// X
	if char == 'x' {
		exp.nextChar()
		return Token{string(char), X}
	}
	// NUM
	if unicode.IsDigit(char) {
		res := ""
		for ; !exp.isEnd() && unicode.IsDigit(exp.curChar()); exp.nextChar() {
			res += string(exp.curChar())
		}
		return Token{res, NUM}
	}
	// SIGN
	if char == '+' || char == '-' {
		exp.nextChar()
		return Token{string(char), SIGN}
	}
	// UNKNOWN
	return Token{"\000", UNK}
}

func (exp *Expr) Parse() []Term {
	tok := exp.GetToken()

	var terms []Term
	var coeffFlag = true
	var coeff = 1.0
	var sign = 1.0

	newTerm := func(pow float64) {
		terms = append(terms, Term{sign, coeff, pow})
		coeffFlag = true
		sign = 1
		coeff = 1
	}

	for true {
		switch tok.typ {
		case SIGN:
			if tok.val == "+" {
				sign = 1
			} else {
				sign = -1
			}
		case NUM:
			num, _ := strconv.ParseFloat(tok.val, 64)
			if coeffFlag {
				coeff = num
			} else {
				newTerm(num)
			}
		case X:
			coeffFlag = false
		case END:
			newTerm(0)
			return terms
		}

		tok = exp.GetToken()
	}

	return terms
}

func (exp *Expr) PrintSolving(mod int) {
	terms := exp.Parse()

	for i := 0; i < mod; i++ {
		result := 0.0

		fmt.Printf("f(%v)= ", i)
		for _, term := range terms {
			val := term.Eval(float64(i))
			result += val
			if term.sign == 1 {
				fmt.Print("+", val)
			} else {
				fmt.Print("-", val)
			}
		}

		fmt.Print(" = ", result)
		fmt.Printf(" ≡ %v (mod %v)", int(result)%mod, mod)
		fmt.Println()
	}
}
