package httpapi

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/imgart/pkg/imgart"
)

type effectsController struct {
	service imgart.ImageService
}

func newEffectsController(service imgart.ImageService) *effectsController {
	return &effectsController{
		service: service,
	}
}

func (c *effectsController) get(w http.ResponseWriter, r *http.Request, params httprouter.Params) appResponse {
	id := params.ByName("id")

	effect, err := c.service.Effect(id)

	if err != nil {
		return errResponse(err)
	}

	return response(http.StatusOK, newEffectModel(effect))
}

func (c *effectsController) getAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) appResponse {
	effects, err := c.service.Effects()

	if err != nil {
		return errResponse(err)
	}

	var effectDesc = make([]effectModel, 0, len(effects))

	for _, e := range effects {
		effectDesc = append(effectDesc, newEffectModel(e))
	}

	return response(http.StatusOK, effectDesc)
}
