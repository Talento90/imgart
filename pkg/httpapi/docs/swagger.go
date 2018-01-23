package docs

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

func GenerateOpenApi() (http.Handler, error) {
	//spec, err := loads.Analyzed(json.RawMessage([]byte(specJSON)), "")

	// if err != nil {
	// 	return nil, err
	// }

	redoc := middleware.Redoc(
		middleware.RedocOpts{SpecURL: "/static/swagger.json", Title: "gorpo"}, nil)

	return redoc, nil
}

const specJSON = `
{
	"swagger": "2.0",
	"info": {
	  "version": "1.0.0",
	  "title": "Swagger Gorpo",
	  "description": "Gorpo API documentation.",
	  "contact": {
		"email": "marcotalento90@gmail.com"
	  },
	  "license": {
		"name": "MIT"
	  }
	},
	"basePath": "/v1/api",
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
		"name": "profile",
		"description": "Image profiles"
	  },
	  {
		"name": "effect",
		"description": "Available effects"
	  },
	  {
		"name": "image",
		"description": "Image transform"
	  }
	],
	"paths": {
	  "/profiles": {
		"get": {
		  "tags": [
			"profile"
		  ],
		  "parameters": [
			{
			  "name": "limit",
			  "in": "query",
			  "description": "How many profiles to return",
			  "required": false,
			  "type": "integer",
			  "format": "int32"
			},
			{
			  "name": "skip",
			  "in": "query",
			  "description": "How many profiles to skip",
			  "required": false,
			  "type": "integer",
			  "format": "int32"
			}
		  ],
		  "description": "Returns all profiles",
		  "produces": [
			"application/json"
		  ],
		  "responses": {
			"200": {
			  "description": "A list of profiles.",
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
			"profile"
		  ],
		  "description": "Create a new profile",
		  "produces": [
			"application/json"
		  ],
		  "parameters": [
			{
			  "name": "profile",
			  "in": "body",
			  "description": "Create Profile body",
			  "required": true,
			  "schema": {
				"$ref": "#/definitions/CreateProfile"
			  }
			}
		  ],
		  "responses": {
			"201": {
			  "description": "Created profile",
			  "schema": {
				"$ref": "#/definitions/Profile"
			  }
			}
		  }
		}
	  },
	  "/profiles/{id}": {
		"get": {
		  "tags": [
			"profile"
		  ],
		  "parameters": [
			{
			  "name": "id",
			  "in": "path",
			  "required": true,
			  "description": "The id of profile",
			  "type": "string"
			}
		  ],
		  "description": "Returns a single profile",
		  "produces": [
			"application/json"
		  ],
		  "responses": {
			"200": {
			  "description": "A list of profiles.",
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
			"profile"
		  ],
		  "parameters": [
			{
			  "name": "id",
			  "in": "path",
			  "required": true,
			  "description": "The id of profile",
			  "type": "string"
			},
			{
			  "name": "profile",
			  "in": "body",
			  "description": "Update Profile body",
			  "required": true,
			  "schema": {
				"$ref": "#/definitions/UpdateProfile"
			  }
			}
		  ],
		  "description": "Returns the updated profile",
		  "produces": [
			"application/json"
		  ],
		  "responses": {
			"200": {
			  "description": "A list of profiles.",
			  "schema": {
				"type": "array",
				"items": {
				  "$ref": "#/definitions/Profile"
				}
			  }
			},
			"404": {
			  "description": "Profile not found"
			}
		  }
		},
		"delete": {
		  "tags": [
			"profile"
		  ],
		  "parameters": [
			{
			  "name": "id",
			  "in": "path",
			  "required": true,
			  "description": "The id of profile",
			  "type": "string"
			}
		  ],
		  "description": "Deletes profile by the given ID",
		  "produces": [
			"application/json"
		  ],
		  "responses": {
			"204": {
			  "description": "No content to show."
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
			"effect"
		  ],
		  "description": "Returns all available effects",
		  "produces": [
			"application/json"
		  ],
		  "responses": {
			"200": {
			  "description": "A list of effects.",
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
			"effect"
		  ],
		  "parameters": [
			{
			  "name": "id",
			  "in": "path",
			  "required": true,
			  "description": "The id of effect",
			  "type": "string"
			}
		  ],
		  "description": "Returns a single effect",
		  "produces": [
			"application/json"
		  ],
		  "responses": {
			"200": {
			  "description": "Return the effect",
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
			"image"
		  ],
		  "parameters": [
			{
			  "name": "imgSrc",
			  "in": "query",
			  "description": "Image Source",
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
		  "description": "Returns all profiles",
		  "produces": [
			"application/json"
		  ],
		  "responses": {
			"200": {
			  "description": "A list of profiles.",
			  "schema": {
				"type": "array",
				"items": {
				  "$ref": "#/definitions/Profile"
				}
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
			  "additionalProperties": true
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
	  }
	}
  }
`
