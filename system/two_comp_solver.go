package system

import (
	"algebraHelper/comp"
	"fmt"
)

type System []comp.Expr

func (s *System) AddExpr(expr comp.Expr) {
	*s = append(*s, expr)
}

func (s *System) Solve() {

	var initSol comp.Solution
	alphabet := "xyztklmnijghfrqvwpseouabcd"

	// down
	for i := 0; i < len(*s); i++ {
		exp := (*s)[i]

		fmt.Printf("=== Solve comparsion [%v]: ===\n", i+1)

		if i == 0 {
			if !exp.Simplify() {
				fmt.Println("!!! Check error !!!")
				return
			}
			step := comp.CreateSolveStep(exp)
			initSol = step.Solve()
		} else {
			// update with x solution
			fmt.Printf("%v*(%v + %v*%v) = %v (mod %v)\n",
				exp.A, initSol.Base, initSol.Step, string(alphabet[i]),
				exp.B, exp.M)

			fmt.Printf("%v + %v*%v = %v (mod %v)\n",
				exp.A*initSol.Base, exp.A*initSol.Step, string(alphabet[i]),
				exp.B, exp.M)

			fmt.Printf("%v*%v = %v (mod %v)\n\n",
				exp.A*initSol.Step, string(alphabet[i]),
				exp.B-exp.A*initSol.Base, exp.M)

			newA := (exp.A * initSol.Step) % exp.M
			newB := exp.B - (exp.A*initSol.Base)%exp.M
			// negative values to positive
			for newB < 0 {
				newB += exp.M
			}

			exp.A, exp.B = newA, newB
			// solve
			if !exp.Simplify() {
				fmt.Println("!!! Check error !!!")
				return
			}
			step := comp.CreateSolveStep(exp)
			sol := step.Solve()
			// update x solution with new
			fmt.Println("New X:")
			fmt.Printf("%v = %v + %v*(%v + %v*%v)\n",
				string(alphabet[i]),
				initSol.Base, initSol.Step,
				sol.Base, sol.Step,
				string(alphabet[i+1]))

			initSol.Base += initSol.Step * sol.Base
			initSol.Step *= sol.Step
		}

		fmt.Println()
	}

	fmt.Println("=== System solution: ===")
	// result
	fmt.Printf("x = %v (mod %v)\n", initSol.Base, initSol.Step)
}
