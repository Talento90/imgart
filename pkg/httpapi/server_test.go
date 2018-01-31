package httpapi

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/talento90/imgart/pkg/health"
	"github.com/talento90/imgart/pkg/image"
	"github.com/talento90/imgart/pkg/log"
	"github.com/talento90/imgart/pkg/mock"
	"github.com/talento90/imgart/pkg/profile"
	"github.com/talento90/imgart/pkg/repository/memory"
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
		Health:         health.New("imgart"),
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
