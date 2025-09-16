package service

import (
	"database/sql"
	"path/filepath"
	"refactory-simpers-au/config"
	"refactory-simpers-au/constants"
	"refactory-simpers-au/internal/dto"
	"refactory-simpers-au/internal/model"
	"refactory-simpers-au/internal/repository"
)

type SimpersServiceImpl struct {
	env          *config.EnvironmentVariable
	db           *sql.DB
	personelRepo repository.PersonelRepository
	npwpRepo     repository.NPWPRepository
	asabriRepo   repository.AsabriRepository
	pasporRepo   repository.PasporRepository
	lampiranRepo repository.LampiranRepository
}

func NewSimpersService(
	env *config.EnvironmentVariable,
	db *sql.DB,
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

	// var (
	// 	tanggalLahir string
	// 	tmtMasuk     string
	// 	tmtPerwira   string
	// )
	var pangkatNama string
	if personel.PangkatNama.Valid {
		pangkatNama = personel.PangkatNama.String
	} else {
		pangkatNama = "-"
	}

	var pangkatID *int32
	if personel.PangkatID.Valid {
		pangkatID = &personel.PangkatID.Int32
	} else {
		pangkatID = nil
	}

	var korpsNama string
	if personel.KorpsNama.Valid {
		korpsNama = personel.KorpsNama.String
	} else {
		korpsNama = "-"
	}

	var korpsID *int32
	if personel.KorpsID.Valid {
		korpsID = &personel.KorpsID.Int32
	} else {
		korpsID = nil
	}

	var profesiID *int32
	if personel.ProfesiID.Valid {
		profesiID = &personel.ProfesiID.Int32
	} else {
		profesiID = nil
	}

	var JabatanNamaPanjang string
	if personel.JabatanNamaPanjang.Valid {
		JabatanNamaPanjang = personel.JabatanNamaPanjang.String
	} else {
		JabatanNamaPanjang = "-"
	}

	// if personel.TmtMasuk != nil {
	// 	tmtMasuk = personel.TmtMasuk.Format(constants.LayoutDDMMYYYY)
	// }

	// if personel.TmtPerwira != nil {
	// 	tmtPerwira = personel.TmtPerwira.Format(constants.LayoutDDMMYYYY)
	// }

	out = dto.GetPersonelByNRPResponse{
		PersonelID:         uint64(personel.PersonelID),
		PersonelNama:       personel.PersonelNama,
		PangkatNama:        pangkatNama,
		PangkatID:          pangkatID,
		KorpsNama:          korpsNama,
		KorpsID:            korpsID,
		ProfesiID:          profesiID,
		JabatanNamaPanjang: JabatanNamaPanjang,
		// TanggalLahir: tanggalLahir,
		// TmtMasuk:     tmtMasuk,
		// TmtPerwira:   tmtPerwira,
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
		PersonelID: int(npwp.PersonelID),
		NPWP:       npwp.NPWP,
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
		PersonelID: int(asabri.PersonelID),
		Asabri:     asabri.NomorAsabri,
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
		PersonelID:  int(paspor.PersonelID),
		NomorPaspor: paspor.NomorPassport,
	}

	return
}

func (s *SimpersServiceImpl) CreateLampiran(req dto.CreateLampiranRequest, docID string, docType string) (out dto.CreateLampiranResponse, err error) {
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
	out.PersonelID = id

	return
}
