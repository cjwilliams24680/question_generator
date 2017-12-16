package main

import (
	"github.com/labstack/echo"
	"time"
	"math/rand"
	"net/http"
)

var questions []*Question
var numberOfQuestions int

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	questions = (&QuestionSource{}).buildQuestionsList()

	// Since the questions can't change during runtime, I can count it once and leave it.
	// If I decide to allow runtime modifications, I'll need to revisit this however
	numberOfQuestions = len(questions)

	e := echo.New()
	e.GET("generate", generateQuestion)

	e.Logger.Fatal(e.Start(":5000"))
}

func generateQuestion(c echo.Context) error {
	index := rand.Intn(numberOfQuestions)
	return c.String(http.StatusOK, questions[index].generateQuestion())
}


