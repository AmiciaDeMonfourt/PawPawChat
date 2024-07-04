// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

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
        "/api/user": {
            "get": {
                "description": "User info",
                "summary": "User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.GetUserInfoResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.HTTPError"
                        }
                    }
                }
            }
        },
        "/signin": {
            "post": {
                "description": "Authorization",
                "summary": "Sign in",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Credentials",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/web.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.SignInResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.HTTPError"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "Registration",
                "summary": "Sign up",
                "parameters": [
                    {
                        "description": "Credentials",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/web.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/web.SignUpResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/response.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.HTTPError"
                        }
                    }
                }
            }
        },
        "/{username}": {
            "get": {
                "description": "The page received basic information about the user",
                "summary": "Profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.ProfileResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "second_name": {
                    "type": "string"
                }
            }
        },
        "response.HTTPError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "web.GetUserInfoResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/domain.User"
                }
            }
        },
        "web.ProfileResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/domain.User"
                }
            }
        },
        "web.SignInRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "web.SignInResponse": {
            "type": "object",
            "properties": {
                "token_string": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/domain.User"
                }
            }
        },
        "web.SignUpRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "second_name": {
                    "type": "string"
                }
            }
        },
        "web.SignUpResponse": {
            "type": "object",
            "properties": {
                "token_string": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/domain.User"
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
