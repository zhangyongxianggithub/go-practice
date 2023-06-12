package builtin

import (
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/connector"
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/source"
)

var _ source.Source = new(bieSource)

type bieSource struct {
}

func (b *bieSource) GetKind() source.Kind {
	return source.MQTT
}

func (b *bieSource) GetConnector() (connector.Connector[any], error) {
	//TODO implement me
	panic("implement me")
}

func (b *bieSource) GetDescription() string {
	return "标准的BIE数据源"
}

func (b *bieSource) GetID() string {
	return "BIE"
}

func (b *bieSource) GetName() string {
	return "BIE"
}

func (b *bieSource) GetDisplayName() string {
	return "BIE"
}

func (b *bieSource) GetSourceFields() []*source.Field {
	//TODO implement me, 设置字段数据
	return []*source.Field{}
}

func (b *bieSource) GetSourceField(ID string) (*source.Field, error) {
	//TODO implement me
	panic("implement me")
}

func NewBIESource() source.Source {
	return &bieSource{}
}
