package main

import (
	"github.com/anggardagasta/some_product/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app := config.New()
	config.Catch(app.InitEnv())
	config.Catch(app.InitMysql())
	config.Catch(app.InitService())
	config.Catch(app.Start())
}
