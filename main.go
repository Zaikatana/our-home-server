package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"

	"our-home-server/db"
	"our-home-server/routers"
)

func main() {
	db.InitPostgresDb()

	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                                                 // Specify allowed origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},           // Specify allowed HTTP methods
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"}, // Specify allowed request headers
		AllowCredentials: true,                                                          // Allow cookies and authentication headers to be sent
		Debug:            true,                                                          // Enable debug logging for CORS issues
	})

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		corsConfig.HandlerFunc(c.Writer, c.Request)
		if c.IsAborted() { // If CORS handled the request (e.g., preflight OPTIONS), abort further processing
			return
		}
		c.Next() // Otherwise, continue to the next middleware or handler
	})
	routers.InitItemsRouter(r)
	routers.InitRoomsRouter(r)
	routers.InitCommentsRouter(r)
	r.Run(":3001")
}
