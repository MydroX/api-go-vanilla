package models

import "time"

// Medal represents a medal.
type Medal struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
