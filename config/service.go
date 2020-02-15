package config

import (
	"github.com/anggardagasta/mini_wallet/service/delivery"
	"github.com/anggardagasta/mini_wallet/service/repository/mysql"
	"github.com/anggardagasta/mini_wallet/service/usecase"
)

func (c *Config) InitService() error {
	serviceUserRepo := mysql.NewServiceUsersRepository(c.DB)
	serviceUserUseCase := usecase.NewServiceUsersUsecase(serviceUserRepo)

	c.Route = delivery.Router(serviceUserUseCase)

	return nil
}