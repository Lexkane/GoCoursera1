package main

import (
	"encoding/xml"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"net/http"
)

const (
	BotToken   = "12345678:XXXXXXXXXXXXXXXXXXXXXX"
	WebhookURL = "https://XXXXXX.ngrok.io"
)

var rss = map[string]string{
	"Habr": "https://habrahabr.ru/rss/best/",
}

type RSS struct {
	Items []item `xml:"channel>item`
}
type Item struct {
	URL   string `xml :"guid"`
	Title string `xml:"title`
}

func getNews(url string) (*RSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	rss := new(RSS)
	err = xml.Unmarshal(body, rss)
	if err != nil {
		return nil, err
	}
	return rss, nil
}
func main() {
	bot, err := tgbotapi.NewBotApi(BotToken)
	if err != nil {
		panic(err)
	}
	//bot.Debug=true
	fmt.Printf("Authorized on account %s\\n", bot.Self.UserName)
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
	if err != nil {
		panic(err)
	}
	updates := bot.ListenForWebhook("/")
	go htt.ListenAndServe(":8080", nil)
	fmt.Println("start listen:8080")
	//get all messages from updates

	for update := range updates {
		if url, ok := rss[update.Message.Text]; ok {
			rss, err := getNews(url)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(
					update.Message.ChatID,
					"sorry,error happend",
				))
			}
			for _, item := range rss.Items {
				bot.Send(tgbotapi.NewMessage(
					update.Message.Chat.ID,
					item.URL+"\n"+item.Title,
				))
			}
		} else {
			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				`there is only Habr feed available`,
			))

		}
	}
}
