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

func (c *effectsController) GetEffectById(w http.ResponseWriter, r *http.Request, params httprouter.Params) AppResponse {
	id := params.ByName("id")

	effect, err := c.service.GetEffect(id)

	if err != nil {
		return Response(http.StatusInternalServerError, err)
	}

	if effect == nil {
		return Response(http.StatusNotFound, fmt.Sprintf("Effect %s does not exists.", id))
	}

	return Response(http.StatusOK, effect)
}

func (c *effectsController) GetAllEffects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) AppResponse {
	effects, err := c.service.GetEffects()

	if err != nil {
		return Response(http.StatusInternalServerError, err)
	}

	return Response(http.StatusOK, effects)
}
