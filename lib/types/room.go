package types

import "time"

type Room struct {
	Id string
	Name string
	Owner string
	CreatedAt time.Time
}
