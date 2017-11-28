package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/merlin/merlin"
	"github.com/talento90/merlin/effects"
)

type EffectsControler struct {
	manager: effects.manager 
}

func NewEffectsControler(router *httprouter.Router) {
	controller := &EffectsControler{
		manager: effects.NewEffectManager()
	}

	router.GET("/api/v1/effect/:id", controller.GetEffectByIdHandler)
	router.GET("/api/v1/effects", controller.GetAllEffectsHandler)
}

func (c *EffectsControler) GetEffectByIdHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.Header().Set("Content-Type", "image/application/json")

	body, _ := json.Marshal(gomage.GetAllEffects)

	w.Write(body)
}

func (c *EffectsControler) GetAllEffectsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, _ := json.Marshal(gomage.GetAllEffects())

	w.Write(body)
	w.Header().Set("Content-Type", "image/application/json")
}
