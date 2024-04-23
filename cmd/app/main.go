package main

import (
	"github.com/zhikariz/go-commerce/configs"
	"github.com/zhikariz/go-commerce/internal/builder"
	"github.com/zhikariz/go-commerce/pkg/postgres"
	"github.com/zhikariz/go-commerce/pkg/server"
)

func main() {
	cfg, err := configs.NewConfig(".env")
	checkError(err)

	db, err := postgres.InitPostgres(&cfg.Postgres)
	checkError(err)

	publicRoutes := builder.BuildAppPublicRoutes(db)
	privateRoutes := builder.BuildAppPrivateRoutes()

	srv := server.NewServer("app", publicRoutes, privateRoutes)
	srv.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
