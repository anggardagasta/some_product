package config

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"net/http"
)

type Config struct {
	Route *chi.Mux
	DB    *sql.DB
}

func New() Config {
	return Config{}
}

func (c *Config) InitEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}

func Catch(err error) {
	if err != nil {
		panic(err)
	}
}

func (c *Config) Start() error {
	fmt.Print("Listening to port 8082")
	err := http.ListenAndServe(":8082", c.Route)
	if err != nil {
		return err
	}
	return nil
}
