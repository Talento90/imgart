package httpservice

import (
	"bytes"
	"fmt"
	"go-mage/downloaders"
	"image/png"
	"log"
	"net/http"
)

func CreateServer() {
	downloader := downloaders.NewHTTPDownloader()

	http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		imgSrc := r.URL.Query().Get("imgSrc")

		img, imgType, err := downloader.DownloadImage(imgSrc)

		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", fmt.Sprintf("image/%s", imgType))

		buf := new(bytes.Buffer)
		png.Encode(buf, img)

		w.Write(buf.Bytes())
	})

	log.Fatal(http.ListenAndServe(":4005", nil))
}
