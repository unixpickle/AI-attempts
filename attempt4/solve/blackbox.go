package solve

// A black box provides a stream of questions and their assorted answers.
type BlackBox interface {
	AnswerLen() int
	Ask(last interface{}) (question []bool, state interface{})
	Check(answer []bool, last interface{}) (right bool, state interface{})
	Initial() interface{}
	QuestionLen() int
}
