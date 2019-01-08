package config

import (
	"log"
	"sync"

	"github.com/galuhpradipta/service-package-layout/src/controllers"
	"github.com/galuhpradipta/service-package-layout/src/infrastructures"
	"github.com/galuhpradipta/service-package-layout/src/repositories"
	"github.com/galuhpradipta/service-package-layout/src/services"
)

type IServiceContainer interface {
	InjectKwController() controllers.KwController
}

type kernel struct{}

func (k *kernel) InjectKwController() controllers.KwController {
	pgConn, err := PgInit()
	if err != nil {
		log.Panic(err)
	}
	pgHandler := &infrastructures.PgHandler{}
	pgHandler.Conn = pgConn

	kwRepository := &repositories.KwRepository{pgHandler}
	kwService := &services.KwService{kwRepository}
	kwController := controllers.KwController{kwService}

	return kwController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
