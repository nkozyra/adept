package adept

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func quizQuestionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	o := GetQuestionsByKey(key)
	fmt.Fprintln(w, o)
}
