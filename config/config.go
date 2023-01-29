package config

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetDB() (db *sql.DB, err error) {
	con, err := sql.Open("mssql", "server=DESKTOP-5VQTP5K\\SAGAR2019;user id=sa;password=123456;database=lucifer;")
	return con, err
}
