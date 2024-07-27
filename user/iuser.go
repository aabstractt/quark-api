package user

import (
	"github.com/aabstractt/quark-api/grant"
	"github.com/aabstractt/quark-api/user/source"
	"sync"
)

var (
	usersMu sync.Mutex
	users   = make(map[string]IUser)
)

type IUser interface {
	// Source returns the user's source
	// contains the user's XUID and name
	Source() source.ISource

	// ActiveGrants returns the user's active grants
	// usually, this is fetched everytime the user logs in
	ActiveGrants() []grant.IGrant
	// AddActiveGrant adds the grant to the user
	AddActiveGrant(g grant.IGrant)
	// RemoveActiveGrant removes the grant from the user
	RemoveActiveGrant(grantId string) error

	// AddExpiredGrant adds the grant to the user
	AddExpiredGrant(g grant.IGrant)
	// ExpiredGrants returns the user's expired grants
	// usually, this is needed to fetch from the database
	// and store here to avoid fetching it again
	ExpiredGrants() []grant.IGrant
	// Grants returns the user's grants
	Grants() []grant.IGrant

	// HasGroup returns true if the user has the given group
	// and the group is not expired
	HasGroup(groupId string) bool

	// SetPermissions sets the user's permissions
	SetPermissions(perms []string)

	// HasPermission returns true if the user has the given permission
	HasPermission(perm string) bool
}

// Lookup returns the user with the given XUID
func Lookup(xuid string) IUser {
	usersMu.Lock()
	defer usersMu.Unlock()

	return users[xuid]
}

// Store stores the user into the cache
func Store(u IUser) {
	usersMu.Lock()
	defer usersMu.Unlock()

	users[u.Source().XUID()] = u
}

// Delete deletes the user from the cache
func Delete(xuid string) {
	usersMu.Lock()
	defer usersMu.Unlock()

	delete(users, xuid)
}
