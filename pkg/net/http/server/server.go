package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Debug   bool   `yaml:"debug"`
	Address string `yaml:"address"`
}

type Engine struct {
	*gin.Engine
}

func NewEngine(conf *Config) *Engine {

	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// TODO
	engine := &Engine{gin.New()}
	engine.Use(gin.Logger(), gin.Recovery())

	return engine
}

func (engine *Engine) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return engine.Engine.Use(middleware...)
}

func (engine *Engine) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{engine.Engine.Group(relativePath, transferHandlers(handlers...)...)}
}

func (engine *Engine) RunServer(conf *Config) {

	server := &http.Server{
		Addr:    conf.Address,
		Handler: engine.Engine,
	}

	go func() {
		log.Printf("Server listen on %v\n", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("Listen: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Success graceful shutdown server.")
}

type RouterGroup struct {
	*gin.RouterGroup
}

type Context struct {
	*gin.Context
}

type HandlerFunc func(*Context)

func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{group.RouterGroup.Group(relativePath, transferHandlers(handlers...)...)}
}

func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return group.RouterGroup.GET(relativePath, transferHandlers(handlers...)...)
}

func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return group.RouterGroup.POST(relativePath, transferHandlers(handlers...)...)
}

func transferHandlers(handlers ...HandlerFunc) []gin.HandlerFunc {
	var ginHandlers []gin.HandlerFunc
	for _, handler := range handlers {
		ginHandlers = append(ginHandlers, func(context2 *gin.Context) { handler(&Context{context2}) })
	}
	return ginHandlers
}
