package comp

import (
	"testing"
)

func TestSolveStep_Solve(t *testing.T) {
	e := NewExpr("10x = 40 (mod 25)")

	if e.Simplify() {
		s := CreateSolveStep(e)
		s.Solve()
	}
}
