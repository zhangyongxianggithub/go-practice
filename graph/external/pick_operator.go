package external

type PickAttributes struct {
	Fields []string `json:"fields"`
	Ignore bool     `json:"ignore"`
}
type Pick struct {
	Operator[PickAttributes]
}
