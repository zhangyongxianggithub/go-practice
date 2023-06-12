package graph

import (
	"context"

	"icode.baidu.com/baidu/bce-soe/rule-engine/graph/node"
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler"
	"icode.baidu.com/baidu/bce-soe/rule-engine/handler/source"
	"icode.baidu.com/baidu/themis/logger-go"
)

type InternalRule struct {

	// data handler
	handler           handler.RuleHandler
	graphRule         *Rule
	backendRuleEngine BackendRuleEngine

	// data def
	Name        string `json:"name"        binding:"max=200"`  // graph rule name Required
	DisplayName string `json:"displayName" binding:"max=200"`  // chinese name Required
	Description string `json:"description" binding:"max=1024"` // description Required

	// source def
	Sources map[string]*node.SourceNode[any]
}

func NewInternalRule(graphRule *Rule, handler handler.RuleHandler,
	backendRuleEngine BackendRuleEngine) (*InternalRule, error) {
	rule := &InternalRule{
		handler:           handler,
		graphRule:         graphRule,
		backendRuleEngine: backendRuleEngine,
	}
	rule.Name = rule.graphRule.Name
	rule.DisplayName = rule.graphRule.DisplayName
	rule.Description = rule.graphRule.Description
	return rule, nil
}

func (g *InternalRule) Create(ctx context.Context) error {
	go func() {
		_, err := g.CreateSourceNode(ctx)
		if err != nil {
			logger.Errorf(ctx, "error creating source node, reason: %s", err)
		}
	}()
	return nil
}

func (g *InternalRule) CreateSourceNode(ctx context.Context) ([]*node.SourceNode[any],
	error) {
	graphRule := g.graphRule
	g.Sources = make(map[string]*node.SourceNode[any])
	sourceNodes := make([]*node.SourceNode[any], 0)
	if len(graphRule.Sources) > 0 {
		for nodeName, s := range graphRule.Sources {
			// get source
			curSource, err := g.handler.GetSource(ctx, s.SourceName)
			if err != nil {
				logger.Errorf(
					ctx,
					"error getting source, source name: %s, reason: %s",
					s.NodeName,
					err,
				)
				return sourceNodes, err
			}
			sourceNode, _ := ConvertSourceToSourceNode(nodeName, curSource)
			g.Sources[nodeName] = sourceNode
			sourceNodes = append(sourceNodes, sourceNode)
		}
	}
	return sourceNodes, nil
}

func ConvertSourceToSourceNode(
	nodeName string,
	source source.Source,
) (*node.SourceNode[any], error) {
	connector, err := source.GetConnector()
	if err != nil {
		return nil, err
	}
	return node.NewSourceNode[any](nodeName, source.GetName(), connector.GetProperties()), nil
}
