package source

import "time"

type ISource interface {

	// XUID returns the user's XUID
	XUID() string
	// Name returns the user's name
	Name() string

	// FirstJoined returns the user's first joined time
	FirstJoined() time.Time
	// LastJoined returns the user's last joined time
	LastJoined() time.Time
}
