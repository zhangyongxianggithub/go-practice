package builtin

import (
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/connector"
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/source"
)

var _ source.Source = new(rawDataSource)

type rawDataSource struct {
}

func (r *rawDataSource) GetKind() source.Kind {
	return source.MQTT
}

func (r *rawDataSource) GetConnector() (connector.Connector[any], error) {
	//TODO implement me
	panic("implement me")
}

func (r *rawDataSource) GetDescription() string {
	return "标准的raw data数据源"
}

func (r *rawDataSource) GetID() string {
	return "raw_data"
}

func (r *rawDataSource) GetName() string {
	return "raw_data"
}

func (r *rawDataSource) GetDisplayName() string {
	return "raw_data"
}

func (r *rawDataSource) GetSourceFields() []*source.Field {
	//TODO implement me
	panic("implement me")
}

func (r *rawDataSource) GetSourceField(ID string) (*source.Field, error) {
	//TODO implement me
	panic("implement me")
}

func NewRawDataSource() source.Source {
	return &rawDataSource{}
}
