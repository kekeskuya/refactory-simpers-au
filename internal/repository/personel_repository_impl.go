package repository

import (
	"context" // Import the initializers package
	"dummy-simpers-au/config"
	"dummy-simpers-au/database"
	"dummy-simpers-au/internal/model"
	"errors"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
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

func (r *PersonelRepoImpl) GetByID(id int) (out model.Personel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT id, created_at, updated_at, nama, nrp, tmt_masuk, tmt_perwira, pangkat, korps, profesi, spesialisasi, tempat_lahir, tanggal_lahir, kesatuan, jabatan, status_keaktifan FROM personel WHERE id = $1`

	err = pgxscan.Get(ctx, r.db.Postgres.Conn, &out, query, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return out, nil
	}
	if err != nil {
		return out, err
	}

	return out, nil
}

func (r *PersonelRepoImpl) GetByNRP(nrp string) (out model.Personel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT id, created_at, updated_at, nama, nrp, tmt_masuk, tmt_perwira, pangkat, korps, profesi, spesialisasi, tempat_lahir, tanggal_lahir, kesatuan, jabatan, status_keaktifan FROM personel WHERE nrp = $1`

	err = pgxscan.Get(ctx, r.db.Postgres.Conn, &out, query, nrp)
	if errors.Is(err, pgx.ErrNoRows) {
		return out, nil
	}
	if err != nil {
		return out, err
	}

	return out, nil
}

func (r *PersonelRepoImpl) Create(in model.Personel) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `INSERT INTO personel (nama, nrp, tmt_masuk, tmt_perwira, pangkat, korps, profesi, spesialisasi, tempat_lahir, tanggal_lahir, kesatuan, jabatan, status_keaktifan) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, in.Nama, in.NRP, in.TmtMasuk, in.TmtPerwira, in.Pangkat, in.Korps, in.Profesi, in.Spesialisasi, in.TempatLahir, in.TanggalLahir, in.Kesatuan, in.Jabatan, in.StatusKeaktifan)
	if err != nil {
		return err
	}
	return
}

func (r *PersonelRepoImpl) GetAll() (out []model.Personel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT id, created_at, updated_at, nama, nrp, tmt_masuk, tmt_perwira, pangkat, korps, profesi, spesialisasi, tempat_lahir, tanggal_lahir, kesatuan, jabatan, status_keaktifan FROM personel ORDER BY id DESC`

	err = pgxscan.Select(ctx, r.db.Postgres.Conn, &out, query)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (r *PersonelRepoImpl) Update(in model.Personel) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `UPDATE personel SET 
	nama = $2,
	nrp = $3,
	tmt_masuk = $4,
	tmt_perwira = $5,
	pangkat = $6,
	korps = $7,
	profesi = $8,
	spesialisasi = $9,
	tempat_lahir = $10,
	tanggal_lahir = $11,
	kesatuan = $12,
	jabatan = $13,
	status_keaktifan = $14  
	WHERE id=$1`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, in.ID, in.Nama, in.NRP, in.TmtMasuk, in.TmtPerwira, in.Pangkat, in.Korps, in.Profesi, in.Spesialisasi, in.TempatLahir, in.TanggalLahir, in.Kesatuan, in.Jabatan, in.StatusKeaktifan)
	if err != nil {
		return err
	}
	return
}

func (r *PersonelRepoImpl) DeleteByID(id int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `DELETE FROM personel WHERE id = $1`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return
}
