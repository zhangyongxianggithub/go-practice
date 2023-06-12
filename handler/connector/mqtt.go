package connector

var _ Connector[MQTTProperties] = new(MQTTConnector[MQTTProperties])

type MQTTProperties struct {
	Server   string `json:"server"`
	Topic    string `json:"topic"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type MQTTConnector[MQTTProperties comparable] struct {
	Name       string         `json:"name"`
	Properties MQTTProperties `json:"properties"`
}

func (m *MQTTConnector[MQTTProperties]) GetKind() Kind {
	return ConnectorMQTT
}

func (m *MQTTConnector[MQTTProperties]) GetName() string {
	return m.Name
}

func (m *MQTTConnector[MQTTProperties]) GetProperties() MQTTProperties {
	return m.Properties
}
