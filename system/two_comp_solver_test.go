package system

import (
	"algebraHelper/comp"
	"testing"
)

func TestSystemComp_Solve(t *testing.T) {
	s := System{}

	s.AddExpr(comp.NewExpr("2x = 1 (mod 3)"))
	s.AddExpr(comp.NewExpr("5x = 4 (mod 6)"))
	s.AddExpr(comp.NewExpr("6x = 2 (mod 5)"))

	s.Solve()
}
