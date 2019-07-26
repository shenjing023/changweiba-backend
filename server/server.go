package main

import (
	"changweiba-backend/graphql"
	"context"
	"github.com/99designs/gqlgen/handler"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const defaultPort = ":8088"
var accountConn *grpc.ClientConn

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/graphql")

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
	initRpcConnection()
	registerSignalHandler()
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
	r.Run(port)
}

func initRpcConnection(){
	var err error
	accountConn,err=grpc.Dial("localhost:9112",grpc.WithInsecure())
	if err!=nil{
		log.Fatal("fail to dial: %+v",err)
	}
}

func registerSignalHandler() {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		for {
			sig := <-c
			logs.Info("Signal %d received", sig)
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				exitServer()
				time.Sleep(time.Second)
				os.Exit(0)
			}
		}
	}()
}

func exitServer(){
	//关闭rpc连接
	if accountConn!=nil{
		accountConn.Close()
	}
}
