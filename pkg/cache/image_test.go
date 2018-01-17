package cache

import (
	"testing"

	"github.com/talento90/gorpo/pkg/gorpo"
)

func TestImageHashIgnoringOrder(t *testing.T) {
	url := "http://image2.com/test.png"

	filters1 := []gorpo.Filter{
		gorpo.Filter{ID: "rotate", Parameters: map[string]interface{}{"angle": 90, "bgcolor": "black"}},
		gorpo.Filter{ID: "resize", Parameters: map[string]interface{}{"width": 390, "height": 500}},
	}

	filters2 := []gorpo.Filter{
		gorpo.Filter{ID: "resize", Parameters: map[string]interface{}{"height": 500, "width": 390}},
		gorpo.Filter{ID: "rotate", Parameters: map[string]interface{}{"bgcolor": "black", "angle": 90}},
	}

	hash, err1 := generateHash(url, filters1)
	hash2, err2 := generateHash(url, filters2)

	if err1 != nil || err2 != nil || hash != hash2 {
		t.Errorf("Hash must be equal even when order it's not the same")
	}
}

func TestImageHashDifferentUrls(t *testing.T) {
	url1 := "http://image2.com/test.png"
	url2 := "http://image.com/test1.png"

	filters1 := []gorpo.Filter{
		gorpo.Filter{ID: "rotate", Parameters: map[string]interface{}{"angle": 90, "bgcolor": "black"}},
		gorpo.Filter{ID: "resize", Parameters: map[string]interface{}{"width": 390, "height": 500}},
	}

	filters2 := []gorpo.Filter{
		gorpo.Filter{ID: "resize", Parameters: map[string]interface{}{"height": 500, "width": 390}},
		gorpo.Filter{ID: "rotate", Parameters: map[string]interface{}{"bgcolor": "black", "angle": 90}},
	}

	hash, err1 := generateHash(url1, filters1)
	hash2, err2 := generateHash(url2, filters2)

	if err1 != nil || err2 != nil || hash == hash2 {
		t.Errorf("Hash must be different when urls are different")
	}
}
