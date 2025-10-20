package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"our-home-server/db"
	"our-home-server/routers"
)

func main() {
	db.InitPostgresDb()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type"}}))
	routers.InitItemsRouter(r)
	routers.InitRoomsRouter(r)
	routers.InitCommentsRouter(r)
	r.Run(":3001")
}
