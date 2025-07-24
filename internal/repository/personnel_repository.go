package repository

import (
	"dummy-simpers-au/internal/model"
)

type PersonnelRepository interface {
	GetByID(id int) (out model.Personnel, err error)
	Create(in model.Personnel) (err error)
	GetAll() (out []model.Personnel, err error)
	Update(in model.Personnel) (err error)
	DeleteByID(id int) (err error)
}
