package models

import (
	"encoding/json"
	"hss/internal/utils"
	"hss/internal/utils/validation"
)

type Customer struct {
	ID        int    `json:"id"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required,sha256"`
	CompanyID int    `json:"user_id" validate:"required"`
}

func (c Customer) ToJSON() []byte {
	jsonData, _ := json.Marshal(c)
	return jsonData
}

func (c *Customer) FromJSON(jsonData []byte) error {
	return json.Unmarshal(jsonData, c)
}

func (c *Customer) ValidateInput() error {
	return validation.GetValidator().StructExcept(c, "ID")
}

func (c *Customer) ValidateOutput() error {
	return validation.GetValidator().Struct(c)
}

func (c Customer) Hash() string {
	return utils.Hash(string(c.ToJSON()))
}

func (c Customer) Equals(other Customer) bool {
	return c.Hash() == other.Hash()
}

func (c Customer) WithForeignKey(foreignKey int) *Customer {
	c.CompanyID = foreignKey
	return &c
}
