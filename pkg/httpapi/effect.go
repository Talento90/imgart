package httpapi

import (
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
		toJSON(w, err, http.StatusInternalServerError)
		return
	}

	if effect != nil {
		toJSON(w, effect.Descriptor(), http.StatusOK)
	} else {
		toJSON(w, nil, http.StatusNotFound)
	}
}

func (c *effectsController) GetAllEffects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	effects, err := c.service.GetEffects()

	if err != nil {
		toJSON(w, err, http.StatusInternalServerError)
	} else {
		descriptors := make([]gorpo.EffectDescriptor, len(effects))

		for _, e := range effects {
			descriptors = append(descriptors, e.Descriptor())
		}

		toJSON(w, descriptors, http.StatusOK)
	}
}
