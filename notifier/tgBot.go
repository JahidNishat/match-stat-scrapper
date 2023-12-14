package notifier

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"log"
	"match_statistics_scrapper/scrapper"
	"reflect"
)

func TgBot() {
	token := viper.GetString("telegram.token")
	chatIDList := viper.GetIntSlice("telegram.chatIDList")
	fmt.Println(chatIDList)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println("Token Problem")
		log.Panic(err)
	}

	scrapData := []interface{}{}

	arrayBnxt := scrapper.ScrapsBnxt(`https://bnxtleague.com/en/player-statistics/?player_id=2882&amp;team\_id=162`)
	arrayNbl := scrapper.NblScrap("https://nbl.com.au/player/3713/853140/lachlan-anderson")
	arrayB3league := scrapper.ScrapsB3league("http://210.140.77.209/player/?key=93&amp;team=715&amp;player=43239") // Working one : https://www.b3league.jp/player/?key=93&team=2725&player=9208

	scrapData = append(scrapData, arrayBnxt, arrayNbl, arrayB3league)

	for _, data := range scrapData {
		for _, chatID := range chatIDList {
			stringified := stringifyArrayOfStructs(data)

			message := tgbotapi.NewMessage(int64(chatID), stringified)
			_, err = bot.Send(message)
			if err != nil {
				log.Println(err)
			}
		}
	}
	//message := tgbotapi.NewMessage(chatIDList, "Hello There ! ")
	//_, err = bot.Send(message)
	//if err != nil {
	//	log.Println(err)
	//}
}

func stringifyArrayOfStructs(data interface{}) string {
	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Slice {
		return "Not a slice"
	}

	var str string
	for i := 0; i < val.Len(); i++ {
		structElement := val.Index(i)
		str += stringifyStruct(structElement.Interface()) + "\n"
	}

	return str
}

func stringifyStruct(data interface{}) string {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	if val.Kind() != reflect.Struct {
		return "Not a struct"
	}

	var str string
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		str += fmt.Sprintf("%s: %v\n", fieldType.Name, field.Interface())
	}

	return str
}
