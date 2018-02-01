package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/imgart/pkg/errors"
	"github.com/talento90/imgart/pkg/imgart"
)

type imagesController struct {
	service        imgart.ImageService
	profileService imgart.ProfileService
}

func newImagesController(service imgart.ImageService, profile imgart.ProfileService) *imagesController {
	return &imagesController{
		service:        service,
		profileService: profile,
	}
}

func getQuality(r *http.Request) int {
	const defaultJpegQuality = 100

	h := r.Header.Get("accept")
	values := strings.Split(h, ";")

	for _, v := range values {
		if i := strings.Index(v, "q="); i > -1 {
			q, err := strconv.Atoi(v[i+2:])

			if err != nil {
				return defaultJpegQuality
			}

			return q
		}
	}

	return defaultJpegQuality
}

func (c *imagesController) transformImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) appResponse {
	var filters []imgart.Filter
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

	img, format, err := c.service.Process(imgSrc, filters)

	if err != nil {
		return errResponse(err)
	}

	q := getQuality(r)

	w.Header().Set("Content-Type", fmt.Sprintf("image/%s", format))

	bytes, err := imgart.Encode(format, img, q)

	if err != nil {
		return errResponse(err)
	}

	w.Write(bytes)

	return response(http.StatusOK, nil)
}
