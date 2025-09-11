package repository

import (
	"refactory-simpers-au/internal/entity"
	"refactory-simpers-au/internal/model"
)

type PasporRepository interface {
	GetByID(id int) (out entity.PasporWithNRP, err error)
	GetByNRP(nrp string) (out entity.PasporWithNRP, err error)
	Create(in model.Paspor) (err error)
	GetAll() (out []model.Paspor, err error)
	Update(in model.Paspor) (err error)
	DeleteByID(id int) (err error)
}
