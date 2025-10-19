package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"our-home-server/db"
)

func getRooms(c *gin.Context) {
	rooms, err := db.GetRooms()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to load rooms"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"rooms": rooms})
}

func InitRoomsRouter(r *gin.Engine) {
	rooms := r.Group("/api/rooms")
	{
		rooms.GET("/all", getRooms)
	}
}
