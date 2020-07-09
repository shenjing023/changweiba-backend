package main

import (
	"changweiba-backend/common"
	"changweiba-backend/conf"
	"changweiba-backend/dao"
	mygraphql "changweiba-backend/graphql"
	"changweiba-backend/graphql/dataloader"
	"changweiba-backend/graphql/generated"
	"changweiba-backend/pkg/middleware"
	"flag"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	//"github.com/99designs/gqlgen/handler"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

const defaultPort = ":8088"

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	//h := handler.Playground("GraphQL", "/graphql")
	h := playground.Handler("Graphql", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.GraphQL(
		generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}),
		handler.ComplexityLimit(20),
	)

	h := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}),
		handler.ComplexityLimit(20))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	//设置可同时使用的CPU数目，
	runtime.GOMAXPROCS(runtime.NumCPU())
	//命令行解析
	pwd, _ := os.Getwd()
	execDir := flag.String("d", pwd, "execute directory")
	flag.Parse()
	fmt.Println("Current execute directory is:", *execDir)
	conf.InitConfig(*execDir)
	//graphql.InitRPCConnection()
	dao.Init()
	registerSignalHandler()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Setting up Gin
	r := gin.Default()
	r.Use(common.GinContextToContextMiddleware())
	r.Use(middleware.JWTMiddleware(conf.Cfg.SignKey, conf.Cfg.QueryDeep))
	r.Use(dataloader.LoaderMiddleware())

	r.POST("/graphql", graphqlHandler())
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
				//graphql.StopRPCConnection()
				time.Sleep(time.Second)
				os.Exit(0)
			}
		}
	}()
}
