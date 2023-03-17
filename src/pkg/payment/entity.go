package payment

import "time"

type Payment struct {
	PaymentId     string    `json:"id"`
	ReservationId string    `json:"reservation_idid"`
	PaymentAmount float64   `json:"payment_amount"`
	PaymentDate   time.Time `json:"payment_date"`
	PaymentStatus string    `json:"payment_status"`
}
