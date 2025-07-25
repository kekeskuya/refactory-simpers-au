package service

import (
	"dummy-simpers-au/config"
	"dummy-simpers-au/constants"
	"dummy-simpers-au/internal/dto"
	"dummy-simpers-au/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SimpersServiceImpl struct {
	env          *config.EnvironmentVariable
	db           *pgxpool.Pool
	personelRepo repository.PersonelRepository
	npwpRepo     repository.NPWPRepository
}

func NewSimpersService(
	env *config.EnvironmentVariable,
	db *pgxpool.Pool,
	personelRepo repository.PersonelRepository,
	npwpRepo repository.NPWPRepository,
) SimpersService {
	return &SimpersServiceImpl{
		env:          env,
		db:           db,
		personelRepo: personelRepo,
		npwpRepo:     npwpRepo,
	}
}

func (s *SimpersServiceImpl) GetPersonelByNRP(nrp string) (out dto.GetPersonelByNRPResponse, err error) {
	personel, err := s.personelRepo.GetByNRP(nrp)
	if err != nil {
		return out, err
	}

	if personel.IsEmpty() {
		return out, constants.ErrorMessageDataNotFound
	}

	var (
		tanggalLahir string
		tmtMasuk     string
		tmtPerwira   string
	)

	if personel.TanggalLahir != nil {
		tanggalLahir = personel.TanggalLahir.Format(constants.LayoutDDMMYYYY)
	}

	if personel.TmtMasuk != nil {
		tmtMasuk = personel.TmtMasuk.Format(constants.LayoutDDMMYYYY)
	}

	if personel.TmtPerwira != nil {
		tmtPerwira = personel.TmtPerwira.Format(constants.LayoutDDMMYYYY)
	}

	out = dto.GetPersonelByNRPResponse{
		ID:              int(personel.ID),
		NRP:             personel.NRP,
		Nama:            personel.Nama,
		Pangkat:         personel.Pangkat,
		Korps:           personel.Korps,
		StatusKeaktifan: personel.StatusKeaktifan,
		TempatLahir:     personel.TempatLahir,
		TanggalLahir:    tanggalLahir,
		Jabatan:         personel.Jabatan,
		Profesi:         personel.Profesi,
		Spesialisasi:    personel.Spesialisasi,
		SatuanKerja:     personel.Kesatuan,
		TmtMasuk:        tmtMasuk,
		TmtPerwira:      tmtPerwira,
	}

	return
}

func (s *SimpersServiceImpl) GetNPWPByNRP(nrp string) (out dto.GetNPWPByNRPResponse, err error) {
	npwp, err := s.npwpRepo.GetByNRP(nrp)
	if err != nil {
		return out, err
	}

	if npwp.IsEmpty() {
		err = constants.ErrorMessageDataNotFound
		return out, err
	}

	out = dto.GetNPWPByNRPResponse{
		ID:   int(npwp.ID),
		NRP:  npwp.NRP,
		NPWP: npwp.NPWP,
	}

	return
}
