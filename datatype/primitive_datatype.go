package datatype

type PrimitiveDataType interface {
	Name() string
	Description() string
}

type BasePrimitiveDataType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (b *BasePrimitiveDataType) GetDescription() string {
	return b.Description
}

func (b *BasePrimitiveDataType) GetName() string {
	return b.Name
}

type Int struct {
	BasePrimitiveDataType
}

type Float struct {
	BasePrimitiveDataType
}

type Boolean struct {
	BasePrimitiveDataType
}

type String struct {
	BasePrimitiveDataType
}
