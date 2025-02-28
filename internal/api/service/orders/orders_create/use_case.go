package orders_create

import (
	"context"
	"fmt"
	"time"

	"booking/internal/pkg/constant"
	"booking/internal/pkg/convert"
	dto "booking/internal/pkg/dto/orders"
	"booking/internal/pkg/entity"
	"booking/internal/pkg/errors"
	"booking/internal/pkg/logger"
	timeUtils "booking/internal/pkg/utils/time"
)

/* Для случаев подключения внешнего хранилища

type roomsAvailability interface {
	List(ctx context.Context, ...roomsparams) (entity.RoomsAvailability, error)
	Update(ctx context.Context, entity.RoomsAvailability) error
}

type orders interface {
	Create(ctx context.Context, order entity.Order) (uint64, error)
}
*/

var Orders = entity.Orders{}

var Availability = entity.RoomsAvailability{
	{constant.ReddisonHotelID, constant.LuxRoomID, timeUtils.Date(2025, 3, 1), 1},
	{constant.ReddisonHotelID, constant.LuxRoomID, timeUtils.Date(2025, 3, 2), 1},
	{constant.ReddisonHotelID, constant.LuxRoomID, timeUtils.Date(2025, 3, 3), 1},
	{constant.ReddisonHotelID, constant.LuxRoomID, timeUtils.Date(2025, 3, 4), 1},
	{constant.ReddisonHotelID, constant.LuxRoomID, timeUtils.Date(2025, 3, 5), 0},
}

type UseCase struct {
	//roomsAvailability roomsAvailability,
	//orders orders,
	lg *logger.Logger
}

func NewUseCase(lg *logger.Logger) *UseCase {
	return &UseCase{lg: lg}
}

func (u *UseCase) Do(_ context.Context, req dto.CreateRequest) (*dto.CreateResponse, *errors.HTTPError) {
	if len(Availability) == 0 {
		u.lg.LogErrorf("Not found available hotel rooms now for order: \n%v", req.Order)
		return nil, errors.NewInternalError(fmt.Sprintf("Not found available hotel rooms now for order: %v", req.Order))
	}
	dateRoomsAvailabilityMap := Availability.ToDateMap()

	daysToBookMap := timeUtils.DaysBetween(req.From, req.To)
	unavailableDaysMap := make(map[time.Time]struct{}, len(Availability))

	for dayToBook := range daysToBookMap {
		if roomAvailable, ok := dateRoomsAvailabilityMap[dayToBook]; ok {
			if req.HotelID == roomAvailable.HotelID && req.RoomID == roomAvailable.RoomID && roomAvailable.Quota > 0 {
				continue
			}
		}
		unavailableDaysMap[dayToBook] = struct{}{}
	}

	if len(unavailableDaysMap) != 0 {
		u.lg.LogErrorf("Hotel room is not available for selected dates: \n%v", unavailableDaysMap)
		return nil, errors.NewInternalError(fmt.Sprintf("Hotel room is not available for selected dates: %v", unavailableDaysMap))
	}

	for i, roomAvailable := range Availability {
		if _, ok := daysToBookMap[roomAvailable.Date]; ok {
			if req.HotelID == roomAvailable.HotelID && req.RoomID == roomAvailable.RoomID {
				Availability[i].Quota -= 1
			}
		}
	}

	Orders = append(Orders, convert.OrdersFromDTOToEntity(req.Order))

	u.lg.LogInfo("Order successfully created: %v", req.Order)

	return &dto.CreateResponse{}, nil
}
