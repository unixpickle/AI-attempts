package main

import (
	"fmt"
	"github.com/unixpickle/AI-attempts/attempt4/solve"
	"math/rand"
	"time"
)

type Box struct{}

func (b Box) AnswerLen() int {
	return 2
}

func (b Box) Ask(last interface{}) ([]bool, interface{}) {
	res := []bool{rand.Intn(2) == 0, rand.Intn(2) == 0}
	answer := []bool{false, false}
	if res[0] && res[1] {
		answer = []bool{true, false}
	} else if res[0] || res[1] {
		answer = []bool{false, true}
	}
	return res, answer
}

func (b Box) Check(answer []bool, last interface{}) (bool, interface{}) {
	realAnswer := last.([]bool)
	res := realAnswer[0] == answer[0] && realAnswer[1] == answer[1]
	return res, last
}

func (b Box) Initial() interface{} {
	return false
}

func (b Box) QuestionLen() int {
	return 2
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Solving...")
	n := solve.Solve(Box{}, 128)
	fmt.Println("Solved:\n", n)
}
