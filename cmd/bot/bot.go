package main

import (
	"flag"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/leosaysger/go-bot/internal/bot"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	debug := flag.Bool("v", false, "Debug")
	flag.Parse()

	//Set up log file
	file, err := os.OpenFile("bot.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer file.Close()
	setupLog(*debug)
	bot.StartRTM()
}

func setupLog(debug bool) {

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	//Set log size/rotation parameters
	lumberjack := &lumberjack.Logger{
		Filename:   "./bot.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	}
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(lumberjack)
}
