package main

import (
	"flag"
	"telegram-bot/bot"
	"telegram-bot/handler"
	"telegram-bot/repo/inmemory"
	webapi "telegram-bot/repo/webapi/vk"
	"telegram-bot/usecase"
)

const (
	apiURL = "https://api.telegram.org/bot"
)

func main() {
	var telegramToken string
	var vkToken string

	flag.StringVar(&telegramToken, "telegram-token", "", "Telegram bot token")
	flag.StringVar(&vkToken, "vk-token", "", "VK API token")

	flag.Parse()

	if telegramToken == "" || vkToken == "" {
		panic("Telegram and VK tokens are required")
	}

	stateRepo := inmemory.NewInMemoryStateRepo()
	requestRepo := inmemory.NewInMemoryRequestRepo()
	userVKRepo := webapi.NewUserWebAPI(vkToken)

	aboutBotUsecase := usecase.NewAboutBotUsecase(requestRepo, stateRepo)
	requestUsecase := usecase.NewRequestUsecase(requestRepo, stateRepo)
	stateRouterUsecase := usecase.NewStateRouterUsecase(stateRepo)
	userVKUsecase := usecase.NewVKUserUsecase(userVKRepo, requestRepo, stateRepo)

	aboutBotStateHandler := handler.NewAboutBotHandler(aboutBotUsecase)
	requestStateHandler := handler.NewRequestHandler(requestUsecase)
	mainStateHandler := handler.NewMainStateHandler(stateRouterUsecase, requestUsecase, aboutBotUsecase)
	entryStateHandler := handler.NewEntryStateHandler(stateRouterUsecase)
	userStateHandler := handler.NewUserHandler(userVKUsecase)

	router := bot.NewRouter(stateRepo, aboutBotStateHandler, requestStateHandler, mainStateHandler, userStateHandler, entryStateHandler)
	bot := bot.New(apiURL, telegramToken, router)

	bot.AskAndServe()
}
