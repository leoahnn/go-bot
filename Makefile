target=go-bot

build:
	@go build -o bin/$(target) cmd/bot/bot.go

clean:
	@rm -rf bin/
	@rm bot.db
