package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (server *Server) SetupRoutes(version string) http.Handler {
	r := gin.Default()

	/* CORS */
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH", "OPTIONS", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
	}))

	v1 := r.Group("/api/v1")
	{
		v1.POST("/upload", server.uploadFile)
		v1.GET("/notes", server.checkGrammar)
		v1.GET("/render-html/:id", server.getNote)
		v1.GET("/check-grammar/:id", server.checkGrammar)
	}

	return r
}
