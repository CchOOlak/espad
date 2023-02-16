package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Path	string
	Method  string
	Handler	func(c *gin.Context)
}

func InitRouter() *gin.Engine {
	gin.SetMode("release")
	r := gin.Default()
	return r
}

func AttachHandler(r *gin.Engine, routes []Route) {
	for _, route := range routes {
		if route.Method == "GET" {
			r.GET(route.Path, route.Handler)
		}
		if route.Method == "POST" {
			r.POST(route.Path, route.Handler)
		}
	}
}

func Start(ctx context.Context, router *gin.Engine) {
	srv := &http.Server{
        Addr:    ":8585",
        Handler: router,
    }
    
	go srv.ListenAndServe()

	for {
		select {
		case <-ctx.Done():
			srv.Shutdown(ctx)
			return
		}
	}
}
