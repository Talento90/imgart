package httpapi

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/talento90/gorpo/pkg/health"
	"github.com/talento90/gorpo/pkg/image"
	"github.com/talento90/gorpo/pkg/log"
	"github.com/talento90/gorpo/pkg/mock"
	"github.com/talento90/gorpo/pkg/profile"
	"github.com/talento90/gorpo/pkg/repository/memory"
)

func mockDependencies() *ServerDependencies {
	imgRepository := mock.NewImageRepository()
	effectRepo := memory.NewImageRepository(imgRepository)
	imgService := image.NewService(imgRepository, effectRepo)
	profileService := profile.NewService(mock.NewProfileRepository())
	logger, _ := log.NewLogger(log.Configuration{Output: ioutil.Discard})

	dep := &ServerDependencies{
		ImgService:     imgService,
		ProfileService: profileService,
		Logger:         logger,
		Health:         health.New(),
	}

	return dep
}

func createMockServer() *httptest.Server {
	dep := mockDependencies()
	handler := createRouter(dep)

	return httptest.NewServer(handler)
}

func TestNewServer(t *testing.T) {
	dep := mockDependencies()

	srv := NewServer(&Configuration{}, dep)

	if srv.Handler == nil {
		t.Error("Expect Handler to have a valid http.Hanlder")
	}
}
