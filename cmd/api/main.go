package main

import (
	"awesomeProject/cmd/api/handlers/player"
	mongo "awesomeProject/pkg/repositories/mongo"
	playerMongo "awesomeProject/pkg/repositories/mongo/player"
	playerService "awesomeProject/pkg/services/player"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ginEngine := gin.Default()
	client, err := mongo.ConnectClient(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal(err.Error())
	}

	playerRepo := playerMongo.Repository{
		Client: client,
	}

	playerSrv := playerService.Service{
		Repo: playerRepo,
	}

	playerHandler := player.Handler{
		PlayerService: playerSrv, // Podemos asignarle al atributo PlayerService (de tipo ports.PlayerService) una variable de tipo playerService.Service porque esta ultima ya implementa la interfaz ports.PlayerService
	}

	ginEngine.POST("/players", playerHandler.CreatePlayer)

	log.Fatalln(ginEngine.Run(":8001"))
}
