package adept

type Breadcrumbs struct {
	Breadcrumbs []Breadcrumb
}

type Breadcrumb struct {
	Active bool
	Text   string
	Link   string
}

func (bc *Breadcrumbs) Add(link string, text string, active bool) {
	bc.Breadcrumbs = append(bc.Breadcrumbs, Breadcrumb{Text: text, Link: link, Active: active})
}

func NewBreadcrumbs() Breadcrumbs {
	var bc Breadcrumbs
	return bc
}
