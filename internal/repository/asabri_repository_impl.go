package repository

import (
	"context"
	"refactory-simpers-au/config"
	"refactory-simpers-au/database"
	"refactory-simpers-au/internal/entity"
	"refactory-simpers-au/internal/model"

	"github.com/georgysavva/scany/v2/sqlscan"
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

func (r *AsabriRepoImpl) GetByID(id string) (out entity.AsabriWithNRP, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT n.id, n.created_at, n.updated_at, n.personel_id, n.nomor_asabri, p.nrp
              FROM asabri n
              LEFT JOIN personel p ON p.id = n.personel_id
              WHERE n.id = @p1`

	err = sqlscan.Get(ctx, r.db.SQLserver.Conn, &out, query, id)
	if sqlscan.NotFound(err) {
		return out, nil
	}
	return out, err
}

func (r *AsabriRepoImpl) GetByNRP(nrp string) (out entity.AsabriWithNRP, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT n.id, n.created_at, n.updated_at, n.personel_id, n.nomor_asabri, p.nrp
			  FROM asabri n
			  LEFT JOIN personel p ON p.id = n.personel_id
			  WHERE p.nrp = @p1`

	err = sqlscan.Get(ctx, r.db.SQLserver.Conn, &out, query, nrp)
	if sqlscan.NotFound(err) {
		return out, nil
	}
	return out, err
}

func (r *AsabriRepoImpl) Create(in model.Asabri) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `INSERT INTO asabri (personel_id, nomor_asabri) VALUES (@p1, @p2)`
	_, err = r.db.SQLserver.Conn.ExecContext(ctx, query, in.PersonelID, in.NomorAsabri)
	return err
}

func (r *AsabriRepoImpl) GetAll() (out []model.Asabri, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT id, created_at, updated_at, personel_id, nomor_asabri FROM asabri ORDER BY id DESC`

	err = sqlscan.Select(ctx, r.db.SQLserver.Conn, &out, query)
	return out, err
}

func (r *AsabriRepoImpl) Update(in model.Asabri) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `UPDATE asabri
			  SET nomor_asabri = @p3
			  WHERE id = @p1`

	_, err = r.db.SQLserver.Conn.ExecContext(ctx, query, in.PersonelID, in.NomorAsabri)
	return err
}

func (r *AsabriRepoImpl) DeleteByID(id int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `DELETE FROM asabri WHERE id = @p1`
	_, err = r.db.SQLserver.Conn.ExecContext(ctx, query, id)
	return err
}
