package main

import (
	"log"

	"github.com/abhaylpu/CIGO0123/tree/main/(Week-2%20Assignment)/routers"
	"github.com/abhaylpu/CIGO0123/tree/main/(Week-2%20Assignment)/database"
)

func main() {
	database.Setup()                    //establish the datatbase connection
	engine := routers.Router()          //get the customized engine
	err := engine.Run("127.0.0.1:8080") //start the engine
	if err != nil {
		log.Fatal(err)
	}
}
