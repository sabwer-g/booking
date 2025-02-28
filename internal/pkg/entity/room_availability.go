package entity

import (
	"time"
)

type RoomAvailability struct {
	HotelID string    `json:"hotel_id"`
	RoomID  string    `json:"room_id"`
	Date    time.Time `json:"date"`
	Quota   int       `json:"quota"`
}

type RoomsAvailability []RoomAvailability

func (rs RoomsAvailability) ToDateMap() map[time.Time]RoomAvailability {
	result := make(map[time.Time]RoomAvailability, len(rs))
	for _, r := range rs {
		result[r.Date] = r
	}
	return result
}
