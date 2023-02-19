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
        "/hello": {
            "get": {
                "description": "Hello",
                "tags": [
                    "Hello"
                ],
                "summary": "Hello",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "description": "GetOrders",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetOrders"
                ],
                "summary": "GetOrders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.orderResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.orderResponse": {
            "type": "object",
            "properties": {
                "consigneeAddress": {
                    "type": "string"
                },
                "consigneeCity": {
                    "type": "string"
                },
                "consigneeCountry": {
                    "type": "string"
                },
                "consigneePostalCode": {
                    "type": "string"
                },
                "consigneeProvince": {
                    "type": "string"
                },
                "height": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "length": {
                    "type": "number"
                },
                "trackingNumber": {
                    "type": "string"
                },
                "weight": {
                    "type": "number"
                },
                "width": {
                    "type": "number"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
