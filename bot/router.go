package bot

import (
	"regexp"
	"telegram-bot/entity"
	"telegram-bot/handler"
	"telegram-bot/handler/helpers"
)

type StateInformerRepo interface {
	ChatState(chatID int) entity.State
}

type CommandHandler func(chatID int, msg string) entity.Response

type Router struct {
	handlers map[entity.State]map[entity.Command]CommandHandler
	repo     StateInformerRepo
}

func restartBot(chatID int, msg string) entity.Response {
	return helpers.CreateResponseWithTwoButton(chatID, "Добро пожаловать!", entity.ToMain, entity.ToMain)
}

func NewRouter(repo StateInformerRepo, aboutBotHandler *handler.AboutBotHandler, requestHandler *handler.RequestHandler,
	mainStateHandler *handler.MainStateHandler, userStateHandler *handler.UserHandler, entryStateHandler *handler.EntryStateHandler) *Router {
	return &Router{
		handlers: map[entity.State]map[entity.Command]CommandHandler{
			entity.EntryPoint: {
				entity.ToMain: entryStateHandler.GoToMain,
			},
			entity.Main: {
				entity.InfoAboutUser:        mainStateHandler.GoToGetUser,
				entity.InfoAboutUserFriends: mainStateHandler.GoToGetUserFriends,
				entity.RequestHistory:       mainStateHandler.GoToGetRequestHistory,
				entity.InfoAboutAuthor:      mainStateHandler.GoToGetInfAboutAuthor,
			},
			entity.GetUser: {
				entity.InfoAboutCommand: userStateHandler.GetInfoAboutGeUser,
				entity.ToMain:           entryStateHandler.GoToMain,
				entity.VKUserID:         userStateHandler.GetUser,
			},
			entity.GetUserFriends: {
				entity.InfoAboutCommand: userStateHandler.GetInfoAboutGeUserFriends,
				entity.ToMain:           entryStateHandler.GoToMain,
				entity.VKUserID:         userStateHandler.GetUserFriends,
			},
			entity.GetHistory: {
				entity.InfoAboutCommand: requestHandler.GetInfoAboutChatRequests,
				entity.ToMain:           entryStateHandler.GoToMain,
			},
			entity.GetInfoAboutAuthor: {
				entity.InfoGitHub: aboutBotHandler.GetAuthorGitHub,
				entity.ToMain:     entryStateHandler.GoToMain,
			},
			entity.Any: {
				entity.RestartBot:     entryStateHandler.RestartBot,
				entity.UnknownCommand: entryStateHandler.UnknownCommand,
			},
		},
		repo: repo,
	}
}

func (r *Router) HandleUpdate(msg Message) entity.Response {
	command := entity.Command(msg.Text)
	state := entity.State(r.repo.ChatState(msg.Chat.ID))

	// checks the input matches the VK user ID
	if match, _ := regexp.MatchString(entity.VKIDPattern, string(command)); match {
		command = entity.VKUserID
	}

	if command == entity.RestartBot {
		return r.handlers[entity.Any][entity.RestartBot](msg.Chat.ID, msg.Text)
	}

	if handlers, ok := r.handlers[state]; ok {
		if handler, ok := handlers[command]; ok {
			return handler(msg.Chat.ID, msg.Text)
		}
	}

	return r.handlers[entity.Any][entity.UnknownCommand](msg.Chat.ID, msg.Text)
}
