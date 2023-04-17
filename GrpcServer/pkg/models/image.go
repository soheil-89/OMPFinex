package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Sha256    string `gorm:"size:256"`
	ChunkSize int32  `gorm:"not null"`
	Data      string `gorm:"not null"`
	Number    int32  `gorm:"not null"`
}
