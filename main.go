package main

import (
	"TouristAppSolution/api/routers"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	healthcheck "github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	"github.com/tavsec/gin-healthcheck/config"
)

func main() {
    	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Server is getting started...")

	r := gin.Default()

	setupRouter := routers.SetupRouter()
	setupGroup := r.Group("/recommendation")
	setupGroup.Any("/*path", gin.WrapH(setupRouter))

	healthcheck.New(r, config.DefaultConfig(), []checks.Check{})

	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Listening on port 8000")
}
