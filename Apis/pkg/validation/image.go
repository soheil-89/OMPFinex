package validation

type Register struct {
	Sha256    string `json:"sha256" from:"sha256" binding:"required,exist_in_db=image.sha256"`
	Size      int32  `json:"size" from:"size" binding:"required"`
	ChunkSize int32  `json:"chunk_size" from:"chunk_size" binding:"required"`
}

type Chunk struct {
	Id   int32  `json:"id" from:"id" binding:"required"`
	Size int32  `json:"size" from:"size" binding:"required"`
	Data string `json:"data" from:"data" binding:"required"`
}

type Download struct {
}
