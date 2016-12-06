package adept

import (
	"fmt"
	"net/http"
)

type HomeView struct {
	Foo         string
	Courses     Courses
	Breadcrumbs Breadcrumbs
	UserSection UserSection
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	u := getSession(r)
	LoadTemplates()

	c := NewCourses()
	c.UserID = u.ID
	c.Get()
	dash := HomeView{Courses: c}
	dash.UserSection.User = u
	dash.Foo = "bar"
	dash.Breadcrumbs = NewBreadcrumbs()

	fmt.Println(dash)
	Templates.ExecuteTemplate(w, "dashboard.html", dash)
}
