package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Sha256    string `gorm:"size:256;unique"`
	Size      int32  `gorm:"not null"`
	ChunkSize int32
}
