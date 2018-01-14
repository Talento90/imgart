package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/pkg/errors"
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
	query := r.URL.Query()
	limit := intBinder(query.Get("limit"), 5)
	skip := intBinder(query.Get("skip"), 0)

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

func (c *profilesController) delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) appResponse {
	id := params.ByName("id")

	err := c.service.Delete(id)

	if err != nil {
		return errResponse(err)
	}

	return response(http.StatusNoContent, nil)
}

func (c *profilesController) create(w http.ResponseWriter, r *http.Request, params httprouter.Params) appResponse {
	var model createProfileModel

	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		return errResponse(errors.EMalformed("Profile body is malformed", err))
	}

	defer r.Body.Close()

	profile, err := model.toProfile()

	if err != nil {
		return errResponse(err)
	}

	err = c.service.Create(profile)

	if err != nil {
		return errResponse(err)
	}

	w.Header().Set("location", fmt.Sprintf("/profiles/%s", profile.ID))

	return response(http.StatusCreated, profile)
}

func (c *profilesController) update(w http.ResponseWriter, r *http.Request, params httprouter.Params) appResponse {
	id := params.ByName("id")

	var model updateProfileModel

	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		return errResponse(errors.EMalformed("Profile body is malformed", err))
	}

	oldProfile, err := c.service.Get(id)

	if err != nil {
		return errResponse(err)
	}

	profile, err := model.toProfile(oldProfile)

	if err != nil {
		return errResponse(err)
	}

	err = c.service.Update(profile)

	if err != nil {
		return errResponse(err)
	}

	return response(http.StatusOK, profile)
}
