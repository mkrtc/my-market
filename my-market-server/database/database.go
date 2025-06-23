package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func Init() *gorm.DB {
	user := os.Getenv("PG_USER")
	passwd := os.Getenv("PG_PASSWD")
	dbName := os.Getenv("PG_DB")
	host := os.Getenv("PG_HOST")
	port, portConvErr := strconv.Atoi(os.Getenv("PG_PORT"))
	if portConvErr != nil {
		panic(fmt.Sprintf("DB Port not inited: %v", portConvErr))
	}

	pgConfig := postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, passwd, dbName, port),
	})

	var connectionError error
	Connection, connectionError := gorm.Open(pgConfig, &gorm.Config{})

	if connectionError != nil {
		panic(fmt.Sprintf("\033[31m[error] \033[0mCannot connect to db: %v", connectionError))
	}
	fmt.Printf("\033[32m[success] \033[0mPg successfuly connected to %s\n", dbName)

	return Connection
}
