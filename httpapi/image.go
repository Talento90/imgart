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

func getParameters(srv imgart.ProfileService, r *http.Request) (string, []imgart.Filter, error) {
	var filters []imgart.Filter
	imgSrc := r.URL.Query().Get("imgSrc")
	filtersJSON := r.URL.Query().Get("filters")
	profileID := r.URL.Query().Get("profile")

	if imgSrc == "" {
		return imgSrc, filters, errors.EMalformed("Missing imgSrc query parameter", nil)
	}

	if filtersJSON != "" {
		err := json.Unmarshal([]byte(filtersJSON), &filters)

		if err != nil {
			return imgSrc, filters, errors.EMalformed("effects query parameter is malformed", err)
		}
	}

	if profileID != "" {
		profile, err := srv.Get(profileID)

		if err == nil {
			filters = append(profile.Filters, filters...)
		}
	}

	return imgSrc, filters, nil
}

type imageResult struct {
	img    image.Image
	format string
	err    error
}

func processImage(ctx context.Context, srv imgart.ImageService, imgSrc string, filters []imgart.Filter) chan imageResult {
	ch := make(chan imageResult)

	go func() {
		defer close(ch)

		img, format, err := srv.Process(ctx, imgSrc, filters)

		ch <- imageResult{img: img, format: format, err: err}
	}()

	return ch
}

func (c *imagesController) transformImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) appResponse {
	imgSrc, filters, err := getParameters(c.profileService, r)

	if err != nil {
		return errResponse(err)
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		return errResponse(ctx.Err())
	case result := <-processImage(ctx, c.service, imgSrc, filters):
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
