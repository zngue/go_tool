package db

import time2 "time"

type TimeStampModel struct {
	CreatedAt int64 `gorm:"column:created_at" json:"created_at" `
	UpdatedAt int64 `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time2.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}
type GormModel struct {
	CreatedAt time2.Time
	UpdatedAt time2.Time
	DeletedAt *time2.Time `sql:"index"`
}
