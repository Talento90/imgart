package httpapi

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/pkg/gorpo"
)

type profilesController struct {
	service gorpo.ProfileService
}

func newProfilesController(service gorpo.ProfileService) *profilesController {
	return &profilesController{
		service: service,
	}
}

func (c *profilesController) getAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) appResponse {
	limit := intBinder(params.ByName("limit"), 5)
	skip := intBinder(params.ByName("skip"), 0)

	profiles, err := c.service.GetAll(limit, skip)

	if err != nil {
		return errResponse(err)
	}

	return response(http.StatusOK, profiles)
}

func (c *profilesController) get(w http.ResponseWriter, r *http.Request, params httprouter.Params) appResponse {
	id := params.ByName("id")

	profile, err := c.service.Get(id)

	if err != nil {
		return errResponse(err)
	}

	return response(http.StatusOK, profile)
}
