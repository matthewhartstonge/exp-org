package org

import (
	"encoding/json"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
	Email     string    `json:"email"`
	Timezone  string    `json:"timezone"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`

	Memberships []*Membership `json:"memberships"`
	Teams       []*Team       `json:"teams"`
}

func (u *User) AddMembership(membership *Membership) *User {
	for i, m := range u.Memberships {
		if m.ID == membership.ID {
			u.Memberships[i] = membership
			return u
		}
	}

	u.Memberships = append(u.Memberships, membership)
	return u
}

func (u *User) AddTeam(team *Team) *User {
	for i, t := range u.Teams {
		if t.ID == team.ID {
			u.Teams[i] = team
			return u
		}
	}

	u.Teams = append(u.Teams, team)
	return u
}

func (u *User) MarshalJSON() ([]byte, error) {
	type userCopy struct {
		User

		Memberships []string `json:"memberships"`
		Teams       []string `json:"teams"`
	}

	c := userCopy{User: *u}
	for _, m := range u.Memberships {
		c.Memberships = append(c.Memberships, m.ID)
	}
	for _, t := range u.Teams {
		c.Teams = append(c.Teams, t.ID)
	}

	return json.Marshal(c)
}
