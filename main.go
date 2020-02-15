package main

import (
	"github.com/anggardagasta/mini_wallet/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app := config.New()
	config.Catch(app.InitEnv())
	config.Catch(app.InitMysql())
	config.Catch(app.InitService())
	config.Catch(app.Start())
}
