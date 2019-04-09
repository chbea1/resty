package cmd

import (
	"fmt"
	"os"

	"github.com/chbea1/resty/app/crud"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "resty",
	Short: " Resty is a CRUD rest service generator for databases",
	Long: `
          Resty is a CRUD rest service generator for database tables.
          Each operation is created based off of the schema of a DB table
          C - "POST"; R - "GET"; U - "PUT"; D - "DELETE". Resty uses openAPI
          standards and automatically generates api documentation.

          Resty can be configured using the config file ~/.resty.json. For
          an example config file run rest with  --config-example.

          `,
	Run: func(cmd *cobra.Command, args []string) {
		tableShema := config.Table.TableDescription()
		apispec := crud.CreateReadFromTableSchema(tableShema)
		openapi := crud.CreateOpenApi(*apispec, &crud.Spec{
			APIDescription: " Crud api ",
			Version:        "1.0.0",
			TermsOfService: "",
			EmailContact:   "christianbeasley0@gmail.com",
			LicenseName:    "license",
			LicenseURL:     "http://localhost:8080",
			Title:          "CRUD service",
			Schemes:        []string{"http"},
			Port:           "8080",
			BasePath:       "/v1",
		})

		fmt.Println(openapi)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
