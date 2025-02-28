package convert

import (
	dto "booking/internal/pkg/dto/orders"
	"booking/internal/pkg/entity"
)

func OrdersFromDTOToEntity(o dto.Order) entity.Order {
	return entity.Order{
		HotelID:   o.HotelID,
		RoomID:    o.RoomID,
		UserEmail: o.UserEmail,
		From:      o.From,
		To:        o.To,
	}
}
