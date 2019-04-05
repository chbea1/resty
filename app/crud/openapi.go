package crud

import "github.com/fatih/structs"

type OpenApi struct {
	Swagger      string        `yaml:"swagger"`
	Info         ApiInfo       `yaml:"info"`
	Host         string        `yaml:"host"`
	BasePath     string        `yaml:"basePath"`
	Tags         []Tag         `yaml:"tags"`
	Schemes      []string      `yaml:"schemes"`
	Paths        Request       `yaml:"paths"`
	ExternalDocs *ExternalDocs `yaml:"externalDocs"`
}

type ExternalDocs struct {
	Description string `yaml:"description"`
	URL         string `yaml:"url"`
}

type Tag struct {
	Name         string       `yaml:"name"`
	Description  string       `yaml:"description"`
	ExternalDocs ExternalDocs `yaml:"externalDocs,omitempty"`
}

type ApiInfo struct {
	Description    string      `yaml:"description"`
	Version        string      `yaml:"version"`
	Title          string      `yaml:"title"`
	TermsOfService string      `yaml:"termsOfService"`
	Contact        ContactInfo `yaml:"contact"`
	License        struct {
		Name string `yaml:"name"`
		URL  string `yaml:"url"`
	} `yaml:"license"`
}

type ContactInfo struct {
	Email string `yaml:"email"`
}

type License struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type CrudSpec struct {
	ApiDescription string `yaml:",flow"`
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

func CreateOpenApi(paths []ApiPath, spec *CrudSpec) *OpenApi {
	pathsMap := make(map[string]Request)
	for _, path := range paths {
		pathsMap[path.path()] = structs.Map(path)
	}
	return &OpenApi{
		Swagger: "2.0",
		Info: ApiInfo{
			Description:    spec.ApiDescription,
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
		Host:         "127.0.01:" + spec.Port,
		BasePath:     spec.BasePath,
		Tags:         nil,
		Schemes:      spec.Schemes,
		Paths:        pathsMap,
		ExternalDocs: nil,
	}
}
