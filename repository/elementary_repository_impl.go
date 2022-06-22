package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"go-sisko/helper"
	"go-sisko/model/domain"
)

type ElementaryRepositoryImpl struct {
}

func NewElementaryRepository() ElementaryRepository {
	return &ElementaryRepositoryImpl{}
}

func (e ElementaryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, elementary domain.Elementary) domain.Elementary {
	SQL := "insert into elementarys(name,alamat,tanggal_lahir,tempat_lahir,jenis_kelamin,agama,golongan_darah,no_telepon) values (?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, elementary.Name, elementary.Alamat, elementary.Tanggal_lahir, elementary.Tempat_lahir, elementary.Jenis_kelamin, elementary.Agama, elementary.Golongan_darah, elementary.No_telepon)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	elementary.Id = int(id)
	return elementary
}

func (e ElementaryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, elementary domain.Elementary) domain.Elementary {
	SQL := "update elementarys set name = ?,alamat = ?,tanggal_lahir = ?,tempat_lahir = ?,jenis_kelamin = ?,agama = ?,golongan_darah = ?,no_telepon = ?  where id = ?"
	_, err := tx.ExecContext(ctx, SQL, elementary.Name, elementary.Alamat, elementary.Tanggal_lahir, elementary.Tempat_lahir, elementary.Jenis_kelamin, elementary.Agama, elementary.Golongan_darah, elementary.No_telepon, elementary.Id)
	helper.PanicIfError(err)

	return elementary
}

func (e ElementaryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, elementary domain.Elementary) {
	SQL := "delete from elementarys where id = ?"
	_, err := tx.ExecContext(ctx, SQL, elementary.Id)
	helper.PanicIfError(err)
}

func (e ElementaryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, elementaryId int) (domain.Elementary, error) {
	logrus.Info("Elementary Repository Find By Id Start")
	SQL := "select id, name ,alamat, tanggal_lahir,tempat_lahir,jenis_kelamin,agama,golongan_darah,no_telepon from elementarys where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, elementaryId)
	helper.PanicIfError(err)
	defer rows.Close()

	elementary := domain.Elementary{}
	if rows.Next() {
		err := rows.Scan(&elementary.Id, &elementary.Name, &elementary.Alamat, &elementary.Tanggal_lahir, &elementary.Tempat_lahir, &elementary.Jenis_kelamin, &elementary.Agama, &elementary.Golongan_darah, &elementary.No_telepon)
		helper.PanicIfError(err)
		return elementary, nil
	} else {
		return elementary, errors.New("elementary is not found")
	}
}

func (e ElementaryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Elementary {
	logrus.Info("Elementary Repository Find All Start")
	SQL := "select id,name,alamat,tanggal_lahir,tempat_lahir,jenis_kelamin,agama,golongan_darah,no_telepon from elementarys"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var elementarys []domain.Elementary
	for rows.Next() {
		elementary := domain.Elementary{}
		err := rows.Scan(&elementary.Id, &elementary.Name, &elementary.Alamat, &elementary.Tanggal_lahir, &elementary.Tempat_lahir, &elementary.Jenis_kelamin, &elementary.Agama, &elementary.Golongan_darah, &elementary.No_telepon)
		helper.PanicIfError(err)
		elementarys = append(elementarys, elementary)
	}
	return elementarys
}
