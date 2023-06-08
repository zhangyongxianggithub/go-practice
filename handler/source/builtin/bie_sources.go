package builtin

import (
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/source"
)

var _ source.Source = new(bieSource)

type bieSource struct {
}

func (B *bieSource) GetID() string {
	return "BIE"
}

func (B *bieSource) GetName() string {
	return "BIE"
}

func (B *bieSource) GetDisplayName() string {
	return "BIE"
}

func (B *bieSource) GetSourceFields() []*source.Field {
	//TODO implement me
	panic("implement me")
}

func (B *bieSource) GetSourceField(ID string) (*source.Field, error) {
	//TODO implement me
	panic("implement me")
}

func NewBIESource() source.Source {
	return &bieSource{}
}
