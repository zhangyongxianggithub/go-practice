package connector

type Kind string

const ConnectorMQTT Kind = "mqtt"

const ConnectorHTTP Kind = "http"

type Connector[T any] interface {
	GetKind() Kind

	GetName() string

	GetProperties() T
}
