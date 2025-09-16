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

	// query := `SELECT personel.Personel_Id, personel.Personel_Nama,  M_Pangkat.Pangkat_Nama,  M_Korps.Korps_Nama, M_Profesi.Profesi_Nama, ` +
	// 	`M_LembagaPendidikan.LembagaPendidikan_Nama, FORMAT(R_Pendidikan.Pendidikan_TMT,'dd-MM-yyyy'), FORMAT(personel.Personel_TanggalLahir,'dd-MM-yyyy'), M_Jabatan.Jabatan_Nama_Panjang as jabatan ` +
	// 	`FROM personel ` +
	// 	`LEFT JOIN R_Pangkat on R_Pangkat.Personel_Id = Personel.Personel_Id and R_Pangkat.Pangkat_Status = '1' ` +
	// 	`left join M_Pangkat on M_Pangkat.Pangkat_Id = R_Pangkat.Pangkat_Id ` +
	// 	`left join R_Korps on R_Korps.Personel_Id = Personel.Personel_Id and R_Korps.Korps_Status = '1' ` +
	// 	`left join M_Korps on M_Korps.Korps_Id = R_Korps.Korps_Id ` +
	// 	`LEFT JOIN M_Profesi on M_profesi.Profesi_Id = personel.Profesi_Id ` +
	// 	`left join R_Pendidikan on R_Pendidikan.Personel_Id = personel.Personel_Id ` +
	// 	`LEFT JOIN M_LembagaPendidikan on M_LembagaPendidikan.LembagaPendidikan_Id = R_Pendidikan.LembagaPendidikan_Id ` +
	// 	`left join R_Jabatan on R_Jabatan.Personel_Id = personel.Personel_Id and R_Jabatan.R_Jabatan_Status = '1' ` +
	// 	`left JOIN M_Jabatan on M_Jabatan.Jabatan_Id = R_Jabatan.Jabatan_Id ` +
	// 	`where personel.personel_id = '21718912546621' and M_LembagaPendidikan.PendidikanAsalMasuk = '1' ` +
	// 	`order by M_pangkat.Pangkat_Id desc`
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

	query := `SELECT Personel_Nama, personel_id, StatusPersonel_Id
              FROM personel
              WHERE personel_id = @p1`

	fmt.Println("DEBUG QUERY 2:", query, " NRP:", personel_id)

	err = sqlscan.Get(ctx, r.db.SQLserver.Conn, &out, query, personel_id)
	if sqlscan.NotFound(err) {
		return out, nil
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
