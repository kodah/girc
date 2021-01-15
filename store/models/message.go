package models

import (
	"time"
)

type Message struct {
	ID      uint `gorm:"primaryKey"`
	Time    time.Time
	Name    string
	User    string
	Host    string
	Command string
}
