package entity

type Command string

const (
	ToMain               Command = "В меню"
	InfoAboutUser        Command = "Информация о пользователе"
	InfoAboutUserFriends Command = "Информация о друзьях"
	RequestHistory       Command = "История запросов"
	InfoAboutAuthor      Command = "Информация об авторе"
	InfoGitHub           Command = "GitHub"
	InfoAboutCommand     Command = "Информация о команде"
	RestartBot           Command = "/start"
	UnknownCommand       Command = "?"
	VKUserID             Command = "ID пользователя"
	VKIDPattern                  = `^[1-9]\d{0,8}$` // regular expression. it's not a command, it's used only to check the input for compliance with VKID
)
