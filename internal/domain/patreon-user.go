package domain

import "errors"

var (
	NotFound = errors.New("patreon user not found")
)

type PatreonUser struct {

	// Patreon user id
	UserId string `json:"user_id"`
}
