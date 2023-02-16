package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode("release")
	r := gin.Default()
	return r
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
