package controller

import (
	"context"
	"icode.baidu.com/baidu/themis/server-go"
	"icode.baidu.com/baidu/themis/server-go/pkg/binder"

	"icode.baidu.com/baidu/bce-soe/rule-engine/handler"
	"icode.baidu.com/baidu/themis/logger-go"
	servergo "icode.baidu.com/baidu/themis/server-go/pkg"
)

type SOERuleController struct {
	handler handler.RuleHandler
}

func (c *SOERuleController) RegisterRoute(engineGroup *servergo.RouterGroup) {
	logger.Noticef(context.Background(), "register rule engine http endpoints")
	engineGroup.GET("/rule-engine/scalar-functions", c.GetScalarFunctions)
	engineGroup.GET("/rule-engine/agg-functions", c.GetAggFunctions)
	engineGroup.GET("/rule-engine/sources", c.GetSources)
	engineGroup.GET("/rule-engine/sources/{sourceName}", c.GetSource)
	engineGroup.POST("/rule-engine/rules", c.CreateRule)
}

// @Summary		普通函数
// @Description	普通函数
// @Tags			function
// @Accept			json
// @Produce		json
// @Param			workspaceID	path		string	true	"workspaceId"
// @Param			projectName	path		string	true	"projectName"
// @Success		200			{object}	string
// @Failure		500			{string}	string	"internal server error"
// @Router			/v1/workspaces/{workspaceID}/projects/{projectName}/analyticsjobs [post]
func (c *SOERuleController) GetScalarFunctions(context *servergo.Context) {

	// TODO: implement logic
}

func (c *SOERuleController) GetAggFunctions(context *servergo.Context) {
	// TODO: implement logic
}

// GetSources
//
// @Summary		自定义Source列表
// @Description	自定义Source列表
// @Tags		source
// @Accept		json
// @Produce		json
// @Param		listRequest	query		handler.GetSourcesRequest	true	"数据源信息"
// @Success		200			{object}	SourceListElement
// @Failure		500			{string}	string	"internal server error"
// @Router		/soe/rule-engine/sources [get]
func (c *SOERuleController) GetSources(context *servergo.Context) {
	var request handler.GetSourcesRequest
	if err := binder.Bind(context, &request).All().Error(); err != nil {
		_ = context.Error(err)
	} else if sources, err := c.handler.GetSources(request); err != nil {
		_ = context.Error(err)
	} else {
		sourceList := make([]*SourceListElement, 0)
		for _, source := range sources {
			sourceList = append(sourceList, &SourceListElement{
				ID:   source.GetID(),
				Name: source.GetName(),
			})
		}
		server.OK(context, sourceList)
	}
}

type SourceListElement struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *SOERuleController) GetSource(context *servergo.Context) {
	// TODO: implement logic
}

func (c *SOERuleController) CreateRule(context *servergo.Context) {
	// TODO: implement logic
}

func NewRuleController(handler handler.RuleHandler) (*SOERuleController, error) {
	return &SOERuleController{handler: handler}, nil
}
