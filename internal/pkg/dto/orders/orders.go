package orders

import (
	"time"

	"booking/internal/pkg/errors"
	"booking/internal/pkg/utils"
)

type Order struct {
	HotelID   string    `json:"hotel_id" validate:"required"`
	RoomID    string    `json:"room_id" validate:"required"`
	UserEmail string    `json:"email" validate:"required"`
	From      time.Time `json:"from" validate:"required"`
	To        time.Time `json:"to" validate:"required"`
}

type CreateRequest struct {
	Order
}

func (cr *CreateRequest) Validate() error {
	now := utils.EmptyTimezone(utils.BeginningOfDay(time.Now()))
	if cr.From.Before(now) || cr.To.Before(now) {
		return errors.NewRequestValidationError("Date from or to cant be early then date now")
	}

	if cr.To.Before(cr.From) {
		return errors.NewRequestValidationError("Date to cant be early then date from")
	}

	return nil
}

type CreateResponse struct {
}
