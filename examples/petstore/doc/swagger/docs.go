// Code generated by swaggo/swag. DO NOT EDIT.

package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v2/pets": {
            "put": {
                "description": "Update an existing pet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update an existing pet",
                "parameters": [
                    {
                        "description": "Pet object that needs to be added to the store",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Pet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Pet"
                        }
                    },
                    "400": {
                        "description": "Invalid ID supplied",
                        "schema": {
                            "$ref": "#/definitions/golithgin.HttpError"
                        }
                    },
                    "404": {
                        "description": "Pet not found",
                        "schema": {
                            "$ref": "#/definitions/golithgin.HttpError"
                        }
                    },
                    "405": {
                        "description": "Validation exception",
                        "schema": {
                            "$ref": "#/definitions/golithgin.HttpError"
                        }
                    }
                }
            },
            "post": {
                "description": "Adds a new pet to the store",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Adds a new pet to the store",
                "parameters": [
                    {
                        "description": "Pet object that needs to be added to the store",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Pet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Pet"
                        }
                    },
                    "405": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/golithgin.HttpError"
                        }
                    }
                }
            }
        },
        "/v2/pets/findByStatus": {
            "get": {
                "description": "Multiple status values can be provided with comma separated strings",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Finds Pets by status",
                "parameters": [
                    {
                        "enum": [
                            "available",
                            "pending",
                            "sold"
                        ],
                        "type": "string",
                        "description": "Status values that need to be considered for filter",
                        "name": "status",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Pet"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid ID supplied",
                        "schema": {
                            "$ref": "#/definitions/golithgin.HttpError"
                        }
                    },
                    "404": {
                        "description": "Pet not found",
                        "schema": {
                            "$ref": "#/definitions/golithgin.HttpError"
                        }
                    }
                }
            }
        },
        "/v2/pets/{petId}": {
            "get": {
                "description": "Returns a single pet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Find pet by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of pet to return",
                        "name": "petId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation",
                        "schema": {
                            "$ref": "#/definitions/entity.Pet"
                        }
                    },
                    "400": {
                        "description": "Invalid ID supplied",
                        "schema": {
                            "$ref": "#/definitions/golithgin.HttpError"
                        }
                    },
                    "404": {
                        "description": "Pet not found",
                        "schema": {
                            "$ref": "#/definitions/golithgin.HttpError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a pet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Deletes a pet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "api_key",
                        "name": "api_key",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "Pet id to delete",
                        "name": "petId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Invalid ID supplied",
                        "schema": {
                            "$ref": "#/definitions/golithgin.HttpError"
                        }
                    },
                    "404": {
                        "description": "Pet not found",
                        "schema": {
                            "$ref": "#/definitions/golithgin.HttpError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Category": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.Pet": {
            "type": "object",
            "properties": {
                "category": {
                    "$ref": "#/definitions/entity.Category"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "photoUrls": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "status": {
                    "$ref": "#/definitions/entity.PetStatus"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Tag"
                    }
                }
            }
        },
        "entity.PetStatus": {
            "type": "string",
            "enum": [
                "available",
                "pending",
                "sold"
            ],
            "x-enum-varnames": [
                "AVAILABLE",
                "PENDING",
                "SOLD"
            ]
        },
        "entity.Tag": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "golithgin.HttpError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "meta": {},
                "path": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "timestamp": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
