package services

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/r08521610/personal_backend/mongo"
	// "go.mongodb.org/mongo-driver/bson"
)

// GetFile get file
func GetFile(c *gin.Context) {
	db := c.Param("db")
	id := c.Param("id")
	fmt.Println(db)

	content, size := mongo.GetInstance().DownloadFile(db, id)

	c.Writer.WriteHeader(http.StatusOK)
	// c.Header("Content-Disposition", "attachment; filename=hello.txt")
	// c.Header("Content-Type", "application/text/plain")
	c.Header("Accept-Length", fmt.Sprintf("%d", size))
	c.Writer.Write(content)
}