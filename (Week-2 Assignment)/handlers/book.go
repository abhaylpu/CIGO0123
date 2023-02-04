package handler

import (
	"log"
	"net/http"

	"github.com/abhaylpu/CIGO0123/tree/main/(Week-2%20Assignment)/database"
	"github.com/abhaylpu/CIGO0123/tree/main/(Week-2%20Assignment)/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) GetBooks(in *gin.Context) {
	books, err := database.GetBooks(h.DB)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, books)

}
func (h *Handler) SaveBook(in *gin.Context) {
	book := models.Book{}
	err := in.BindJSOn(&book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}

	err = database.SaveBook(h.DB, &book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, book)
}
