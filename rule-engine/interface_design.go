package rule_engine

type Source interface {
	GetFields() ([]string, error)

	GetCanonicalField(fieldName string) (string, error)

	GetProperties() (map[string]string, error)
}
type GraphRule string

type DAG string

type GraphAnalyzer interface {
	convertGraphRule(dag DAG) (GraphRule, error)
}
