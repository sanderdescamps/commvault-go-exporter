package client

import "time"

type ClientStatus struct {
	IsActive        bool
	LastStatusCheck time.Time
}
