package adept

import (
	"fmt"
	"net/http"
)

type DashboardView struct {
	Foo         string
	Courses     Courses
	NoCourses   bool
	Breadcrumbs Breadcrumbs
	UserSection UserSection
	Inverse     bool
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	u := getSession(r)
	LoadTemplates()

	fmt.Println(u)
	fmt.Println(u.ID)
	c := NewCourses()
	c.UserID = u.ID
	c.Get()
	dash := DashboardView{Courses: c}
	dash.UserSection.User = u
	dash.Foo = "bar"
	if len(c.Courses) < 1 {
		dash.NoCourses = true
	}
	dash.Breadcrumbs = NewBreadcrumbs()
	dash.Breadcrumbs.Add("/dashboard", "Dashboard", false)

	fmt.Println(dash)
	Templates.ExecuteTemplate(w, "dashboard.html", dash)
}
