package main

import (
	"github.com/zhikariz/go-commerce/internal/builder"
	"github.com/zhikariz/go-commerce/pkg/server"
)

func main() {
	publicRoutes := builder.BuildAuthPublicRoutes()
	privateRoutes := builder.BuildAuthPrivateRoutes()

	srv := server.NewServer("auth", publicRoutes, privateRoutes)
	srv.Run()
}
