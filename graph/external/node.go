package external

type Style struct {
	// style of node
}
type Node[T any] struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Style      Style  `json:"style"`
	Type       string `json:"type"`
	Attributes T      `json:"attributes"`
}
