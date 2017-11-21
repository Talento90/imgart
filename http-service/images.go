package httpservice

import (
	"bytes"
	"fmt"
	"go-mage/downloaders"
	"go-mage/gomage"
	"image/png"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ImagesControler struct {
	downloader gomage.Downloader
}

func NewImagesController(router *httprouter.Router) {
	controller := &ImagesControler{
		downloader: downloaders.NewHTTPDownloader(),
	}

	router.GET("/api/v1/images", controller.ImageHandler)
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
