package crud

import "encoding/json"

// OpenApi - Top level struct. This is the actual open Api
// specification for the generated api.
type OpenApi struct {
	Swagger      string        `yaml:"swagger,omitempty" json:"swagger,omitempty"`
	Info         ApiInfo       `yaml:"info,omitempty" json:"info,omitempty"`
	Host         string        `yaml:"host,omitempty" json:"host,omitempty"`
	BasePath     string        `yaml:"basePath,omitempty" json:"basePath,omitempty"`
	Tags         []Tag         `json:"tags,omitempty" yaml:"tags,omitempty"`
	Schemes      []string      `yaml:"schemes,omitempty" json:"schemes,omitempty"`
	Paths        ApiPath       `yaml:"paths,omitempty" json:"paths,omitempty"`
	ExternalDocs *ExternalDocs `yaml:"externalDocs,omitempty" json:"externalDocs,omitempty"`
}

// ExternalDocs - External docs struct for openapi documentation
type ExternalDocs struct {
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
	URL         string `yaml:"url,omitempty" json:"url,omitempty"`
}

//Tag - Tag struct for open api documentation
type Tag struct {
	Name         string       `yaml:"name,omitempty" json:"name,omitempty"`
	Description  string       `yaml:"description,omitempty" json:"description,omitempty"`
	ExternalDocs ExternalDocs `yaml:"externalDocs,omitempty" json:"externalDocs,omitempty"`
}

// ApiInfo - Api Info struct for openapi documentation
type ApiInfo struct {
	Description    string      `yaml:"description,omitempty" json:"description,omitempty"`
	Version        string      `yaml:"version,omitempty" json:"version,omitempty"`
	Title          string      `yaml:"title,omitempty" json:"title,omitempty"`
	TermsOfService string      `yaml:"termsOfService,omitempty" json:"termsOfService,omitempty"`
	Contact        ContactInfo `yaml:"contact,omitempty" json:"contact,omitempty"`
	License        License     `yaml:"license,omitempty" json:"license,omitempty"`
}

// License - License struct for openapi documentation
type License struct {
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	URL  string `yaml:"url,omitempty" json:"url,omitempty"`
}

// ContactInfo - Contact info struct for openapi documentation
type ContactInfo struct {
	Email string `yaml:"email,omitempty" json:"email,omitempty"`
}

// Spec - Crud specification struct for openpi documentation
type Spec struct {
	APIDescription string `yaml:",flow"`
	Version        string
	TermsOfService string
	EmailContact   string
	LicenseName    string
	LicenseURL     string
	Title          string
	Schemes        []string
	Port           string
	BasePath       string
}

func (o OpenApi) String() string {
	ret, _ := json.Marshal(o)
	return string(ret)
}

// CreateOpenApi - Creates an openAPI specification from a list of api paths.
func CreateOpenApi(paths ApiPath, spec *Spec) *OpenApi {
	return &OpenApi{
		Swagger: "2.0",
		Info: ApiInfo{
			Description:    spec.APIDescription,
			Version:        spec.Version,
			Title:          spec.Title,
			TermsOfService: spec.TermsOfService,
			Contact: ContactInfo{
				Email: spec.EmailContact,
			},
			License: License{
				Name: spec.LicenseName,
				URL:  spec.LicenseURL,
			},
		},
		Host:     "127.0.01:" + spec.Port,
		BasePath: spec.BasePath,
		Tags: []Tag{Tag{
			Name:        "get",
			Description: "Desc",
			ExternalDocs: ExternalDocs{
				Description: "no",
				URL:         "http://localhost:8080",
			},
		}},
		Schemes:      spec.Schemes,
		Paths:        paths,
		ExternalDocs: nil,
	}
}
