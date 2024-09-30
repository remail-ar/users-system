package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint
	Username     string `gorm:"type:varchar(100);unique_index"`
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:milli"`
}
