package graph

import "icode.baidu.com/baidu/bce-soe/rule-engine/graph/external"

// Rule request body content type
type Rule struct {
	Name        string                            `json:"name"        binding:"max=200"`  // graph rule name Required
	DisplayName string                            `json:"displayName" binding:"max=200"`  // chinese name Required
	Description string                            `json:"description" binding:"max=1024"` // description Required
	Sources     map[string]external.Source        `json:"sources"`                        // source定义, 前端传递只需要传递source的ID or name
	Operators   map[string]external.Operator[any] `json:"operators"`                      // 所有的中间计算节点的operator定义
	Sink        map[string]external.Sink[any]     `json:"sink"`                           // Sink data output
	Topo        external.Topo                     `json:"topo"`                           // topo
}
