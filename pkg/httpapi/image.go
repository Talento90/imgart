package httpapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"log"
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
	imgSrc := r.URL.Query().Get("imgSrc")
	profileId := r.URL.Query().Get("profile")
	jsonProfile := r.URL.Query().Get("json")

	var profile gorpo.Profile

	println(imgSrc)
	println(profileId)
	println(jsonProfile)

	if err := json.Unmarshal([]byte(jsonProfile), &profile); err != nil {
		toJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	img, imgType, err := c.downloader.DownloadImage(imgSrc)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", fmt.Sprintf("image/%s", imgType))

	buf := new(bytes.Buffer)
	png.Encode(buf, img)

	w.Write(buf.Bytes())
}
