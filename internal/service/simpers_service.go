package service

import "refactory-simpers-au/internal/dto"

type SimpersService interface {
	GetPersonelByNRP(nrp string) (out dto.GetPersonelByNRPResponse, err error)
	GetFamilyCardByNRP(nrp string) ([]dto.GetFamilyCardByNRPResponse, error)
	GetDikMilByNRP(nrp string) ([]dto.GetDikMilByNRPResponse, error)
	GetDikUmByNRP(nrp string) ([]dto.GetDikUmByNRPResponse, error)
	GetJabatanByNRP(nrp string) ([]dto.GetJabatanByNRPResponse, error)
	GetTanhorByNRP(nrp string) ([]dto.GetTanhorByNRPResponse, error)
	GetPangkatByNRP(nrp string) ([]dto.GetPangkatByNRPResponse, error)
	GetLampiranKKByNRP(nrp string) ([]dto.GetLampiranKKByNRPResponse, error)
	GetNPWPByNRP(nrp string) (out dto.GetNPWPByNRPResponse, err error)
	GetAsabriByNRP(nrp string) (out dto.GetAsabriByNRPResponse, err error)
	GetPasporByNRP(nrp string) (out dto.GetPasporByNRPResponse, err error)
	CreateLampiran(req dto.CreateLampiranRequest, docID string, docType string) (out dto.CreateLampiranResponse, err error)
}
