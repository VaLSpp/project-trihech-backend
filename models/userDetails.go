package models

import "gorm.io/gorm"

type UserDetails struct {
	gorm.Model
	UserID       uint   `gorm:"primaryKey"`
	NamaLengkap  string `gorm:"size:100"`
	JenisKelamin string `gorm:"type:enum('L','P')"`
	Bio          string
	User         User `gorm:"foreignKey:UserID"`
}
