package httpservice

import (
	"bytes"
	"fmt"
	"go-mage/downloaders"
	"image/png"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterImagesController(router *httprouter.Router) {

	router.GET("/api/v1/images", imageHandler)
}

func imageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	downloader := downloaders.NewHTTPDownloader()
	imgSrc := r.URL.Query().Get("imgSrc")

	img, imgType, err := downloader.DownloadImage(imgSrc)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", fmt.Sprintf("image/%s", imgType))

	buf := new(bytes.Buffer)
	png.Encode(buf, img)

	w.Write(buf.Bytes())
}
