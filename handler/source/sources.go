package source

import (
	"icode.baidu.com/baidu/bce-soe/rule-engine/datatype"
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/connector"
)

type Kind string

const MQTT Kind = "mqtt"

const HTTP Kind = "http"

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
	GetDescription() string
	GetKind() Kind
	GetSourceFields() []*Field
	GetSourceField(ID string) (*Field, error)
	GetConnector() (connector.Connector[any], error)
}
