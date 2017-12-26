package httpapi

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/effect"
)

type effectsController struct {
	service effect.Service
}

func newEffectsController(service effect.Service) *effectsController {
	return &effectsController{
		service: service,
	}
}

func (c *effectsController) GetEffectByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) appResponse {
	id := params.ByName("id")

	effect, err := c.service.GetEffect(id)

	if err != nil {
		return errResponse(err)
	}

	return response(http.StatusOK, effect.Descriptor())
}

func (c *effectsController) GetAllEffects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) appResponse {
	effects, err := c.service.GetEffects()

	if err != nil {
		return errResponse(err)
	}

	var effectDesc = make([]effect.Descriptor, 0, len(effects))

	for _, e := range effects {
		effectDesc = append(effectDesc, e.Descriptor())
	}

	return response(http.StatusOK, effectDesc)
}
