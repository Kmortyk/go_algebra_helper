package comp

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/* solve comparsions like
 * ax ≡ B (mod M)
 */

type Expr struct {
	Fx  string
	cur int

	A int
	B int
	M int

	SolutionsCount int
}

func (exp *Expr) nextChar() {
	exp.cur++
}

func (exp *Expr) curChar() rune {
	return rune(exp.Fx[exp.cur])
}

func (exp *Expr) isEnd() bool {
	return exp.cur >= len(exp.Fx)
}

func (exp *Expr) getNum() (int, error) {
	if unicode.IsDigit(exp.curChar()) {
		res := ""
		for ; !exp.isEnd() && unicode.IsDigit(exp.curChar()); exp.nextChar() {
			res += string(exp.curChar())
		}
		return strconv.Atoi(res)
	}

	return 0, errors.New("parse error: not A number")
}

func (exp *Expr) getChar() rune {
	defer exp.nextChar()
	return exp.curChar()
}

func (exp *Expr) getStr() string {
	res := ""
	if unicode.IsLetter(exp.curChar()) {
		for ; !exp.isEnd() && unicode.IsLetter(exp.curChar()); exp.nextChar() {
			res += string(exp.curChar())
		}
	}

	return res
}

func NewExpr(fx string) Expr {
	// remove whitespaces
	fx = strings.Replace(fx, " ", "", -1)

	e := Expr{fx, 0, 0, 0, 0, 0}
	e.A, _ = e.getNum()

	e.getChar() // x
	e.getChar() // ≡

	e.B, _ = e.getNum()

	e.getChar() // (
	e.getStr()  // mod

	e.M, _ = e.getNum()

	e.getChar() // )

	e.SolutionsCount = GCD(e.A, e.M)

	return e
}

func (exp *Expr) Check() bool {
	gcd := GCD(exp.A, exp.M)

	fmt.Printf("GCD(%v, %v) = %v (solutions) \n\n", exp.A, exp.M, gcd)
	fmt.Printf("Check:\n%v ⁝ %v\n\n", exp.B, gcd)

	return (exp.B % gcd) == 0
}

func (exp *Expr) RemoveMultiplicity() bool {
	gcd := GCD3(exp.A, exp.B, exp.M)

	if gcd == 1 {
		return false
	}

	exp.A /= gcd
	exp.B /= gcd
	exp.M /= gcd

	return true
}

func (exp *Expr) RemoveSemiMultiplicity() bool {
	gcd := GCD(exp.A, exp.B)

	if gcd != 1 && GCD(gcd, exp.M) == 1 {
		exp.A /= gcd
		exp.B /= gcd
		return true
	}

	return false
}

func (exp *Expr) MakeMod() bool {
	prevA := exp.A
	prevB := exp.B

	exp.A %= exp.M
	exp.B %= exp.M

	return prevA != exp.A || prevB != exp.B
}

func (exp *Expr) Print() {
	fmt.Printf("%vx ≡ %v (mod %v)\n", exp.A, exp.B, exp.M)
}

func (exp *Expr) Println() {
	exp.Print()
	fmt.Println()
}

func (exp *Expr) Simplify() bool {

	maxSteps := 10
	curStep := 0

	fmt.Println("Initial:")
	exp.Print()

	if !exp.Check() {
		fmt.Println("!!! Check error: no solutions !!!")
		exp.Println()
		return false
	}

	for exp.A != 1 && curStep < maxSteps {
		curStep++

		if exp.RemoveMultiplicity() {
			fmt.Println("RemoveMultiplicity:")
			exp.Println()
		}

		if exp.RemoveSemiMultiplicity() {
			fmt.Println("RemoveSemiMultiplicity:")
			exp.Println()
		}

		if exp.MakeMod() {
			fmt.Println("MakeMod:")
			exp.Println()
		}
	}

	return true
}

func GCD3(a int, b int, m int) int {
	gcd1 := GCD(a, b)
	return GCD(gcd1, m)
}

func GCD(a int, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
