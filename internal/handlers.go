package bot

import (
	"fmt"
	"strings"

	tele "gopkg.in/telebot.v4"
)

func RegisterHandlers(b *tele.Bot) {
	b.Handle("/start", handleStart)
	b.Handle("/help", handleHelp)
	b.Handle("/settings", handleSettings)
	b.Handle(tele.OnText, handleText)
}

func handleStart(c tele.Context) error {
	return c.Send("Привіт 👋\n\nЯ kbot.\n\nКоманди:\n/start\n/help\n/settings")
}

func handleHelp(c tele.Context) error {
	return c.Send("Напиши мені hello, ping або будь-який текст.")
}

func handleSettings(c tele.Context) error {
	return c.Send("Налаштування поки що базові.")
}

func handleText(c tele.Context) error {
	text := strings.TrimSpace(strings.ToLower(c.Text()))

	switch text {
	case "hello", "hi", "привіт":
		return c.Send("Hello! 👋")
	case "ping":
		return c.Send("pong")
	default:
		return c.Send(fmt.Sprintf("Ти написав: %s", c.Text()))
	}
}