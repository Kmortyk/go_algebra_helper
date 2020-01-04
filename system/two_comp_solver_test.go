package system

import (
	"algebraHelper/comp"
	"testing"
)

func TestSystemComp_Solve(t *testing.T) {
	s := System{}
	//s.AddExpr(comp.NewExpr("10x = 40 (mod 25)"))
	//s.AddExpr(comp.NewExpr("25x = 8 (mod 11)"))

	s.AddExpr(comp.NewExpr("3x = 2 (mod 5)"))
	s.AddExpr(comp.NewExpr("2x = 1 (mod 3)"))
	s.AddExpr(comp.NewExpr("4x = 3 (mod 7)"))

	s.Solve()
}
