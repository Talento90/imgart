package httpapi

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDocumentationSpec(t *testing.T) {
	server := createMockServer()
	defer server.Close()

	res, err := http.Get(fmt.Sprintf("%s/api/v1/docs/swagger.json", server.URL))

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expect status code 200 and got: %d", res.StatusCode)
	}
}

func TestDocumentationRedoc(t *testing.T) {
	server := createMockServer()
	defer server.Close()

	res, err := http.Get(fmt.Sprintf("%s/api/v1/docs", server.URL))

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expect status code 200 and got: %d", res.StatusCode)
	}
}
