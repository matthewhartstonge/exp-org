package main

import (
	"encoding/json"
	"fmt"
	"github.com/matthewhartstonge/exp/org"
	"time"
)

func main() {
	s1 := &org.Subscription{
		ID:             "1",
		Created:        time.Now(),
		Updated:        time.Now(),
		Active:         true,
		BillingDetails: "",
	}

	t1 := &org.Team{
		ID:           "1",
		Created:      time.Now(),
		Updated:      time.Now(),
		Timezone:     "nz",
		Locale:       "en-uk",
		Type:         org.TeamTypeTeam,
		Name:         "Builders Inc.",
		Subscription: s1,
		Invitations:  nil,
		Memberships:  nil,
	}

	s2 := &org.Subscription{
		ID:             "2",
		Created:        time.Now(),
		Updated:        time.Now(),
		Active:         true,
		BillingDetails: "",
	}

	t2 := &org.Team{
		ID:           "1",
		Created:      time.Now(),
		Updated:      time.Now(),
		Timezone:     "nz",
		Locale:       "en-uk",
		Type:         org.TeamTypeSupport,
		Name:         "Cement Mixers inc.",
		Subscription: s2,
		Invitations:  nil,
		Memberships:  nil,
	}

	u1 := &org.User{
		ID:          "1",
		Created:     time.Now(),
		Updated:     time.Now(),
		Email:       "bob@example.com",
		Timezone:    "nz",
		FirstName:   "Bob",
		LastName:    "TheBuilder",
		Memberships: nil,
		Teams:       nil,
	}

	u2 := &org.User{
		ID:          "2",
		Created:     time.Now(),
		Updated:     time.Now(),
		Email:       "dizzy@example.com",
		Timezone:    "nz",
		FirstName:   "Dizzy",
		LastName:    "Rascal",
		Memberships: nil,
		Teams:       nil,
	}

	m1 := &org.Membership{
		ID:      "1",
		Created: time.Now(),
		Updated: time.Now(),
		Active:  true,
		Roles: []org.Role{
			org.RoleTeamAdmin,
		},
		Invitation: nil,
		Team:       nil,
		User:       nil,
	}

	m2 := &org.Membership{
		ID:      "2",
		Created: time.Now(),
		Updated: time.Now(),
		Active:  false,
		Roles: []org.Role{
			org.RoleTeamUser,
		},
		Invitation: nil,
		Team:       t1,
		User:       nil,
	}

	m3 := &org.Membership{
		ID:      "3",
		Created: time.Now(),
		Updated: time.Now(),
		Active:  false,
		Roles: []org.Role{
			org.RoleTeamAdmin,
		},
		Invitation: nil,
		Team:       t2,
		User:       u2,
	}

	i1 := &org.Invitation{
		ID:      "1",
		Created: time.Now(),
		Updated: time.Now(),
		Team:    t1.ID,
		From:    m1.ID,
		To:      m2.ID,
		Email:   "dizzy@example.com",
		State:   org.InvitationSent,
	}
	m2.Invitation = i1

	// Bind team and user to membership
	m1.Team = t1
	m1.User = u1

	// Bind team to invitations and memberships
	t1.AddMembership(m1)
	t1.AddMembership(m2).AddInvitation(i1)
	t2.AddMembership(m3)

	// Bind user to membership and team
	u1.AddMembership(m1).AddTeam(t1)
	u2.AddMembership(m3).AddTeam(t2)
	u2.AddMembership(m2)

	fmt.Println("Team 1:", jsonPrettyPrint(t1))
	fmt.Println("Team 2:", jsonPrettyPrint(t2))

	fmt.Println("Membership 1:", jsonPrettyPrint(m1))
	fmt.Println("Membership 2:", jsonPrettyPrint(m2))
	fmt.Println("Membership 3:", jsonPrettyPrint(m3))

	fmt.Println("User 1:", jsonPrettyPrint(u1))
	fmt.Println("User 2:", jsonPrettyPrint(u2))

	fmt.Println("Invitation 1:", jsonPrettyPrint(i1))
}

func jsonPrettyPrint(v any) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
