package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB,error) {
	
	db,err := sql.Open("mysql","root:@tcp(localhost:3036)/go_crud")
	if err != nil {
		panic(err)
	}
	return db,err
}