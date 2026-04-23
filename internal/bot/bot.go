package bot

import (
	"fmt"
	"os"
	"strings"
	"time"

	tele "gopkg.in/telebot.v4"
)

func Start() error {
	token := strings.TrimSpace(os.Getenv("TELE_TOKEN"))
	if token == "" {
		return fmt.Errorf("TELE_TOKEN environment variable is not set")
	}

	b, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return fmt.Errorf("failed to create bot: %w", err)
	}

	RegisterHandlers(b)

	fmt.Println("Bot is running...")
	b.Start()
	return nil
}

func RegisterHandlers(b *tele.Bot) {
	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Привіт 👋")
	})

	b.Handle("/help", func(c tele.Context) error {
		return c.Send("Допомога")
	})

	b.Handle("/settings", func(c tele.Context) error {
		return c.Send("Налаштування")
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send("Ти написав: " + c.Text())
	})
}
