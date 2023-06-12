package service

import (
	"context"
	"icode.baidu.com/baidu/bce-soe/rule-engine/controller"
	"icode.baidu.com/baidu/themis/logger-go"
	"icode.baidu.com/baidu/themis/server-go/host"
	servergo "icode.baidu.com/baidu/themis/server-go/pkg"
)

type RuleEngineService struct {
	host.Service
	Engine     *servergo.Engine
	Controller *controller.SOERuleController
	Prefix     string
}

var _ host.Service = &RuleEngineService{}

func NewRuleEngineService(
	controller *controller.SOERuleController,
	prefix string,
) *RuleEngineService {
	logger.Noticef(context.Background(), "create a rule engine service, config: %v", controller)
	return &RuleEngineService{Controller: controller, Prefix: prefix}
}

func (s *RuleEngineService) Initialize(e *servergo.Engine) {
	s.Engine = e
}

func (s *RuleEngineService) StartAsync() {
	s.Controller.RegisterRoute(s.Engine.Group(s.Prefix))
}

func (s *RuleEngineService) StopAsync() {
	logger.Noticef(context.Background(), "stop serve rule engine request")
}
