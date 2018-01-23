package httpapi

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/pkg/gorpo"
	"github.com/talento90/gorpo/pkg/httpapi/docs"
	"github.com/talento90/gorpo/pkg/log"
)

// ServerDependencies contains all dependencies
type ServerDependencies struct {
	Logger         log.Logger
	ImgService     gorpo.ImageService
	ProfileService gorpo.ProfileService
}

// NewServer creates an http server
func NewServer(config *Configuration, dep *ServerDependencies) http.Server {
	router := httprouter.New()

	imgCtrl := newImagesController(dep.ImgService)
	effectCtrl := newEffectsController(dep.ImgService)
	profileCtrl := newProfilesController(dep.ProfileService)

	docs, err := docs.GenerateOpenApi()

	if err != nil {
		dep.Logger.Error("Error generating api documenation:", err)
	}

	router.ServeFiles("/static/*filepath", http.Dir("pkg/httpapi/docs"))
	router.Handler("GET", "/docs", docs)

	router.GET("/api/v1/images", loggerMiddleware(dep.Logger, responseMiddleware(imgCtrl.transformImage)))

	router.GET("/api/v1/effects", loggerMiddleware(dep.Logger, responseMiddleware(effectCtrl.getAll)))
	router.GET("/api/v1/effects/:id", loggerMiddleware(dep.Logger, responseMiddleware(effectCtrl.get)))

	router.GET("/api/v1/profiles", loggerMiddleware(dep.Logger, responseMiddleware(profileCtrl.getAll)))
	router.GET("/api/v1/profiles/:id", loggerMiddleware(dep.Logger, responseMiddleware(profileCtrl.get)))
	router.DELETE("/api/v1/profiles/:id", loggerMiddleware(dep.Logger, responseMiddleware(profileCtrl.delete)))
	router.PUT("/api/v1/profiles/:id", loggerMiddleware(dep.Logger, responseMiddleware(profileCtrl.update)))
	router.POST("/api/v1/profiles", loggerMiddleware(dep.Logger, responseMiddleware(profileCtrl.create)))

	return http.Server{
		Addr:         config.Address,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		Handler:      router,
	}
}
