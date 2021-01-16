package models

type Network struct {
	ID          uint `gorm:"primaryKey"`
	NetworkName string
	Servers     []Server `gorm:"foreignKey:NetworkID"`
	Nick        string
	Alt         string
	Password    string
	RealName    string
}
