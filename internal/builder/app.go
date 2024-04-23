package builder

import (
	"github.com/zhikariz/go-commerce/internal/http/handler"
	"github.com/zhikariz/go-commerce/internal/http/router"
	"github.com/zhikariz/go-commerce/internal/repository"
	"github.com/zhikariz/go-commerce/internal/service"
	"github.com/zhikariz/go-commerce/pkg/route"
	"gorm.io/gorm"
)

func BuildAppPublicRoutes(db *gorm.DB) []*route.Route {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	return router.AppPublicRoutes(userHandler)
}

func BuildAppPrivateRoutes() []*route.Route {
	return router.AppPrivateRoutes()
}
