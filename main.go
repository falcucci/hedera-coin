package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/falcucci/maga-coin-api/api/routes"
	"github.com/falcucci/maga-coin-api/models"

	utils "github.com/falcucci/maga-coin-api/utils/database"
)

func main() {
	router := routes.Configure()

	utils.DB.AutoMigrate(&models.AccountStatement{})

	log.Printf("Api started at port :%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}
