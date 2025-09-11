package repository

import (
	"refactory-simpers-au/internal/entity"
	"refactory-simpers-au/internal/model"
)

type NPWPRepository interface {
	GetByID(id int) (out entity.NPWPWithNRP, err error)
	GetByNRP(nrp string) (out entity.NPWPWithNRP, err error)
	Create(in model.NPWP) (err error)
	GetAll() (out []model.NPWP, err error)
	Update(in model.NPWP) (err error)
	DeleteByID(id int) (err error)
}
