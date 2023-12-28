package models

import (
	"encoding/json"
	"strconv"
	"time"

	"gorm.io/gorm"
	"mxzero.top/utils"
)

type Model struct {
	ID        int            `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (m User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		*Alias
		IDFormatted        string `json:"id"`
		CreatedAtFormatted string `json:"created_at"`
		UpdatedAtFormatted string `json:"updated_at"`
	}{
		Alias:              (*Alias)(&m),
		IDFormatted:        strconv.Itoa(m.ID),
		CreatedAtFormatted: utils.FormartDateTime(m.CreatedAt),
		UpdatedAtFormatted: utils.FormartDateTime(m.UpdatedAt),
	})
}
