package repository

import (
	"refactory-simpers-au/internal/entity"
)

type AsabriRepository interface {
	GetByID(id string) (out entity.AsabriWithNRP, err error)
	GetByNRP(nrp string) (out entity.AsabriWithNRP, err error)
	// Create(in model.Asabri) (err error)
	// GetAll() (out []model.Asabri, err error)
	// Update(in model.Asabri) (err error)
	// DeleteByID(id int) (err error)
}
