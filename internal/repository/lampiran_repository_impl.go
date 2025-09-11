package repository

import (
	"context"
	"refactory-simpers-au/config"
	"refactory-simpers-au/database"
	"refactory-simpers-au/internal/entity"
	"refactory-simpers-au/internal/model"

	"github.com/georgysavva/scany/v2/sqlscan"
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
	//sqlscan, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT l.id, l.created_at, l.updated_at, l.personel_id, l.dokumen_id, l.kategori, l.link, l.nama, l.keterangan, l.tipe , p.nrp 
	FROM lampiran l	 
	LEFT JOIN personel p ON p.id = l.personel_id 
	WHERE l.id = @p1`

	err = sqlscan.Get(ctx, r.db.SQLserver.Conn, &out, query, id)
	if sqlscan.NotFound(err) {
		return out, nil
	}
	return out, err
}

func (r *LampiranRepoImpl) GetByNRP(nrp string) (out entity.LampiranWithNRP, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT l.id, l.created_at, l.updated_at, l.personel_id, l.dokumen_id, l.kategori, l.link, l.nama, l.keterangan, l.tipe , p.nrp 
	FROM lampiran l 
	LEFT JOIN personel p ON p.id = l.personel_id 
	WHERE p.nrp = @p1`

	err = sqlscan.Get(ctx, r.db.SQLserver.Conn, &out, query, nrp)
	if sqlscan.NotFound(err) {
		return out, nil
	}
	return out, err
}

func (r *LampiranRepoImpl) Create(in model.Lampiran) (id int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `
		INSERT INTO lampiran (personel_id, dokumen_id, kategori, link, nama, keterangan, tipe)
		VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7)
		RETURNING id`

	_, err = r.db.SQLserver.Conn.ExecContext(ctx, query, in.PersonelID, in.DokumenID, in.Kategori, in.Link, in.Nama, in.Keterangan, in.Tipe)
	return 0, err

}

func (r *LampiranRepoImpl) GetAll() (out []model.Lampiran, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `SELECT id, created_at, updated_at, personel_id, dokumen_id, kategori, link, nama, keterangan, tipe FROM lampiran ORDER BY id DESC`

	err = sqlscan.Select(ctx, r.db.SQLserver.Conn, &out, query)
	return out, err
}

func (r *LampiranRepoImpl) Update(in model.Lampiran) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `UPDATE lampiran SET 
	personel_id = @p2,
	dokumen_id = @p3,
	kategori = @p4, 
	link = @p5, 
	nama = @p6, 
	keterangan = @p7, 
	tipe = @p8 
	WHERE id=@p1`

	_, err = r.db.SQLserver.Conn.ExecContext(ctx, query, in.ID, in.PersonelID, in.DokumenID, in.Kategori, in.Link, in.Nama, in.Keterangan, in.Tipe)
	return err
}

func (r *LampiranRepoImpl) DeleteByID(id int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.env.DB.Timeout)
	defer cancel()

	query := `DELETE FROM lampiran WHERE id = @p1`
	_, err = r.db.SQLserver.Conn.ExecContext(ctx, query, id)
	return err
}
