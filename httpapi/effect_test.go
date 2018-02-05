package httpapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetAllEffects(t *testing.T) {
	server := createMockServer()
	defer server.Close()

	res, err := http.Get(fmt.Sprintf("%s/api/v1/effects", server.URL))

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expect status code 200 and got: %d", res.StatusCode)
	}

	effects := []effectModel{}

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(bytes, &effects)

	if err != nil {
		t.Error(err)
	}

	numOfEffects := len(effects)

	if numOfEffects != 8 {
		t.Errorf("Expect 8 effects and got: %d", numOfEffects)
	}
}

func TestGetEffectById(t *testing.T) {
	server := createMockServer()
	defer server.Close()

	res, err := http.Get(fmt.Sprintf("%s/api/v1/effects/rotate", server.URL))

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expect status code 200 and got: %d", res.StatusCode)
	}

	effect := effectModel{}

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(bytes, &effect)

	if err != nil {
		t.Error(err)
	}

	if effect.ID != "rotate" {
		t.Errorf("Expect rotate effect and got: %s", effect.ID)
	}
}

func TestGetEffectWithWrongId(t *testing.T) {
	server := createMockServer()
	defer server.Close()

	res, err := http.Get(fmt.Sprintf("%s/api/v1/effects/xpto", server.URL))

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Expect status code 404 and got: %d", res.StatusCode)
	}
}
