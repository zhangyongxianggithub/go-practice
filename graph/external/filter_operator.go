package external

type FilterAttributes struct {
	Expr Expression `json:"expr,omitempty"`
}
type Filter struct {
	Operator[FilterAttributes]
}

type Expression struct {
	Field string        `json:"field,omitempty"`
	Op    string        `json:"op,omitempty"`
	Value string        `json:"value,omitempty"`
	Not   *Expression   `json:"not,omitempty"` // use in conjunction with single field expression or single value expression
	And   []*Expression `json:"and,omitempty"`
	Or    []*Expression `json:"or,omitempty"`
}

func (e *Expression) IsAnd() bool {
	return len(e.And) > 0 && len(e.Or) <= 0 && e.Not == nil && e.Value == "" && e.Op == "" &&
		e.Field == ""
}

func (e *Expression) IsOr() bool {
	return len(e.Or) > 0 && len(e.And) <= 0 && e.Not == nil && e.Value == "" && e.Op == "" &&
		e.Field == ""
}

func (e *Expression) IsNot() bool {
	return e.Not != nil && len(e.Or) <= 0 && len(e.And) <= 0 && e.Value == "" && e.Op == "" &&
		e.Field == ""
}

// IsField field value can be converted to a boolean value
func (e *Expression) IsField() bool {
	return !e.IsAnd() && !e.IsNot() && !e.IsOr() && e.Value == "" && e.Op == "" && e.Field != ""
}

// IsValue value belongs a type that can be converted to a boolean
func (e *Expression) IsValue() bool {
	return !e.IsAnd() && !e.IsNot() && !e.IsOr() && e.Value != "" && e.Op == "" && e.Field == ""
}

func (e *Expression) IsBinaryExpr() bool {
	return !e.IsAnd() && !e.IsNot() && !e.IsOr() && e.Value != "" && e.Op != "" && e.Field != ""
}

func (e *Expression) IsLegal() bool {
	legal := e.IsBinaryExpr() || e.IsField() || e.IsValue()
	if !legal {
		if e.IsAnd() {
			legal = true
			for _, expr := range e.And {
				legal = legal && expr.IsLegal()
			}
		}
		if e.IsOr() {
			legal = true
			for _, expr := range e.Or {
				legal = legal && expr.IsLegal()
			}
		}
		if e.IsNot() {
			legal = true
			legal = legal && e.Not.IsLegal()
		}
	}
	return legal
}
