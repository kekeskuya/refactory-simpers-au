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

	var profesiNama string
	if personel.ProfesiNama.Valid {
		profesiNama = personel.ProfesiNama.String
	} else {
		profesiNama = "-"
	}

	var JabatanNamaPanjang string
	if personel.JabatanNamaPanjang.Valid {
		JabatanNamaPanjang = personel.JabatanNamaPanjang.String
	} else {
		JabatanNamaPanjang = "-"
	}

	var statusPersonelID *string
	if personel.StatusPersonelID.Valid {
		statusPersonelID = &personel.StatusPersonelID.String
	} else {
		statusPersonelID = nil
	}

	var nik string
	if personel.NIK.Valid {
		nik = personel.NIK.String
	} else {
		nik = "XXX"
	}

	var pendidikan_asalmasuk_id string
	if personel.PendidikanAsalMasukID.Valid {
		pendidikan_asalmasuk_id = personel.PendidikanAsalMasukID.String
	} else {
		pendidikan_asalmasuk_id = "-"
	}

	var pendidikan_asalmasuk_nama string
	if personel.PendidikanAsalMasukNama.Valid {
		pendidikan_asalmasuk_nama = personel.PendidikanAsalMasukNama.String
	} else {
		pendidikan_asalmasuk_nama = "-"
	}

	var pendidikan_militer_id string
	if personel.PendidikanMiliterID.Valid {
		pendidikan_militer_id = personel.PendidikanMiliterID.String
	} else {
		pendidikan_militer_id = "-"
	}

	var pendidikan_militer_nama string
	if personel.PendidikanMiliterNama.Valid {
		pendidikan_militer_nama = personel.PendidikanMiliterNama.String
	} else {
		pendidikan_militer_nama = "-"
	}

	var pendidikan_umum_id string
	if personel.PendidikanUmumID.Valid {
		pendidikan_umum_id = personel.PendidikanUmumID.String
	} else {
		pendidikan_umum_id = "-"
	}

	var pendidikan_umum_nama string
	if personel.PendidikanUmumNama.Valid {
		pendidikan_umum_nama = personel.PendidikanUmumNama.String
	} else {
		pendidikan_umum_nama = "-"
	}

	var satuan_kerja_id string
	if personel.SatuanKerjaID.Valid {
		satuan_kerja_id = personel.SatuanKerjaID.String
	} else {
		satuan_kerja_id = "-"
	}

	var satuan_kerja_nama string
	if personel.SatuanKerjaNama.Valid {
		satuan_kerja_nama = personel.SatuanKerjaNama.String
	} else {
		satuan_kerja_nama = "-"
	}

	// if personel.TmtMasuk != nil {
	// 	tmtMasuk = personel.TmtMasuk.Format(constants.LayoutDDMMYYYY)
	// }

	// if personel.TmtPerwira != nil {
	// 	tmtPerwira = personel.TmtPerwira.Format(constants.LayoutDDMMYYYY)
	// }

	out = dto.GetPersonelByNRPResponse{
		PersonelID:              uint64(personel.PersonelID),
		PersonelNama:            personel.PersonelNama,
		PangkatNama:             pangkatNama,
		PangkatID:               pangkatID,
		KorpsNama:               korpsNama,
		KorpsID:                 korpsID,
		ProfesiID:               profesiID,
		ProfesiNama:             profesiNama,
		StatusPersonelID:        statusPersonelID,
		JabatanNamaPanjang:      JabatanNamaPanjang,
		NIK:                     nik,
		PendidikanAsalMasukID:   pendidikan_asalmasuk_id,
		PendidikanAsalMasukNama: pendidikan_asalmasuk_nama,
		PendidikanMiliterID:     pendidikan_militer_id,
		PendidikanMiliterNama:   pendidikan_militer_nama,
		PendidikanUmumID:        pendidikan_umum_id,
		PendidikanUmumNama:      pendidikan_umum_nama,
		SatuanKerjaID:           satuan_kerja_id,
		SatuanKerjaNama:         satuan_kerja_nama,
	}

	return
}

func (s *SimpersServiceImpl) GetFamilyCardByNRP(nrp string) ([]dto.GetFamilyCardByNRPResponse, error) {
	familyCards, err := s.personelRepo.GetFamilyCardByNRP(nrp)
	if err != nil {
		return nil, err
	}
	if len(familyCards) == 0 {
		return nil, constants.ErrorMessageDataNotFound
	}

	out := make([]dto.GetFamilyCardByNRPResponse, 0, len(familyCards))
	for _, fc := range familyCards {
		resp := dto.GetFamilyCardByNRPResponse{
			PersonelID:           fc.PersonelID,
			NamaKeluarga:         nullableToString(fc.NamaKeluarga, "-"),
			AlamatKeluarga:       nullableToString(fc.AlamatKeluarga, "-"),
			TempatLahirKeluarga:  nullableToString(fc.TempatLahirKeluarga, "-"),
			TanggalLahirKeluarga: nullableToString(fc.TanggalLahirKeluarga, "-"),
			JenisKelaminKeluarga: nullableToString(fc.JenisKelaminKeluarga, "-"),
			StatusNikahKeluarga:  nullableToString(fc.StatusNikahKeluarga, "-"),
			PekerjaanKeluarga:    nullableToString(fc.PekerjaanKeluarga, "-"),
			HubunganKeluarga:     nullableToString(fc.HubunganKeluarga, "-"),
		}
		out = append(out, resp)
	}

	return out, nil
}

func (s *SimpersServiceImpl) GetDikMilByNRP(nrp string) ([]dto.GetDikMilByNRPResponse, error) {
	dikMils, err := s.personelRepo.GetDikMilByNRP(nrp)
	if err != nil {
		return nil, err
	}
	if len(dikMils) == 0 {
		return nil, constants.ErrorMessageDataNotFound
	}
	out := make([]dto.GetDikMilByNRPResponse, 0, len(dikMils))
	for _, dm := range dikMils {
		resp := dto.GetDikMilByNRPResponse{
			PersonelID:          dm.PersonelID,
			RPendidikanId:       dm.RPendidikanId,
			JenisPendidikanNama: nullableToString(dm.JenisPendidikanNama, "-"),
			SekolahNama:         nullableToString(dm.SekolahNama, "-"),
			TahunLulus:          nullableToString(dm.TahunLulus, "-"),
			Gelar:               nullableToString(dm.Gelar, "-"),
			// Add other fields as necessary
		}
		out = append(out, resp)
	}

	return out, nil
}

func (s *SimpersServiceImpl) GetDikUmByNRP(nrp string) ([]dto.GetDikUmByNRPResponse, error) {
	dikUms, err := s.personelRepo.GetDikUmByNRP(nrp)
	if err != nil {
		return nil, err
	}
	if len(dikUms) == 0 {
		return nil, constants.ErrorMessageDataNotFound
	}
	out := make([]dto.GetDikUmByNRPResponse, 0, len(dikUms))
	for _, du := range dikUms {
		resp := dto.GetDikUmByNRPResponse{
			PersonelID:          du.PersonelID,
			RPendidikanId:       du.RPendidikanId,
			JenisPendidikanNama: nullableToString(du.JenisPendidikanNama, "-"),
			SekolahNama:         nullableToString(du.SekolahNama, "-"),
			TahunLulus:          nullableToString(du.TahunLulus, "-"),
			Gelar:               nullableToString(du.Gelar, "-"),
			// Add other fields as necessary
		}
		out = append(out, resp)
	}

	return out, nil
}

func (s *SimpersServiceImpl) GetJabatanByNRP(nrp string) ([]dto.GetJabatanByNRPResponse, error) {
	jabatans, err := s.personelRepo.GetJabatanByNRP(nrp)
	if err != nil {
		return nil, err
	}
	if len(jabatans) == 0 {
		return nil, constants.ErrorMessageDataNotFound
	}

	out := make([]dto.GetJabatanByNRPResponse, 0, len(jabatans))
	for _, jb := range jabatans {
		resp := dto.GetJabatanByNRPResponse{
			PersonelID:         jb.PersonelID,
			RJabatanId:         nullableToString(jb.RJabatanId, "-"),
			JabatanId:          nullableToString(jb.JabatanId, "-"),
			JabatanNama:        toStringPtr(jb.JabatanNama),
			JabatanNamaPanjang: toStringPtr(jb.JabatanNamaPanjang),
			SuratKeputusan:     nullableToString(jb.SuratKeputusan, "-"),
			SuratKeputusanTgl:  nullableToString(jb.SuratKeputusanTgl, "-"),
			JabatanTMT:         nullableToString(jb.JabatanTMT, "-"),
			JabatanTipe:        nullableToString(jb.JabatanTipe, "-"),
			JabatanKeterangan:  nullableToString(jb.JabatanKeterangan, "-"),
			SprintlakTgl:       nullableToString(jb.SprintlakTgl, "-"),
		}
		out = append(out, resp)
	}

	return out, nil
}

func (s *SimpersServiceImpl) GetTanhorByNRP(nrp string) ([]dto.GetTanhorByNRPResponse, error) {
	tanhors, err := s.personelRepo.GetTanhorByNRP(nrp)
	if err != nil {
		return nil, err
	}
	if len(tanhors) == 0 {
		return nil, constants.ErrorMessageDataNotFound
	}
	out := make([]dto.GetTanhorByNRPResponse, 0, len(tanhors))
	for _, th := range tanhors {
		resp := dto.GetTanhorByNRPResponse{
			PersonelID:        th.PersonelID,
			RJasaId:           nullableToString(th.RJasaId, "-"),
			JasaId:            nullableToString(th.JasaId, "-"),
			JasaNama:          toStringPtr(th.JasaNama),
			SuratKeputusan:    nullableToString(th.SuratKeputusan, "-"),
			SuratKeputusanTgl: nullableToString(th.SuratKeputusanTgl, "-"),
			JasaTMT:           nullableToString(th.JasaTMT, "-"),
			JasaTST:           nullableToString(th.JasaTST, "-"),
			Keterangan:        nullableToString(th.Keterangan, "-"),
		}
		out = append(out, resp)
	}

	return out, nil
}

func (s *SimpersServiceImpl) GetPangkatByNRP(nrp string) ([]dto.GetPangkatByNRPResponse, error) {
	pangkats, err := s.personelRepo.GetPangkatByNRP(nrp)
	if err != nil {
		return nil, err
	}
	if len(pangkats) == 0 {
		return nil, constants.ErrorMessageDataNotFound
	}
	out := make([]dto.GetPangkatByNRPResponse, 0, len(pangkats))
	for _, pk := range pangkats {
		resp := dto.GetPangkatByNRPResponse{
			PersonelID:          pk.PersonelID,
			RPangkatId:          nullableToString(pk.RPangkatId, "-"),
			PangkatId:           nullableToString(pk.PangkatId, "-"),
			PangkatNama:         toStringPtr(pk.PangkatNama),
			SuratKeputusan:      nullableToString(pk.SuratKeputusan, "-"),
			SuratKeputusanTgl:   nullableToString(pk.SuratKeputusanTgl, "-"),
			SuratKeputusanJenis: nullableToString(pk.SuratKeputusanJenis, "-"),
			PangkatTMT:          nullableToString(pk.PangkatTMT, "-"),
			SprintlakTgl:        nullableToString(pk.SprintlakTgl, "-"),
			SprintlakNo:         nullableToString(pk.SprintlakNo, "-"),
		}
		out = append(out, resp)
	}

	return out, nil
}

func toStringPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func nullableToString(ns sql.NullString, fallback string) string {
	if ns.Valid {
		return ns.String
	}
	return fallback
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
