package builder

import (
	"github.com/zhikariz/go-commerce/internal/http/handler"
	"github.com/zhikariz/go-commerce/internal/http/router"
	"github.com/zhikariz/go-commerce/internal/repository"
	"github.com/zhikariz/go-commerce/internal/service"
	"github.com/zhikariz/go-commerce/pkg/route"
	"github.com/zhikariz/go-commerce/pkg/token"
	"gorm.io/gorm"
)

func BuildAppPublicRoutes(db *gorm.DB, tokenUseCase token.TokenUseCase) []*route.Route {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, tokenUseCase)
	userHandler := handler.NewUserHandler(userService)
	return router.AppPublicRoutes(userHandler)
}

func BuildAppPrivateRoutes(db *gorm.DB) []*route.Route {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, nil)
	userHandler := handler.NewUserHandler(userService)
	return router.AppPrivateRoutes(userHandler)
}
