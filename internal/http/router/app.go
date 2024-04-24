package router

import (
	"net/http"

	"github.com/zhikariz/go-commerce/internal/http/handler"
	"github.com/zhikariz/go-commerce/pkg/route"
)

func AppPublicRoutes(userHandler handler.UserHandler) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: userHandler.Login,
		},
	}
}

func AppPrivateRoutes(userHandler handler.UserHandler) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: userHandler.FindAllUser,
		},
	}
}
