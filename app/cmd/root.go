package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/chbea1/resty/app/crud"
	"github.com/chbea1/resty/app/data"
	_ "github.com/go-sql-driver/mysql"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

var config data.Config

func init() {
	var host string
	var db string
	var pass string
	var user string
	var port int

	rootCmd.PersistentFlags().StringVar(&host, "host", "", "Host path for the DB connection")
	rootCmd.PersistentFlags().StringVar(&db, "db", "", "Database name")
	rootCmd.PersistentFlags().StringVarP(&pass, "password", "p", "", "Password for database authentication")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "Username for database authentication")
	rootCmd.PersistentFlags().IntVar(&port, "port", -1, "Port number for connection")

	viper.BindPFlag("table.host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("table.db", rootCmd.PersistentFlags().Lookup("db"))
	viper.BindPFlag("table.pass", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("table.user", rootCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("table.port", rootCmd.PersistentFlags().Lookup("port"))
	initConfig()

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Panic(fmt.Errorf("unable to decode config into appropriate struct: %v", err))
	}

	if err := config.Table.Valid(); err != nil {
		rootCmd.Help()
		os.Exit(0)
	}
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	viper.AddConfigPath(home)
	viper.SetConfigName(".resty")
	viper.SetConfigType("json")
	viper.MergeInConfig()
}

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
		run(cmd, args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Table.User,
		config.Table.Pass,
		config.Table.Host,
		config.Table.Port,
		config.Table.Db,
	)
	conn, dbErr := sql.Open("mysql", connection)
	if dbErr != nil {
		panic(dbErr)
	}

	statement, stErr := conn.Prepare("describe store_sgln_tag") //needs to be changed
	if stErr != nil {
		fmt.Println(stErr)
		os.Exit(1)
	}
	rows, err := statement.Query() // execute select statement
	defer statement.Close()
	if err != nil {
		os.Exit(1)
	}
	var tableSchema data.TableSchema
	var rowSchemas []data.RowSchema
	for rows.Next() {
		var field string
		var typeV string
		var null string
		var key string
		var defaultV *string
		var extra string
		err := rows.Scan(&field, &typeV, &null, &key, &defaultV, &extra)
		if err != nil {
			panic(err.Error())
		}
		schema := data.RowSchema{
			Field:   field,
			Type:    typeV,
			Null:    data.GetNull(null),
			Key:     data.GetKey(key),
			Default: defaultV,
			Extra:   extra,
		}
		rowSchemas = append(rowSchemas, schema)
	}
	tableSchema.Fields = &rowSchemas
	requestDoc := crud.CreateReadFromTableSchema(tableSchema)
	crudSpec := &crud.CrudSpec{
		ApiDescription: "Description",
		Version:        "1.0.0",
		TermsOfService: "",
		EmailContact:   "christianbeasley0@gmail.com",
		LicenseName:    "",
		LicenseURL:     "",
		Title:          "CRUD Service",
		Schemes:        []string{"http", "https"},
		Port:           "8080",
		BasePath:       "/v1",
	}
	apiSpec := crud.CreateOpenApi([]crud.ApiPath{requestDoc}, crudSpec)

	data, err := yaml.Marshal(apiSpec)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", data)

}
