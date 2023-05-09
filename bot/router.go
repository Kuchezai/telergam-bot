package bot

import (
	"regexp"
	"telegram-bot/entity"
	"telegram-bot/entity/command"
	"telegram-bot/handler"
	"telegram-bot/handler/helpers"
)

type StateInformerRepo interface {
	ChatState(chatID int) entity.State
}

type CommandHandler func(chatID int, msg string) entity.Response

type Router struct {
	handlers map[entity.State]map[command.Command]CommandHandler
	repo     StateInformerRepo
}

func restartBot(chatID int, msg string) entity.Response {
	return helpers.CreateResponseWithTwoButton(chatID, "Добро пожаловать!", command.ToMain, command.ToMain)
}

func NewRouter(repo StateInformerRepo, aboutBotHandler *handler.AboutBotHandler, requestHandler *handler.RequestHandler,
	mainStateHandler *handler.MainStateHandler, userStateHandler *handler.UserHandler, entryStateHandler *handler.EntryStateHandler) *Router {
	return &Router{
		handlers: map[entity.State]map[command.Command]CommandHandler{
			entity.EntryPoint: {
				command.ToMain: entryStateHandler.GoToMain,
			},
			entity.Main: {
				command.InfoAboutUser:        mainStateHandler.GoToGetUser,
				command.InfoAboutUserFriends: mainStateHandler.GoToGetUserFriends,
				command.RequestHistory:       mainStateHandler.GoToGetRequestHistory,
				command.InfoAboutAuthor:      mainStateHandler.GoToGetInfAboutAuthor,
			},
			entity.GetUser: {
				command.InfoAboutCommand: userStateHandler.GetInfoAboutGeUser,
				command.ToMain:           entryStateHandler.GoToMain,
				command.VKUserID:         userStateHandler.GetUser,
			},
			entity.GetUserFriends: {
				command.InfoAboutCommand: userStateHandler.GetInfoAboutGeUserFriends,
				command.ToMain:           entryStateHandler.GoToMain,
				command.VKUserID:         userStateHandler.GetUserFriends,
			},
			entity.GetHistory: {
				command.InfoAboutCommand: requestHandler.GetInfoAboutChatRequests,
				command.ToMain:           entryStateHandler.GoToMain,
			},
			entity.GetInfoAboutAuthor: {
				command.InfoGitHub: aboutBotHandler.GetAuthorGitHub,
				command.ToMain:     entryStateHandler.GoToMain,
			},
			entity.Any: {
				command.RestartBot:     entryStateHandler.RestartBot,
				command.UnknownCommand: entryStateHandler.UnknownCommand,
			},
		},
		repo: repo,
	}
}

func (r *Router) HandleUpdate(msg Message) entity.Response {
	cmd := command.Command(msg.Text)
	state := entity.State(r.repo.ChatState(msg.Chat.ID))

	// checks the input matches the VK user ID
	if match, _ := regexp.MatchString(command.VKIDPattern, string(cmd)); match {
		cmd = command.VKUserID
	}

	if cmd == command.RestartBot {
		return r.handlers[entity.Any][command.RestartBot](msg.Chat.ID, msg.Text)
	}

	if handlers, ok := r.handlers[state]; ok {
		if handler, ok := handlers[cmd]; ok {
			return handler(msg.Chat.ID, msg.Text)
		}
	}

	return r.handlers[entity.Any][command.UnknownCommand](msg.Chat.ID, msg.Text)
}
