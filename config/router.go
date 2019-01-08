package config

import (
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type RouterConfig interface {
	Init() *chi.Mux
}

type router struct{}

var (
	m          *router
	routerOnce sync.Once
)

func Router() RouterConfig {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}

func (router *router) Init() *chi.Mux {
	kwController := ServiceContainer().InjectKwController()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/api", kwController.Welcome)

	return r
}
