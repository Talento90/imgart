package httpapi

import (
	"bytes"
	"fmt"
	"image/png"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/pkg/gorpo"
)

type ImagesControler struct {
	downloader gorpo.Downloader
}

func NewImagesController(downloader gorpo.Downloader) ImagesControler {
	return &ImagesControler{
		downloader: downloader,
	}
}

func (c *ImagesControler) ImageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	imgSrc := r.URL.Query().Get("imgSrc")

	img, imgType, err := c.downloader.DownloadImage(imgSrc)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", fmt.Sprintf("image/%s", imgType))

	buf := new(bytes.Buffer)
	png.Encode(buf, img)

	w.Write(buf.Bytes())
}
