package gorm

import "time"

type Model struct {
	ID        uint `gorm:"primary_key"`
	Created_At time.Time
	Updated_At time.Time
	Deleted_At *time.Time `sql:"index"`
}
