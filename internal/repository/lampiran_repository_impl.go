package repository

import (
	"context" // Import the initializers package
	"dummy-simpers-au/config"
	"dummy-simpers-au/database"
	"dummy-simpers-au/internal/entity"
	"dummy-simpers-au/internal/model"
	"errors"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type LampiranRepoImpl struct {
	db  *database.WrapDB
	env *config.EnvironmentVariable
}

func NewLampiranRepository(db *database.WrapDB, env *config.EnvironmentVariable) LampiranRepository {
	return &LampiranRepoImpl{
		db:  db,
		env: env,
	}
}

func (r *LampiranRepoImpl) GetByID(id int) (out entity.LampiranWithNRP, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT l.id, l.created_at, l.updated_at, l.personel_id, l.dokumen_id, l.kategori, l.link, l.nama, l.keterangan, l.tipe , p.nrp 
	FROM lampiran l	 
	LEFT JOIN personel p ON p.id = l.personel_id 
	WHERE l.id = $1`

	err = pgxscan.Get(ctx, r.db.Postgres.Conn, &out, query, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return out, nil
	}
	if err != nil {
		return out, err
	}

	return out, nil
}

func (r *LampiranRepoImpl) GetByNRP(nrp string) (out entity.LampiranWithNRP, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT l.id, l.created_at, l.updated_at, l.personel_id, l.dokumen_id, l.kategori, l.link, l.nama, l.keterangan, l.tipe , p.nrp 
	FROM lampiran l 
	LEFT JOIN personel p ON p.id = l.personel_id 
	WHERE p.nrp = $1`

	err = pgxscan.Get(ctx, r.db.Postgres.Conn, &out, query, nrp)
	if errors.Is(err, pgx.ErrNoRows) {
		return out, nil
	}
	if err != nil {
		return out, err
	}

	return out, nil
}

func (r *LampiranRepoImpl) Create(in model.Lampiran) (id int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `
		INSERT INTO lampiran (personel_id, dokumen_id, kategori, link, nama, keterangan, tipe)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	err = r.db.Postgres.Conn.
		QueryRow(ctx, query, in.PersonelID, in.DokumenID, in.Kategori, in.Link, in.Nama, in.Keterangan, in.Tipe).
		Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *LampiranRepoImpl) GetAll() (out []model.Lampiran, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT id, created_at, updated_at, personel_id, dokumen_id, kategori, link, nama, keterangan, tipe FROM lampiran ORDER BY id DESC`

	err = pgxscan.Select(ctx, r.db.Postgres.Conn, &out, query)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (r *LampiranRepoImpl) Update(in model.Lampiran) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `UPDATE lampiran SET 
	personel_id = $2,
	dokumen_id = $3,
	kategori = $4, 
	link = $5, 
	nama = $6, 
	keterangan = $7, 
	tipe = $8 
	WHERE id=$1`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, in.ID, in.PersonelID, in.DokumenID, in.Kategori, in.Link, in.Nama, in.Keterangan, in.Tipe)
	if err != nil {
		return err
	}
	return
}

func (r *LampiranRepoImpl) DeleteByID(id int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `DELETE FROM lampiran WHERE id = $1`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return
}
