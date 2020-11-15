package httpapi

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/julienschmidt/httprouter"
)

var redoc = middleware.Redoc(middleware.RedocOpts{
	BasePath: "/api/v1",
	Path:     "/docs",
	SpecURL:  "/api/v1/docs/swagger.json",
	Title:    "imgart API",
}, nil)

// RedocSpec returns handler for API Redoc documentation
func RedocSpec() http.Handler {
	return redoc
}

// Spec returns the swagger 2.0 spec
func Spec(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	_, err := w.Write([]byte(specJSON))

	if err != nil {
		errResponse(err)
	}
}

const specJSON = `
{
  "swagger": "2.0",
  "info": {
    "version": "1.4.0",
    "title": "IMGART",
    "description": "IMGART it is an HTTP service for image processing based on filters and profiles. Supported filters: overlay, rotate, blur, contrast, brightness, crop, gamma.",
    "contact": {
      "email": "marcotalento90@gmail.com"
    },
    "license": {
      "name": "MIT"
    }
  },
  "basePath": "/api/v1",
  "schemes": [
    "http"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "tags": [
    {
      "name": "profiles",
      "description": "Profile is a configured set of filters that can be applied when processing images."
    },
    {
      "name": "effects",
      "description": "Effects are used to transform images."
    },
    {
      "name": "images",
      "description": "Processes images based on filters and profiles."
    }
  ],
  "paths": {
    "/profiles": {
      "get": {
        "tags": [
          "profiles"
        ],
        "parameters": [
          {
            "name": "limit",
            "in": "query",
            "description": "Number of profiles to return (max 10, default 5)",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "skip",
            "in": "query",
            "description": "Number of profiles to skip",
            "required": false,
            "type": "integer",
            "format": "int32"
          }
        ],
        "description": "Return a list of profiles",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "List of profiles",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Profile"
              }
            }
          }
        }
      },
      "post": {
        "tags": [
          "profiles"
        ],
        "description": "Creates a new profile",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "profile",
            "in": "body",
            "description": "Profile properties",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateProfile"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Returns the created profile",
            "schema": {
              "$ref": "#/definitions/Profile"
            }
          },
          "400": {
            "description": "Body is not a valid json",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Profile not found"
          },
          "422": {
            "description": "Validation error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/profiles/{id}": {
      "get": {
        "tags": [
          "profiles"
        ],
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "description": "Profile ID",
            "type": "string"
          }
        ],
        "description": "Returns a profile with the given ID",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "Return a profile",
            "schema": {
              "$ref": "#/definitions/Profile"
            }
          },
          "404": {
            "description": "Profile not found"
          }
        }
      },
      "put": {
        "tags": [
          "profiles"
        ],
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "description": "Profile ID",
            "type": "string"
          },
          {
            "name": "profile",
            "in": "body",
            "description": "Profile properties to update",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UpdateProfile"
            }
          }
        ],
        "description": "Update profile with the given ID",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "Returns the updated profile",
            "schema": {
              "$ref": "#/definitions/Profile"
            }
          },
          "400": {
            "description": "Body is not a valid json",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Profile not found"
          },
          "422": {
            "description": "Validation error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "profiles"
        ],
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "description": "Profile ID",
            "type": "string"
          }
        ],
        "description": "Delete profile with the given ID",
        "produces": [
          "application/json"
        ],
        "responses": {
          "204": {
            "description": "Deleted profile successfully"
          },
          "404": {
            "description": "Profile not found"
          }
        }
      }
    },
    "/effects": {
      "get": {
        "tags": [
          "effects"
        ],
        "description": "Returns all available effects",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "A list of effects",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Effect"
              }
            }
          }
        }
      }
    },
    "/effects/{id}": {
      "get": {
        "tags": [
          "effects"
        ],
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "description": "Effect ID",
            "type": "string"
          }
        ],
        "description": "Returns an effect with the given ID",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "Return an effect",
            "schema": {
              "$ref": "#/definitions/Effect"
            }
          },
          "404": {
            "description": "Effect not found"
          }
        }
      }
    },
    "/images": {
      "get": {
        "tags": [
          "images"
        ],
        "parameters": [
          {
            "name": "imgSrc",
            "in": "query",
            "description": "Image source url",
            "required": true,
            "type": "string"
          },
          {
            "name": "profile",
            "in": "query",
            "description": "Profile to apply",
            "required": false,
            "type": "string"
          },
          {
            "name": "filters",
            "in": "query",
            "description": "Json with filters",
            "required": false,
            "type": "string"
          }
        ],
        "description": "Process image applying the given filters",
        "produces": [
          "image/png",
          "image/jpeg"
        ],
        "responses": {
          "400": {
            "description": "Filters are not a valid json",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Image not found"
          },
          "422": {
            "description": "Filters are not valid",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Filter": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "parameters": {
          "type": "array",
          "items": {
            "type": "object",
            "additionalProperties": {
              "type": "object"
            }
          }
        }
      }
    },
    "Profile": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "created": {
          "type": "string",
          "format": "date-time"
        },
        "updated": {
          "type": "string",
          "format": "date-time"
        },
        "filters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Filter"
          }
        }
      }
    },
    "CreateProfile": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "filters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Filter"
          }
        }
      }
    },
    "UpdateProfile": {
      "type": "object",
      "properties": {
        "filters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Filter"
          }
        }
      }
    },
    "Effect": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "parameters": {
          "type": "object",
          "properties": {
            "description": {
              "type": "string"
            },
            "required": {
              "type": "boolean"
            },
            "type": {
              "type": "string"
            },
            "example": {
              "type": "string"
            },
            "default": {
              "type": "string"
            },
            "values": {
              "type": "string"
            }
          }
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "error_type": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    }
  }
}
`
