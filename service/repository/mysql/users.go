package mysql

import (
	"database/sql"
	"github.com/anggardagasta/some_product/models"
	"github.com/anggardagasta/some_product/service"
)

type serviceUsersRepository struct {
	DB *sql.DB
}

func NewServiceUsersRepository(db *sql.DB) service.IServiceUsersRepository {
	return serviceUsersRepository{DB: db}
}

func (repo serviceUsersRepository) GetUserByID(id int64) (result models.GetUserScanner, err error) {
	rows, err := repo.DB.Query(queryGetUserByID, id)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		if err := rows.Scan(&result.ID, &result.Username, &result.FullName, &result.Password, &result.Picture); err != nil {
			return result, err
		}
	}
	_ = rows.Close()
	return result, nil
}

func (repo serviceUsersRepository) GetUserByUsername(username string) (result models.GetUserScanner, err error) {
	rows, err := repo.DB.Query(queryGetUserByUsername, username)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		if err := rows.Scan(&result.ID, &result.Username, &result.FullName, &result.Password, &result.Picture); err != nil {
			return result, err
		}
	}
	_ = rows.Close()
	return result, nil
}

func (repo serviceUsersRepository) InsertUser(input models.FormRegister) (id int64, err error) {
	stmt, err := repo.DB.Prepare(queryInsertUser)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(input.Username, input.Password, input.FullName, input.Picture)
	if err != nil {
		return 0, err
	}
	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repo serviceUsersRepository) UpdateUser(id int64, input models.FormUpdateProfile) (err error) {
	stmt, err := repo.DB.Prepare(queryUpdateUser)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(input.Picture, id)
	if err != nil {
		return err
	}
	id, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
