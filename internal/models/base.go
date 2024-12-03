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

type Comparable interface {
	Hash() string
	Equals(Comparable) bool
}

type BaseEntity interface {
	Serializable
	Deserializable
	Validatable
	Comparable
}
