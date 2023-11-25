package entity

import "time"

type Task struct {
	Id          int
	Title       string
	Description string
	Status      bool
	User_id     int
	Category_id int
	Created_at  time.Time
	Updated_at  time.Time
}
