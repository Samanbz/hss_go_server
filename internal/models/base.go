package models

type Serializable interface {
	ToJSON() []byte
}

type Deserializable interface {
	FromJSON([]byte) error
}

type Validatable interface {
	ValidateInput() error
	ValidateOutput() error
}

type BaseEntity interface {
	Serializable
	Deserializable
	Validatable
}
