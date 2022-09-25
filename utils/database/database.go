package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	// Importing for open connection with mysql
	"github.com/falcucci/maga-coin-api/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	env = config.Env
	DB  = initDatabase()
)

type configDatabase struct {
	name        string
	host        string
	port        string
	username    string
	password    string
	maxOpenCon  int
	maxIdleCon  int
	maxLifetime int
}

func initDatabase() *gorm.DB {
	configDatabase, err := getEnvsDatabase()
	if err != nil {
		log.Fatal("Invalid environments")
	}

	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configDatabase.username,
		configDatabase.password,
		configDatabase.host,
		configDatabase.port,
		configDatabase.name)

	db, err := gorm.Open("mysql", dbString)
	if err != nil {
		log.Fatal("Fail to open database")
	}

	db.DB().SetMaxIdleConns(configDatabase.maxIdleCon)
	db.DB().SetMaxOpenConns(configDatabase.maxOpenCon)

	if configDatabase.maxLifetime > 0 {
		db.DB().SetConnMaxLifetime(time.Millisecond * time.Duration(configDatabase.maxLifetime))
	} else {
		db.DB().SetConnMaxLifetime(0)
	}

	err = db.DB().Ping()
	if err != nil {
		log.Fatal("Fail ping in Database", err)
	}

	return db
}

func getEnvsDatabase() (configDatabase, error) {
	var err error
	var configDatabase configDatabase

	configDatabase.host = env.DatabaseHost
	configDatabase.port = env.DatabasePort
	configDatabase.name = env.DatabaseName
	configDatabase.password = env.DatabasePassword
	configDatabase.username = env.DatabaseUsername
	configDatabase.maxIdleCon = env.DatabaseMaxIdle
	configDatabase.maxLifetime = env.DatabaseIdleTimeout
	configDatabase.maxOpenCon = env.DatabaseMaxCon

	if err != nil {
		return configDatabase, err
	}
	return configDatabase, nil
}

func validateEnv(env string) string {
	value := os.Getenv(env)
	if value == "" {
		log.Fatal(fmt.Sprintf("Invalid environment: %s", env))
	}
	return value
}
