package models

import "time"

type Schedule struct {
	ID          uint      `json:"schedule_id"`
	Aid_name    string    `json:"aid_name"`
	Aid_per_day int       `json:"aid_per_day"`
	UserID      uint      `json:"user_id"`
	Duration    int       `json:"duration"`
	Created_at  time.Time `json:"created_at"`
}
