package main

import (
	"context"
	// swagger import
	_ "icode.baidu.com/baidu/bce-soe/rule-engine/exmaples/example/doc"
	"icode.baidu.com/baidu/bce-soe/rule-engine/exmaples/example/service"
	"icode.baidu.com/baidu/themis/logger-go"
	"icode.baidu.com/baidu/themis/server-go/conf"
	"icode.baidu.com/baidu/themis/server-go/host"
	"icode.baidu.com/baidu/themis/server-go/swagger"
)

// $(SWAG) init -g server.go -d $(HOMEDIR)/cmd/server,$(SWAG_SCAN_DIRS) --instanceName=${SWAG_DOC_PREFIX}
// --parseDependency --parseInternal --generatedTime -o $(HOMEDIR)/doc/openapi
//
//	@title			rule engine swagger api specification
//	@version		1.0
//	@description	rule engine api document，restful api
//	@BasePath		/
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

	config := conf.Default()
	h := host.NewHTTPHost(config)
	swagger.DocPrefix = "RuleEngine"

	// add service
	svc := service.NewRuleEngineService(config)
	h.AddService(svc)
	h.Run()
}
