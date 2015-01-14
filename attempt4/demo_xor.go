package main

import (
	"fmt"
	"github.com/unixpickle/AI-attempts/attempt4/solve"
	"math/rand"
	"time"
)

type Box struct{}

func (b Box) AnswerLen() int {
	return 1
}

func (b Box) Ask(last interface{}) ([]bool, interface{}) {
	res := []bool{rand.Intn(2) == 0, rand.Intn(2) == 0}
	return res, (res[0] != res[1])
}

func (b Box) Check(answer []bool, last interface{}) (bool, interface{}) {
	return answer[0] == last.(bool), last
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
