package state

type State int

// state of the state machine
const (
	EntryPoint State = iota
	Main
	GetUser
	GetUserFriends
	GetHistory
	GetInfoAboutAuthor
	Any
)
