package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/pkg/gorpo"
	"github.com/talento90/gorpo/pkg/effect"
)

type EffectsController struct {
	service: gorpo.EffectService 
}

func NewEffectsController(gorpo.EffectService) EffectsController{
	return &EffectsController{
		service: gorpo
	}
}

func (c *EffectsController) GetEffectById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	
	effect, err := c.service.GetEffectById(id)

	if err != nil {
		toJSON(w, err, http.StatusInternalServerError)
		return
	} 
	
	if effect != nil {
		toJSON(w, effect, http.StatusOK)
	} else {
		toJSON(w, nil, http.NotFound)
	}
}

func (c *EffectsController) GetAllEffects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	effects, err := c.service.GetAllEffects()

	if err != nil {
		toJSON(w, err, http.StatusInternalServerError)
	} else if effect != nil {
		toJSON(w, effects, http.StatusOK)
	}
}
