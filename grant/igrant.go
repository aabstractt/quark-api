package grant

import (
	"github.com/aabstractt/quark-api/user/source"
	"time"
)

// IGrant represents a group granted to a user
type IGrant interface {

	// UniqueId returns the grant's unique identifier
	UniqueId() string
	// GroupId returns the grant's user identifier
	GroupId() string
	// AddedAt returns the grant's added at time
	AddedAt() time.Time
	// AddedBy returns the grant's added by user identifier
	AddedBy() source.ISource // TODO: Return a Source object

	// ExpiresAt returns the grant's expiration time
	// If the grant does not expire, this should return a zero time
	ExpiresAt() time.Time
	// Expired returns true if the grant has expired
	// If expires at is zero, this should return false
	Expired() bool

	// RemovedAt returns the grant's removed at time
	// If the grant is not removed, this should return a zero time
	RemovedAt() time.Time
	// RemovedBy returns the grant's removed by user identifier
	// If the grant is not removed, this should return a nil source
	RemovedBy() source.ISource // TODO: Return a Source object
	// Remove is used to remove the grant from the user
	// this going to set the RemovedAt and RemovedBy fields
	Remove(src source.ISource)

	// Marshal returns the grant's JSON representation
	Marshal() // TODO: Implement marshal to MongoDB
}

func Unmarshal() {
	panic("Not implemented")

	// TODO: Implement unmarshal from MongoDB
	// to fetch the source I going to check if it is stored into our cache...
	// If not, I will fetch it from the database and store it into the cache
}
