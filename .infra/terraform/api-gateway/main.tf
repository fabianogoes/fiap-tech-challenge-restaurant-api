resource "aws_api_gateway_rest_api" "restaurant-api" {
  body = jsonencode({
  "openapi" : "3.0.1",
  "info" : {
    "title" : "restaurant-api",
    "version" : timestamp()
  },
  "paths" : {
    "/orders/{id}/sent-for-delivery" : {
      "put" : {
        "parameters" : [ {
          "name" : "id",
          "in" : "path",
          "required" : true,
          "schema" : {
            "type" : "string"
          }
        } ],
        "responses" : {
          "200" : {
            "description" : "200 response",
            "content" : {
              "application/json" : {
                "schema" : {
                  "$ref" : "#/components/schemas/Empty"
                }
              }
            }
          }
        },
        "security" : [ {
          "restaurant-gateway-auth" : [ ]
        } ],
        "x-amazon-apigateway-integration" : {
          "httpMethod" : "PUT",
          "uri" : "${var.api_url}/orders/{id}/sent-for-delivery",
          "responses" : {
            "default" : {
              "statusCode" : "200"
            }
          },
          "requestParameters" : {
            "integration.request.path.id" : "method.request.path.id"
          },
          "passthroughBehavior" : "when_no_match",
          "type" : "http_proxy"
        }
      }
    },
    "/orders/{id}/cancel" : {
      "put" : {
        "parameters" : [ {
          "name" : "id",
          "in" : "path",
          "required" : true,
          "schema" : {
            "type" : "string"
          }
        } ],
        "responses" : {
          "200" : {
            "description" : "200 response",
            "content" : {
              "application/json" : {
                "schema" : {
                  "$ref" : "#/components/schemas/Empty"
                }
              }
            }
          }
        },
        "security" : [ {
          "restaurant-gateway-auth" : [ ]
        } ],
        "x-amazon-apigateway-integration" : {
          "httpMethod" : "PUT",
          "uri" : "${var.api_url}/orders/{id}/cancel",
          "responses" : {
            "default" : {
              "statusCode" : "200"
            }
          },
          "requestParameters" : {
            "integration.request.path.id" : "method.request.path.id"
          },
          "passthroughBehavior" : "when_no_match",
          "type" : "http_proxy"
        }
      }
    },
    "/orders/{id}" : {
      "get" : {
        "parameters" : [ {
          "name" : "authorizationToken",
          "in" : "header",
          "required" : true,
          "schema" : {
            "type" : "string"
          }
        }, {
          "name" : "id",
          "in" : "path",
          "required" : true,
          "schema" : {
            "type" : "string"
          }
        } ],
        "responses" : {
          "200" : {
            "description" : "200 response",
            "content" : {
              "application/json" : {
                "schema" : {
                  "$ref" : "#/components/schemas/Empty"
                }
              }
            }
          }
        },
        "security" : [ {
          "restaurant-gateway-auth" : [ ]
        } ],
        "x-amazon-apigateway-integration" : {
          "httpMethod" : "GET",
          "uri" : "${var.api_url}/orders/{id}",
          "responses" : {
            "default" : {
              "statusCode" : "200"
            }
          },
          "requestParameters" : {
            "integration.request.header.authorizationToken" : "method.request.header.authorizationToken",
            "integration.request.path.id" : "method.request.path.id"
          },
          "passthroughBehavior" : "when_no_templates",
          "type" : "http_proxy"
        }
      }
    },
    "/sign-in" : {
      "post" : {
        "responses" : {
          "200" : {
            "description" : "200 response",
            "content" : {
              "application/json" : {
                "schema" : {
                  "$ref" : "#/components/schemas/Empty"
                }
              }
            }
          }
        },
        "x-amazon-apigateway-integration" : {
          "httpMethod" : "POST",
          "uri" : "${var.api_url}/customers/sign-in",
          "responses" : {
            "default" : {
              "statusCode" : "200"
            }
          },
          "passthroughBehavior" : "when_no_match",
          "type" : "http_proxy"
        }
      }
    },
    "/env" : {
      "get" : {
        "responses" : {
          "200" : {
            "description" : "200 response",
            "content" : {
              "application/json" : {
                "schema" : {
                  "$ref" : "#/components/schemas/Empty"
                }
              }
            }
          }
        },
        "security" : [ {
          "restaurant-gateway-auth" : [ ]
        } ],
        "x-amazon-apigateway-integration" : {
          "httpMethod" : "GET",
          "uri" : "${var.api_url}/env",
          "responses" : {
            "default" : {
              "statusCode" : "200"
            }
          },
          "passthroughBehavior" : "when_no_match",
          "type" : "http_proxy"
        }
      }
    },
    "/orders/{id}/ready-for-delivery" : {
      "put" : {
        "parameters" : [ {
          "name" : "id",
          "in" : "path",
          "required" : true,
          "schema" : {
            "type" : "string"
          }
        } ],
        "responses" : {
          "200" : {
            "description" : "200 response",
            "content" : {
              "application/json" : {
                "schema" : {
                  "$ref" : "#/components/schemas/Empty"
                }
              }
            }
          }
        },
        "security" : [ {
          "restaurant-gateway-auth" : [ ]
        } ],
        "x-amazon-apigateway-integration" : {
          "httpMethod" : "PUT",
          "uri" : "${var.api_url}/orders/{id}/ready-for-delivery",
          "responses" : {
            "default" : {
              "statusCode" : "200"
            }
          },
          "requestParameters" : {
            "integration.request.path.id" : "method.request.path.id"
          },
          "passthroughBehavior" : "when_no_match",
          "type" : "http_proxy"
        }
      }
    },
    "/orders/{id}/item/{idItem}" : {
      "delete" : {
        "parameters" : [ {
          "name" : "idItem",
          "in" : "path",
          "required" : true,
          "schema" : {
            "type" : "string"
          }
        }, {
          "name" : "id",
          "in" : "path",
          "required" : true,
          "schema" : {
            "type" : "string"
          }
        } ],
        "responses" : {
          "200" : {
            "description" : "200 response",
            "content" : {
              "application/json" : {
                "schema" : {
                  "$ref" : "#/components/schemas/Empty"
                }
              }
            }
          }
        },
        "security" : [ {
          "restaurant-gateway-auth" : [ ]
        } ],
        "x-amazon-apigateway-integration" : {
          "httpMethod" : "DELETE",
          "uri" : "${var.api_url}/orders/{id}/item/{idItem}",
          "responses" : {
            "default" : {
              "statusCode" : "200"
            }
          },
          "requestParameters" : {
            "integration.request.path.idItem" : "method.request.path.idItem",
            "integration.request.path.id" : "method.request.path.id"
          },
          "passthroughBehavior" : "when_no_match",
          "type" : "http_proxy"
        }
      }
    },
    "/orders/{id}/item" : {
      "post" : {
        "parameters" : [ {
          "name" : "id",
          "in" : "path",
          "required" : true,
          "schema" : {
            "type" : "string"
          }
        } ],
        "responses" : {
          "200" : {
            "description" : "200 response",
            "content" : {
              "application/json" : {
                "schema" : {
                  "$ref" : "#/components/schemas/Empty"
                }
              }
            }
          }
        },
        "security" : [ {
          "restaurant-gateway-auth" : [ ]
        } ],
        "x-amazon-apigateway-integration" : {
          "httpMethod" : "POST",
          "uri" : "${var.api_url}/orders/{id}/item",
          "responses" : {
            "default" : {
              "statusCode" : "200"
            }
          },
          "requestParameters" : {
            "integration.request.path.id" : "method.request.path.id"
          },
          "passthroughBehavior" : "when_no_match",
          "type" : "http_proxy"
        }
      }
    },
    "/" : {
      "get" : {
        "responses" : {
          "200" : {
            "description" : "200 response",
            "content" : {
              "application/json" : {
                "schema" : {
                  "$ref" : "#/components/schemas/Empty"
                }
              }
            }
          }
        },
        "x-amazon-apigateway-integration" : {
          "httpMethod" : "GET",
          "uri" : "${var.api_url}/",
          "responses" : {
            "default" : {
              "statusCode" : "200"
            }
          },
          "passthroughBehavior" : "when_no_match",
          "type" : "http"
        }
      }
    },
    "/orders" : {
      "post" : {
        "responses" : {
          "200" : {
            "description" : "200 response",
            "content" : {
              "application/json" : {
                "schema" : {
                  "$ref" : "#/components/schemas/Empty"
                }
              }
            }
          }
        },
        "security" : [ {
          "restaurant-gateway-auth" : [ ]
        } ],
        "x-amazon-apigateway-integration" : {
          "httpMethod" : "POST",
          "uri" : "${var.api_url}/orders/",
          "responses" : {
            "default" : {
              "statusCode" : "200"
            }
          },
          "passthroughBehavior" : "when_no_match",
          "type" : "http_proxy"
        }
      }
    },
    "/orders/{id}/payment/webhook" : {
      "put" : {
        "parameters" : [ {
          "name" : "id",
          "in" : "path",
          "required" : true,
          "schema" : {
            "type" : "string"
          }
        } ],
        "responses" : {
          "200" : {
            "description" : "200 response",
            "content" : {
              "application/json" : {
                "schema" : {
                  "$ref" : "#/components/schemas/Empty"
                }
              }
            }
          }
        },
        "security" : [ {
          "restaurant-gateway-auth" : [ ]
        } ],
        "x-amazon-apigateway-integration" : {
          "httpMethod" : "PUT",
          "uri" : "${var.api_url}/orders/{id}/payment/webhook",
          "responses" : {
            "default" : {
              "statusCode" : "200"
            }
          },
          "requestParameters" : {
            "integration.request.path.id" : "method.request.path.id"
          },
          "passthroughBehavior" : "when_no_match",
          "type" : "http_proxy"
        }
      }
    },
    "/orders/{id}/in-preparation" : {
      "put" : {
        "parameters" : [ {
          "name" : "id",
          "in" : "path",
          "required" : true,
          "schema" : {
            "type" : "string"
          }
        } ],
        "responses" : {
          "200" : {
            "description" : "200 response",
            "content" : {
              "application/json" : {
                "schema" : {
                  "$ref" : "#/components/schemas/Empty"
                }
              }
            }
          }
        },
        "security" : [ {
          "restaurant-gateway-auth" : [ ]
        } ],
        "x-amazon-apigateway-integration" : {
          "httpMethod" : "PUT",
          "uri" : "${var.api_url}/orders/{id}/in-preparation",
          "responses" : {
            "default" : {
              "statusCode" : "200"
            }
          },
          "requestParameters" : {
            "integration.request.path.id" : "method.request.path.id"
          },
          "passthroughBehavior" : "when_no_match",
          "type" : "http_proxy"
        }
      }
    },
    "/health" : {
      "get" : {
        "responses" : {
          "200" : {
            "description" : "200 response",
            "content" : {
              "application/json" : {
                "schema" : {
                  "$ref" : "#/components/schemas/Empty"
                }
              }
            }
          }
        },
        "x-amazon-apigateway-integration" : {
          "httpMethod" : "GET",
          "uri" : "${var.api_url}/health",
          "responses" : {
            "default" : {
              "statusCode" : "200"
            }
          },
          "passthroughBehavior" : "when_no_match",
          "type" : "http"
        }
      }
    },
    "/orders/{id}/payment" : {
      "put" : {
        "parameters" : [ {
          "name" : "id",
          "in" : "path",
          "required" : true,
          "schema" : {
            "type" : "string"
          }
        } ],
        "responses" : {
          "200" : {
            "description" : "200 response",
            "content" : {
              "application/json" : {
                "schema" : {
                  "$ref" : "#/components/schemas/Empty"
                }
              }
            }
          }
        },
        "security" : [ {
          "restaurant-gateway-auth" : [ ]
        } ],
        "x-amazon-apigateway-integration" : {
          "httpMethod" : "PUT",
          "uri" : "${var.api_url}/orders/{id}/payment",
          "responses" : {
            "default" : {
              "statusCode" : "200"
            }
          },
          "requestParameters" : {
            "integration.request.path.id" : "method.request.path.id"
          },
          "passthroughBehavior" : "when_no_match",
          "type" : "http_proxy"
        }
      }
    },
    "/orders/{id}/confirmation" : {
      "put" : {
        "parameters" : [ {
          "name" : "id",
          "in" : "path",
          "required" : true,
          "schema" : {
            "type" : "string"
          }
        } ],
        "responses" : {
          "200" : {
            "description" : "200 response",
            "content" : {
              "application/json" : {
                "schema" : {
                  "$ref" : "#/components/schemas/Empty"
                }
              }
            }
          }
        },
        "security" : [ {
          "restaurant-gateway-auth" : [ ]
        } ],
        "x-amazon-apigateway-integration" : {
          "httpMethod" : "PUT",
          "uri" : "${var.api_url}/orders/{id}/confirmation",
          "responses" : {
            "default" : {
              "statusCode" : "200"
            }
          },
          "requestParameters" : {
            "integration.request.path.id" : "method.request.path.id"
          },
          "passthroughBehavior" : "when_no_match",
          "type" : "http_proxy"
        }
      }
    },
    "/orders/{id}/delivered" : {
      "put" : {
        "parameters" : [ {
          "name" : "id",
          "in" : "path",
          "required" : true,
          "schema" : {
            "type" : "string"
          }
        } ],
        "responses" : {
          "200" : {
            "description" : "200 response",
            "content" : {
              "application/json" : {
                "schema" : {
                  "$ref" : "#/components/schemas/Empty"
                }
              }
            }
          }
        },
        "security" : [ {
          "restaurant-gateway-auth" : [ ]
        } ],
        "x-amazon-apigateway-integration" : {
          "httpMethod" : "PUT",
          "uri" : "${var.api_url}/orders/{id}/delivered",
          "responses" : {
            "default" : {
              "statusCode" : "200"
            }
          },
          "requestParameters" : {
            "integration.request.path.id" : "method.request.path.id"
          },
          "passthroughBehavior" : "when_no_match",
          "type" : "http_proxy"
        }
      }
    }
  },
  "components" : {
    "schemas" : {
      "Empty" : {
        "title" : "Empty Schema",
        "type" : "object"
      }
    },
    "securitySchemes" : {
      "restaurant-gateway-auth" : {
        "type" : "apiKey",
        "name" : "authorizationToken",
        "in" : "header",
        "x-amazon-apigateway-authtype" : "custom",
        "x-amazon-apigateway-authorizer" : {
          "authorizerUri" : "arn:aws:apigateway:${var.region}:lambda:path/2015-03-31/functions/arn:aws:lambda:${var.region}:${var.account_id}:function:fiap-tech-challenge-authorizer-lambda/invocations",
          "authorizerResultTtlInSeconds" : 0,
          "identitySource" : "method.request.header.authorizationToken",
          "type" : "request"
        }
      }
    }
  }
})


  name = "restaurant-api"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

resource "aws_api_gateway_deployment" "restaurant-api-deployment" {
  rest_api_id = aws_api_gateway_rest_api.restaurant-api.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.restaurant-api.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "restaurant-api-stage" {
  deployment_id = aws_api_gateway_deployment.restaurant-api-deployment.id
  rest_api_id   = aws_api_gateway_rest_api.restaurant-api.id
  stage_name    = "default"
}