package main

import (
	"changweiba-backend/graphql"
	"context"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"os"
)

const defaultPort = "8080"

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Setting up Gin
	r := gin.Default()
	r.Use(GinContextToContextMiddleware())
	//r.Use(jwt.JWTMiddleware())
	r.POST("/graphql", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
