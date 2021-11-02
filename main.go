package main

import (
	"api/src/database"
	"api/src/repositories"
	"api/src/router"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	mongoURI, ok := os.LookupEnv("MONGO_URI")
	if !ok {
		mongoURI = "mongodb://localhost:27017"
	}

	client, err := database.Connect(mongoURI)
	if err != nil {
		log.Fatalf("could not connect to database: %v\n", err)
	}

	mongoDB, ok := os.LookupEnv("MONGO_DB")
	if !ok {
		mongoDB = "devbook"
	}

	repositories.Database = client.Database(mongoDB)
	repositories.UserCollection = repositories.NewUserCollection()

	apiPort, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		apiPort = 4000
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	switch os.Getenv("SERVICE") {
	case "users":
		r = router.AddUserRoutes(r)
	case "posts":
		r = router.AddPostRoutes(r)
	default:
		log.Fatal("unable to identify which service to start")
	}

	log.Printf("API listening on port %d\n", apiPort)

	if err := r.Run(fmt.Sprintf(":%d", apiPort)); err != nil {
		log.Fatalf("could not start api: %v\n", err)
	}
}
