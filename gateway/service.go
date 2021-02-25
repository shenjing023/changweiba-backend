package main

import (
	"flag"
	"fmt"
	"gateway/conf"
	"gateway/generated"
	"gateway/middleware"
	"os"
	"os/signal"
	"syscall"
	"time"

	service_handler "gateway/handler"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	log "github.com/shenjing023/llog"
)

// runGatewayService create gateway service
func runGatewayService(configPath string) {
	conf.Init(configPath)
	service_handler.InitGRPCConn()
	// dao.Init()
	middleware.InitAuth()
	registerSignalHandler()

	// Setting up Gin
	r := gin.Default()
	r.Use(middleware.GinContextToContextMiddleware())
	r.Use(middleware.QueryDeepMiddleware(conf.Cfg.QueryDeep))
	r.Use(middleware.AuthMiddleware())
	// r.Use(middleware.JWTMiddleware(conf.Cfg.SignKey, conf.Cfg.QueryDeep))
	// r.Use(dataloader.LoaderMiddleware())

	r.POST("/graphql", graphqlHandler())
	r.GET("/", playgroundHandler())

	r.Run(fmt.Sprintf(":%d", conf.Cfg.Port))
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	//h := handler.Playground("GraphQL", "/graphql")
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &Resolver{},
		},
	))
	srv.Use(extension.FixedComplexityLimit(20))

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func registerSignalHandler() {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		for {
			sig := <-c
			log.Infof("Signal %d received", sig)
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				service_handler.StopGRPCConn()
				time.Sleep(time.Second)
				os.Exit(0)
			}
		}
	}()
}

func main() {
	pwd, _ := os.Getwd()
	execDir := flag.String("d", pwd, "execute directory")
	flag.Parse()
	runGatewayService(*execDir + "/conf/config.yaml")
}
