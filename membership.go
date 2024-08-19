package org

import (
	"encoding/json"
	"time"
)

// Membership creates a binding between a team and a user, also allows for
// tracking of membership invitation status.
type Membership struct {
	ID         string      `json:"id"`
	Created    time.Time   `json:"created"`
	Updated    time.Time   `json:"updated"`
	Active     bool        `json:"active"`
	Roles      []Role      `json:"roles"`
	Invitation *Invitation `json:"invitation"`
	Team       *Team       `json:"team"`
	User       *User       `json:"user"`
}

func (m *Membership) AddedBy() string {
	if m.Team == nil || m.Invitation == nil {
		return ""
	}

	return m.Invitation.From
}

func (m *Membership) MarshalJSON() ([]byte, error) {
	type membershipCopy struct {
		Membership
		Invitation string `json:"invitation"`
		Team       string `json:"membership"`
		User       string `json:"user"`
	}

	c := membershipCopy{Membership: *m}
	if c.Membership.Invitation != nil {
		c.Invitation = c.Membership.Invitation.ID
	}
	if c.Membership.Team != nil {
		c.Team = c.Membership.Team.ID
	}
	if c.Membership.User != nil {
		c.User = c.Membership.User.ID
	}

	return json.Marshal(c)
}

type Role string

const (
	RoleTeamAdmin Role = "admin"
	RoleTeamUser  Role = "user"
)
