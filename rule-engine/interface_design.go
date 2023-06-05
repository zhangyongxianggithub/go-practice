package rule_engine

type Source interface {
	GetFields() ([]string, error)

	GetCanonicalField(fieldName string) (string, error)

	GetProperties() (map[string]string, error)
}

type GraphRule string

type DAG string

type GraphAnalyzer interface {
	ConvertGraphRule(dag DAG) (GraphRule, error)
}
type DAGNode string

type GraphNodeRule string

type RuleNodeAnalyzer interface {
	ConvertGraphNodeRule(node DAGNode) (GraphNodeRule, error)
}

type Sink interface {
	ConvertGraphNodeRule(node DAGNode) (GraphNodeRule, error)

	SilentSetting(silentSetting SilentSetting)
}

type SilentSetting interface {
	// TODO: 定义支持的几种静默条件设置
}
type Field struct {
	Code        string `编码`
	DisplayName string `显示的名字`
	Type        string `数据类型`
}

type Type string

type Engine interface {
	RegisterSource(source Source) error
	RemoveSource(source Source) error
	RegisterSink(sink Sink) error
	RemoveSink(sink Sink) error
	GetGraphAnalyzer(typo Type) GraphAnalyzer
	GetRuleNodeAnalyzer(typo Type, nodeType string) RuleNodeAnalyzer
	Submit(rule GraphRule) error
}
