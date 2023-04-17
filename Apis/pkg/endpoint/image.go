package endpoint

import (
	"time"
)

type Register struct {
	ID        uint32    `json:"id"`
	Sha256    string    `json:"sha256"`
	Size      int32     `json:"size"`
	ChunkSize int32     `json:"chunk_size"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
