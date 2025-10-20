package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"our-home-server/db"
	"our-home-server/routers"

	"time"
)

func main() {
	db.InitPostgresDb()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour}))
	routers.InitItemsRouter(r)
	routers.InitRoomsRouter(r)
	routers.InitCommentsRouter(r)
	r.Run(":3001")
}
