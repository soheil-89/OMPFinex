package models

import (
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Sha256    string `gorm:"size:256"`
	Size      int    `gorm:"not null"`
	ChunkSize int
	Data      string
}

func (i *Image) handler(request interface{}) {
	_ = copier.Copy(i, request)
}
