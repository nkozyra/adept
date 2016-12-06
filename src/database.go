package adept

import (
	"strings"
)

func getConnectString() string {
	cs := Conf.DB.User + ":" + Conf.DB.Password + "@tcp(" + Conf.DB.Host + ")/" + "adept" + "?charset=utf8mb4"
	return cs
}

func connect() {

}

func NewParams() (Wheres, Vals) {
	var w Wheres
	var v Vals

	return w, v
}

func ParamRoute(W *Wheres, V *Vals, params ...[]map[string]string) {
	for i := range params {
		for j := range params[i] {
			for k, v := range params[i][j] {
				W.Push(k + "=?")
				V.Push(v)
			}
		}
	}
}

type Wheres struct {
	params []string
}

func (w *Wheres) Push(s string) {
	w.params = append(w.params, s)
}

func (w *Wheres) Compile() string {
	var output string
	if len(w.params) > 0 {
		output = " WHERE "
		output += strings.Join(w.params, ", ")
	}
	return output
}

type Vals struct {
	params []interface{}
}

func (v *Vals) Push(s string) {
	v.params = append(v.params, s)
}
