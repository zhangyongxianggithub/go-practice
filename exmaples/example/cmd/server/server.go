package main

import (
	"context"

	"icode.baidu.com/baidu/bce-soe/rule-engine/controller"
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler"
	"icode.baidu.com/baidu/themis/server-go"
	"icode.baidu.com/baidu/themis/server-go/db"

	"icode.baidu.com/baidu/bce-soe/rule-engine/examples/example/service"
	"icode.baidu.com/baidu/themis/logger-go"
	"icode.baidu.com/baidu/themis/server-go/conf"
	"icode.baidu.com/baidu/themis/server-go/host"
	"icode.baidu.com/baidu/themis/server-go/swagger"
	// swagger import
)

// $(SWAG) init -g server.go -d $(HOMEDIR)/cmd/server,$(SWAG_SCAN_DIRS)
// --instanceName=${SWAG_DOC_PREFIX}
// --parseDependency --parseInternal --generatedTime -o $(HOMEDIR)/doc/openapi
//
//	@title			rule engine swagger api specification
//	@version		1.0
//	@description	rule engine api document，restful api
//	@BasePath		/soe
//	@contact.name	zhangyongxiang
//	@contact.email	zhangyongxiang@baidu.com
//	@Accept			json,xml,plain,html,x-www-form-urlencoded
func main() {
	// 服务停止前将日志落盘
	defer func() {
		err := logger.TryCloseDefaultLogger()
		if err != nil {
			logger.Errorf(context.Background(), "failed close logger: %s", err)
		}
	}()
	module := &db.Module{}
	bootstrap := server.NewBootstrap(module)
	err := bootstrap.Initialize()
	if err != nil {
		logger.Fatalf(context.Background(), "init bootstrap error: %v", err)
	}
	config := conf.Default()
	h := host.NewHTTPHost(config)
	swagger.DocPrefix = "RuleEngine"
	ruleHandler, _ := handler.NewRuleEngineHandler(config, module)

	ruleController, _ := controller.NewRuleController(ruleHandler)
	// add service & set a rule engine prefix
	svc := service.NewRuleEngineService(ruleController, "/soe")
	h.AddService(svc)
	h.Run()
}
