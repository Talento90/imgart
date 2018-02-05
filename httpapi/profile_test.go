package httpapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/talento90/imgart/imgart"
)

func TestCreateProfile(t *testing.T) {
	server := createMockServer()
	defer server.Close()

	tt := []struct {
		name       string
		body       createProfileModel
		statusCode int
	}{
		{
			name: "Create profile",
			body: createProfileModel{ID: "test", Filters: []imgart.Filter{
				imgart.Filter{
					ID:         "rotate",
					Parameters: map[string]interface{}{"angle": 90.0},
				},
			}},
			statusCode: 201,
		},
		{
			name:       "Create invalid profile",
			body:       createProfileModel{ID: "test", Filters: []imgart.Filter{}},
			statusCode: 422,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			body, err := json.Marshal(tc.body)

			if err != nil {
				t.Error(err)
			}

			res, err := http.Post(fmt.Sprintf("%s/api/v1/profiles", server.URL), "application/json", bytes.NewBuffer(body))

			if err != nil {
				t.Error(err)
			}

			if res.StatusCode != tc.statusCode {
				t.Errorf("Expect status code %d and got: %d", tc.statusCode, res.StatusCode)
			}

			if res.StatusCode == 201 {
				res, err := http.Get(fmt.Sprintf("%s/api/v1/profiles/%s", server.URL, "test"))

				if err != nil {
					t.Error(err)
				}

				p := imgart.Profile{}

				bytes, err := ioutil.ReadAll(res.Body)

				if err != nil {
					t.Error(err)
				}

				err = json.Unmarshal(bytes, &p)

				if err != nil {
					t.Error(err)
				}

				if p.ID != "test" || p.Filters[0].ID != "rotate" {
					t.Error("Wrong profile")
				}
			}
		})
	}
}

func TestUpdateProfile(t *testing.T) {
	dep := mockDependencies()

	dep.ProfileService.Create(&imgart.Profile{
		ID: "test",
		Filters: []imgart.Filter{
			imgart.Filter{
				ID:         "rotate",
				Parameters: map[string]interface{}{"angle": 90.0},
			},
		},
	})

	handler := createRouter(dep)
	server := httptest.NewServer(handler)

	defer server.Close()

	tt := []struct {
		name       string
		profile    string
		body       updateProfileModel
		statusCode int
	}{
		{
			name:    "Update profile",
			profile: "test",
			body: updateProfileModel{
				Filters: []imgart.Filter{
					imgart.Filter{
						ID:         "gamma",
						Parameters: map[string]interface{}{"sigma": 90.0},
					},
				}},
			statusCode: 200,
		},
		{
			name:    "Update invalid profile",
			profile: "test",
			body: updateProfileModel{
				Filters: []imgart.Filter{}},
			statusCode: 422,
		},
		{
			name:       "Update profile does not exist",
			profile:    "notexists",
			statusCode: 404,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			body, err := json.Marshal(tc.body)

			if err != nil {
				t.Error(err)
			}

			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/v1/profiles/%s", server.URL, tc.profile), bytes.NewBuffer(body))

			if err != nil {
				t.Error(err)
			}

			res, err := http.DefaultClient.Do(req)

			if err != nil {
				t.Error(err)
			}

			if res.StatusCode != tc.statusCode {
				t.Errorf("Expect status code %d and got: %d", tc.statusCode, res.StatusCode)
			}

			if res.StatusCode == 200 {
				res, err := http.Get(fmt.Sprintf("%s/api/v1/profiles/%s", server.URL, "test"))

				if err != nil {
					t.Error(err)
				}

				p := imgart.Profile{}

				bytes, err := ioutil.ReadAll(res.Body)

				if err != nil {
					t.Error(err)
				}

				err = json.Unmarshal(bytes, &p)

				if err != nil {
					t.Error(err)
				}

				if p.ID != "test" || p.Filters[0].ID != "gamma" {
					t.Error("Wrong profile")
				}
			}
		})
	}
}

func TestDeleteProfile(t *testing.T) {
	dep := mockDependencies()

	dep.ProfileService.Create(&imgart.Profile{
		ID: "test",
		Filters: []imgart.Filter{
			imgart.Filter{
				ID:         "rotate",
				Parameters: map[string]interface{}{"angle": 90.0},
			},
		},
	})

	handler := createRouter(dep)
	server := httptest.NewServer(handler)

	defer server.Close()

	tt := []struct {
		name       string
		profile    string
		statusCode int
	}{
		{
			name:       "Delete profile",
			profile:    "test",
			statusCode: 200,
		},
		{
			name:       "Delete profile does not exist",
			profile:    "notexists",
			statusCode: 404,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/v1/profiles/%s", server.URL, tc.profile), nil)

			if err != nil {
				t.Error(err)
			}

			res, err := http.DefaultClient.Do(req)

			if err != nil {
				t.Error(err)
			}

			if res.StatusCode != tc.statusCode {
				t.Errorf("Expect status code %d and got: %d", tc.statusCode, res.StatusCode)
			}

			if res.StatusCode == 200 {
				res, err := http.Get(fmt.Sprintf("%s/api/v1/profiles/%s", server.URL, "test"))

				if err != nil {
					t.Error(err)
				}

				if res.StatusCode != 404 {
					t.Errorf("Expect status code 404 and got: %d", res.StatusCode)
				}
			}
		})
	}
}

func TestGetProfileById(t *testing.T) {
	dep := mockDependencies()

	dep.ProfileService.Create(&imgart.Profile{
		ID: "test",
		Filters: []imgart.Filter{
			imgart.Filter{
				ID:         "rotate",
				Parameters: map[string]interface{}{"angle": 90.0},
			},
		},
	})

	handler := createRouter(dep)
	server := httptest.NewServer(handler)

	defer server.Close()

	tt := []struct {
		name       string
		profile    string
		statusCode int
	}{
		{
			name:       "Get profile",
			profile:    "test",
			statusCode: 200,
		},
		{
			name:       "Get profile does not exist",
			profile:    "notexists",
			statusCode: 404,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := http.Get(fmt.Sprintf("%s/api/v1/profiles/%s", server.URL, tc.profile))

			if err != nil {
				t.Error(err)
			}

			if res.StatusCode != tc.statusCode {
				t.Errorf("Expect status code %d and got: %d", tc.statusCode, res.StatusCode)
			}

			if res.StatusCode == 200 {
				if err != nil {
					t.Error(err)
				}

				p := imgart.Profile{}

				bytes, err := ioutil.ReadAll(res.Body)

				if err != nil {
					t.Error(err)
				}

				err = json.Unmarshal(bytes, &p)

				if err != nil {
					t.Error(err)
				}

				if p.ID != "test" {
					t.Error("Wrong profile")
				}
			}
		})
	}
}

func TestGetProfiles(t *testing.T) {
	dep := mockDependencies()

	for i := 1; i <= 10; i++ {
		dep.ProfileService.Create(&imgart.Profile{
			ID: fmt.Sprintf("test-%d", i),
			Filters: []imgart.Filter{
				imgart.Filter{
					ID:         "rotate",
					Parameters: map[string]interface{}{"angle": 90.0},
				},
			},
		})
	}

	handler := createRouter(dep)
	server := httptest.NewServer(handler)

	defer server.Close()

	tt := []struct {
		name       string
		limit      int
		skip       int
		statusCode int
		expected   string
	}{
		{
			name:       "Get profiles",
			limit:      5,
			skip:       0,
			statusCode: 200,
			expected:   "test-5",
		},
		{
			name:       "Get profiles skip 5",
			limit:      5,
			skip:       5,
			statusCode: 200,
			expected:   "test-10",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := http.Get(fmt.Sprintf("%s/api/v1/profiles?limit=%d&skip=%d", server.URL, tc.limit, tc.skip))

			if err != nil {
				t.Error(err)
			}

			if res.StatusCode != tc.statusCode {
				t.Errorf("Expect status code %d and got: %d", tc.statusCode, res.StatusCode)
			}

			p := []imgart.Profile{}

			bytes, err := ioutil.ReadAll(res.Body)

			if err != nil {
				t.Error(err)
			}

			err = json.Unmarshal(bytes, &p)

			if err != nil {
				t.Error(err)
			}

			if len(p) != 5 {
				t.Errorf("Expected 5 profiles and got %d", len(p))
			}

			if p[tc.limit-1].ID != tc.expected {
				t.Errorf("Expected %s and got %s", tc.expected, p[tc.limit-1].ID)
			}
		})
	}
}
