package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PlatformStats struct {
	OneDay    float64 `json:"1d"`
	ThreeDay  float64 `json:"3d"`
	FiveDay   float64 `json:"5d"`
	OneWeek   float64 `json:"1w"`
	TwoWeek   float64 `json:"2w"`
	ThreeWeek float64 `json:"3w"`
	OneMonth  float64 `json:"1m"`
	Latest    float64 `json:"latest"`
}

type StatsResponse struct {
	Data map[string]PlatformStats `json:"data"`
}

var platformNames = map[string]string{
	"1": "BINANCE",
	"2": "OKX",
	"3": "BYBIT",
	"4": "BITGET",
	"5": "GATEIO",
}

func fetchStats() (*StatsResponse, error) {
	resp, err := http.Get("https://launchpooler.fly.dev/api/earn_apy/USDT/stats")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var stats StatsResponse
	if err := json.Unmarshal(body, &stats); err != nil {
		return nil, err
	}

	return &stats, nil
}

func formatStatsMessage(stats *StatsResponse) string {
	var msg strings.Builder
	msg.WriteString("📊 USDT APY Stats\n\n")

	// Define platform order
	platformOrder := []string{"1", "2", "3", "4", "5"}

	for _, platformID := range platformOrder {
		stat, exists := stats.Data[platformID]
		if !exists {
			continue
		}

		platformName := platformNames[platformID]
		msg.WriteString(fmt.Sprintf("🏦 %s:\n", platformName))
		msg.WriteString(fmt.Sprintf("📍 Latest: %.2f%%\n", stat.Latest))
		msg.WriteString(fmt.Sprintf("1d: %.2f%% | 3d: %.2f%% | 5d: %.2f%%\n", stat.OneDay, stat.ThreeDay, stat.FiveDay))
		msg.WriteString(fmt.Sprintf("1w: %.2f%% | 2w: %.2f%% | 3w: %.2f%%\n", stat.OneWeek, stat.TwoWeek, stat.ThreeWeek))
		if stat.OneMonth > 0 {
			msg.WriteString(fmt.Sprintf("1m: %.2f%%\n", stat.OneMonth))
		}
		msg.WriteString("\n")
	}

	msg.WriteString("Source: https://coinsider.app/earn")
	return msg.String()
}

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is not set")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "stats":
			stats, err := fetchStats()
			if err != nil {
				msg.Text = "Sorry, failed to fetch stats. Please try again later."
				log.Printf("Error fetching stats: %v", err)
			} else {
				msg.Text = formatStatsMessage(stats)
			}
		default:
			msg.Text = "Available commands:\n/stats - Get USDT stats"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}
