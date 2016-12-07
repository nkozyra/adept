package adept

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
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

	vars := mux.Vars(r)

	qz := NewQuizzes()
	quizGUID := vars["quiz"]

	p := make(map[string]string)
	p["quiz_guid"] = quizGUID
	qz.Get(p)

	var ids []string
	var questions Questions
	for i := 0; i < 10; i++ {
		newID := GetRandom(qz.Quizzes[0].CourseID, qz.Quizzes[0].Level, 1, ids)
		ids = append(ids, strconv.FormatInt(newID.ID, int(10)))
		questions.Questions = append(questions.Questions, newID)
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
