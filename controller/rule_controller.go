package controller

import (
	"context"

	"icode.baidu.com/baidu/bce-soe/rule-engine/graph"
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/source"
	"icode.baidu.com/baidu/bce-soe/rule-engine/tools/printer"
	"icode.baidu.com/baidu/themis/server-go"
	"icode.baidu.com/baidu/themis/server-go/pkg/binder"

	"icode.baidu.com/baidu/bce-soe/rule-engine/handler"
	"icode.baidu.com/baidu/themis/logger-go"
	servergo "icode.baidu.com/baidu/themis/server-go/pkg"
)

type SOERuleController struct {
	handler handler.RuleHandler
}

type GetSourceRequest struct {
	SourceName string `uri:"sourceName"`
}

type GetSourceResponse struct {
	Name         string          `json:"sourceName"`
	ID           string          `json:"id"`
	DisplayName  string          `json:"displayName"`
	Description  string          `json:"description"`
	Kind         source.Kind     `json:"kind"`
	SourceFields []*source.Field `json:"sourceFields"`
}

func (c *SOERuleController) RegisterRoute(engineGroup *servergo.RouterGroup) {
	logger.Noticef(context.Background(), "register rule engine http endpoints")
	engineGroup.GET("/rule-engine/scalar-functions", c.GetScalarFunctions)
	engineGroup.GET("/rule-engine/agg-functions", c.GetAggFunctions)
	engineGroup.GET("/rule-engine/sources", c.GetSources)
	engineGroup.GET("/rule-engine/sources/:sourceName", c.GetSource)
	engineGroup.POST("/rule-engine/rules", c.CreateRule)
	engineGroup.PUT("/rule-engine/rules/:ruleName", nil)
	engineGroup.DELETE("/rule-engine/rules/:ruleName", nil)
	engineGroup.GET("/rule-engine/rules/:ruleName", nil)
}

// GetScalarFunctions
//
//	@Summary		普通函数
//	@Description	普通函数
//	@Tags			function
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string
//	@Failure		500	{string}	string	"internal server error"
//	@Router			/rule-engine/scalar-functions [get]
func (c *SOERuleController) GetScalarFunctions(ctx *servergo.Context) {
	logger.Noticef(ctx, "get scalar functions")
	// TODO: implement logic
}

func (c *SOERuleController) GetAggFunctions(ctx *servergo.Context) {
	// TODO: implement logic
	logger.Noticef(ctx, "get aggregation functions")
}

type GetSourcesRequest struct {
	Filter string `json:"filter"`
}

// GetSources
//
//	@Summary		自定义Source列表
//	@Description	自定义Source列表
//	@Tags			source
//	@Accept			json
//	@Produce		json
//	@Param			listRequest	query		handler.GetSourcesRequest	true	"数据源信息"
//	@Success		200			{object}	SourceListElement
//	@Failure		500			{string}	string	"internal server error"
//	@Router			/rule-engine/sources [get]
func (c *SOERuleController) GetSources(ctx *servergo.Context) {
	var request GetSourcesRequest
	if err := binder.Bind(ctx, &request).All().Error(); err != nil {
		_ = ctx.Error(err)
	} else if sources, err := c.handler.GetSources(ctx, request); err != nil {
		_ = ctx.Error(err)
	} else {
		sourceList := make([]*SourceListElement, 0)
		for _, s := range sources {
			sourceList = append(sourceList, &SourceListElement{
				ID:          s.GetID(),
				Name:        s.GetName(),
				Description: s.GetDescription(),
			})
		}
		server.OK(ctx, sourceList)
	}
}

type SourceListElement struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetSource
//
//	@Summary		获取Source详情
//	@Description	获取Source详情
//	@Tags			source
//	@Accept			json
//	@Produce		json
//	@Param			request	query		GetSourceRequest	true	"数据源名字"
//	@Success		200		{object}	source.Source
//	@Failure		500		{string}	string	"internal server error"
//	@Router			/rule-engine/sources [get]
func (c *SOERuleController) GetSource(ctx *servergo.Context) {
	var request GetSourceRequest
	if err := binder.Bind(ctx, &request).All().Error(); err != nil {
		_ = ctx.Error(err)
	} else if s, err := c.handler.GetSource(ctx, request.SourceName); err != nil {
		_ = ctx.Error(err)
	} else {
		server.OK(ctx, GetSourceResponse{
			Name:         s.GetName(),
			ID:           s.GetID(),
			DisplayName:  s.GetDisplayName(),
			Description:  s.GetDescription(),
			Kind:         s.GetKind(),
			SourceFields: s.GetSourceFields(),
		})
	}
}

// CreateRule
//
//	@Summary		创建GraphRule
//	@Description	创建GraphRule
//	@Tags			rule
//	@Accept			json
//	@Produce		json
//	@Param			request	query		graph.Rule	true	"graph rule body"
//	@Success		200		{object}	string		"success"
//	@Failure		500		{string}	string		"internal server error"
//	@Router			/rule-engine/rules [post]
func (c *SOERuleController) CreateRule(ctx *servergo.Context) {
	var graphRule graph.Rule
	if err := binder.Bind(ctx, &graphRule).All().Error(); err != nil {
		_ = ctx.Error(err)
	}
	logger.Noticef(ctx, "submit graph rule: %s", printer.PrintAny(graphRule))
	rule, err := graph.NewInternalRule(&graphRule, c.handler, nil)
	if err != nil {
		_ = ctx.Error(err)
	} else {
		err := rule.Create(ctx)
		if err != nil {
			_ = ctx.Error(err)
		}
	}
	server.OK(ctx, graphRule)
}

func NewRuleController(handler handler.RuleHandler) (*SOERuleController, error) {
	return &SOERuleController{handler: handler}, nil
}
