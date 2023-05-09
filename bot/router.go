package bot

import (
	"regexp"
	"telegram-bot/entity"
	"telegram-bot/entity/command"
	"telegram-bot/entity/state"
	"telegram-bot/handler"
	"telegram-bot/handler/helpers"
)

type StateInformerRepo interface {
	ChatState(chatID int) state.State
}

type CommandHandler func(chatID int, msg string) entity.Response

type Router struct {
	handlers map[state.State]map[command.Command]CommandHandler
	repo     StateInformerRepo
}

func restartBot(chatID int, msg string) entity.Response {
	return helpers.CreateResponseWithTwoButton(chatID, "Добро пожаловать!", command.ToMain, command.ToMain)
}

func NewRouter(repo StateInformerRepo, aboutBotHandler *handler.AboutBotHandler, requestHandler *handler.RequestHandler,
	mainStateHandler *handler.MainStateHandler, userStateHandler *handler.UserHandler, entryStateHandler *handler.EntryStateHandler) *Router {
	return &Router{
		handlers: map[state.State]map[command.Command]CommandHandler{
			state.EntryPoint: {
				command.ToMain: entryStateHandler.GoToMain,
			},
			state.Main: {
				command.InfoAboutUser:        mainStateHandler.GoToGetUser,
				command.InfoAboutUserFriends: mainStateHandler.GoToGetUserFriends,
				command.RequestHistory:       mainStateHandler.GoToGetRequestHistory,
				command.InfoAboutAuthor:      mainStateHandler.GoToGetInfAboutAuthor,
			},
			state.GetUser: {
				command.InfoAboutCommand: userStateHandler.GetInfoAboutGeUser,
				command.ToMain:           entryStateHandler.GoToMain,
				command.VKUserID:         userStateHandler.GetUser,
			},
			state.GetUserFriends: {
				command.InfoAboutCommand: userStateHandler.GetInfoAboutGeUserFriends,
				command.ToMain:           entryStateHandler.GoToMain,
				command.VKUserID:         userStateHandler.GetUserFriends,
			},
			state.GetHistory: {
				command.InfoAboutCommand: requestHandler.GetInfoAboutChatRequests,
				command.ToMain:           entryStateHandler.GoToMain,
			},
			state.GetInfoAboutAuthor: {
				command.InfoGitHub: aboutBotHandler.GetAuthorGitHub,
				command.ToMain:     entryStateHandler.GoToMain,
			},
			state.Any: {
				command.RestartBot:     entryStateHandler.RestartBot,
				command.UnknownCommand: entryStateHandler.UnknownCommand,
			},
		},
		repo: repo,
	}
}

func (r *Router) HandleUpdate(msg Message) entity.Response {
	cmd := command.Command(msg.Text)
	st := state.State(r.repo.ChatState(msg.Chat.ID))

	// checks the input matches the VK user ID
	if match, _ := regexp.MatchString(command.VKIDPattern, string(cmd)); match {
		cmd = command.VKUserID
	}

	if cmd == command.RestartBot {
		return r.handlers[state.Any][command.RestartBot](msg.Chat.ID, msg.Text)
	}

	if handlers, ok := r.handlers[st]; ok {
		if handler, ok := handlers[cmd]; ok {
			return handler(msg.Chat.ID, msg.Text)
		}
	}

	return r.handlers[state.Any][command.UnknownCommand](msg.Chat.ID, msg.Text)
}
