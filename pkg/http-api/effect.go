package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/pkg/gorpo"
	"github.com/talento90/gorpo/pkg/effect"
)

type EffectsControler struct {
	service: gorpo.EffectService 
}

func NewEffectsControler(router *httprouter.Router, gorpo.EffectService) {
	controller := &EffectsControler{
		service: gorpo
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
