package config

import (
	"sync"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type RouterConfig interface {
	Run() *echo.Echo
}

type router struct{}

var (
	m          *router
	routerOnce sync.Once
)

func EchoRouter() RouterConfig {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}

func (router *router) Run() *echo.Echo {
	kwController := ServiceContainer().InjectKwController()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderCookie},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	r := e.Group("/api")

	r.GET("", kwController.Welcome)

	e.Logger.Fatal(e.Start(":45004"))

	return e
}
