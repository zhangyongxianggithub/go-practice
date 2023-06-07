package controller

import (
	"context"

	"icode.baidu.com/baidu/bce-soe/rule-engine/handler"
	"icode.baidu.com/baidu/themis/logger-go"
	servergo "icode.baidu.com/baidu/themis/server-go/pkg"
)

func RegisterRoute(engineGroup *servergo.RouterGroup, ruleHandler handler.RuleHandler) {
	logger.Noticef(context.Background(), "register rule engine http endpoints")
	engineGroup.GET("/rule-engine/scalar-functions", GetScalarFunctions)
	engineGroup.GET("/rule-engine/agg-functions", GetAggFunctions)
	engineGroup.GET("/rule-engine/sources", GetSources)
	engineGroup.GET("/rule-engine/sources/{sourceName}", GetSource)
	engineGroup.POST("/rule-engine/rules", CreateRule)
}

// @Summary		创建流/批分析JOB
// @Description	创建流/批分析JOB
// @Tags			analytic
// @Accept			json
// @Produce		json
// @Param			workspaceID	path		string	true	"workspaceId"
// @Param			projectName	path		string	true	"projectName"
// @Success		200			{object}	string
// @Failure		500			{string}	string	"internal server error"
// @Router			/v1/workspaces/{workspaceID}/projects/{projectName}/analyticsjobs [post]
func GetScalarFunctions(context *servergo.Context) {

	// TODO: implement logic
}

func GetAggFunctions(context *servergo.Context) {
	// TODO: implement logic
}

func GetSources(context *servergo.Context) {
	// TODO: implement logic
}

func GetSource(context *servergo.Context) {
	// TODO: implement logic
}

func CreateRule(context *servergo.Context) {
	// TODO: implement logic
}
