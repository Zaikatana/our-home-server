package db

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Room struct {
	RoomID uuid.UUID `json:"roomId" gorm:"primaryKey"`
	Name   string    `json:"name"`
}

func GetRooms() ([]Room, error) {
	var rooms []Room

	res := db.Model(&Room{}).Find(&rooms)
	if res.Error != nil {
		return nil, errors.New("error fetching rooms")
	}

	return rooms, nil
}

func GetRoom(roomId uuid.UUID) (*Room, error) {
	var room Room

	res := db.First(&room, "room_id = ?", roomId)
	if res.Error != nil {
		return nil, fmt.Errorf("error fetching room: %s", roomId)
	}

	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("room of id(%s) was not found", roomId)
	}

	return &room, nil
}
