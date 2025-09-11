package repository

import (
	"context" // Import the initializers package
	"refactory-simpers-au/config"
	"refactory-simpers-au/database"
	"refactory-simpers-au/internal/entity"
	"refactory-simpers-au/internal/model"

	"github.com/georgysavva/scany/v2/sqlscan"
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
	WHERE n.id = @p1`

	err = sqlscan.Get(ctx, r.db.SQLserver.Conn, &out, query, id)
	if sqlscan.NotFound(err) {
		return out, nil
	}
	return out, err
}

func (r *PasporRepoImpl) GetByNRP(nrp string) (out entity.PasporWithNRP, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT n.id, n.created_at, n.updated_at, n.personel_id, n.nomor_passport, p.nrp 
	FROM paspor n 
	LEFT JOIN personel p ON p.id = n.personel_id 
	WHERE p.nrp = @p1`

	err = sqlscan.Get(ctx, r.db.SQLserver.Conn, &out, query, nrp)
	if sqlscan.NotFound(err) {
		return out, nil
	}
	return out, err
}

func (r *PasporRepoImpl) Create(in model.Paspor) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `INSERT INTO paspor (personel_id, nomor_passport) VALUES (@p1, @p2)`
	_, err = r.db.SQLserver.Conn.ExecContext(ctx, query, in.PersonelID, in.NomorPassport)
	return err
}

func (r *PasporRepoImpl) GetAll() (out []model.Paspor, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT id, created_at, updated_at, personel_id, nomor_passport FROM paspor ORDER BY id DESC`

	err = sqlscan.Select(ctx, r.db.SQLserver.Conn, &out, query)
	return out, err
}

func (r *PasporRepoImpl) Update(in model.Paspor) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `UPDATE paspor SET 
	personel_id = @p2,
	nomor_passport = @p3 
	WHERE id=@p1`

	_, err = r.db.SQLserver.Conn.ExecContext(ctx, query, in.ID, in.PersonelID, in.NomorPassport)
	return err
}

func (r *PasporRepoImpl) DeleteByID(id int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `DELETE FROM paspor WHERE id = @p1`
	_, err = r.db.SQLserver.Conn.ExecContext(ctx, query, id)
	return err
}
