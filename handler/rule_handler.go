package handler

import (
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/sink"
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/source"
)

type GetSourcesRequest struct {
}

type RuleHandler interface {
	GetSources(request GetSourcesRequest) ([]source.Source, error)
	GetSource(name string) (source.Source, error)
	GetSinks() ([]sink.Sink, error)
	GetSink(name string) (sink.Sink, error)
	CreateRule() error
}
