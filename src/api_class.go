package adept

import (
	"fmt"
	"strconv"
)

type Course struct {
	ID               int64  `json:"course_id"`
	Name             string `json:"course_name"`
	GUID             string `json:"course_guid"`
	ImageURL         string `json:"course_image_url"`
	OrganizationName string `json:"organization_name"`
	ExternalID       string

	RegCount              int64
	CourseRegisteredCount int64
	Registered            bool
}

type Courses struct {
	Courses []Course

	UserID int
}

func NewCourses() Courses {
	var c Courses

	return c
}

func (cs *Courses) AddUser(u int) {
	DB.Exec("INSERT INTO users_courses SET course_id=?, user_id=?", cs.Courses[0].ID, u)
}

func (cs *Courses) Get(params ...map[string]string) {

	W, V := NewParams()
	ParamRoute(&W, &V, params)

	if cs.UserID != 0 {
		W.params = append(W.params, "course_id IN (SELECT course_id FROM users_courses WHERE user_id="+strconv.FormatInt(int64(cs.UserID), 16)+")")
	}

	WhereString := W.Compile()

	rows, err := DB.Query("SELECT course_id, course_name, course_guid, course_image_url, organization_name, (SELECT count(*) FROM users_courses WHERE course_id=c.course_id) as 'rcount', course_external_id FROM courses c LEFT JOIN organizations o ON o.organization_id=c.organization_id "+WhereString+" ", V.params...)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var c Course
		err = rows.Scan(&c.ID, &c.Name, &c.GUID, &c.ImageURL, &c.OrganizationName, &c.CourseRegisteredCount, &c.ExternalID)

		if c.CourseRegisteredCount > 0 {
			c.Registered = true
		}

		if err != nil {
		}

		cs.Courses = append(cs.Courses, c)
	}
}
