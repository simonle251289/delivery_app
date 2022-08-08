package common

import "time"

type BaseModel struct {
	Id          int       `json:"id" gorm:"column:id;"`
	CreatedDate time.Time `json:"created_at" gorm:"column:created_at;"`
}
