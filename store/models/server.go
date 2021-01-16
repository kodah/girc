package models

type Server struct {
	ID        uint `gorm:"primaryKey"`
	Address   string
	Port      string
	NetworkID uint
}
