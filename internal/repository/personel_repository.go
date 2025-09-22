package repository

import (
	"refactory-simpers-au/internal/model"
)

type PersonelRepository interface {
	GetByID(id string) (out model.Personel, err error)
	GetByNRP(nrp string) (out model.Personel, err error)
	GetFamilyCardByNRP(nrp string) ([]model.FamilyCard, error)
	GetDikMilByNRP(nrp string) ([]model.DikMil, error)
	GetDikUmByNRP(nrp string) ([]model.DikUm, error)
	GetJabatanByNRP(nrp string) ([]model.Jabatan, error)
	GetTanhorByNRP(nrp string) ([]model.Tanhor, error)
	GetPangkatByNRP(nrp string) ([]model.Pangkat, error)
	//	Create(in model.Personel) (err error)
	//	GetAll() (out []model.Personel, err error)
	//	Update(in model.Personel) (err error)
	//	DeleteByID(id int) (err error)
}
