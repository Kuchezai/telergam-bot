package bot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"telegram-bot/entity"
)

const (
	getUpdatesEndpoint  = "/getUpdates"
	sendMessageEndpoint = "/sendMessage"
)

type TelegramBot struct {
	apiURL string
	token  string
	router *Router
}

func New(apiURL, token string, router *Router) *TelegramBot {
	return &TelegramBot{apiURL, token, router}
}

func (b *TelegramBot) AskAndServe() {
	offset := 0
	for {
		updates, err := b.getUpdates(offset)
		if err != nil {
			log.Printf("Error getting updates: %v", err)
			continue
		}
		for _, update := range updates {
			b.sendMessage(b.router.HandleUpdate(update.Message))
			offset = update.UpdateID + 1
		}

	}
}

func (b *TelegramBot) getUpdates(offset int) ([]update, error) {

	u := b.apiURL + b.token + getUpdatesEndpoint
	query := "?offset=" + strconv.Itoa(offset)

	resp, err := http.Get(u + query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result getUpdatesResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Updates, nil
}

func (b *TelegramBot) sendMessage(msg entity.Response) {
	rspn := responseWithJsonTags(msg)
	body, err := json.Marshal(rspn)
	if err != nil {
		log.Printf("Error marshaling message: %v", err)
		return
	}
	u := b.apiURL + b.token + sendMessageEndpoint
	resp, err := http.Post(u, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error marshaling message: %v", err)
		return
	}

	defer resp.Body.Close()
}

type update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type getUpdatesResult struct {
	Updates []update `json:"result"`
}

type responseWithJsonTags struct {
	ChatID      int         `json:"chat_id"`
	Text        string      `json:"text"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
