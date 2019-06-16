package domain

import (
	"time"
)

type from struct {
	value time.Time
}
type to struct {
	value time.Time
}

type Period struct {
	from from
	to   to
}
