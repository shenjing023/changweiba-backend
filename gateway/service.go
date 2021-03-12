package main

import (
	"context"
	"flag"
	"fmt"
	"gateway/conf"
	"gateway/dataloader"
	"gateway/generated"
	"gateway/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	middleware.InitAuth()
	dataloader.Init()

	// Setting up Gin
	engine := gin.Default()
	engine.Use(middleware.GinContextToContextMiddleware())
	engine.Use(middleware.QueryDeepMiddleware(conf.Cfg.QueryDeep))
	engine.Use(middleware.AuthMiddleware())
	// r.Use(middleware.JWTMiddleware(conf.Cfg.SignKey, conf.Cfg.QueryDeep))
	// r.Use(dataloader.LoaderMiddleware())

	engine.POST("/graphql", graphqlHandler())
	engine.GET("/", playgroundHandler())

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Cfg.Port),
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("run service fatal: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infof("signal %d received and shutdown service", quit)
	srv.Shutdown(context.Background())
	service_handler.StopGRPCConn()
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
	// srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
	// 	log.Errorf("service panic: %+v", err)
	// 	log.Error(string(debug.Stack()))
	// 	return errors.New("Internal system error")
	// })
	srv.Use(extension.FixedComplexityLimit(20))

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	pwd, _ := os.Getwd()
	execDir := flag.String("d", pwd, "execute directory")
	flag.Parse()
	runGatewayService(*execDir + "/conf/config.yaml")
}
