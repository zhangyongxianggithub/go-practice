package source

import "icode.baidu.com/baidu/bce-soe/rule-engine/datatype"

type Field struct {
	ID          string                     `json:"id"`
	DisplayName string                     `json:"display"`
	Name        string                     `json:"name"`
	DataType    datatype.PrimitiveDataType `json:"dataType"`
	RuleField   string                     `json:"ruleField"`
}
type Source interface {
	GetID() string
	GetName() string
	GetDisplayName() string
	GetSourceFields() []*Field
	GetSourceField(ID string) (*Field, error)
}
