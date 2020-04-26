package services

import (
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"

	"github.com/r08521610/personal_backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

// UploadHandler handle the file upload
func UploadHandler(c *gin.Context)  {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	data, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	fileID := mongo.GetInstance().SaveFile("ImagesHolder", data, filename)

	id := mongo.GetInstance().InsertOne("ImagesHolder", "files", bson.M{
		"name": file.Filename,
		"fileID": fileID,
	})

	c.JSON(http.StatusOK, gin.H{
		"fileId": id,
	})
}