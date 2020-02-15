package service

import "github.com/anggardagasta/some_product/models"

type IServiceUsersUseCase interface {
	RegisterUser(form models.FormRegister) (result models.AuthResult, err error)
	Auth(form models.FormAuth) (result models.AuthResult, err error)
	GetProfile(id int64) (result models.GetProfileResult, err error)
	UpdateProfile(id int64, form models.FormUpdateProfile) (err error)
}
