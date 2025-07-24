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

type PersonnelRepoImpl struct {
	db  *database.WrapDB
	env *config.EnvironmentVariable
}

func NewPersonnelRepository(db *database.WrapDB, env *config.EnvironmentVariable) PersonnelRepository {
	return &PersonnelRepoImpl{
		db:  db,
		env: env,
	}
}

func (r *PersonnelRepoImpl) GetByID(id int) (out model.Personnel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT id, created_at, updated_at, nama, nrp, tmt_masuk, tmt_perwira, pangkat, korps, profesi, spesialisasi, tanggal_lahir, kesatuan, jabatan FROM personnel WHERE id = $1`

	err = pgxscan.Get(ctx, r.db.Postgres.Conn, &out, query, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return out, nil
	}
	if err != nil {
		return out, err
	}

	return out, nil
}

func (r *PersonnelRepoImpl) Create(in model.Personnel) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `INSERT INTO personnel (nama, nrp, tmt_masuk, tmt_perwira, pangkat, korps, profesi, spesialisasi, tanggal_lahir, kesatuan, jabatan) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, in.Nama, in.NRP, in.TmtMasuk, in.TmtPerwira, in.Pangkat, in.Korps, in.Profesi, in.Spesialisasi, in.TanggalLahir, in.Kesatuan, in.Jabatan)
	if err != nil {
		return err
	}
	return
}

func (r *PersonnelRepoImpl) GetAll() (out []model.Personnel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT id, created_at, updated_at, nama, nrp, tmt_masuk, tmt_perwira, pangkat, korps, profesi, spesialisasi, tanggal_lahir, kesatuan, jabatan FROM personnel ORDER BY id DESC`

	err = pgxscan.Select(ctx, r.db.Postgres.Conn, &out, query)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (r *PersonnelRepoImpl) Update(in model.Personnel) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `UPDATE personnel SET 
	nama = $2,
	nrp = $3,
	tmt_masuk = $4,
	tmt_perwira = $5,
	pangkat = $6,
	korps = $7,
	profesi = $8,
	spesialisasi = $9,
	tanggal_lahir = $10,
	kesatuan = $11,
	jabatan = $12 
	WHERE id=$1`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, in.ID, in.Nama, in.NRP, in.TmtMasuk, in.TmtPerwira, in.Pangkat, in.Korps, in.Profesi, in.Spesialisasi, in.TanggalLahir, in.Kesatuan, in.Jabatan)
	if err != nil {
		return err
	}
	return
}

func (r *PersonnelRepoImpl) DeleteByID(id int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `DELETE FROM personnel WHERE id = $1`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return
}
