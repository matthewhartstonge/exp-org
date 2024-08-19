package org

import "time"

// Subscription enables billing for the team.
type Subscription struct {
	ID             string    `json:"id"`
	Created        time.Time `json:"created"`
	Updated        time.Time `json:"updated"`
	Active         bool      `json:"active"`
	BillingDetails string    `json:"billing"`
}
