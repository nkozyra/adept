package adept

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CourseView struct {
	Foo           string
	Course        Course
	Breadcrumbs   Breadcrumbs
	Quizzes       Quizzes
	Organizations []Organization
	UserSection   UserSection
	Inverse       bool
}

type CoursesView struct {
	Courses       Courses
	Breadcrumbs   Breadcrumbs
	UserSection   UserSection
	Inverse       bool
	Organizations []Organization
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	u := getSession(r)
	var csv CoursesView
	LoadTemplates()

	c := NewCourses()
	c.Get()

	csv.Courses = c
	csv.UserSection.User = u
	csv.Breadcrumbs = NewBreadcrumbs()
	csv.Breadcrumbs.Add("/courses", "Browse Catalog", false)
	Templates.ExecuteTemplate(w, "courses.html", csv)
}

func courseJoinHandler(w http.ResponseWriter, r *http.Request) {
	u := getSession(r)
	vars := mux.Vars(r)
	params := make(map[string]string)
	params["course_guid"] = vars["course"]

	course := NewCourses()
	course.Get(params)

	course.AddUser(u.ID)
	http.Redirect(w, r, "/dashboard", 302)
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	var cv CourseView
	LoadTemplates()

	vars := mux.Vars(r)
	params := make(map[string]string)
	params["course_guid"] = vars["course"]

	course := NewCourses()
	course.Get(params)

	// Get class quizzes
	qparams := make(map[string]string)
	qparams["course_id"] = strconv.FormatInt(course.Courses[0].ID, 10)
	cvs := NewQuizzes()
	cvs.Get(qparams)

	cv.Breadcrumbs = NewBreadcrumbs()
	cv.Breadcrumbs.Add("/dashboard", "Dashboard", true)
	cv.Breadcrumbs.Add("", course.Courses[0].Name, false)
	fmt.Println(cv.Breadcrumbs)

	cv.Quizzes = cvs
	cv.Course = course.Courses[0]
	fmt.Println(cv)

	Templates.ExecuteTemplate(w, "course.html", cv)
}
