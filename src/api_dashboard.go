package adept

type Dashboard struct {
	Courses []Course `json:"courses"`
}

func NewDashboard() Dashboard {
	var d Dashboard
	return d
}
