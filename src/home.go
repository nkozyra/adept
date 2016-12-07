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
	Inverse     bool
	UserCount   int
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	u := getSession(r)
	LoadTemplates()

	c := NewCourses()
	c.UserID = u.ID
	c.Get()
	home := HomeView{Courses: c}
	home.UserSection.User = u
	home.Foo = "bar"
	home.Inverse = true
	home.Breadcrumbs = NewBreadcrumbs()
	home.UserCount = CountUsers()

	fmt.Println(home)
	Templates.ExecuteTemplate(w, "home.html", home)
}
