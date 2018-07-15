package httpapi

import (
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/imgart/errors"
	"github.com/talento90/imgart/imgart"
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
	h := r.Header.Get("accept")
	values := strings.Split(h, ";")

	for _, v := range values {
		if i := strings.Index(v, "q="); i > -1 {
			q, err := strconv.Atoi(v[i+2:])

			if err != nil {
				return jpeg.DefaultQuality
			}

			return q
		}
	}

	return jpeg.DefaultQuality
}

type imageResult struct {
	img    image.Image
	format string
	err    error
}

func processImage(srv imgart.ImageService, imgSrc string, filters []imgart.Filter) chan imageResult {
	ch := make(chan imageResult)

	go func() {
		defer close(ch)
		img, format, err := srv.Process(imgSrc, filters)

		ch <- imageResult{
			img:    img,
			format: format,
			err:    err,
		}
	}()

	return ch
}

func (c *imagesController) transformImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) appResponse {
	var filters []imgart.Filter
	imgSrc := r.URL.Query().Get("imgSrc")
	filtersJSON := r.URL.Query().Get("filters")
	profileID := r.URL.Query().Get("profile")

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

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

	process := processImage(c.service, imgSrc, filters)

	select {
	case <-ctx.Done():
		return errResponse(ctx.Err())
	case result := <-process:
		if result.err != nil {
			return errResponse(result.err)
		}

		q := getQuality(r)

		w.Header().Set("Content-Type", fmt.Sprintf("image/%s", result.format))

		bytes, err := imgart.Encode(result.format, result.img, q)

		if err != nil {
			return errResponse(err)
		}

		w.Write(bytes)

		return response(http.StatusOK, nil)
	}
}
