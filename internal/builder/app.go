package builder

import (
	"github.com/redis/go-redis/v9"
	"github.com/zhikariz/go-commerce/internal/http/handler"
	"github.com/zhikariz/go-commerce/internal/http/router"
	"github.com/zhikariz/go-commerce/internal/repository"
	"github.com/zhikariz/go-commerce/internal/service"
	"github.com/zhikariz/go-commerce/pkg/cache"
	"github.com/zhikariz/go-commerce/pkg/encrypt"
	"github.com/zhikariz/go-commerce/pkg/route"
	"github.com/zhikariz/go-commerce/pkg/token"
	"gorm.io/gorm"
)

func BuildAppPublicRoutes(db *gorm.DB, tokenUseCase token.TokenUseCase, encryptTool encrypt.EncryptTool) []*route.Route {
	userRepository := repository.NewUserRepository(db, nil)
	userService := service.NewUserService(userRepository, tokenUseCase, encryptTool)
	userHandler := handler.NewUserHandler(userService)
	return router.AppPublicRoutes(userHandler)
}

func BuildAppPrivateRoutes(db *gorm.DB, redisDB *redis.Client, encryptTool encrypt.EncryptTool) []*route.Route {
	cacheable := cache.NewCacheable(redisDB)
	userRepository := repository.NewUserRepository(db, cacheable)
	userService := service.NewUserService(userRepository, nil, encryptTool)
	userHandler := handler.NewUserHandler(userService)
	return router.AppPrivateRoutes(userHandler)
}
