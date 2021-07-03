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
	t1 := time.Date(2018, time.February, 28, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2018, time.March, 4, 0, 0, 0, 0, time.UTC)
	tNow := time.Now()
	anniversarySinceMet := strconv.Itoa(int(tNow.Sub(t1).Hours() / 24))
	anniversarySinceStartRelationship := strconv.Itoa(int(tNow.Sub(t2).Hours() / 24))
	// テキストメッセージを生成する
	message := linebot.NewTextMessage(anniversarySinceMet + " days since met, " + anniversarySinceStartRelationship + " days since started dating")
	// テキストメッセージを友達登録しているユーザー全員に配信する
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
