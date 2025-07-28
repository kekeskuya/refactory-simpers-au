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

type PasporRepoImpl struct {
	db  *database.WrapDB
	env *config.EnvironmentVariable
}

func NewPasporRepository(db *database.WrapDB, env *config.EnvironmentVariable) PasporRepository {
	return &PasporRepoImpl{
		db:  db,
		env: env,
	}
}

func (r *PasporRepoImpl) GetByID(id int) (out entity.PasporWithNRP, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT n.id, n.created_at, n.updated_at, n.personel_id, n.nomor_passport, p.nrp 
	FROM paspor n 
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

func (r *PasporRepoImpl) GetByNRP(nrp string) (out entity.PasporWithNRP, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT n.id, n.created_at, n.updated_at, n.personel_id, n.nomor_passport, p.nrp 
	FROM paspor n 
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

func (r *PasporRepoImpl) Create(in model.Paspor) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `INSERT INTO paspor (personel_id, nomor_passport) VALUES ($1, $2)`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, in.PersonelID, in.NomorPassport)
	if err != nil {
		return err
	}
	return
}

func (r *PasporRepoImpl) GetAll() (out []model.Paspor, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT id, created_at, updated_at, personel_id, nomor_passport FROM paspor ORDER BY id DESC`

	err = pgxscan.Select(ctx, r.db.Postgres.Conn, &out, query)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (r *PasporRepoImpl) Update(in model.Paspor) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `UPDATE paspor SET 
	personel_id = $2,
	nomor_passport = $3 
	WHERE id=$1`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, in.ID, in.PersonelID, in.NomorPassport)
	if err != nil {
		return err
	}
	return
}

func (r *PasporRepoImpl) DeleteByID(id int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `DELETE FROM paspor WHERE id = $1`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return
}
