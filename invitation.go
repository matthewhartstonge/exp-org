package org

import "time"

type Invitation struct {
	ID      string          `json:"id"`
	Created time.Time       `json:"created"`
	Updated time.Time       `json:"updated"`
	Team    string          `json:"team"`
	From    string          `json:"from"` // membership who invited
	To      string          `json:"to"`
	Email   string          `json:"email"` // new/existing member
	State   InvitationState `json:"state"`
}

type InvitationState string

const (
	InvitationSent     InvitationState = "sent"
	InvitationAccepted InvitationState = "accepted"
	InvitationDeclined InvitationState = "declined"
)
