package repository

import (
	"context" // Import the initializers package
	"fmt"
	"refactory-simpers-au/config"
	"refactory-simpers-au/database"
	"refactory-simpers-au/internal/model"
	"strings"

	"github.com/georgysavva/scany/v2/sqlscan"
)

type PersonelRepoImpl struct {
	db  *database.WrapDB
	env *config.EnvironmentVariable
}

func NewPersonelRepository(db *database.WrapDB, env *config.EnvironmentVariable) PersonelRepository {
	return &PersonelRepoImpl{
		db:  db,
		env: env,
	}
}

func (r *PersonelRepoImpl) GetByID(id string) (out model.Personel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()
	fmt.Println("DEBUG ID:", id)
	// query := `SELECT id, created_at, updated_at, nama, nrp, tmt_masuk, tmt_perwira, pangkat, korps, profesi, spesialisasi, tempat_lahir, tanggal_lahir, kesatuan, jabatan, StatusPersonel_Id FROM personel WHERE personel_id = @p1`
	query := `SELECT Personel_Nama, personel_id, StatusPersonel_Id FROM personel WHERE personel_id = '@p1' `

	fmt.Println("DEBUG QUERY 1:", query)
	err = sqlscan.Get(ctx, r.db.SQLserver.Conn, &out, query, id)
	if sqlscan.NotFound(err) {
		return out, nil
	}
	return out, err
}

// func (r *PersonelRepoImpl) GetByNRP(nrp string) (out model.Personel, err error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
// 	defer cancel()

// 	query := `SELECT Personel_Nama, personel_id, StatusPersonel_Id FROM personel WHERE personel_id = @p1 `
// 	fmt.Println("DEBUG QUERY 2:", query)
// 	err = sqlscan.Get(ctx, r.db.SQLserver.Conn, &out, query, nrp)
// 	if sqlscan.NotFound(err) {
// 		return out, nil
// 	}
// 	return out, err
// }

func (r *PersonelRepoImpl) GetByNRP(personel_id string) (out model.Personel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	personel_id = strings.TrimSpace(personel_id)

	query := `SELECT 
        Personel_Id        AS personel_id,
        Personel_Nama      AS personel_nama,
        Pangkat_Nama       AS pangkat_nama,
        Pangkat_Id         AS pangkat_id,
        StatusPersonel_Id  AS statuspersonel_id,
		Korps_Nama        AS korps_nama,
		Korps_Id          AS korps_id,
		Profesi_Nama      AS profesi_nama,
		Profesi_Id        AS profesi_id,
		Jabatan_Nama_Panjang AS jabatan_nama_panjang,
		NIK AS nik,
		Pendidikan_AsalMasuk_Id as pendidikan_asalmasuk_id,
		Pendidikan_AsalMasuk_Nama as pendidikan_asalmasuk_nama,
		Pendidikan_Militer_Id as pendidikan_militer_id,
		Pendidikan_Militer_Nama as pendidikan_militer_nama,
		Pendidikan_Umum_Id as pendidikan_umum_id,
		Pendidikan_Umum_Nama as pendidikan_umum_nama,
		Satuan_Kerja_Id as satuan_kerja_id,
		Satuan_Kerja_Nama as satuan_kerja_nama
    FROM V_personel_twp
    WHERE Personel_Id = @p1`

	fmt.Println("DEBUG QUERY 2:", query, " NRP:", personel_id)

	err = sqlscan.Get(ctx, r.db.SQLserver.Conn, &out, query, personel_id)
	if sqlscan.NotFound(err) {
		return out, nil
	}
	return out, err
}

// func (r *PersonelRepoImpl) GetFamilyCardByNRP(personel_id string) (out []model.FamilyCard, err error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
// 	defer cancel()

// 	personel_id = strings.TrimSpace(personel_id)

// 	query := `SELECT Personel_Id AS personel_id,
// 					 nama AS nama_keluarga,
// 					 alamat AS alamat_keluarga,
// 					 tempat_lahir AS tempat_lahir_keluarga,
// 					 tanggal_lahir AS tanggal_lahir_keluarga,
// 					 jenis_kelamin AS jenis_kelamin_keluarga,
// 					 status_nikah AS status_nikah_keluarga,
// 					 hubungan_keluarga AS hubungan_keluarga,
// 					 pekerjaan AS pekerjaan_keluarga
// 			  FROM M_keluarga
// 	          WHERE Personel_Id = @p1`

// 	fmt.Println("DEBUG QUERY Family :", query, " NRP:", personel_id)
// 	err = sqlscan.Select(ctx, r.db.SQLserver.Conn, &out, query, personel_id)
// 	if sqlscan.NotFound(err) {
// 		return out, nil
// 	}
// 	return out, err
// }

func (r *PersonelRepoImpl) GetFamilyCardByNRP(nrp string) ([]model.FamilyCard, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	var out []model.FamilyCard
	query := `
        SELECT 
            Personel_Id AS personel_id,
			nama AS nama_keluarga,
			alamat AS alamat_keluarga,
			tempat_lahir AS tempat_lahir_keluarga,
			tanggal_lahir AS tanggal_lahir_keluarga,
			jenis_kelamin AS jenis_kelamin_keluarga,
			status_nikah AS status_nikah_keluarga,
			hubungan_keluarga AS hubungan_keluarga,
			pekerjaan AS pekerjaan_keluarga
        FROM M_keluarga
        WHERE personel_id = @p1
    `

	err := sqlscan.Select(ctx, r.db.SQLserver.Conn, &out, query, nrp)
	if sqlscan.NotFound(err) {
		return []model.FamilyCard{}, nil
	}
	return out, err
}

func (r *PersonelRepoImpl) GetDikMilByNRP(nrp string) ([]model.DikMil, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()
	var out []model.DikMil
	query := `
		SELECT
			M_JenisPendidikan.JenisPendidikan_Nama,R_Pendidikan.Sekolah_Nama,R_Pendidikan.LembagaPendidikan_Nama
		  , R_Pendidikan.SuratKeputusan, R_Pendidikan.SuratKeputusan_Tgl, R_Pendidikan.Pendidikan_TMT
		  , R_Pendidikan.Tahun_Masuk, R_Pendidikan.Tahun_Lulus, R_Pendidikan.Gelar
		FROM R_Pendidikan
		LEFT JOIN M_LembagaPendidikan on M_LembagaPendidikan.LembagaPendidikan_Id = R_Pendidikan.LembagaPendidikan_Id 
	    LEFT JOIN M_JenisPendidikan on M_JenisPendidikan.JenisPendidikan_Id = M_LembagaPendidikan.JenisPendidikan_Id
		WHERE M_JenisPendidikan.JenisPendidikan_Tipe = 'M'
		  AND Personel_Id = @p1
	`

	err := sqlscan.Select(ctx, r.db.SQLserver.Conn, &out, query, nrp)
	if sqlscan.NotFound(err) {
		return []model.DikMil{}, nil
	}
	return out, err
}

func (r *PersonelRepoImpl) GetDikUmByNRP(nrp string) ([]model.DikUm, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()
	var out []model.DikUm
	query := `
		SELECT
			M_JenisPendidikan.JenisPendidikan_Nama,R_Pendidikan.Sekolah_Nama,R_Pendidikan.LembagaPendidikan_Nama
		  , R_Pendidikan.SuratKeputusan, R_Pendidikan.SuratKeputusan_Tgl, R_Pendidikan.Pendidikan_TMT
		  , R_Pendidikan.Tahun_Masuk, R_Pendidikan.Tahun_Lulus, R_Pendidikan.Gelar
		FROM R_Pendidikan
		LEFT JOIN M_LembagaPendidikan on M_LembagaPendidikan.LembagaPendidikan_Id = R_Pendidikan.LembagaPendidikan_Id 
	    LEFT JOIN M_JenisPendidikan on M_JenisPendidikan.JenisPendidikan_Id = M_LembagaPendidikan.JenisPendidikan_Id
		WHERE M_JenisPendidikan.JenisPendidikan_Tipe = 'U'
		  AND Personel_Id = @p1
	`

	err := sqlscan.Select(ctx, r.db.SQLserver.Conn, &out, query, nrp)
	if sqlscan.NotFound(err) {
		return []model.DikUm{}, nil
	}
	return out, err
}

func (r *PersonelRepoImpl) GetJabatanByNRP(nrp string) ([]model.Jabatan, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()
	var out []model.Jabatan
	query := `
		select RJ.R_Jabatan_Id as r_jabatan_id
			 , RJ.Jabatan_Id as jabatan_id
			 , RJ.Personel_Id as personel_id
			 , MJ.Jabatan_Nama as jabatan_nama
			 , MJ.Jabatan_Nama_Panjang as jabatan_nama_panjang
			 , RJ.Jabatan_TMT as jabatan_tmt
			 , RJ.R_Jabatan_Tipe as r_jabatan_tipe
		     , RJ.SuratKeputusan as suratkeputusan
		     , RJ.SuratKeputusan_Tgl as suratkeputusan_tgl
			 , RJ.R_Jabatan_Keterangan as r_jabatan_keterangan
			 , MJ.Jabatan_Status as jabatan_status
		from R_Jabatan RJ
		LEFT JOIN M_Jabatan MJ on MJ.Jabatan_Id = RJ.Jabatan_Id
		where RJ.Personel_Id = @p1
	`

	fmt.Println("DEBUG QUERY Jabatan :", query, " NRP:", nrp)
	err := sqlscan.Select(ctx, r.db.SQLserver.Conn, &out, query, nrp)
	if sqlscan.NotFound(err) {
		return []model.Jabatan{}, nil
	}
	return out, err
}

// func (r *PersonelRepoImpl) Create(in model.Personel) (err error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
// 	defer cancel()

// 	query := `INSERT INTO personel (nama, nrp, tmt_masuk, tmt_perwira, pangkat, korps, profesi, spesialisasi, tempat_lahir, tanggal_lahir, kesatuan, jabatan, StatusPersonel_Id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
// 	_, err = r.db.SQLserver.Conn.ExecContext(ctx, query, in.Nama, in.NRP, in.TmtMasuk, in.TmtPerwira, in.Pangkat, in.Korps, in.Profesi, in.Spesialisasi, in.TempatLahir, in.TanggalLahir, in.Kesatuan, in.Jabatan, in.StatusPersonel_Id)
// 	return err
// }

func (r *PersonelRepoImpl) GetAll() (out []model.Personel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	// query := `SELECT id, created_at, updated_at, nama, nrp, tmt_masuk, tmt_perwira, pangkat, korps, profesi, spesialisasi, tempat_lahir, tanggal_lahir, kesatuan, jabatan, StatusPersonel_Id FROM personel ORDER BY id DESC`
	query := `SELECT Personel_Nama, personel_id, StatusPersonel_Id FROM personel ORDER BY personel_id DESC`
	err = sqlscan.Select(ctx, r.db.SQLserver.Conn, &out, query)
	return out, err
}

// func (r *PersonelRepoImpl) Update(in model.Personel) (err error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
// 	defer cancel()

// 	query := `UPDATE personel SET
// 	nama = @p2,
// 	nrp = @p3,
// 	tmt_masuk = @p4,
// 	tmt_perwira = @p5,
// 	pangkat = @p6,
// 	korps = @p7,
// 	profesi = @p8,
// 	spesialisasi = @p9,
// 	tempat_lahir = @p10,
// 	tanggal_lahir = @p11,
// 	kesatuan = @p12,
// 	jabatan = @p13,
// 	StatusPersonel_Id = @p14
// 	WHERE id=@p1`
// 	_, err = r.db.SQLserver.Conn.ExecContext(ctx, query, in.ID, in.Nama, in.NRP, in.TmtMasuk, in.TmtPerwira, in.Pangkat, in.Korps, in.Profesi, in.Spesialisasi, in.TempatLahir, in.TanggalLahir, in.Kesatuan, in.Jabatan)
// 	return err
// }

func (r *PersonelRepoImpl) DeleteByID(id int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `DELETE FROM personel WHERE id = @p1`
	_, err = r.db.SQLserver.Conn.ExecContext(ctx, query, id)
	return err
}
