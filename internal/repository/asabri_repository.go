package repository

import (
	"dummy-simpers-au/internal/entity"
	"dummy-simpers-au/internal/model"
)

type AsabriRepository interface {
	GetByID(id int) (out entity.AsabriWithNRP, err error)
	GetByNRP(nrp string) (out entity.AsabriWithNRP, err error)
	Create(in model.Asabri) (err error)
	GetAll() (out []model.Asabri, err error)
	Update(in model.Asabri) (err error)
	DeleteByID(id int) (err error)
}
