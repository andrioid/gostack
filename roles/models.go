// Package roles helps the application understand what little information the authentication provider hands over.
// In a sense, it handles what is the basis of our authorisation
package roles

// RoleEnum is an IOTA type used to set roles
type RoleEnum string

const RoleCustomer RoleEnum = "customer"
const RoleOverlord RoleEnum = "overlord"

// Role of a user
type Role struct {
	// UUID given to us by the authentication provider. Used to identify a specific user.
	UUID string
	// Role Name.
	Role RoleEnum
}
