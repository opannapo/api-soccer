package app

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

//SetupConfig --> Function Setup Config
func SetupConfig(arg string) {
	var configFile string
	switch arg {
	case "local":
		configFile = "./config/config.json"
	case "dev":
		configFile = "./config/config-dev.json"
	case "prod":
		configFile = "./config/config-prod.json"
	}

	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	println("Config Env : " + arg)
	println("Config File : " + configFile)
	println("Config Port : " + viper.GetString("server.address"))
}

//SetupDbConnection --> Function to Setup DB Connection
func SetupDbConnection() *gorm.DB {
	println("setupDbConnection")
	user := viper.GetString("database.user")
	password := viper.GetString("database.pass")
	dbname := viper.GetString("database.name")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open("mysql", connectionString)
	db.LogMode(true)
	if err != nil {
		panic(err)
	}

	return db
}
