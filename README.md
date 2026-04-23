# kbot

Telegram bot written in Go using Cobra and Telebot.

## Features

- `/start` — start command
- `/help` — help command
- `/settings` — settings command
- text message handling

## Requirements

- Go
- Telegram bot token in environment variable `TELE_TOKEN`

## Run locally

```bash
go mod tidy
export TELE_TOKEN="your_token"
go run .git add .