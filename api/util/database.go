package util

import (
	"database/sql"

	"github.com/spf13/viper"
)

// ConnectDB function
func ConnectDB() (*sql.DB, error) {
	dbusername := viper.GetString("DB_USERNAME")
	dbpassword := viper.GetString("DB_PASSWORD")
	dbhost := viper.GetString("DB_HOST")
	dbport := viper.GetString("DB_PORT")
	dbname := viper.GetString("DB_NAME")

	db, err := sql.Open("mysql", dbusername+":"+dbpassword+"@tcp("+dbhost+":"+dbport+")/"+dbname)

	return db, err
}
