# Telegram USDT Stats Bot

A Telegram bot that provides USDT APY and TVL statistics.

## Features

- `/stats` command to fetch current USDT statistics
- Displays APY, TVL, and last update time
- Easy to use interface

## Setup

1. Install Go (version 1.21 or higher)
2. Clone this repository
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Set your Telegram bot token as an environment variable:
   ```bash
   export TELEGRAM_BOT_TOKEN="your_bot_token_here"
   ```
5. Run the bot:
   ```bash
   go run main.go
   ```

## Usage

1. Start a chat with your bot on Telegram
2. Send `/stats` to get current USDT statistics

## Environment Variables

- `TELEGRAM_BOT_TOKEN`: Your Telegram bot token (get it from [@BotFather](https://t.me/botfather))
