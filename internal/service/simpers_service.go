package service

import "dummy-simpers-au/internal/dto"

type SimpersService interface {
	GetPersonelByNRP(nrp string) (out dto.GetPersonelByNRPResponse, err error)
	GetNPWPByNRP(nrp string) (out dto.GetNPWPByNRPResponse, err error)
}
