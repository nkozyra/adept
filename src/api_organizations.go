package adept

import (
	"fmt"
)

type Organizations struct {
	Organizations []Organization

	GUID string
}

type Organization struct {
	ID   int64
	Name string
	GUID string
}

func NewOrganizations() Organizations {
	var o Organizations
	return o
}

func (os *Organizations) Get() {
	W, V := NewParams()

	if os.GUID != "" {
		W.params = append(W.params, "organization_guid=?")
		V.Push(os.GUID)
	}

	WhereString := W.Compile()
	rows, _ := DB.Query("SELECT o.organization_id, o.organization_name, o.organization_guid FROM organizations o "+WhereString, V.params...)
	defer rows.Close()
	for rows.Next() {
		var o Organization
		err := rows.Scan(&o.ID, &o.Name, &o.GUID)
		fmt.Println(o)
		if err != nil {

		}
		os.Organizations = append(os.Organizations, o)
	}
}
