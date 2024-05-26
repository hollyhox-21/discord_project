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
        "/v1/message/send": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "chat_service"
                ],
                "summary": "отправка сообщения",
                "operationId": "MessageSend",
                "parameters": [
                    {
                        "enum": [
                            "unknown",
                            "server",
                            "user"
                        ],
                        "type": "string",
                        "description": "string enums",
                        "name": "typeMsg",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.MessageSendRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errorResp"
                        }
                    }
                }
            }
        },
        "/v1/message/server/history/{server_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat_service"
                ],
                "summary": "Получение истории сообщений сервера",
                "operationId": "GetServerChatHistory",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "server id",
                        "name": "server_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.GetServerChatHistoryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errorResp"
                        }
                    }
                }
            }
        },
        "/v1/message/user/history/{user_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat_service"
                ],
                "summary": "Получение истории приватного чата",
                "operationId": "GetPrivetChatHistory",
                "parameters": [
                    {
                        "type": "string",
                        "description": " id user recipient",
                        "name": "recipient_user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.GetPrivetChatHistoryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errorResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.GetPrivetChatHistoryResponse": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "api.GetServerChatHistoryResponse": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "api.MessageSendRequest": {
            "type": "object",
            "properties": {
                "consumerId": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "producerId": {
                    "type": "integer"
                }
            }
        },
        "api.errorResp": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "ChatService",
	Description:      "Сервис отправки сообщений",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}