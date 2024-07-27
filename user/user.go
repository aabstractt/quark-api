package user

import (
	"fmt"
	"github.com/aabstractt/quark-api/grant"
	"github.com/aabstractt/quark-api/user/source"
	"slices"
	"strings"
)

type User struct {
	source source.ISource

	activeGrants  []grant.IGrant // Active grants are fetched everytime the struct is created
	expiredGrants []grant.IGrant // Expired grants are fetched only when needed

	permissions []string // Permissions are calculated from the grants and stored here
}

// Source returns the user's source
// contains the user's XUID and name
func (u *User) Source() source.ISource {
	return u.source
}

// ActiveGrants returns the user's active grants
// usually, this is fetched everytime the user logs in
func (u *User) ActiveGrants() []grant.IGrant {
	return u.activeGrants
}

// AddActiveGrant adds the grant to the user
func (u *User) AddActiveGrant(g grant.IGrant) {
	if g.Expired() {
		panic("cannot add expired grant")
	}

	u.activeGrants = append(u.activeGrants, g)
}

// RemoveActiveGrant removes the grant from the user
func (u *User) RemoveActiveGrant(grantId string) error {
	for i, g := range u.activeGrants {
		if g.UniqueId() != grantId {
			continue
		}

		u.activeGrants = append(u.activeGrants[:i], u.activeGrants[i+1:]...)

		return nil
	}

	return fmt.Errorf("grant not found")
}

// AddExpiredGrant adds the grant to the user
func (u *User) AddExpiredGrant(g grant.IGrant) {
	u.expiredGrants = append(u.expiredGrants, g)
}

// ExpiredGrants returns the user's expired grants
// usually, this is needed to fetch from the database
// and store here to avoid fetching it again
func (u *User) ExpiredGrants() []grant.IGrant {
	return u.expiredGrants
}

// Grants returns the user's grants
func (u *User) Grants() []grant.IGrant {
	return append(u.activeGrants, u.expiredGrants...)
}

// HasGroup returns true if the user has the given group
// and the group is not expired
func (u *User) HasGroup(groupId string) bool {
	for _, g := range u.activeGrants {
		if g.GroupId() != groupId || g.Expired() {
			continue
		}

		return true
	}

	return false
}

// SetPermissions sets the user's permissions
func (u *User) SetPermissions(perms []string) {
	u.permissions = perms
}

// Permissions returns the user's permissions
func (u *User) Permissions() []string {
	return u.permissions
}

// HasPermission returns true if the user has the given permission
func (u *User) HasPermission(perm string) bool {
	if u.permissions == nil {
		return false
	}

	if len(u.permissions) == 0 {
		return false
	}

	if slices.Contains(u.permissions, "-"+perm) {
		return false
	}

	if slices.Contains(u.permissions, perm) {
		return true
	}

	split := strings.Split(perm, ".")
	for i := 1; i < len(split); i++ {
		if slices.Contains(u.permissions, "-"+strings.Join(split[:i], ".")+".*") {
			return false
		}
	}

	for _, p := range u.permissions {
		if !strings.HasSuffix(p, ".*") {
			continue
		}

		tempSplit := strings.Split(p, ".")
		if len(tempSplit) > len(split) {
			continue
		}

		for i := 0; i < len(tempSplit); i++ {
			if slices.Contains(u.permissions, strings.Join(split[:i], ".")+".*") {
				return true
			}
		}
	}

	return false
}
