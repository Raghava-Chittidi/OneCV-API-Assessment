package main

import (
	"Raghava/OneCV-Assignment/internal/database"
	"Raghava/OneCV-Assignment/internal/router"
	"Raghava/OneCV-Assignment/internal/util"

	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Setup()
	fmt.Println("Listening on port 8000 at http://localhost:8000!")

	_, err := database.ConnectToDB()
	if err != nil {
		log.Fatalln("Cannot connect to database", err)
	}

	err = util.Migrate()
	if err != nil {
		log.Fatalln("Cannot migrate models", err)
	}

	// Start server after connecting to db and migrating
	log.Fatalln(http.ListenAndServe(":8000", r))
}
