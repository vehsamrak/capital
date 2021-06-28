package main

import (
	"bytes"
	"os"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"

	"github.com/vehsamrak/capital/internal/logger"
	"github.com/vehsamrak/capital/internal/telegram"
)

func main() {
	log.SetFormatter(&logger.TextFormatter{})

	token := os.Getenv("TELEGRAM_APPLICATION_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Infof("Authorized on account @%s", bot.Self.UserName)

	introUpdateConfig := tgbotapi.NewUpdate(0)
	introUpdateConfig.Timeout = 60
	introUpdates, _ := bot.GetUpdatesChan(introUpdateConfig)

	for update := range introUpdates {
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		var inputText string
		var chatId int64

		if update.Message != nil {
			inputText = update.Message.Text
			chatId = update.Message.Chat.ID
		}

		if update.CallbackQuery != nil {
			inputText = update.CallbackQuery.Data
			chatId = update.CallbackQuery.Message.Chat.ID
			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
		}

		// user := userProvider.FromTelegramUpdate(update)

		commandName, commandParameters := parseCommand(inputText)
		log.Infof(
			"[%s] %s %s",
			update.Message.From.UserName,
			commandName,
			commandParameters,
		)

		// commandResult := commandHandler.HandleCommand(user, commandName, commandParameters)

		output := telegram.Output{Text: "test"}
		output.ChatID = chatId

		if output.ReplyMarkup == nil {
			output.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}

		_, err = bot.Send(output.GenerateChattable())
		if err != nil {
			// TODO[petr]: handle error
			log.WithError(err).Warning("Output error occured")
			return
		}
	}
}

func parseCommand(rawCommand string) (commandName string, commandParameters []string) {
	rawCommand = strings.TrimSpace(string(bytes.Trim([]byte(rawCommand), "\r\n\x00")))
	commandWithParameters := strings.Fields(rawCommand)

	if len(commandWithParameters) == 0 {
		return
	}

	commandName = strings.ToLower(commandWithParameters[0])

	return commandName, commandWithParameters[1:]
}
