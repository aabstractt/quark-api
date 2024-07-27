package group

import (
	"errors"
	"slices"
	"sync"
)

var (
	groups   = make(map[string]*Group)
	groupsMu sync.Mutex
)

// Lookup returns the group with the given ID
func Lookup(id string) *Group {
	groupsMu.Lock()
	defer groupsMu.Unlock()

	return groups[id]
}

// Store stores the group into the cache
func Store(g *Group) {
	groupsMu.Lock()
	defer groupsMu.Unlock()

	groups[g.id] = g
}

// Delete deletes the group from the cache
func Delete(id string) {
	groupsMu.Lock()
	defer groupsMu.Unlock()

	delete(groups, id)
}

// All returns all the groups
func All() []*Group {
	groupsMu.Lock()
	defer groupsMu.Unlock()

	var all []*Group
	for _, g := range groups {
		all = append(all, g)
	}

	return all
}

// New creates a new group
func New(id, name string) *Group {
	return &Group{
		id:   id,
		name: name,
	}
}

func Load() {
	// Load groups from database
}

type Group struct {
	id   string
	name string

	prefixColor string // This is just a color string
	prefix      string // This is the prefix shown before the name
	suffix      string // This is the suffix shown after the name

	permissions []string
}

// Id returns the group's unique identifier
func (g *Group) Id() string {
	return g.id
}

// Name returns the group's name
func (g *Group) Name() string {
	return g.name
}

// SetName sets the group's name
func (g *Group) SetName(name string) {
	g.name = name
}

// PrefixColor returns the group's prefix color
func (g *Group) PrefixColor() string {
	return g.prefixColor
}

// SetPrefixColor sets the group's prefix color
func (g *Group) SetPrefixColor(color string) {
	g.prefixColor = color
}

// Prefix returns the group's prefix
func (g *Group) Prefix() string {
	return g.prefix
}

// SetPrefix sets the group's prefix
func (g *Group) SetPrefix(prefix string) {
	g.prefix = prefix
}

// Suffix returns the group's suffix
func (g *Group) Suffix() string {
	return g.suffix
}

// SetSuffix sets the group's suffix
func (g *Group) SetSuffix(suffix string) {
	g.suffix = suffix
}

// Permissions returns the group's permissions
func (g *Group) Permissions() []string {
	return g.permissions
}

// SetPermissions sets the group's permissions
func (g *Group) SetPermissions(perms []string) {
	g.permissions = perms
}

// AddPermission adds the permission to the group
func (g *Group) AddPermission(perm string) {
	g.permissions = append(g.permissions, perm)
}

// RemovePermission removes the permission from the group
func (g *Group) RemovePermission(perm string) error {
	if i := slices.Index(g.permissions, perm); i != -1 {
		g.permissions = append(g.permissions[:i], g.permissions[i+1:]...)

		return nil
	}

	return errors.New("permission not found")
}

// HasPermission returns true if the group has the given permission
func (g *Group) HasPermission(perm string) bool {
	return slices.Contains(g.permissions, perm)
}
