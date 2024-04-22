package main

import (
	"github.com/zhikariz/go-commerce/configs"
	"github.com/zhikariz/go-commerce/pkg/server"
)

func main() {
	_, err := configs.NewConfig(".env")
	checkError(err)

	srv := server.NewServer("app", nil, nil)
	srv.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
