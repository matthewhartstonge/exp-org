package org

import (
	"testing"
)

func TestTeam_SubscriptionQuantity(t1 *testing.T) {
	type fields struct {
		ID           string
		Name         string
		Subscription *Subscription
		Memberships  []*Membership
	}
	type args struct {
		activeOnly bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "All memberships within an active subscription",
			fields: fields{
				Subscription: &Subscription{
					Active: true,
				},
				Memberships: []*Membership{
					{
						Active: true,
						Roles:  nil,
					},
					{
						Active: false,
						Roles:  nil,
					},
					{
						Active: true,
						Roles:  nil,
					},
				},
			},
			args: args{
				activeOnly: false,
			},
			want: 3,
		},
		{
			name: "Only active memberships within an active subscription",
			fields: fields{
				Subscription: &Subscription{
					Active: true,
				},
				Memberships: []*Membership{
					{
						Active: true,
						Roles:  nil,
					},
					{
						Active: false,
						Roles:  nil,
					},
					{
						Active: true,
						Roles:  nil,
					},
				},
			},
			args: args{
				activeOnly: true,
			},
			want: 2,
		},
		{
			name: "No subscriptions when subscription is not active",
			fields: fields{
				Subscription: &Subscription{
					Active: false,
				},
				Memberships: []*Membership{
					{
						Active: true,
						Roles:  nil,
					},
					{
						Active: false,
						Roles:  nil,
					},
					{
						Active: true,
						Roles:  nil,
					},
				},
			},
			args: args{
				activeOnly: false,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Team{
				ID:           tt.fields.ID,
				Name:         tt.fields.Name,
				Subscription: tt.fields.Subscription,
				Memberships:  tt.fields.Memberships,
			}
			if got := t.SubscriptionQuantity(tt.args.activeOnly); got != tt.want {
				t1.Errorf("SubscriptionQuantity() = %v, want %v", got, tt.want)
			}
		})
	}
}
