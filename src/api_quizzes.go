package adept

import (
	"encoding/json"
	"fmt"
)

type Quiz struct {
	ID               int64  `json:"quiz_id"`
	Name             string `json:"quiz_name"`
	ShortDescription string `json:"short_description"`
	GUID             string `json:"quiz_guid"`
	CourseID         int64
	Level            int64
}

type Quizzes struct {
	Quizzes []Quiz
}

func NewQuizzes() Quizzes {
	var q Quizzes

	return q
}

type QuizQuestions struct {
	IDs []int64 `json:"ids"`
}

func CreateQuiz(q Questions, t int64, uuid string) {
	var qi QuizQuestions
	for i := range q.Questions {
		qi.IDs = append(qi.IDs, q.Questions[i].ID)
	}
	d, _ := json.Marshal(qi)
	quests := string(d)
	DB.Exec("INSERT INTO quizzes_keys set quiz_timestamp=?, quiz_key=?, quiz_questions=? ", t, uuid, quests)

}

func (qzs *Quizzes) Get(params ...map[string]string) {

	W, V := NewParams()
	ParamRoute(&W, &V, params)

	WhereString := W.Compile()

	rows, err := DB.Query("SELECT quiz_id, quiz_name, quiz_guid, quiz_short_description, course_id, quiz_level FROM quizzes "+WhereString+" LIMIT 100", V.params...)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var q Quiz
		err = rows.Scan(&q.ID, &q.Name, &q.GUID, &q.ShortDescription, &q.CourseID, &q.Level)
		fmt.Println(q)
		if err != nil {
		}

		qzs.Quizzes = append(qzs.Quizzes, q)
	}
}
