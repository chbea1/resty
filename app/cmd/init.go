package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/chbea1/resty/app/data"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var config data.Config

// init - Reads both the config file and command line arguments
// and merges them togeather.
func init() {
	var host string
	var db string
	var pass string
	var user string
	var port int
	initConfig()
	rootCmd.PersistentFlags().StringVar(&host, "host", "", "Host path for the DB connection")
	rootCmd.PersistentFlags().StringVar(&db, "db", "", "Database name")
	rootCmd.PersistentFlags().StringVar(&db, "name", "", "table name")
	rootCmd.PersistentFlags().StringVarP(&pass, "password", "p", "", "Password for database authentication")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "Username for database authentication")
	rootCmd.PersistentFlags().IntVar(&port, "port", -1, "Port number for connection")

	viper.BindPFlag("table.host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("table.db", rootCmd.PersistentFlags().Lookup("db"))
	viper.BindPFlag("table.pass", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("table.user", rootCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("table.port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("table.name", rootCmd.PersistentFlags().Lookup("name"))
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Panic(fmt.Errorf("unable to decode config into appropriate struct: %v", err))
	}

	fmt.Println(fmt.Sprintf("Config : %+v", viper.Get("table.name")))
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
