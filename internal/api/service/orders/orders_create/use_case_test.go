package orders_create

import (
	"context"
	"reflect"
	"testing"
	"time"

	"booking/internal/pkg/constant"
	dto "booking/internal/pkg/dto/orders"
	"booking/internal/pkg/errors"
	"booking/internal/pkg/logger"
)

func TestUseCase_Do(t *testing.T) {
	lg, _ := logger.New(logger.Config{})

	type fields struct {
		lg *logger.Logger
	}
	type args struct {
		in0 context.Context
		req dto.CreateRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *dto.CreateResponse
		want1  *errors.HTTPError
	}{
		{
			name:   "Success",
			fields: fields{lg: lg},
			args: args{
				in0: context.Background(),
				req: dto.CreateRequest{
					Order: dto.Order{
						HotelID:   constant.ReddisonHotelID,
						RoomID:    constant.LuxRoomID,
						UserEmail: "test@test.test",
						From:      time.Date(2025, 03, 02, 0, 0, 0, 0, time.UTC),
						To:        time.Date(2025, 03, 04, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			want:  &dto.CreateResponse{},
			want1: nil,
		},
		{
			name:   "Hotel room is not available",
			fields: fields{lg: lg},
			args: args{
				in0: context.Background(),
				req: dto.CreateRequest{
					Order: dto.Order{
						HotelID:   constant.ReddisonHotelID,
						RoomID:    constant.LuxRoomID,
						UserEmail: "test@test.test",
						From:      time.Date(2025, 12, 12, 0, 0, 0, 0, time.UTC),
						To:        time.Date(2025, 12, 14, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			want:  nil,
			want1: &errors.HTTPError{Description: "Hotel room is not available for selected dates: map[2025-12-12 00:00:00 +0000 UTC:{} 2025-12-13 00:00:00 +0000 UTC:{} 2025-12-14 00:00:00 +0000 UTC:{}]"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				lg: tt.fields.lg,
			}
			got, got1 := u.Do(tt.args.in0, tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() got = %v, want %v", got, tt.want)
			}
			if got1 != nil && got1.Description != tt.want1.Description {
				t.Errorf("Do() got1 = %v, want %v", got1.Description, tt.want1.Description)
			}
		})
	}
}
