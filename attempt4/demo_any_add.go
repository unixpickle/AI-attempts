package main

import (
	"fmt"
	"github.com/unixpickle/AI-attempts/attempt4/solve"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Box struct {
	InputSize  uint
	OutputSize uint
}

func (b Box) AnswerLen() int {
	return int(b.OutputSize)
}

func (b Box) Ask(last interface{}) ([]bool, interface{}) {
	num1 := rand.Intn(1 << b.InputSize)
	num2 := rand.Intn(1 << b.InputSize)
	res := []bool{}
	for _, num := range []int{num1, num2} {
		for i := uint(0); i < b.InputSize; i++ {
			mask := 1 << i
			res = append(res, (num&mask) != 0)
		}
	}
	answer := (num1 + num2) % (1 << b.OutputSize)
	answerList := []bool{}
	for i := uint(0); i < b.OutputSize; i++ {
		mask := 1 << i
		answerList = append(answerList, (answer&mask) != 0)
	}
	return res, answerList
}

func (b Box) Check(answer []bool, last interface{}) (bool, interface{}) {
	realAnswer := last.([]bool)
	for i, x := range answer {
		if realAnswer[i] != x {
			return false, last
		}
	}
	return true, last
}

func (b Box) Initial() interface{} {
	return []bool{}
}

func (b Box) QuestionLen() int {
	return int(b.InputSize * 2)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run demo_any_adder.go <input size> " +
			"<output size>")
		os.Exit(1)
	}
	inSize, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input size")
		os.Exit(1)
	}
	outSize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid output size")
		os.Exit(1)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Solving...")
	sampleCount := 128 << uint(inSize*2)
	n := solve.Solve(Box{uint(inSize), uint(outSize)}, sampleCount)
	fmt.Println("Solved:\n", n)
}
