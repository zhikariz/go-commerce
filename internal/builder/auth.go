package builder

import (
	"github.com/zhikariz/go-commerce/internal/http/router"
	"github.com/zhikariz/go-commerce/pkg/route"
)

func BuildAuthPublicRoutes() []*route.Route {
	return router.AuthPublicRoutes()
}

func BuildAuthPrivateRoutes() []*route.Route {
	return router.AuthPrivateRoutes()
}
