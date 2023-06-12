package node

type Kind string

const SourceNodeKind Kind = "source"

type GraphNode interface {
	GetID() string
	GetName() string
	GetType() Kind
}
