package poly

import (
	"reflect"
	"testing"
)

func TestGetToken(t *testing.T) {
	e := NewExpr("x5+x2+3")
	var got []string

	tok := e.GetToken()
	for tok.typ != END {
		got = append(got, tok.val)
		tok = e.GetToken()
	}

	want := []string{"x", "5", "+", "x", "2", "+", "3"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetToken() = %v, want %v", got, want)
	}
}

func TestExpr_Parse(t *testing.T) {
	e := NewExpr("x5+x2+3")

	got := e.Parse()
	want := []Term{
		{1, 1, 5},
		{1, 1, 2},
		{1, 3, 0}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expr_Parse() = %v, want %v", got, want)
	}
}

func TestExpr_PrintSolving(t *testing.T) {
	e := NewExpr("3x6+4x20+2")
	e.PrintSolving(6)
}

func TestSuperscript(t *testing.T) {
	num := 64
	got := Superscript(num)
	want := "⁶⁴"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Superscript() = %v, want %v", got, want)
	}
}
