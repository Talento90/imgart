package httpservice

import (
	"encoding/json"
	"go-mage/gomage"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type EffectsControler struct {
}

func NewEffectsControler(router *httprouter.Router) {
	controller := &EffectsControler{}

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
