package comp

import (
	"reflect"
	"testing"
)

func TestNewExpr(t *testing.T) {

	e := NewExpr("20x = 24 (mod 82)")

	want := Expr{"20x=24(mod82)", 13, 20, 24, 82, 2}

	if !reflect.DeepEqual(e, want) {
		t.Errorf("NewExpr() = %v, want %v", e, want)
	}
}

func TestExpr_RemoveMultiplicity(t *testing.T) {

	e := NewExpr("20x = 24 (mod 82)")
	e.RemoveMultiplicity()

	want := Expr{"20x=24(mod82)", 13, 10, 12, 41, 2}

	if !reflect.DeepEqual(e, want) {
		t.Errorf("RemoveMultiplicity() = %v, want %v", e, want)
	}
}

func TestExpr_Simplify(t *testing.T) {
	e := NewExpr("28x = 14 (mod 210)")
	e.Simplify()
}
