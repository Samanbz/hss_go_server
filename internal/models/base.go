package models

type Serializable interface {
	ToJSON() []byte
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

type BaseModel interface {
	Serializable
	Validatable
	Comparable
}
