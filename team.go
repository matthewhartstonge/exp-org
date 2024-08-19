package org

import (
	"encoding/json"
	"time"
)

// Team or rather, an organisation.
type Team struct {
	ID           string        `json:"id"`
	Created      time.Time     `json:"created"`
	Updated      time.Time     `json:"updated"`
	Timezone     string        `json:"timezone"`
	Locale       string        `json:"locale"`
	Type         TeamType      `json:"type"`
	Name         string        `json:"name"`
	Subscription *Subscription `json:"subscription"`
	Invitations  []*Invitation `json:"invitations"`
	Memberships  []*Membership `json:"memberships"`
}

func (t *Team) MarshalJSON() ([]byte, error) {
	type teamCopy struct {
		Team
		Invitations []string `json:"invitations"`
		Memberships []string `json:"memberships"`
	}

	c := teamCopy{Team: *t}
	for _, i := range t.Invitations {
		c.Invitations = append(c.Invitations, i.ID)
	}
	for _, m := range t.Memberships {
		c.Memberships = append(c.Memberships, m.ID)
	}

	return json.Marshal(c)
}

func (t *Team) AddInvitation(invitation *Invitation) *Team {
	for i, invite := range t.Invitations {
		if invitation.ID == invite.ID {
			t.Invitations[i] = invitation // yolo overwrite it with the update
			return t
		}
	}

	t.Invitations = append(t.Invitations, invitation)
	return t
}

func (t *Team) AddMembership(membership *Membership) *Team {
	for i, member := range t.Memberships {
		if membership.ID == member.ID {
			t.Memberships[i] = membership // yolo overwrite it with the update
			return t
		}
	}

	t.Memberships = append(t.Memberships, membership)
	return t
}

type TeamType string

const (
	TeamTypeTeam    TeamType = "organisation"
	TeamTypeSupport TeamType = "support"
)

// SubscriptionQuantity provides a way to calculate subscription fees for a
// team, based on active membership - or not.
func (t *Team) SubscriptionQuantity(activeOnly bool) int {
	if !t.Subscription.Active {
		return 0
	}

	count := 0
	for _, m := range t.Memberships {
		if m.Active || !activeOnly {
			count++
		}
	}

	return count
}
