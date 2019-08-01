package main

import (
	"changweiba-backend/common"
	"changweiba-backend/conf"
	"changweiba-backend/graphql"
	"changweiba-backend/graphql/generated"
	"flag"
	"fmt"
	"github.com/99designs/gqlgen/handler"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
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
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}


func main() {
	//命令行解析
	pwd, _ := os.Getwd()
	execDir := flag.String("d", pwd, "execute directory")
	flag.Parse()
	fmt.Println("Current execute directory is:", *execDir)
	conf.InitConfig(*execDir)
	graphql.InitRPCConnection()
	registerSignalHandler()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Setting up Gin
	r := gin.Default()
	au:=r.Group("/auth")
	au.Use(common.GinContextToContextMiddleware())
	{
		au.POST("/",graphqlHandler())
	}
	authorized :=r.Group("/")
	authorized.Use(middleware.JWTMiddleware())
	authorized.Use(common.GinContextToContextMiddleware())
	{
		authorized.POST("/graphql",graphqlHandler())
	}
	
	//r.POST("/graphql", graphqlHandler())
	r.GET("/", playgroundHandler())
	
	r.Run(port)
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
				graphql.StopRPCConnection()
				time.Sleep(time.Second)
				os.Exit(0)
			}
		}
	}()
}


