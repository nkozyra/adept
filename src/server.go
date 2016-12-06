package adept

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "sup")
}

func Serve() {
	r := mux.NewRouter()

	assets := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	r.PathPrefix("/assets").Handler(assets)
	r.HandleFunc("/", homeHandler)

	r.HandleFunc("/api/quiz/{key}", quizQuestionHandler)

	r.HandleFunc("/register", registerHandler).Methods("GET")
	r.HandleFunc("/register", registerProcess).Methods("POST")
	r.HandleFunc("/login", authHandler).Methods("GET")
	r.HandleFunc("/login", authProcess).Methods("POST")
	r.HandleFunc("/logout", logoutHandler)
	r.HandleFunc("/dashboard", dashboardHandler)
	r.HandleFunc("/courses", coursesHandler)
	r.HandleFunc("/course/{course}", courseHandler)
	r.HandleFunc("/course/{course}/join", courseJoinHandler)
	r.HandleFunc("/course/{course}/quizzes", fooHandler)
	r.HandleFunc("/course/{course}/quizzes/{quiz}", quizGenerator)
	r.HandleFunc("/course/{course}/quizzes/{quiz}/{key}", quizHandler)
	r.HandleFunc("/course/{course}/leaderboard", fooHandler)

	http.Handle("/", r)
	http.ListenAndServe(":7500", r)
}
