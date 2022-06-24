// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/api/v1/auth/line": {
            "get": {
                "description": "Get LINE login url and register user's host to redirect back",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Get LINE login url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "localhost:3000",
                        "name": "host",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "https://access.line.me/oauth2/v2.1/authorize",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/line/callback": {
            "get": {
                "description": "Redirect user back to the website they're coming from",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Receive callback from LINE and redirect user back to the website they're coming from",
                "parameters": [
                    {
                        "type": "string",
                        "description": "123456abcdef",
                        "name": "state",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "Location": {
                                "type": "string",
                                "description": "https://localhost:3000/user"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/line/jwt": {
            "get": {
                "description": "Exchange code from line to jwt",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Exchange code from line to jwt",
                "parameters": [
                    {
                        "type": "string",
                        "description": "123456abcdef",
                        "name": "state",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "123456abcdef",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT TOKEN....",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/login": {
            "post": {
                "description": "Login email password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login email password",
                "parameters": [
                    {
                        "description": "Login body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.loginBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/me": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "JWT claim infomation",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "JWT claim infomation",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.CustomClaims"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/register": {
            "post": {
                "description": "Register with email password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register with email password",
                "parameters": [
                    {
                        "description": "Register body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.registerBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.CustomClaims": {
            "type": "object",
            "properties": {
                "aud": {
                    "description": "the ` + "`" + `aud` + "`" + ` (Audience) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.3",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "exp": {
                    "description": "the ` + "`" + `exp` + "`" + ` (Expiration Time) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.4",
                    "$ref": "#/definitions/jwt.NumericDate"
                },
                "iat": {
                    "description": "the ` + "`" + `iat` + "`" + ` (Issued At) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.6",
                    "$ref": "#/definitions/jwt.NumericDate"
                },
                "id": {
                    "type": "string"
                },
                "is_admin": {
                    "type": "boolean"
                },
                "iss": {
                    "description": "the ` + "`" + `iss` + "`" + ` (Issuer) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.1",
                    "type": "string"
                },
                "jti": {
                    "description": "the ` + "`" + `jti` + "`" + ` (JWT ID) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.7",
                    "type": "string"
                },
                "line_uid": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nbf": {
                    "description": "the ` + "`" + `nbf` + "`" + ` (Not Before) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.5",
                    "$ref": "#/definitions/jwt.NumericDate"
                },
                "password": {
                    "type": "string"
                },
                "picture_url": {
                    "type": "string"
                },
                "sub": {
                    "description": "the ` + "`" + `sub` + "`" + ` (Subject) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.2",
                    "type": "string"
                },
                "tags": {
                    "description": "Custome fields",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "auth.loginBody": {
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
        "auth.registerBody": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "jwt.NumericDate": {
            "type": "object",
            "properties": {
                "time.Time": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Be Friends API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
