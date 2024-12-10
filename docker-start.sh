#!/bin/bash

# Check if .env file exists
if [ ! -f .env ]; then
    echo "Creating .env file from .env.example..."
    cp .env.example .env
    echo "Please edit .env file and set your Telegram Bot Token"
    exit 1
fi

# Build and start the container
docker-compose up -d --build

echo "Bot is running in Docker container"
echo "Check logs with: docker-compose logs -f"
