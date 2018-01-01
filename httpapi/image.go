package httpapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/effect"
	"github.com/talento90/gorpo/errors"
	"github.com/talento90/gorpo/image"
)

type imagesController struct {
	service image.Service
}

func newImagesController(service image.Service) *imagesController {
	return &imagesController{
		service: service,
	}
}

func (c *imagesController) transformImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) appResponse {
	var filters []effect.Filter
	imgSrc := r.URL.Query().Get("imgSrc")
	filtersJSON := r.URL.Query().Get("effects")

	if imgSrc == "" {
		return errResponse(errors.EMalformed("Missing imgSrc query parameter", nil))
	}

	if filtersJSON != "" {
		err := json.Unmarshal([]byte(filtersJSON), &filters)

		if err != nil {
			return errResponse(errors.EMalformed("effects query parameter is malformed", err))
		}
	}

	img, err := c.service.Process(imgSrc, filters)

	if err != nil {
		return errResponse(err)
	}

	w.Header().Set("Content-Type", fmt.Sprintf("image/jpg"))

	buf := new(bytes.Buffer)
	png.Encode(buf, img)

	w.Write(buf.Bytes())

	return response(http.StatusOK, nil)
}
