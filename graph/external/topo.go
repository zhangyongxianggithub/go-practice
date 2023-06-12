package external

type Topo struct {
	Sources []string            `json:"sources"`
	Edge    map[string][]string `json:"edges"`
}
