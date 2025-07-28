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

type AsabriRepoImpl struct {
	db  *database.WrapDB
	env *config.EnvironmentVariable
}

func NewAsabriRepository(db *database.WrapDB, env *config.EnvironmentVariable) AsabriRepository {
	return &AsabriRepoImpl{
		db:  db,
		env: env,
	}
}

func (r *AsabriRepoImpl) GetByID(id int) (out entity.AsabriWithNRP, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT n.id, n.created_at, n.updated_at, n.personel_id, n.nomor_asabri, p.nrp 
	FROM asabri n 
	LEFT JOIN personel p ON p.id = n.personel_id 
	WHERE n.id = $1`

	err = pgxscan.Get(ctx, r.db.Postgres.Conn, &out, query, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return out, nil
	}
	if err != nil {
		return out, err
	}

	return out, nil
}

func (r *AsabriRepoImpl) GetByNRP(nrp string) (out entity.AsabriWithNRP, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT n.id, n.created_at, n.updated_at, n.personel_id, n.nomor_asabri, p.nrp 
	FROM asabri n 
	LEFT JOIN personel p ON p.id = n.personel_id 
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

func (r *AsabriRepoImpl) Create(in model.Asabri) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `INSERT INTO asabri (personel_id, nomor_asabri) VALUES ($1, $2)`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, in.PersonelID, in.NomorAsabri)
	if err != nil {
		return err
	}
	return
}

func (r *AsabriRepoImpl) GetAll() (out []model.Asabri, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT id, created_at, updated_at, personel_id, nomor_asabri FROM asabri ORDER BY id DESC`

	err = pgxscan.Select(ctx, r.db.Postgres.Conn, &out, query)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (r *AsabriRepoImpl) Update(in model.Asabri) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `UPDATE asabri SET 
	personel_id = $2,
	nomor_asabri = $3 
	WHERE id=$1`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, in.ID, in.PersonelID, in.NomorAsabri)
	if err != nil {
		return err
	}
	return
}

func (r *AsabriRepoImpl) DeleteByID(id int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `DELETE FROM asabri WHERE id = $1`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return
}
