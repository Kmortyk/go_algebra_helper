package system

import (
	"algebraHelper/comp"
	"testing"
)

func TestSystemComp_Solve(t *testing.T) {
	s := System{}

	s.AddExpr(comp.NewExpr("3x = 2 (mod 4)"))
	s.AddExpr(comp.NewExpr("5x = 4 (mod 6)"))
	s.AddExpr(comp.NewExpr("4x = 3 (mod 5)"))

	s.Solve()
}
