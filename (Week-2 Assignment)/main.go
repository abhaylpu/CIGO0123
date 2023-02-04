package main

import (
	"log"

	"github.com/11001001/week2-GL1-CipherSchools/routers"
	"github.com/Gaurav11001001/week2-GL1-CipherSchools/database"
)

func main() {
	database.Setup()                    //establish the datatbase connection
	engine := routers.Router()          //get the customized engine
	err := engine.Run("127.0.0.1:8080") //start the engine
	if err != nil {
		log.Fatal(err)
	}
}
