package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"our-home-server/db"
)

func getItems(c *gin.Context) {
	items, err := db.GetItems()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to load items"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"items": items})
}

func getItemsByRoom(c *gin.Context) {
	roomId, err := uuid.Parse(c.Param("roomId"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Room is invalid"})
		return
	}

	// check if room exists
	room, err := db.GetRoom(roomId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Room not found"})
		return
	}

	roomItems, err := db.GetItemsByRoom(room.RoomID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to load items"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"items": roomItems})
}

func getItem(c *gin.Context) {
	itemId, err := uuid.Parse(c.Param("itemId"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item is invalid"})
		return
	}

	item, err := db.GetItem(itemId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"item": item})
}

func addItem(c *gin.Context) {
	var newItem *db.Item
	if err := c.BindJSON(&newItem); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid item body"})
		return
	}

	addedItem, err := db.AddItem(newItem)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to add item"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"itemId": addedItem.ItemID})
}

func updateItem(c *gin.Context) {
	var updatedItem *db.Item
	itemId, err := uuid.Parse(c.Param("itemId"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item is invalid"})
		return
	}

	if err := c.BindJSON(&updatedItem); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid item body"})
		return
	}

	updatedItemRes := db.UpdateItem(itemId, updatedItem)
	if updatedItemRes != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to update item"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"itemId": itemId})
}

func deleteItem(c *gin.Context) {
	itemId, err := uuid.Parse(c.Param("itemId"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item is invalid"})
		return
	}

	deleteItemResult := db.DeleteItem(itemId)
	if deleteItemResult != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to delete item"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Item was deleted"})
}

func InitItemsRouter(r *gin.Engine) {
	items := r.Group("/api/items")
	{
		items.GET("/all", getItems)
		items.GET("/room/:roomId", getItemsByRoom)
		items.GET("/:itemId", getItem)
		items.DELETE("/:itemId", deleteItem)
		items.POST("/add", addItem)
		items.PUT("/:itemId", updateItem)
	}
}
