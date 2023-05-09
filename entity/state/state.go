package state

type State int

const (
	EntryPoint State = iota
	Main
	GetUser
	GetUserFriends
	GetHistory
	GetInfoAboutAuthor
	Any
)
