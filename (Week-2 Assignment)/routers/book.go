package routers

import (
	"github.com/abhaylpu/CIGO0123/tree/main/(Week-2%20Assignment)/database"
	"github.com/abhaylpu/CIGO0123/tree/main/(Week-2%20Assignment)/handlers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default() // get the dfault engine for further customization
	api := handler.Handler{
		DB: database.GetDB(), //set the handler DB
	}
	router.GET("/books", api.GetBooks) //set thefunction for thiss  https://localhost:8080/books : Get Request
	// while calling the handler function , gin will pass *gin.Context in the handler function
	router.POST("/books", api.SaveBook)

	return router
}
