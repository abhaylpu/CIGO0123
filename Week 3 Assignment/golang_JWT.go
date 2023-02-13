//database autho.go
/*
package database
import (
	"errors"
	"log"
	"github.com/jinzhu/gorm"
	"github.com/rahul10-pu/CIGO0122/models"
)
func SignUp(db *gorm.DB, user *models.Users) error {
	//Save update value in database, if the value doesn't have primary key, will insert it
	//1. check whether email id exis or not
	//2.if email is unique then save user
	err := db.Save(user).Error //book here is without id
	if err != nil {
		return err
	}
	return nil
}
func SignIn(db *gorm.DB, user *models.Users) error {
	getUser := models.Users{}
	err := db.Select("users.*").Group("users.id").Where("users.email_id= ?", user.EmailId).First(&getUser).Error
	if err != nil {
		return err
	}
	log.Println(getUser)
	if user.Password != getUser.Password {
		return errors.New("Passowrd Incorrect")
	}
	log.Println("exiting signin")
	return nil
}
*/
//handler - handler.go
/*
package handler
import "github.com/jinzhu/gorm"
type Handler struct {
	DB *gorm.DB
}
*/

//handler- autho.go
/*
package handler
import (
	"errors"
	"log"
	"net/http"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rahul10-pu/CIGO0122/database"
	"github.com/rahul10-pu/CIGO0122/models"
)
func (h *Handler) SignUp(in *gin.Context) {
	user := models.Users{}
	err := in.BindJSON(&user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	err = database.SignUp(h.DB, &user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, user)
}
func (h *Handler) SignIn(in *gin.Context) {
	user := models.Users{}
	err := in.BindJSON(&user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	err = database.SignIn(h.DB, &user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	log.Println("generating token")
	token, err := generateToken(user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	log.Println("generating token done")
	in.JSON(http.StatusOK, token)
}
func generateToken(user models.Users) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["emailid"] = user.EmailId
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString([]byte("rahul"))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}
func ValidateToken(Intoken string) (bool, error) {
	token, err := jwt.Parse(Intoken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("some unexpected error")
		}
		return []byte("rahul"), nil
	})
	if err != nil {
		log.Println(err)
		return false, err
	}
	if !token.Valid {
		return false, errors.New("token is not validdd")
	}
	return true, nil
}
*/

//router -book.go
/*
package routers
import (
	"github.com/gin-gonic/gin"
	"github.com/rahul10-pu/CIGO0122/handler"
)
func BookRouter(router *gin.Engine, api handler.Handler) {
	//get the default engine for further customizatikon
	// api := handler.Handler{
	// 	DB: database.GetDB(), //set the handler DB
	// }
	router.GET("/books", api.GetBooks) //set the function for http://localhost:8080/books : Get request
	//while calling handler function, gin will pass *gin.Context in the handler function
	router.POST("/books", api.SaveBook)
	router.GET("/books/:id", api.GetBookByID)
	router.DELETE("/books/:id", api.DeleteBookByID)
	router.PUT("/books", api.UpdateBookByID)
	//	router.DELETE("/books/:id", api.DeleteBook)
	// return router
}
*/