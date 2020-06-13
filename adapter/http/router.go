package http

import "github.com/labstack/echo"

type router struct {
	handler Handler
}

type ServiceRouter interface {
	Router(e *echo.Echo)
}

func NewServiceRouter(h Handler) ServiceRouter {
	return &router{h}
}

func (r *router) Router(e *echo.Echo) {
	routerGroup := e.Group("/api/v1/stock")
	routerGroup.POST("/create_stock", r.handler.CreateStock)
}
