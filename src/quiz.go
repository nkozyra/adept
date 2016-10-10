package adept

type Quiz struct {
	Questions []Question `json:"questions"`
}

func (q Quiz) Get() {

}
