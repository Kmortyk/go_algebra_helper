package comp

import "fmt"

type SolveStep struct {
	mod   int
	parts []int

	powPart int
	pow     int

	solsCount int
}

type Solution struct {
	// x = base + step * y
	Base int
	Step int
}

func (s *SolveStep) AppendPart(part int) {
	s.parts = append(s.parts, part)
}

func (s *SolveStep) DecOne() {
	s.AppendPart(s.powPart)
	s.pow -= 1
}

func (s *SolveStep) Print() {
	for _, p := range s.parts {
		fmt.Print(p, "*")
	}
	fmt.Printf("(%v)^%v\n", s.powPart, s.pow)
}

func (s *SolveStep) PrintSolutions() {

	sol := s.parts[0]
	for i := 0; i < s.solsCount; i++ {
		fmt.Printf("x%v = %v + %v*%v = %v\n",
			i, sol, s.mod, i, sol+s.mod*i)
	}
}

func (s *SolveStep) Next() bool {

	if s.pow%2 != 0 { // if not even
		s.DecOne()
	}

	sqPart := s.powPart * s.powPart
	sqPart %= s.mod

	s.powPart = sqPart
	s.pow /= 2

	return s.pow > 1
}

func (s *SolveStep) SimplifyParts() {

	res := s.powPart

	for _, p := range s.parts {
		fmt.Printf("%v * %v = %v (mod %v)\n", res, p, (res*p)%s.mod, s.mod)

		res *= p
		res %= s.mod
	}

	s.powPart = 1
	s.parts = []int{res}
}

func (s *SolveStep) Solve() Solution {
	fmt.Println("Power simplification:")
	s.Print()
	for s.Next() {
		s.Print()
	}
	s.Print()
	fmt.Println("\nParts multiplication:")
	s.SimplifyParts()
	fmt.Println("\nx0:")
	s.Print()
	fmt.Println("\nSolutions:")
	s.PrintSolutions()

	return Solution{s.parts[0], s.mod}
}

func CreateSolveStep(exp Expr) SolveStep {
	return SolveStep{
		mod:       exp.M,
		parts:     []int{exp.B},
		powPart:   exp.A,
		pow:       Phi(exp.M) - 1,
		solsCount: exp.SolutionsCount,
	}
}

func Phi(num int) int {
	result := 1
	for i := 2; i < num; i++ {
		if GCD(i, num) == 1 {
			result++
		}
	}

	return result
}
