package restauranthandler

import (
	"bytes"
	"fmt"
	logger "log"

	customerrors "pavan/MAD-Assignment-1/restapplication/packages/errors"

	"github.com/xeipuuv/gojsonschema"
)

func ValidateRestaurantCreateUpdateRequest(rStr string) (bool, error) {
	//logger := loggerutils.GetLogger()
	schemaStr := `
	{
		"$schema": "http://json-schema.org/draft-04/schema#",
		"type": "object",
		"properties": {
			"id": {
				"type": "string",
				"minLength": 1,
				"maxLength": 50
			},
			"name": {
				"type": "string",
				"minLength": 1,
				"maxLength": 100
			},
			"address": {
				"type": "string",
				"minLength": 1,
				"maxLength": 100
			},
			"addressLine2": {
				"type": "string",
				"minLength": 0,
				"maxLength": 100
			},
			"url": {
				"type": "string",
				"minLength": 1,
				"maxLength": 200
			},
			"outcode": {
				"type": "string",
				"minLength": 1,
				"maxLength": 10
			},
			"postcode": {
				"type": "string",
				"minLength": 1,
				"maxLength": 10
			},
			"rating": {
				"type": "number",
				"minLength": 0,
				"maxLength": 5
			},
			"type_of_food": {
				"type": "string",
				"minLength": 1,
				"maxLength": 20
			}
	
		},
		"required": [
			"name",
			"address",
			"addressLine2",
			"url",
			"outcode",
			"postcode",
			"rating",
			"type_of_food"
		]
	}`

	schema := gojsonschema.NewStringLoader(schemaStr)
	content := gojsonschema.NewStringLoader(rStr)
	result, err := gojsonschema.Validate(schema, content)

	if err != nil {
		//logger.Fatalf("Invalid Json Schema Error: %v", err)
		return false, customerrors.InternalError(fmt.Sprintf("Invalid Json Schema Error: %v", err))
		//panic(err)
	}
	if result.Valid() {
		return true, nil
	} else {
		var buffer bytes.Buffer
		for _, resulterr := range result.Errors() {
			logger.Printf("- %s\n", resulterr)
			errString := fmt.Sprintf("Field: %s - %s, ", resulterr.Field(), resulterr.Description())
			buffer.Write([]byte(errString))

		}
		errorDesc := buffer.String()
		return false, customerrors.BadRequest(errorDesc)
	}

}
func ValidateRestaurantUpdateRequest(rStr string) (bool, error) {
	//logger := loggerutils.GetLogger()
	schemaStr := `
	{
		"$schema": "http://json-schema.org/draft-04/schema#",
		"type": "object",
		"properties": {
			"id": {
				"type": "string",
				"minLength": 1,
				"maxLength": 50
			},
			"name": {
				"type": "string",
				"minLength": 1,
				"maxLength": 100
			},
			"address": {
				"type": "string",
				"minLength": 1,
				"maxLength": 100
			},
			"addressLine2": {
				"type": "string",
				"minLength": 0,
				"maxLength": 100
			},
			"url": {
				"type": "string",
				"minLength": 1,
				"maxLength": 200
			},
			"outcode": {
				"type": "string",
				"minLength": 1,
				"maxLength": 10
			},
			"postcode": {
				"type": "string",
				"minLength": 1,
				"maxLength": 10
			},
			"rating": {
				"type": "number",
				"minLength": 0,
				"maxLength": 5
			},
			"type_of_food": {
				"type": "string",
				"minLength": 1,
				"maxLength": 20
			}
	
		},
		"required": [
			"id",
			"name",
			"address",
			"addressLine2",
			"url",
			"outcode",
			"postcode",
			"rating",
			"type_of_food"
		]
	}`

	schema := gojsonschema.NewStringLoader(schemaStr)
	content := gojsonschema.NewStringLoader(rStr)
	result, err := gojsonschema.Validate(schema, content)

	if err != nil {
		//logger.Fatalf("Invalid Json Schema Error: %v", err)
		return false, customerrors.InternalError(fmt.Sprintf("Invalid Json Schema Error: %v", err))
		//panic(err)
	}
	if result.Valid() {
		return true, nil
	} else {
		var buffer bytes.Buffer
		for _, resulterr := range result.Errors() {
			logger.Printf("- %s\n", resulterr)
			errString := fmt.Sprintf("Field: %s - %s, ", resulterr.Field(), resulterr.Description())
			buffer.Write([]byte(errString))

		}
		errorDesc := buffer.String()
		return false, customerrors.BadRequest(errorDesc)
	}

}
