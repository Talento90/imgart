package httpapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/pkg/gorpo"
)

type imagesController struct {
	service gorpo.ImageService
}

func newImagesController(service gorpo.ImageService) imagesController {
	return imagesController{
		service: service,
	}
}

func (c *imagesController) ImageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var filters []gorpo.Filter
	imgSrc := r.URL.Query().Get("imgSrc")
	filtersJSON := r.URL.Query().Get("effects")

	if imgSrc == "" {
		//toJSON(w, http.StatusBadRequest, NewApiResponse(false, "Missing imgSrc parameter", nil))
		return
	}

	if filtersJSON != "" {
		err := json.Unmarshal([]byte(filtersJSON), &filters)

		if err != nil {
			//toJSON(w, http.StatusBadRequest, NewApiResponse(false, "Error parsing filters", nil))
			return
		}
	}

	img, err := c.service.Process(imgSrc, filters)

	if err != nil {
		//toJSON(w, http.StatusInternalServerError, NewApiResponse(false, err.Error(), nil))
		return
	}

	w.Header().Set("Content-Type", fmt.Sprintf("image/png"))

	buf := new(bytes.Buffer)
	png.Encode(buf, img)

	w.Write(buf.Bytes())
}
