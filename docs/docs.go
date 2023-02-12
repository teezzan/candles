// Code generated by swaggo/swag. DO NOT EDIT
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
        "/data": {
            "get": {
                "description": "The endpoint returns the OHLC points for a particular Symbol for  the given time range",
                "produces": [
                    "application/json"
                ],
                "summary": "returns the OHLC points for the given time range",
                "parameters": [
                    {
                        "type": "string",
                        "example": "BTC",
                        "description": "This is the symbol of the OHLC token",
                        "name": "symbol",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "10344553332",
                        "description": "UNIX time representation of the start time",
                        "name": "from",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "101019283847",
                        "description": "UNIX time representation of the end time",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "page of response",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 5,
                        "description": "Number of OHLC datapoints per page",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/data.GetOHLCResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "The endpoint takes a small CSV file upload and processes it. Max file size is 30MB.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Takes a CSV file upload and processes it",
                "parameters": [
                    {
                        "type": "file",
                        "description": "account image",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/generate_url": {
            "get": {
                "description": "The endpoint generates a pre-signed URL for the given file name for uploading on S3, It supports huge files",
                "summary": "Generates a pre-signed URL for the given file name for uploading on S3",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/data.GeneratePresignedURLResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "data.GeneratePresignedURLResponse": {
            "type": "object",
            "properties": {
                "filename": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "data.GetOHLCResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/data.OHLC"
                    }
                },
                "page": {
                    "type": "integer"
                }
            }
        },
        "data.OHLC": {
            "type": "object",
            "properties": {
                "close": {
                    "type": "number"
                },
                "high": {
                    "type": "number"
                },
                "low": {
                    "type": "number"
                },
                "open": {
                    "type": "number"
                },
                "symbol": {
                    "type": "string"
                },
                "unix": {
                    "type": "integer"
                }
            }
        },
        "httputil.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Candles API",
	Description:      "This is API specification for Candels, a OHLC data API platform.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}