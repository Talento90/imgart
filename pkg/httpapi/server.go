package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/imgart/pkg/errors"
	"github.com/talento90/imgart/pkg/health"
	"github.com/talento90/imgart/pkg/imgart"
	"github.com/talento90/imgart/pkg/log"
)

// ServerDependencies contains all dependencies
type ServerDependencies struct {
	Logger         log.Logger
	ImgService     imgart.ImageService
	ProfileService imgart.ProfileService
	Health         health.Health
}

func createRouter(dep *ServerDependencies) *httprouter.Router {
	router := httprouter.New()

	imgCtrl := newImagesController(dep.ImgService, dep.ProfileService)
	effectCtrl := newEffectsController(dep.ImgService)
	profileCtrl := newProfilesController(dep.ProfileService)

	router.Handler("GET", "/health", dep.Health)

	router.GET("/api/v1/docs/swagger.json", Spec)
	router.Handler("GET", "/api/v1/docs", RedocSpec())

	router.GET("/api/v1/images", loggerMiddleware(dep.Logger, responseMiddleware(imgCtrl.transformImage)))

	router.GET("/api/v1/effects", loggerMiddleware(dep.Logger, responseMiddleware(effectCtrl.getAll)))
	router.GET("/api/v1/effects/:id", loggerMiddleware(dep.Logger, responseMiddleware(effectCtrl.get)))

	router.GET("/api/v1/profiles", loggerMiddleware(dep.Logger, responseMiddleware(profileCtrl.getAll)))
	router.GET("/api/v1/profiles/:id", loggerMiddleware(dep.Logger, responseMiddleware(profileCtrl.get)))
	router.DELETE("/api/v1/profiles/:id", loggerMiddleware(dep.Logger, responseMiddleware(profileCtrl.delete)))
	router.PUT("/api/v1/profiles/:id", loggerMiddleware(dep.Logger, responseMiddleware(profileCtrl.update)))
	router.POST("/api/v1/profiles", loggerMiddleware(dep.Logger, responseMiddleware(profileCtrl.create)))

	return router
}

// NewServer creates an http server
func NewServer(config *Configuration, dep *ServerDependencies) http.Server {
	router := createRouter(dep)

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, panic interface{}) {
		dep.Logger.Error("Panic error:", panic)

		err := appError{
			ErrorType: errors.Internal.String(),
			Message:   "Server internal error",
		}

		json, _ := json.Marshal(err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(json)
	}

	return http.Server{
		Addr:         config.Address,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		Handler:      router,
	}
}
