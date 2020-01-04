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
	alphabet := "xyztklmnijghfrqvwpseouabcd"

	// down
	for i := 0; i < len(*s); i++ {
		exp := (*s)[i]

		fmt.Printf("=== Solve comparsion [%v]: ===\n", i+1)
		if i > 0 {
			prevSol := sols[i-1]

			fmt.Printf("%v*(%v + %v*%v) = %v (mod %v)\n",
				exp.A, prevSol.sol.Base, prevSol.sol.Step, string(alphabet[i]),
				exp.B, exp.M)

			fmt.Printf("%v + %v*%v = %v (mod %v)\n",
				exp.A*prevSol.sol.Base, exp.A*prevSol.sol.Step, string(alphabet[i]),
				exp.B, exp.M)

			fmt.Printf("%v*%v = %v (mod %v)\n\n",
				exp.A*prevSol.sol.Step, string(alphabet[i]),
				exp.B-exp.A*prevSol.sol.Base, exp.M)

			newA := (exp.A * prevSol.sol.Step) % exp.M
			newB := exp.B - (exp.A*prevSol.sol.Base)%exp.M
			// negative values to positive
			for newB < 0 {
				newB += exp.M
			}

			exp.A, exp.B = newA, newB
		}
		exp.Simplify()
		step := comp.CreateSolveStep(exp)
		sol := step.Solve()
		sols = append(sols, SystemStep{step, sol})
		fmt.Println()
	}

	fmt.Println("=== System solution: ===")

	// up
	for i := len(sols) - 2; i >= 0; i-- {
		nextSol := sols[i+1].sol
		curSol := sols[i].sol

		fmt.Printf("%v = %v + %v*(%v + %v*%v)\n",
			string(alphabet[i]),
			sols[i].sol.Base, sols[i].sol.Step,
			nextSol.Base, nextSol.Step,
			string(alphabet[i+1]))
		sols[i].sol.Base += curSol.Step * nextSol.Base
		sols[i].sol.Step *= nextSol.Step
	}

	// result
	sol := sols[0].sol
	fmt.Printf("x = %v (mod %v)\n", sol.Base, sol.Step)
}
