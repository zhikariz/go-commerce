package main

import (
	"github.com/zhikariz/go-commerce/configs"
	"github.com/zhikariz/go-commerce/internal/builder"
	"github.com/zhikariz/go-commerce/pkg/cache"
	"github.com/zhikariz/go-commerce/pkg/encrypt"
	"github.com/zhikariz/go-commerce/pkg/postgres"
	"github.com/zhikariz/go-commerce/pkg/server"
	"github.com/zhikariz/go-commerce/pkg/token"
)

func main() {
	cfg, err := configs.NewConfig(".env")
	checkError(err)

	db, err := postgres.InitPostgres(&cfg.Postgres)
	checkError(err)

	redisDB := cache.InitCache(&cfg.Redis)

	tokenUseCase := token.NewTokenUseCase(cfg.JWT.SecretKey)
	encryptTool := encrypt.NewEncryptTool(cfg.Encrypt.SecretKey, cfg.Encrypt.IV)

	publicRoutes := builder.BuildAppPublicRoutes(db, tokenUseCase, encryptTool)
	privateRoutes := builder.BuildAppPrivateRoutes(db, redisDB, encryptTool)

	srv := server.NewServer("app", publicRoutes, privateRoutes, cfg.JWT.SecretKey)
	srv.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
