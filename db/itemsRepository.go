package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ItemID      uuid.UUID `json:"itemId" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Cost        float64   `json:"cost"`
	RoomID      string    `json:"roomId"`
	Room        Room      `json:"type"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func GetItems() ([]Item, error) {
	var items []Item

	res := db.Preload("Room").Find(&items)
	if res.Error != nil {
		return nil, errors.New("error fetching items")
	}

	return items, nil
}

func GetItem(itemId uuid.UUID) (*Item, error) {
	var item Item

	res := db.Preload("Room").First(&item, "item_id = ?", itemId)
	if res.Error != nil {
		return nil, fmt.Errorf("error fetching item: %s", itemId)
	}

	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("item of id(%s) was not found", itemId)
	}

	return &item, nil
}

func GetItemsByRoom(roomId uuid.UUID) ([]*Item, error) {
	var items []*Item
	res := db.Preload("Room").Find(&items, "room_id = ?", roomId)
	if res.Error != nil {
		return nil, fmt.Errorf("could not retrieve items for room: %s", roomId)
	}

	return items, nil
}

func AddItem(item *Item) (*Item, error) {
	item.ItemID = uuid.New()
	item.CreatedAt = time.Now().UTC()
	item.UpdatedAt = time.Now().UTC()
	res := db.Create(&item)
	if res.Error != nil {
		return nil, res.Error
	}

	return item, nil
}

func UpdateItem(itemId uuid.UUID, updateItem *Item) error {
	updateItem.UpdatedAt = time.Now().UTC()
	res := db.Model(&Item{}).Where("item_id = ?", itemId).Updates(updateItem)
	if res.RowsAffected == 0 {
		return errors.New("item not updated")
	}

	return nil
}

func DeleteItem(itemId uuid.UUID) error {
	var deletedItem Item
	res := db.Delete(&deletedItem, "item_id = ?", itemId)
	if res.RowsAffected == 0 {
		return errors.New("item not deleted")
	}

	return nil
}
