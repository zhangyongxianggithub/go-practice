package handler

import (
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/sink"
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/source"
	"icode.baidu.com/baidu/themis/server-go/conf"
	"icode.baidu.com/baidu/themis/server-go/db"
)

type StdRuleEngineHandler struct {
	conf     *conf.Config
	dbModule *db.Module
}

func (d *StdRuleEngineHandler) CreateRule() error {
	//TODO implement me
	panic("implement me")
}

func NewRuleEngineHandler(conf *conf.Config, dbModule *db.Module) (*StdRuleEngineHandler, error) {
	return &StdRuleEngineHandler{conf: conf, dbModule: dbModule}, nil
}

var _ RuleHandler = new(StdRuleEngineHandler)

func (d *StdRuleEngineHandler) GetSources(request GetSourcesRequest) ([]source.Source, error) {
	// TODO implement me
	panic("implement me")
}

func (d *StdRuleEngineHandler) GetSource(name string) (source.Source, error) {
	// TODO implement me
	panic("implement me")
}

func (d *StdRuleEngineHandler) GetSinks() ([]sink.Sink, error) {
	// TODO implement me
	panic("implement me")
}

func (d *StdRuleEngineHandler) GetSink(name string) (sink.Sink, error) {
	// TODO implement me
	panic("implement me")
}
