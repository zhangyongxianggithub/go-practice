package handler

type RuleHandler interface {
	GetSources() ([]*Source, error)
	GetSource(name string) (*Source, error)
	GetSinks() ([]*Sink, error)
	GetSink(name string) (*Sink, error)
}

var _ RuleHandler = new(DefaultRuleHandler)

type DefaultRuleHandler struct {
}

func (d *DefaultRuleHandler) GetSources() ([]*Source, error) {
	// TODO implement me
	panic("implement me")
}

func (d *DefaultRuleHandler) GetSource(name string) (*Source, error) {
	// TODO implement me
	panic("implement me")
}

func (d *DefaultRuleHandler) GetSinks() ([]*Sink, error) {
	// TODO implement me
	panic("implement me")
}

func (d *DefaultRuleHandler) GetSink(name string) (*Sink, error) {
	// TODO implement me
	panic("implement me")
}
