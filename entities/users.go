package entities

import "time"

type Users struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	NationalityID string    `json:"nationality_id"`
	PhoneNumber   string    `json:"phone_number"`
	CreatedAt     time.Time `json:"created_at"`
}
