package service

import "github.com/anggardagasta/mini_wallet/models"

type IServiceUsersRepository interface {
	GetUserByUsername(username string) (result models.GetUserScanner, err error)
	GetUserByID(id int64) (result models.GetUserScanner, err error)
	InsertUser(input models.FormRegister) (id int64, err error)
	UpdateUser(id int64, input models.FormUpdateProfile) (err error)
}
