package httpapi

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/talento90/health"
	"github.com/talento90/imgart/image"
	"github.com/talento90/imgart/log"
	"github.com/talento90/imgart/mock"
	"github.com/talento90/imgart/profile"
	"github.com/talento90/imgart/repository/memory"
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
		Health:         health.New("imgart", health.Options{CheckersTimeout: time.Second}),
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
