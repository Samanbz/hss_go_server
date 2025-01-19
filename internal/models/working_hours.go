package models

import (
	"encoding/json"
	"hss/internal/utils"
	"hss/internal/utils/validation"
	"time"
)

type DayOfWeek int

type WorkingHours struct {
	ID        int       `json:"id" validate:"required"`
	Day       DayOfWeek `json:"day" validate:"required,min=0,max=6"`
	OpenTime  time.Time `json:"open_time" validate:"required"`
	CloseTime time.Time `json:"close_time" validate:"required,gtfield=Open"`
}

func (w WorkingHours) ToJSON() []byte {
	jsonData, _ := json.Marshal(w)
	return jsonData
}

func (w *WorkingHours) FromJSON(jsonData []byte) error {
	return json.Unmarshal(jsonData, w)
}

func (w *WorkingHours) ValidateInput() error {
	return validation.GetValidator().StructExcept(w, "ID")
}

func (w *WorkingHours) ValidateOutput() error {
	return validation.GetValidator().Struct(w)
}

func (wh WorkingHours) Hash() string {
	return utils.Hash(string(wh.ToJSON()))
}

func (wh WorkingHours) Equals(other WorkingHours) bool {
	return wh.Hash() == other.Hash()
}
