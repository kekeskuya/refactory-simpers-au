package repository

import (
	"context" // Import the initializers package
	"refactory-simpers-au/config"
	"refactory-simpers-au/database"
	"refactory-simpers-au/internal/entity"
	"refactory-simpers-au/internal/model"

	"github.com/georgysavva/scany/v2/sqlscan"
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
	WHERE n.id = @p1`

	err = sqlscan.Get(ctx, r.db.SQLserver.Conn, &out, query, id)
	if sqlscan.NotFound(err) {
		return out, nil
	}
	return out, err
}

func (r *NPWPRepoImpl) GetByNRP(nrp string) (out entity.NPWPWithNRP, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT n.id, n.created_at, n.updated_at, n.personel_id, n.npwp, p.nrp 
	FROM npwp n 
	LEFT JOIN personel p ON p.id = n.personel_id 
	WHERE p.nrp = @p1`

	err = sqlscan.Get(ctx, r.db.SQLserver.Conn, &out, query, nrp)
	if sqlscan.NotFound(err) {
		return out, nil
	}
	return out, err
}

func (r *NPWPRepoImpl) Create(in model.NPWP) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `INSERT INTO npwp (personel_id, npwp) VALUES (@p1, @p2)`

	_, err = r.db.SQLserver.Conn.ExecContext(ctx, query, in.PersonelID, in.NPWP)
	return err
}

func (r *NPWPRepoImpl) GetAll() (out []model.NPWP, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT id, created_at, updated_at, personel_id, npwp FROM npwp ORDER BY id DESC`

	err = sqlscan.Select(ctx, r.db.SQLserver.Conn, &out, query)
	return out, err
}

func (r *NPWPRepoImpl) Update(in model.NPWP) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `UPDATE npwp SET 
	personel_id = @p2,
	npwp = @p3 
	WHERE id=@p1`

	_, err = r.db.SQLserver.Conn.ExecContext(ctx, query, in.ID, in.PersonelID, in.NPWP)
	return err
}

func (r *NPWPRepoImpl) DeleteByID(id int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `DELETE FROM npwp WHERE id = @p1`
	_, err = r.db.SQLserver.Conn.ExecContext(ctx, query, id)
	return err
}
