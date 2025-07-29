package service

import (
	"dummy-simpers-au/config"
	"dummy-simpers-au/constants"
	"dummy-simpers-au/internal/dto"
	"dummy-simpers-au/internal/model"
	"dummy-simpers-au/internal/repository"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SimpersServiceImpl struct {
	env          *config.EnvironmentVariable
	db           *pgxpool.Pool
	personelRepo repository.PersonelRepository
	npwpRepo     repository.NPWPRepository
	asabriRepo   repository.AsabriRepository
	pasporRepo   repository.PasporRepository
	lampiranRepo repository.LampiranRepository
}

func NewSimpersService(
	env *config.EnvironmentVariable,
	db *pgxpool.Pool,
	personelRepo repository.PersonelRepository,
	npwpRepo repository.NPWPRepository,
	asabriRepo repository.AsabriRepository,
	pasporRepo repository.PasporRepository,
	lampiranRepo repository.LampiranRepository,
) SimpersService {
	return &SimpersServiceImpl{
		env:          env,
		db:           db,
		personelRepo: personelRepo,
		npwpRepo:     npwpRepo,
		asabriRepo:   asabriRepo,
		pasporRepo:   pasporRepo,
		lampiranRepo: lampiranRepo,
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

func (s *SimpersServiceImpl) GetAsabriByNRP(nrp string) (out dto.GetAsabriByNRPResponse, err error) {
	asabri, err := s.asabriRepo.GetByNRP(nrp)
	if err != nil {
		return out, err
	}

	if asabri.IsEmpty() {
		err = constants.ErrorMessageDataNotFound
		return out, err
	}

	out = dto.GetAsabriByNRPResponse{
		ID:     int(asabri.ID),
		NRP:    asabri.NRP,
		Asabri: asabri.NomorAsabri,
	}

	return
}

func (s *SimpersServiceImpl) GetPasporByNRP(nrp string) (out dto.GetPasporByNRPResponse, err error) {
	paspor, err := s.pasporRepo.GetByNRP(nrp)
	if err != nil {
		return out, err
	}

	if paspor.IsEmpty() {
		err = constants.ErrorMessageDataNotFound
		return out, err
	}

	out = dto.GetPasporByNRPResponse{
		ID:          int(paspor.ID),
		NRP:         paspor.NRP,
		NomorPaspor: paspor.NomorPassport,
	}

	return
}

func (s *SimpersServiceImpl) CreateLampiran(req dto.CreateLampiranRequest, docID int, docType string) (out dto.CreateLampiranResponse, err error) {
	var personelID int
	switch docType {
	case constants.DataAsabri:
		asabri, err := s.asabriRepo.GetByID(docID)
		if err != nil {
			return out, err
		}
		if asabri.IsEmpty() {
			return out, constants.ErrorMessageDataNotFound
		}
		personelID = asabri.PersonelID
	case constants.DataPaspor:
		paspor, err := s.pasporRepo.GetByID(docID)
		if err != nil {
			return out, err
		}
		if paspor.IsEmpty() {
			return out, constants.ErrorMessageDataNotFound
		}
		personelID = paspor.PersonelID
	case constants.DataNPWP:
		npwp, err := s.npwpRepo.GetByID(docID)
		if err != nil {
			return out, err
		}
		if npwp.IsEmpty() {
			return out, constants.ErrorMessageDataNotFound
		}
		personelID = npwp.PersonelID
	default:
		return out, constants.ErrorMessageCategoryNotSupported
	}

	in := model.Lampiran{
		Kategori:   constants.Category(docType),
		PersonelID: personelID,
		DokumenID:  docID,
		Link:       req.FilePath,
		Nama:       filepath.Base(req.FilePath),
		Keterangan: "",
		Tipe:       filepath.Ext(filepath.Base(req.FilePath)),
	}

	id, err := s.lampiranRepo.Create(in)
	if err != nil {
		return out, err
	}
	out.ID = id

	return
}
