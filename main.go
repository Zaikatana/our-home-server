package main

import (
	"github.com/gin-gonic/gin"

	"our-home-server/db"
	"our-home-server/routers"
)

func main() {
	db.InitPostgresDb()

	r := gin.Default()
	routers.InitItemsRouter(r)
	routers.InitRoomsRouter(r)
	routers.InitCommentsRouter(r)
	r.Run(":3001")
}
