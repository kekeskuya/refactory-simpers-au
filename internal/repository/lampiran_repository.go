package repository

import (
	"dummy-simpers-au/internal/entity"
	"dummy-simpers-au/internal/model"
)

type LampiranRepository interface {
	GetByID(id int) (out entity.LampiranWithNRP, err error)
	GetByNRP(nrp string) (out entity.LampiranWithNRP, err error)
	Create(in model.Lampiran) (id int, err error)
	GetAll() (out []model.Lampiran, err error)
	Update(in model.Lampiran) (err error)
	DeleteByID(id int) (err error)
}
