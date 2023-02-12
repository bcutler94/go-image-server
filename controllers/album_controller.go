package controllers

import (
	"example/web-service-gin/configs"
	"example/web-service-gin/models"
	"example/web-service-gin/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database = configs.GetDB(configs.MongoClient, "test")
var albumCollection *mongo.Collection = configs.GetCollection(db, "albums")

var validate = validator.New()

func CreateAlbum() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var album models.AlbumModel

		if err := ctx.BindJSON(&album); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, responses.AlbumResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validateErr := validate.Struct(&album); validateErr != nil {
			ctx.IndentedJSON(http.StatusBadRequest, responses.AlbumResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validateErr.Error()}})
			return
		}

		newAlbum := models.AlbumModel{
			Title:  album.Title,
			Artist: album.Artist,
			Price:  album.Price,
		}

		result, err := albumCollection.InsertOne(ctx, newAlbum)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, responses.AlbumResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		ctx.IndentedJSON(http.StatusCreated, responses.AlbumResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    map[string]interface{}{"data": result},
		})
	}
}

func GetAlbum() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ID := ctx.Param("id")

		if ID == "" {
			ctx.IndentedJSON(http.StatusBadRequest, responses.AlbumResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": nil},
			})
			return
		}

		objId, _ := primitive.ObjectIDFromHex(ID)

		var album models.AlbumModel
		if err := albumCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&album); err != nil {
			ctx.IndentedJSON(http.StatusNotFound, responses.AlbumResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		ctx.IndentedJSON(http.StatusFound, responses.AlbumResponse{
			Status:  http.StatusBadRequest,
			Message: "success",
			Data:    map[string]interface{}{"data": album},
		})
	}
}
