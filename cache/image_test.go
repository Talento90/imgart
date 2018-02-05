package cache

import (
	"image"
	"testing"

	"github.com/talento90/imgart/errors"
	"github.com/talento90/imgart/imgart"
	"github.com/talento90/imgart/mock"
)

func TestImageGet(t *testing.T) {
	c := NewImage(mock.NewCache())

	url := "http://image2.com/test.png"

	filters := []imgart.Filter{
		imgart.Filter{ID: "resize", Parameters: map[string]interface{}{"width": 390, "height": 500}},
		imgart.Filter{ID: "rotate", Parameters: map[string]interface{}{"bgcolor": "black", "angle": 90}},
	}

	img, _, err := c.Get(url, filters)

	if !errors.Is(errors.NotFound, err) {
		t.Error("Cache must return a NotExists error", err)
	}

	err = c.Set(url, filters, "jpeg", image.NewRGBA(image.Rect(0, 0, 100, 50)))

	if err != nil {
		t.Error("Cache must set our value without error", err)
	}

	img, _, err = c.Get(url, filters)

	if err != nil || img == nil {
		t.Error("Should return a valid image", err)
	}
}

func TestGenerateHashImage(t *testing.T) {
	url := "http://image2.com/test.png"

	filters1 := []imgart.Filter{
		imgart.Filter{ID: "resize", Parameters: map[string]interface{}{"width": 390, "height": 500}},
		imgart.Filter{ID: "rotate", Parameters: map[string]interface{}{"bgcolor": "black", "angle": 90}},
	}

	filters2 := []imgart.Filter{
		imgart.Filter{ID: "resize", Parameters: map[string]interface{}{"height": 500, "width": 390}},
		imgart.Filter{ID: "rotate", Parameters: map[string]interface{}{"angle": 90, "bgcolor": "black"}},
	}

	hash, err1 := generateHash(url, filters1)
	hash2, err2 := generateHash(url, filters2)

	if err1 != nil || err2 != nil || hash != hash2 {
		t.Error("Hashes must be equal", hash, hash2)
	}

	t.Log(hash, "==", hash2)
}

func TestGenerateHashImageMustBeDifferent(t *testing.T) {
	url1 := "http://image2.com/test.png"
	url2 := "http://image.com/test1.png"

	filters := []imgart.Filter{
		imgart.Filter{ID: "rotate", Parameters: map[string]interface{}{"angle": 90, "bgcolor": "black"}},
		imgart.Filter{ID: "resize", Parameters: map[string]interface{}{"width": 390, "height": 500}},
	}

	hash, _ := generateHash(url1, filters)
	hash2, _ := generateHash(url2, filters)

	if hash == hash2 {
		t.Error("Hash must be different when urls are different")
	}

	t.Log(hash, "!=", hash2)
}
