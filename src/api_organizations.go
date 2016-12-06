package adept

import ()

type Organizations struct {
	Organizations []Organization
}

type Organization struct {
}

func NewOrganizations() Organizations {
	var o Organizations
	return o
}
