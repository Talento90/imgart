package httpapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/pkg/errors"
	"github.com/talento90/gorpo/pkg/gorpo"
)

type imagesController struct {
	service        gorpo.ImageService
	profileService gorpo.ProfileService
}

func newImagesController(service gorpo.ImageService) *imagesController {
	return &imagesController{
		service: service,
	}
}

func (c *imagesController) transformImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) appResponse {
	var filters []gorpo.Filter
	imgSrc := r.URL.Query().Get("imgSrc")
	filtersJSON := r.URL.Query().Get("filters")
	profileID := r.URL.Query().Get("profile")

	if imgSrc == "" {
		return errResponse(errors.EMalformed("Missing imgSrc query parameter", nil))
	}

	if filtersJSON != "" {
		err := json.Unmarshal([]byte(filtersJSON), &filters)

		if err != nil {
			return errResponse(errors.EMalformed("effects query parameter is malformed", err))
		}
	}

	if profileID != "" {
		profile, err := c.profileService.Get(profileID)

		if err == nil {
			filters = append(profile.Filters, filters...)
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
