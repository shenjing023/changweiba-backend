package gateway

import (
	"fmt"
	"gateway/conf"
	"gateway/generated"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	log "github.com/shenjing023/llog"
)

// NewGatewayService create gateway service
func NewGatewayService(configPath string) {
	conf.Init(configPath)
	//graphql.InitRPCConnection()
	// dao.Init()
	registerSignalHandler()

	// Setting up Gin
	r := gin.Default()
	// r.Use(common.GinContextToContextMiddleware())
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
			log.Info("Signal %d received", sig)
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				//graphql.StopRPCConnection()
				time.Sleep(time.Second)
				os.Exit(0)
			}
		}
	}()
}
