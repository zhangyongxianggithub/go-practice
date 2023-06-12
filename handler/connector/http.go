package connector

type HTTPProperties struct {
	URL    string            `json:"url"`
	Method string            `json:"method"`
	Header map[string]string `json:"header"`
	Body   []byte            `json:"body"`
	// 考虑还需要什么?
}

var _ Connector[HTTPProperties] = new(HTTPConnector[HTTPProperties])

type HTTPConnector[HTTPProperties any] struct {
	Name       string         `json:"name"`
	Properties HTTPProperties `json:"properties"`
}

func (m *HTTPConnector[HTTPProperties]) GetKind() Kind {
	return ConnectorHTTP
}

func (m *HTTPConnector[HTTPProperties]) GetName() string {
	return m.Name
}

func (m *HTTPConnector[HTTPProperties]) GetProperties() HTTPProperties {
	return m.Properties
}
