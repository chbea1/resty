package crud

import (
	"fmt"

	"github.com/chbea1/resty/app/mysqlspec"
)

// type Request interface{}
// type ApiPath interface {
// 	path() string
// }

type ApiPath struct {
	Get GetRequest `json:"/get,omitempty"`
}

type GetRequest struct {
	Get GetRequestBody `json:"get,omitempty"`
}

type GetRequestBody struct {
	Tags        []string       `json:"tags,omitempty"`
	Summary     string         `json:"summary,omitempty"`
	Description string         `json:"description,omitempty"`
	OperationID string         `json:"operationId,omitempty"`
	Produces    []string       `json:"produces,omitempty"`
	Parameters  []APIParameter `json:"parameters,omitempty"`
	Responses   ApiResponses   `json:"responses,omitempty"`
}

type ApiResponses struct {
	Num200 ApiResponse `json:"200,omitempty"`
}

type ApiResponse struct {
	Description string `json:"description,omitempty"`
}

type APIParameter struct {
	Name        string `json:"name,omitempty"`
	In          string `json:"in,omitempty"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Type        string `json:"type,omitempty"`
	Items       Items  `json:"items,omitempty"`
}

type Items struct {
	Type    string `json:"type,omitempty"`
	Default string `json:"default,omitempty"`
}

// CreateReadFromTableSchema - Creates an openAPI documentation READ object from a table schema.
func CreateReadFromTableSchema(schema mysqlspec.MysqlTableSchema) *ApiPath {
	var parameters []APIParameter
	for _, row := range *schema.Fields {
		parameter := APIParameter{
			Name:        row.Field,
			In:          "query",
			Description: fmt.Sprintf("Parameter restricts the result based on '%s'", row.Field),
			Required:    true,
			Type:        "string",
			Items: Items{
				Type:    "string",
				Default: "string",
			},
		}
		parameters = append(parameters, parameter)
	}
	request := GetRequest{
		Get: GetRequestBody{
			Summary:     "Get the requested data from the database",
			OperationID: "get",
			Produces:    []string{"application/json"},
			Tags:        []string{"get"},
			Description: "Get the requested data from the database",
			Parameters:  parameters,
			Responses: ApiResponses{
				Num200: ApiResponse{
					Description: "Successful get",
				},
			},
		},
	}

	read := ApiPath{Get: request}

	return &read
}

// "paths": {
//    "get": {
//       "get": {
//          "tags": [
//             "pet"
//          ],
//          "summary": "Finds Pets by status",
//          "description": "endpoint dec",
//          "operationId": "findPetsByStatus",
//          "produces": [
//             "application/json"
//          ],
//          "parameters": [
//             {
//                "name": "status",
//                "in": "query",
//                "description": "Status values that need to be considered for filter",
//                "required": true,
//                "type": "array",
//                "items": {
//                   "type": "string",
//                   "default": "available"
//                }
//             }
//          ],
//          "responses": {
//             "200": {
//                "description": "Invalid status value"
//             }
//          }
//       }
//    }
// }
