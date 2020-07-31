package main

import (
	"changweiba-backend/common"
	"changweiba-backend/conf"
	"changweiba-backend/dao"
	"changweiba-backend/graphql/dataloader"
	"changweiba-backend/graphql/generated"
	"changweiba-backend/pkg/middleware"
	"context"
	"flag"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"

	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"changweiba-backend/pkg/logs"

	mygraphql "changweiba-backend/graphql"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const defaultPort = ":8088"

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
			Resolvers: &mygraphql.Resolver{},
		},
	))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		// send this panic somewhere
		logs.Error("system panic: ", err)
		logs.Trace(string(debug.Stack()))
		return errors.New("system error")
	})
	srv.Use(extension.FixedComplexityLimit(20))

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
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
	dao.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Setting up Gin
	r := gin.Default()
	r.Use(common.GinContextToContextMiddleware())
	r.Use(middleware.QueryMiddleware(conf.Cfg.QueryDeep))
	r.Use(middleware.JWTMiddleware(conf.Cfg.SignKey))
	r.Use(dataloader.LoaderMiddleware())

	r.POST("/graphql", graphqlHandler())
	r.GET("/", playgroundHandler())

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logs.Error("listen: %s\n", err)
		}
	}()

	registerSignalHandler(srv)

}

func registerSignalHandler(srv *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	for {
		sig := <-c
		logs.Info("Signal %d received", sig)
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			srv.Shutdown(context.Background())
			//todo 关闭数据库
			time.Sleep(time.Second)
			os.Exit(0)
		}
	}
}
