package node

import (
	"github.com/google/uuid"
)

type SourceNode[T any] struct {
	ID         string `json:"id"`         // node ID
	NodeName   string `json:"name"`       // node name
	SourceName string `json:"sourceName"` // source
	Kind       Kind   `json:"kind"`       // source type
	Properties T      `json:"properties"`
}

func (s *SourceNode[T]) GetID() string {
	return s.GetName()
}

func (s *SourceNode[T]) GetName() string {
	return s.GetName()
}

func (s *SourceNode[T]) GetType() Kind {
	return SourceNodeKind
}

var _ GraphNode = new(SourceNode[any])

func NewSourceNode[T any](nodeName, sourceName string, properties T) *SourceNode[T] {
	return &SourceNode[T]{
		ID:         uuid.New().String(),
		NodeName:   nodeName,
		SourceName: sourceName,
		Kind:       SourceNodeKind,
		Properties: properties,
	}
}
