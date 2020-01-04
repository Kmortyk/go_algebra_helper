package system

import (
	"algebraHelper/comp"
	"fmt"
)

type System []comp.Expr

type SystemStep struct {
	step comp.SolveStep
	sol  comp.Solution
}

func (s *System) AddExpr(expr comp.Expr) {
	*s = append(*s, expr)
}

func (s *System) Solve() {

	var sols = []SystemStep{}

	// down
	for i := 0; i < len(*s); i++ {
		exp := (*s)[i]

		if i > 0 {
			prevSol := sols[i-1]
			newA := (exp.A * prevSol.sol.Step) % exp.M
			newB := exp.B - (exp.A*prevSol.sol.Base)%exp.M
			// negative values to positive
			for newB < 0 {
				newB += exp.M
			}
			exp.A, exp.B = newA, newB
		}
		fmt.Printf("=== Solve comparsion [%v]: ===\n", i+1)
		exp.Simplify()
		step := comp.CreateSolveStep(exp)
		sol := step.Solve()
		sols = append(sols, SystemStep{step, sol})
		fmt.Println()
	}

	// up
	for i := len(sols) - 2; i >= 0; i-- {
		nextSol := sols[i+1].sol
		curSol := sols[i].sol

		sols[i].sol.Base += curSol.Step * nextSol.Base
		sols[i].sol.Step *= nextSol.Step
	}

	// result
	sol := sols[0].sol
	fmt.Println("=== System solution: ===")
	fmt.Printf("x = %v (mod %v)\n", sol.Base, sol.Step)
}
