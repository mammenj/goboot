package daos

import (
	"database/sql"
	"fmt"
	"log"
	//mysql reference to sql interfaces

	_ "github.com/go-sql-driver/mysql"
	"github.com/mammenj/goboot/config"
)

func get() *sql.DB {
	config, err := config.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true", config.User, config.Password, config.Server, config.Port, config.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return db
}
