package adept

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Questions struct {
	Questions []Question
}

type Question struct {
	ID         int64  `json:"id"`
	Text       string `json:"text"`
	RawOptions string
	Options    []Option `json:"answers"`
}

type Option struct {
	Text    string `json:"text"`
	Correct bool   `json:"correct"`
}

type Answer struct {
	Text    string
	Correct bool
}

type QuestionIDs struct {
	IDs []int64 `json:"ids"`
}

func NewQuestions() Questions {
	var qs Questions
	return qs
}

func GetQuestionsByKey(key string) string {
	// 	var hm []string
	var qString string
	err := DB.QueryRow("SELECT quiz_questions FROM quizzes_keys WHERE quiz_key=?", key).Scan(&qString)
	if err != nil {

	}
	var ids QuestionIDs
	var qs Questions
	var qStrings []string
	fmt.Println(key, qString)
	json.Unmarshal([]byte(qString), &ids)

	for i := range ids.IDs {
		qStrings = append(qStrings, strconv.FormatInt(ids.IDs[i], 10))
	}

	fullIDs := strings.Join(qStrings, ",")
	fmt.Println(fullIDs)
	rows, _ := DB.Query("SELECT question_id, question_text, question_options FROM questions WHERE question_id IN (" + fullIDs + ")")

	defer rows.Close()
	for rows.Next() {
		var q Question
		err := rows.Scan(&q.ID, &q.Text, &q.RawOptions)
		fmt.Println(q)
		if err != nil {

		}
		json.Unmarshal([]byte(q.RawOptions), &q.Options)
		qs.Questions = append(qs.Questions, q)
	}

	out, _ := json.Marshal(qs)
	return string(out)
}

func GetRandom(cid int, level int, count int) Question {
	var q Question
	DB.QueryRow("SELECT r1.question_id, question_text, question_options FROM questions AS r1 JOIN (SELECT CEIL(RAND() * (SELECT MAX(question_id) FROM questions)) AS question_id) AS r2 WHERE r1.question_id >= r2.question_id ORDER BY r1.question_id ASC LIMIT 1").Scan(&q.ID, &q.Text, &q.RawOptions)
	return q
}
