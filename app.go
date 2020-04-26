package main

import (
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4"
)

type App struct {
	Router *chi.Mux
	DB     *pgx.Conn
}

func (a *App) Initialize(user, password, dbname string) {}

func (a *App) Run(addr string) {}
