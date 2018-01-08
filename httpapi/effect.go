package httpapi

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/gorpo"
)

type effectsController struct {
	service gorpo.EffectService
}

func newEffectsController(service gorpo.EffectService) *effectsController {
	return &effectsController{
		service: service,
	}
}

func (c *effectsController) getEffectByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) appResponse {
	id := params.ByName("id")

	effect, err := c.service.GetEffect(id)

	if err != nil {
		return errResponse(err)
	}

	return response(http.StatusOK, newEffectModel(effect))
}

func (c *effectsController) getAllEffects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) appResponse {
	effects, err := c.service.GetEffects()

	if err != nil {
		return errResponse(err)
	}

	var effectDesc = make([]effectModel, 0, len(effects))

	for _, e := range effects {
		effectDesc = append(effectDesc, newEffectModel(e))
	}

	return response(http.StatusOK, effectDesc)
}
