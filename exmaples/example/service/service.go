package service

import (
	"context"
	"fmt"

	"icode.baidu.com/baidu/bce-soe/rule-engine/controller"
	"icode.baidu.com/baidu/themis/logger-go"
	"icode.baidu.com/baidu/themis/server-go/conf"
	"icode.baidu.com/baidu/themis/server-go/host"
	servergo "icode.baidu.com/baidu/themis/server-go/pkg"
)

type RuleEngineService struct {
	host.Service
	Engine *servergo.Engine
	conf   *conf.Config
	prefix string
}

var _ host.Service = &RuleEngineService{}

func NewRuleEngineService(config *conf.Config) *RuleEngineService {
	logger.Noticef(context.Background(), "create a rule engine service, config: %v", config)
	return &RuleEngineService{}
}

func (s *RuleEngineService) Initialize(e *servergo.Engine) {
	s.Engine = e
	// specify local rule engine api prefix
	s.prefix = fmt.Sprintf("%s", s.conf.AllSettings["rule-engine"])
}

func (s *RuleEngineService) StartAsync() {
	// must add a prefix
	group := s.Engine.Group("/soe")
	controller.RegisterRoute(group, nil)
}

func (s *RuleEngineService) StopAsync() {
	logger.Noticef(context.Background(), "stop serve rule engine request")
}
