package db

import (
	"time"
)

type SnapmateModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
}

type Snapshot struct {
	SnapmateModel
	Name    string
	Comment string
}
