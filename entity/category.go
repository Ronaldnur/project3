package entity

import "time"

type Category struct {
	Id         int
	Type       string
	Created_at time.Time
	Updated_at time.Time
}
