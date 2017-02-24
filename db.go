package config

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func connstr() string {
	connect := fmt.Sprintf(
		"%[1]s:%[2]s@tcp(%[3]s:%[4]s)/%[5]s",
		viper.GetString("sqluser"),
		viper.GetString("sqlpass"),
		viper.GetString("sqlhost"),
		viper.GetString("sqlport"),
		viper.GetString("sqlname"),
	)

	return connect
}

func Dbconn() *sql.DB {
	config()
	connstr := connstr()
	db, err := sql.Open(viper.GetString("sqltype"), connstr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Table() string {
	config()
	return viper.GetString("table")
}
