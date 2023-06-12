package handler

import (
	"context"

	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/sink"
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/source"
)

type RuleHandler interface {
	GetSources(ctx context.Context, param any) ([]source.Source, error)
	GetSource(ctx context.Context, name string) (source.Source, error)
	GetSinks(ctx context.Context) ([]sink.Sink, error)
	GetSink(ctx context.Context, name string) (sink.Sink, error)
	CreateRule(ctx context.Context) error
	UpdateRule(ctx context.Context) error
	DeleteRule(ctx context.Context, name string) error
	GetRule(ctx context.Context, name string) error
}
