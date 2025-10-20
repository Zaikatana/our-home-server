package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"our-home-server/db"
	"our-home-server/routers"
)

func main() {
	db.InitPostgresDb()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")

	r := gin.Default()
	r.Use(cors.New(corsConfig))
	routers.InitItemsRouter(r)
	routers.InitRoomsRouter(r)
	routers.InitCommentsRouter(r)
	r.Run(":3001")
}
