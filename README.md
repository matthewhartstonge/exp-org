# exp-teams

A little mock-up experiment playing with the multi-tenant access concepts found in the ["Teams should be an MVP feature!"](https://blog.bullettrain.co/teams-should-be-an-mvp-feature/) Blog via Bullet Train, but in Go.

`cmd/teams` produces:

```shell
Team 1: {
  "id": "1",
  "created": "2024-08-19T18:22:23.37304+12:00",
  "updated": "2024-08-19T18:22:23.37304+12:00",
  "timezone": "nz",
  "locale": "en-uk",
  "type": "organisation",
  "name": "Builders Inc.",
  "subscription": {
    "id": "1",
    "created": "2024-08-19T18:22:23.373039+12:00",
    "updated": "2024-08-19T18:22:23.373039+12:00",
    "active": true,
    "billing": ""
  },
  "invitations": [
    "1"
  ],
  "memberships": [
    "1",
    "2"
  ]
}

Team 2: {
  "id": "1",
  "created": "2024-08-19T18:22:23.373041+12:00",
  "updated": "2024-08-19T18:22:23.373041+12:00",
  "timezone": "nz",
  "locale": "en-uk",
  "type": "support",
  "name": "Cement Mixers inc.",
  "subscription": {
    "id": "2",
    "created": "2024-08-19T18:22:23.37304+12:00",
    "updated": "2024-08-19T18:22:23.37304+12:00",
    "active": true,
    "billing": ""
  },
  "invitations": null,
  "memberships": [
    "3"
  ]
}

Membership 1: {
  "id": "1",
  "created": "2024-08-19T18:22:23.373044+12:00",
  "updated": "2024-08-19T18:22:23.373044+12:00",
  "active": true,
  "roles": [
    "admin"
  ],
  "team": {
    "id": "1",
    "created": "2024-08-19T18:22:23.37304+12:00",
    "updated": "2024-08-19T18:22:23.37304+12:00",
    "timezone": "nz",
    "locale": "en-uk",
    "type": "organisation",
    "name": "Builders Inc.",
    "subscription": {
      "id": "1",
      "created": "2024-08-19T18:22:23.373039+12:00",
      "updated": "2024-08-19T18:22:23.373039+12:00",
      "active": true,
      "billing": ""
    },
    "invitations": [
      "1"
    ],
    "memberships": [
      "1",
      "2"
    ]
  },
  "invitation": "",
  "membership": "1",
  "user": "1"
}

Membership 2: {
  "id": "2",
  "created": "2024-08-19T18:22:23.373045+12:00",
  "updated": "2024-08-19T18:22:23.373045+12:00",
  "active": false,
  "roles": [
    "user"
  ],
  "team": {
    "id": "1",
    "created": "2024-08-19T18:22:23.37304+12:00",
    "updated": "2024-08-19T18:22:23.37304+12:00",
    "timezone": "nz",
    "locale": "en-uk",
    "type": "organisation",
    "name": "Builders Inc.",
    "subscription": {
      "id": "1",
      "created": "2024-08-19T18:22:23.373039+12:00",
      "updated": "2024-08-19T18:22:23.373039+12:00",
      "active": true,
      "billing": ""
    },
    "invitations": [
      "1"
    ],
    "memberships": [
      "1",
      "2"
    ]
  },
  "invitation": "1",
  "membership": "1",
  "user": ""
}

Membership 3: {
  "id": "3",
  "created": "2024-08-19T18:22:23.373045+12:00",
  "updated": "2024-08-19T18:22:23.373045+12:00",
  "active": false,
  "roles": [
    "admin"
  ],
  "team": {
    "id": "1",
    "created": "2024-08-19T18:22:23.373041+12:00",
    "updated": "2024-08-19T18:22:23.373041+12:00",
    "timezone": "nz",
    "locale": "en-uk",
    "type": "support",
    "name": "Cement Mixers inc.",
    "subscription": {
      "id": "2",
      "created": "2024-08-19T18:22:23.37304+12:00",
      "updated": "2024-08-19T18:22:23.37304+12:00",
      "active": true,
      "billing": ""
    },
    "invitations": null,
    "memberships": [
      "3"
    ]
  },
  "invitation": "",
  "membership": "1",
  "user": "2"
}

User 1: {
  "id": "1",
  "created": "2024-08-19T18:22:23.373041+12:00",
  "updated": "2024-08-19T18:22:23.373041+12:00",
  "email": "bob@example.com",
  "timezone": "nz",
  "firstName": "Bob",
  "lastName": "TheBuilder",
  "memberships": [
    "1"
  ],
  "teams": [
    "1"
  ]
}

User 2: {
  "id": "2",
  "created": "2024-08-19T18:22:23.373044+12:00",
  "updated": "2024-08-19T18:22:23.373044+12:00",
  "email": "dizzy@example.com",
  "timezone": "nz",
  "firstName": "Dizzy",
  "lastName": "Rascal",
  "memberships": [
    "3",
    "2"
  ],
  "teams": [
    "1"
  ]
}

Invitation 1: {
  "id": "1",
  "created": "2024-08-19T18:22:23.373045+12:00",
  "updated": "2024-08-19T18:22:23.373045+12:00",
  "team": "1",
  "from": "1",
  "to": "2",
  "email": "dizzy@example.com",
  "state": "sent"
}
```