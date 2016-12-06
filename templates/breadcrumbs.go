package adept

type Breadcrumbs struct {
	Breadcrumbs []Breadcrumb
}

type Breadcrumb struct {
	Active bool
	Text   string
	Link   string
}

func NewBreadcrumbs() Breadcrumbs {
	var bc Breadcrumbs
	return bc
}
