package msgtxt

//TODO: Move messages from constants to json .ru file
//error text
const (
	SomethingWrongTryLatter = "Что-то пошло не так. Попробуйте позже"
	NotFoundOrBadVKUserID   = "Пользователь не найден, скорее всего введный ID некорректен"
	UserIsPrivat            = "Указаный пользователь имеет приватный профиль"
	FriendsNotFound         = "У пользователя нет друзей"
)

//info text
const (
	Welcome                     = "Добро пожаловать"
	UnknownCommand              = "Я не знаю такой команды"
	ChoseAction                 = "Выберите действие"
	ChooseNextAction            = "Выберите следующее действие"
	RequestHistory              = "История запросов:"
	InputVKIDOrChooseNextAction = "Пожалуйста, введите ID пользователя VK или выберите следующее действие\n\nЧто такое VK ID и как его узнать:\nhttps://vk.com/faq18062"
	AboutAuthor                 = "Надеюсь кто-то дошёл до этой кнопки"
	GitHubLink                  = "https://github.com/Kuchezai"
	InfoAboutGetChatRequests    = "Это команда показывает Вашу историю запросов к этому боту"
	InfoAboutGetUser            = "Это команда выводит имя и фамилию пользователя VK"
	InfoAboutGetUserFriends     = "Это команда выводит 100 друзей публичных пользователей VK"
)
