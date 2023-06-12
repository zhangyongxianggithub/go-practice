package handler

import (
	"context"
	"fmt"

	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/sink"
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/source"
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/source/builtin"
	"icode.baidu.com/baidu/themis/server-go/conf"
	"icode.baidu.com/baidu/themis/server-go/db"
)

type LakeinsightRuleEngineHandler struct {
	conf     *conf.Config
	dbModule *db.Module
}

func (l *LakeinsightRuleEngineHandler) CreateRule(ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (l *LakeinsightRuleEngineHandler) UpdateRule(ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (l *LakeinsightRuleEngineHandler) DeleteRule(ctx context.Context, name string) error {
	// TODO implement me
	panic("implement me")
}

func (l *LakeinsightRuleEngineHandler) GetRule(ctx context.Context, name string) error {
	// TODO implement me
	panic("implement me")
}

func NewRuleEngineHandler(
	conf *conf.Config,
	dbModule *db.Module,
) (*LakeinsightRuleEngineHandler, error) {
	return &LakeinsightRuleEngineHandler{conf: conf, dbModule: dbModule}, nil
}

var _ RuleHandler = new(LakeinsightRuleEngineHandler)

func (l *LakeinsightRuleEngineHandler) GetSources(
	ctx context.Context,
	param any,
) ([]source.Source, error) {
	return []source.Source{builtin.NewRawDataSource(), builtin.NewBIESource()}, nil
}

func (l *LakeinsightRuleEngineHandler) GetSource(
	ctx context.Context,
	name string,
) (source.Source, error) {
	sources, _ := l.GetSources(ctx, struct{}{})
	for _, s := range sources {
		if s.GetName() == name {
			return s, nil
		}
	}
	return nil, fmt.Errorf("source: %s not found", name)
}

func (l *LakeinsightRuleEngineHandler) GetSinks(ctx context.Context) ([]sink.Sink, error) {
	// TODO implement me
	panic("implement me")
}

func (l *LakeinsightRuleEngineHandler) GetSink(
	ctx context.Context,
	name string,
) (sink.Sink, error) {
	// TODO implement me
	panic("implement me")
}
