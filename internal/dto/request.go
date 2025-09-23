package dto

type CreateLampiranRequest struct {
	FilePath string `json:"file_path" binding:"required"`
}

// type PostDokumenLampiranByNRPRequest struct {
// 	Message string `json:"message"`
// }
