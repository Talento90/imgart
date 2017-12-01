package httpapi

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/pkg/gorpo"
)

type effectsController struct {
	service gorpo.EffectService
}

func newEffectsController(service gorpo.EffectService) effectsController {
	return effectsController{
		service: service,
	}
}

func (c *effectsController) GetEffectById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	effect, err := c.service.GetEffect(id)

	if err != nil {
		toJSON(w, http.StatusInternalServerError, NewApiResponse(false, err.Error(), nil))
		return
	}

	if effect == nil {
		toJSON(w, http.StatusNotFound, NewApiResponse(false, fmt.Sprintf("Effect %s does not exists.", id), effect))
		return
	}

	toJSON(w, http.StatusOK, NewApiResponse(true, "", effect))
}

func (c *effectsController) GetAllEffects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	effects, err := c.service.GetEffects()

	if err != nil {
		toJSON(w, http.StatusInternalServerError, NewApiResponse(false, err.Error(), nil))
		return
	}

	toJSON(w, http.StatusOK, NewApiResponse(true, "", effects))
}
