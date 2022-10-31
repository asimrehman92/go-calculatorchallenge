// package Nadellain
package stores

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var (
	db  *gorm.DB
	err error
)

func DataMigration() {

	// var (
	// 	host     = settings.GetString("POSTGRES_CONFIG.POSTGRES_HOST")
	// 	port     = settings.GetString("POSTGRES_CONFIG.POSTGRES_PORT")
	// 	user     = settings.GetString("POSTGRES_CONFIG.POSTGRES_USER")
	// 	password = settings.GetString("POSTGRES_CONFIG.POSTGRES_PASSWORD")
	// 	dbname   = settings.GetString("POSTGRES_CONFIG.POSTGRES_DBNAME")
	// )

	var (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "123456789"
		dbname   = "mydatabase"
	)

	//connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//open database
	db, err = gorm.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	// defer db.Close()

	fmt.Println("Database Connected Successfully")
	db.AutoMigrate(&Clients{})
}
