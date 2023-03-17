package reservation

import "time"

type Reservation struct {
	ReservationId  string    `json:"id"`
	UserId         string    `json:"user_id"`
	ParkingSpaceId string    `json:"parking_space_id"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	TotalCost      float64   `json:"total_cost""`
}
