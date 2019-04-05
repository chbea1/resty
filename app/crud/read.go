package crud

import (
	"fmt"

	"github.com/chbea1/resty/app/data"
)

type Request interface{}
type ApiPath interface {
	path() string
}

// Read - The openAPI documentation for a Read operation.
type Read struct {
	Get GetRequest `json:"get" yaml:"get"`
}

func (r *Read) path() string {
	return "/get"
}

// GetRequest - The openAPI documentation for a GET request.
type GetRequest struct {
	Description string          `json:"description" yaml:"description"`
	Parameters  *[]APIParameter `json:"parameters" yaml:"parameters"`
}

// APIParameter - The openAPI documentation for a API parameter
type APIParameter struct {
	Name        string             `json:"name" yaml:"name"`
	In          string             `json:"in" yaml:"in"`
	Description string             `json:"description" yaml:"description"`
	Schema      APIParameterSchema `json:"schema" yaml:"schema"`
}

// APIParameterSchema - The openAPI documentation for an API parameter schema.
type APIParameterSchema struct {
	Type    string       `json:"type" yaml:"type"`
	Example *interface{} `json:"example" yaml:"example"`
}

// CreateReadFromTableSchema - Creates an openAPI documentation READ object from a table schema.
func CreateReadFromTableSchema(schema data.TableSchema) *Read {
	var parameters []APIParameter
	for _, row := range *schema.Fields {
		parameter := APIParameter{
			Name:        row.Field,
			In:          "Query",
			Description: fmt.Sprintf("Parameter restricts the result based on '%s'", row.Field),
			Schema: APIParameterSchema{
				Type:    row.Type,
				Example: nil,
			},
		}
		parameters = append(parameters, parameter)
	}
	request := GetRequest{
		Description: "Get the requested data from the database",
		Parameters:  &parameters,
	}

	read := Read{Get: request}

	return &read
}
