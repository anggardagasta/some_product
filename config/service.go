package config

import (
	"github.com/anggardagasta/some_product/service/delivery"
	"github.com/anggardagasta/some_product/service/repository/mysql"
	"github.com/anggardagasta/some_product/service/usecase"
)

func (c *Config) InitService() error {
	serviceUserRepo := mysql.NewServiceUsersRepository(c.DB)
	serviceUserUseCase := usecase.NewServiceUsersUsecase(serviceUserRepo)

	c.Route = delivery.Router(serviceUserUseCase)

	return nil
}