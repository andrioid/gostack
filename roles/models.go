// Package roles helps the application understand what little information the authentication provider hands over.
// In a sense, it handles what is the basis of our authorisation
package roles

import (
	"firebase.google.com/go/auth"
)

// RoleEnum is an IOTA type used to set roles
type RoleEnum string

const RoleCustomer RoleEnum = "customer"
const RoleOverlord RoleEnum = "overlord"

const DefaultRole = RoleCustomer

type User struct {
	DisplayName string
	Role        RoleEnum
	UUID        string
}

func TokenToUser(t *auth.Token) *User {
	// This should look up the user_id from Firebase in our tables
	// If no user is found, create a record with default role
	// Always return an user object, unless token failed to parse
	return nil
}

// Role of a user
type Role struct {
	// UUID given to us by the authentication provider. Used to identify a specific user.
	UUID string
	// Role Name.
	Role RoleEnum
}
