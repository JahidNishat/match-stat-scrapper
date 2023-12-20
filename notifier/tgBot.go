package notifier

import (
	"errors"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"match_statistics_scrapper/db"
	"match_statistics_scrapper/db/repos"
	"match_statistics_scrapper/models"
	"strconv"
	"strings"
)

var tgBot *tgbotapi.BotAPI

func GetTgBot() *tgbotapi.BotAPI {
	if tgBot != nil {
		return tgBot
	}
	token := viper.GetString("telegram.token")
	var err error
	tgBot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println("Token Problem")
		log.Panic(err)
		return nil
	}
	return tgBot
}

func LoadAdmin() {
	ss := repos.SubscriberStore{DB: db.ConnectDB()}
	adminId := viper.GetInt64("telegram.adminID")

	admin, err := ss.GetSubscriberData(adminId)
	if errors.Is(err, gorm.ErrRecordNotFound) && admin == nil {
		subs := models.Subscriber{
			ChatID:   adminId,
			Approved: true,
		}
		ss.Save(&subs)
		return
	}
	return
}

func ServeTgBot() error {
	ss := repos.SubscriberStore{DB: db.ConnectDB()}
	bot := GetTgBot()

	// Update Config From TgBOT
	updateConfig := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(updateConfig)

	// Process updates
	for update := range updates {
		if update.Message != nil {
			// Get the message text and chat ID
			messageText := update.Message.Text
			chatID := update.Message.Chat.ID
			fmt.Println(chatID)

			// Check if the message is "/start"
			if messageText == "/start" {
				// Send a welcome message
				message := tgbotapi.NewMessage(chatID, "Hello! ")
				_, err := bot.Send(message)
				if err != nil {
					log.Println(err)
				}
			} else if messageText == "/subscribe" {
				sub := &models.Subscriber{
					ChatID:   chatID,
					Approved: false,
				}
				err := ss.Save(sub)
				if err != nil {
					log.Println("Cannot Save Subscriber Info: ", err)
					continue
				}
				message := tgbotapi.NewMessage(chatID, "Admin will approve your request! Thanks! :D ")
				_, err = bot.Send(message)
				if err != nil {
					log.Println("Subscription msg rply error", err)
				}

				adminId := viper.GetInt64("telegram.adminID")
				message = tgbotapi.NewMessage(adminId,
					fmt.Sprintf("%v wants to join this channel ! To approve type `/approve %v`", chatID, chatID))
				_, err = bot.Send(message)
				if err != nil {
					log.Println("Admin approve request msg: ", err)
				}
			} else if strings.HasPrefix(messageText, "/approve ") {
				parts := strings.Fields(messageText)
				adminId := viper.GetInt64("telegram.adminID")
				if len(parts) != 2 {
					log.Println("Invalid /approve command format")
					continue
				}

				if chatID != adminId {
					log.Println("Invalid approve access")
					continue
				}

				chatIDToApprove := parts[1]
				chatID, _ := strconv.Atoi(chatIDToApprove)

				res, err := ss.GetSubscriberData(int64(chatID))
				if err != nil {
					log.Println("Could Find ChatID : ", err)
					continue
				}

				approvalMessage := tgbotapi.NewMessage(int64(chatID), "Your request has been approved!")
				_, err = bot.Send(approvalMessage)
				if err != nil {
					log.Println("Approval confirmation message error:", err)
					continue
				}

				approvalMessage = tgbotapi.NewMessage(adminId, fmt.Sprintf("ID %s has been approved!", chatIDToApprove))
				_, err = bot.Send(approvalMessage)
				if err != nil {
					log.Println("Approval confirmation message error:", err)
					continue
				}

				res.Approved = true
				err = ss.Save(res)
				if err != nil {
					log.Println("Cannot Save Subscriber Info: ", err)
				}
			}
		}
	}
	return nil
}

func PublishToSubscribers(stringifiedData string) error {
	ss := repos.SubscriberStore{DB: db.ConnectDB()}
	subscribersList, err := ss.GetAllSubscribers()
	bot := GetTgBot()
	for _, subscriber := range subscribersList {
		message := tgbotapi.NewMessage(subscriber.ChatID, stringifiedData)
		_, err = bot.Send(message)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
