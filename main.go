package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// LINE Botクライアント生成する
	// BOT にはチャネルシークレットとチャネルトークンを環境変数から読み込み引数に渡す
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	// エラーに値があればログに出力し終了する
	if err != nil {
		log.Fatal(err)
	}
	relationshipStartDate := time.Date(2018, time.March, 4, 0, 0, 0, 0, time.UTC)
	tNow := time.Now()
	anniversarySinceStartRelationship := strconv.Itoa(int(tNow.Sub(relationshipStartDate).Hours() / 24))
	// テキストメッセージを生成する
	message := linebot.NewTextMessage(anniversarySinceStartRelationship + " days since started dating")
	// テキストメッセージを友達登録しているユーザー全員に配信する
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}

	var monthlyMessage string
	switch tNow.Month() {
	case 2:
		monthlyMessage = "bought TT's Valentine present? Valentine's is on Feb. 14th"
	case 5:
		monthlyMessage = "bought TT's 520 present? Valentine's is on May 20th"
	case 9:
		monthlyMessage = "just september"
	case 10:
		monthlyMessage = "bought TT's birthday present? Her birthday is on Oct. 9th"
	case 12:
		monthlyMessage = "bought TT's Christmas present? Christmas is on Dec. 24th"
	}

	monthlyLineMessage := linebot.NewTextMessage(monthlyMessage)

	if _, err := bot.BroadcastMessage(monthlyLineMessage).Do(); err != nil {
		log.Fatal(err)
	}
}
