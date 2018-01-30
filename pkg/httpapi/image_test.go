package httpapi

import (
	"fmt"
	"net/http"
	"testing"
)

func TestImageEffect(t *testing.T) {
	tt := []struct {
		name       string
		imgSrc     string
		filters    string
		profile    string
		statusCode int
	}{
		{
			name:       "Transform sucessfully",
			imgSrc:     "http://fake-image.png",
			statusCode: 200,
		},
		{
			name:       "Transform with profile",
			imgSrc:     "http://fake-image.png",
			profile:    "fake-profile",
			statusCode: 200,
		},
		{
			name:       "Transform with valid filters",
			imgSrc:     "http://fake-image.png",
			filters:    "{'id': 'rotate', parameters: { 'angle': 90.0 }}",
			statusCode: 400,
		},
		{
			name:       "Missing imgSrc",
			imgSrc:     "",
			statusCode: 400,
		},
		{
			name:       "Malformed filters",
			imgSrc:     "http://fake-image.png",
			filters:    "{'wrong': 'filter'}",
			statusCode: 400,
		},
	}

	server := createMockServer()
	defer server.Close()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := http.Get(fmt.Sprintf("%s/api/v1/images?imgSrc=%s&filters=%s&profile=%s", server.URL, tc.imgSrc, tc.filters, tc.profile))

			if err != nil {
				t.Error(err)
			}

			if res.StatusCode != tc.statusCode {
				t.Errorf("Expect status code %d and got: %d", tc.statusCode, res.StatusCode)
			}
		})
	}
}

func TestJpegQualityHeader(t *testing.T) {
	r, _ := http.NewRequest("GET", "localhost/api/v1/images", nil)
	r.Header.Set("Accept", "image/jpeg;q=60")

	q := getQuality(r)

	if q != 60 {
		t.Errorf("Expect quality 60 and got %d", q)
	}
}
