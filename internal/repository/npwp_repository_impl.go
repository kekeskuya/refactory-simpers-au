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

type NPWPRepoImpl struct {
	db  *database.WrapDB
	env *config.EnvironmentVariable
}

func NewNPWPRepository(db *database.WrapDB, env *config.EnvironmentVariable) NPWPRepository {
	return &NPWPRepoImpl{
		db:  db,
		env: env,
	}
}

func (r *NPWPRepoImpl) GetByID(id int) (out entity.NPWPWithNRP, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT n.id, n.created_at, n.updated_at, n.personel_id, n.npwp, p.nrp 
	FROM npwp n 
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

func (r *NPWPRepoImpl) GetByNRP(nrp string) (out entity.NPWPWithNRP, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT n.id, n.created_at, n.updated_at, n.personel_id, n.npwp, p.nrp 
	FROM npwp n 
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

func (r *NPWPRepoImpl) Create(in model.NPWP) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `INSERT INTO npwp (personel_id, npwp) VALUES ($1, $2)`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, in.PersonelID, in.NPWP)
	if err != nil {
		return err
	}
	return
}

func (r *NPWPRepoImpl) GetAll() (out []model.NPWP, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT id, created_at, updated_at, personel_id, npwp FROM npwp ORDER BY id DESC`

	err = pgxscan.Select(ctx, r.db.Postgres.Conn, &out, query)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (r *NPWPRepoImpl) Update(in model.NPWP) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `UPDATE npwp SET 
	personel_id = $2,
	npwp = $3 
	WHERE id=$1`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, in.ID, in.PersonelID, in.NPWP)
	if err != nil {
		return err
	}
	return
}

func (r *NPWPRepoImpl) DeleteByID(id int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `DELETE FROM npwp WHERE id = $1`
	_, err = r.db.Postgres.Conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return
}
