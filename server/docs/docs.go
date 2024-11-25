// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "GO feature flag configuration API",
            "url": "https://gofeatureflag.org",
            "email": "contact@gofeatureflag.org"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/thomaspoignant/go-feature-flag/blob/main/LICENSE"
        },
        "version": "{{.Version}}",
        "x-logo": {
            "url": "https://raw.githubusercontent.com/thomaspoignant/go-feature-flag/main/logo_128.png"
        }
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "Check if the API is up and running and that the database is available.",
                "tags": [
                    "Feature Monitoring"
                ],
                "summary": "Health endpoint of the API",
                "responses": {
                    "200": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handler.successResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/flags": {
            "get": {
                "description": "GET request to get all the flags available.",
                "tags": [
                    "Feature Flag management API"
                ],
                "summary": "Return all the flags available",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.FeatureFlag"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "POST will insert in the database the new feature flag with all his properties,\nand it will add all the associated rules too.",
                "tags": [
                    "Feature Flag management API"
                ],
                "summary": "Create a new feature flag with the given configuration.",
                "parameters": [
                    {
                        "description": "Payload which represents the flag to insert",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FeatureFlag"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.FeatureFlag"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict - when trying to insert a flag with a name that already exists",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/flags/{id}": {
            "get": {
                "description": "GET all the information about a flag with a specific .",
                "tags": [
                    "Feature Flag management API"
                ],
                "summary": "Return all the information about a flag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the feature flag",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.FeatureFlag"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "PUT - Updates the flag with the given ID with what is in the payload. It will replace completely the feature flag.",
                "tags": [
                    "Feature Flag management API"
                ],
                "summary": "Updates the flag with the given ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the feature flag",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload which represents the flag to update",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FeatureFlag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.FeatureFlag"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "DELETE - Delete the flag with the given ID.",
                "tags": [
                    "Feature Flag management API"
                ],
                "summary": "Delete the flag with the given ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the feature flag",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/model.FeatureFlag"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/flags/{id}/status": {
            "patch": {
                "description": "PATCH - Update the status of the flag with the given ID",
                "tags": [
                    "Feature Flag management API"
                ],
                "summary": "Update the status of the flag with the given ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the feature flag",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The patch query to update the flag status",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FeatureFlagStatusUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.FeatureFlag"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.successResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "message": {
                    "type": "string",
                    "example": "API is up and running"
                }
            }
        },
        "model.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "errorDetails": {
                    "type": "string",
                    "example": "An error occurred"
                }
            }
        },
        "model.FeatureFlag": {
            "type": "object",
            "properties": {
                "LastModifiedBy": {
                    "type": "string"
                },
                "bucketingKey": {
                    "description": "BucketingKey defines a source for a dynamic targeting key",
                    "type": "string"
                },
                "createdDate": {
                    "type": "string"
                },
                "defaultRule": {
                    "description": "DefaultRule is the rule applied after checking that any other rules\nmatched the user.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.Rule"
                        }
                    ]
                },
                "description": {
                    "type": "string"
                },
                "disable": {
                    "description": "Disable is true if the flag is disabled.",
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "lastUpdatedDate": {
                    "type": "string"
                },
                "metadata": {
                    "description": "Metadata is a field containing information about your flag such as an issue tracker link, a description, etc ...",
                    "type": "object",
                    "additionalProperties": true
                },
                "name": {
                    "type": "string"
                },
                "targeting": {
                    "description": "Rules is the list of Rule for this flag.\nThis an optional field.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Rule"
                    }
                },
                "trackEvents": {
                    "description": "TrackEvents is false if you don't want to export the data in your data exporter.\nDefault value is true",
                    "type": "boolean"
                },
                "type": {
                    "$ref": "#/definitions/model.FlagType"
                },
                "variations": {
                    "description": "Variations are all the variations available for this flag. The minimum is 2 variations and, we don't have any max\nlimit except if the variationValue is a bool, the max is 2.",
                    "type": "object",
                    "additionalProperties": true
                },
                "version": {
                    "description": "Version (optional) This field contains the version of the flag.\nThe version is manually managed when you configure your flags and, it is used to display the information\nin the notifications and data collection.",
                    "type": "string"
                }
            }
        },
        "model.FeatureFlagStatusUpdate": {
            "type": "object",
            "properties": {
                "disable": {
                    "type": "boolean"
                }
            }
        },
        "model.FlagType": {
            "type": "string",
            "enum": [
                "boolean",
                "string",
                "integer",
                "double",
                "json"
            ],
            "x-enum-varnames": [
                "FlagTypeBoolean",
                "FlagTypeString",
                "FlagTypeInteger",
                "FlagTypeDouble",
                "FlagTypeJSON"
            ]
        },
        "model.ProgressiveRollout": {
            "type": "object",
            "properties": {
                "end": {
                    "description": "End contains what describes the end status of the rollout.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ProgressiveRolloutStep"
                        }
                    ]
                },
                "initial": {
                    "description": "Initial contains a description of the initial state of the rollout.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ProgressiveRolloutStep"
                        }
                    ]
                }
            }
        },
        "model.ProgressiveRolloutStep": {
            "type": "object",
            "properties": {
                "date": {
                    "description": "Date is the time it starts or ends.",
                    "type": "string"
                },
                "percentage": {
                    "description": "Percentage is the percentage (initial or end) for the progressive rollout",
                    "type": "number"
                },
                "variation": {
                    "description": "Variation - name of the variation for this step",
                    "type": "string"
                }
            }
        },
        "model.Rule": {
            "type": "object",
            "properties": {
                "disable": {
                    "description": "Disable indicates that this rule is disabled.",
                    "type": "boolean"
                },
                "id": {
                    "description": "Id of the rule",
                    "type": "string"
                },
                "name": {
                    "description": "Name is the name of the rule, this field is mandatory if you want\nto update the rule during scheduled rollout",
                    "type": "string"
                },
                "percentage": {
                    "description": "Percentages represents the percentage we should give to each variation.\nexample: variationA = 10%, variationB = 80%, variationC = 10%",
                    "type": "object",
                    "additionalProperties": {
                        "type": "number"
                    }
                },
                "progressiveRollout": {
                    "description": "ProgressiveRollout is your struct to configure a progressive rollout deployment of your flag.\nIt will allow you to ramp up the percentage of your flag over time.\nYou can decide at which percentage you starts with and at what percentage you ends with in your release ramp.\nBefore the start date we will serve the initial percentage and, after we will serve the end percentage.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ProgressiveRollout"
                        }
                    ]
                },
                "query": {
                    "description": "Query represents an antlr query in the nikunjy/rules format",
                    "type": "string"
                },
                "variation": {
                    "description": "VariationResult represents the variation name to use if the rule apply for the user.\nIn case we have a percentage field in the config VariationResult is ignored",
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
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "GO Feature Flag - configuration API",
	Description:      "# Introduction\n\nThis API is documented in **OpenAPI format** and describe the REST API of the **`GO Feature Flag configuration API`**.\n\nThe goal of this micro-service is to offer a way to configure your feature flags in a more centralized and convenient way than a file.\n",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}