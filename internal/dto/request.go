package dto

type CreateLampiranRequest struct {
	FilePath string `json:"file_path" binding:"required"`
}
