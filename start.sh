#!/bin/bash

# 编译 Go 程序
go build -o telegram-bot

# 使用 nohup 和 caffeinate 在后台运行程序
nohup caffeinate -i ./telegram-bot > bot.log 2>&1 &

echo "Bot is running in background. Check bot.log for output."
