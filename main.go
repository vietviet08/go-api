package main

import (
	"fmt"
	"log"

	"vietquoc/connect-db/config"
	"vietquoc/connect-db/database"
	"vietquoc/connect-db/router"
)

func main() {
	database.InitDB()

	appConfig := config.GetConfig()

	r := router.SetupRouter()

	serverAddr := fmt.Sprintf("%s:%s", appConfig.Server.Host, appConfig.Server.Port)
	fmt.Printf("Server running at http://%s\n", serverAddr)

	if err := r.Run(serverAddr); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
