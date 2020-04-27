package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/r08521610/personal_backend/graph/resolver"
	"github.com/r08521610/personal_backend/graph/generated"
	"github.com/r08521610/personal_backend/services"
)

const defaultPort = "8080"

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	r.POST("/upload", services.UploadHandler)

	r.GET("/file/db/:db/:id", services.GetFile)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	r.Use(cors.Default())
	r.Run(":"+port)
}
