package entities

import "time"

type Users struct {
	Name      string    `json:"name,omitempty" gorm:"type:varchar(50)"`
	Email     string    `json:"email,omitempty" gorm:"type:varchar(100);unique_index"`
	Telephone string    `json:"telephone,omitempty" gorm:"type:varchar(20)"`
	ID        string    `json:"id,omitempty" gorm:"primaryKey"`
	Created   time.Time `json:"created,omitempty" gorm:"autoCreateTime"`
}
