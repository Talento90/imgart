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
	downloader gorpo.Downloader
}

func newImagesController(downloader gorpo.Downloader) imagesController {
	return imagesController{
		downloader: downloader,
	}
}

func (c *imagesController) ImageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var filters []gorpo.Filter
	imgSrc := r.URL.Query().Get("imgSrc")
	filtersJSON := r.URL.Query().Get("effects")

	if imgSrc == "" {
		toJSON(w, http.StatusBadRequest, NewApiResponse(false, "Missing imgSrc parameter", nil))
		return
	}

	if filtersJSON != "" {
		err := json.Unmarshal([]byte(filtersJSON), &filters)

		if err != nil {
			toJSON(w, http.StatusBadRequest, NewApiResponse(false, "Error parsing filters", nil))
			return
		}
	}

	img, imgType, err := c.downloader.DownloadImage(imgSrc)

	if err != nil {
		toJSON(w, http.StatusInternalServerError, NewApiResponse(false, err.Error(), nil))
		return
	}

	w.Header().Set("Content-Type", fmt.Sprintf("image/%s", imgType))

	buf := new(bytes.Buffer)
	png.Encode(buf, img)

	w.Write(buf.Bytes())
}
