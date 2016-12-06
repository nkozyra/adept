package adept

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type QuizPage struct {
	Key         string
	Inverse     bool
	UserSection UserSection
}

func quizGenerator(w http.ResponseWriter, r *http.Request) {
	var questions Questions
	for i := 0; i < 10; i++ {
		questions.Questions = append(questions.Questions, GetRandom(1, 1, 1))
	}

	validTimestamp := time.Now().Unix()

	out, err := exec.Command("uuidgen").Output()
	if err != nil {

	}
	out = []byte(strings.Trim(string(out), " \t\n"))

	questionVal, _ := json.Marshal(questions)
	fmt.Println(string(questionVal))
	fmt.Println(validTimestamp)

	CreateQuiz(questions, validTimestamp, string(out))

	http.Redirect(w, r, "/course/computational-photography/quizzes/level-1/"+string(out), http.StatusFound)

}

func quizHandler(w http.ResponseWriter, r *http.Request) {
	u := getSession(r)
	vars := mux.Vars(r)

	Page := QuizPage{Key: vars["key"], Inverse: true}

	Page.UserSection.User = u

	LoadTemplates()
	Templates.ExecuteTemplate(w, "quiz.html", Page)
}
