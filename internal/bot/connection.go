package bot

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/nlopes/slack"
)

// StartRTM initalizes the websocket connection to slack and handles incoming requests
func StartRTM() {
	config := Config()
	api := slack.New(config.BotToken)
	rtm := api.NewRTM()
	var db SQLite
	db.Init()
	go rtm.ManageConnection()
	log.Debug("starting RTM")
	reciever(rtm, api, &db)
}

func reciever(rtm *slack.RTM, api *slack.Client, db *SQLite) {
	var self *slack.UserDetails
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			self = ev.Info.User
			log.Debug(fmt.Sprintf("Connection counter:%v", ev.ConnectionCount))
		case *slack.MessageEvent:
			Parse(self, ev, rtm, api, db)
		case *slack.ReactionAddedEvent:
			PlusKudo(self, ev, rtm, db)
		case *slack.ReactionRemovedEvent:
			MinusKudo(self, ev, rtm, db)
		case *slack.PresenceChangeEvent:
			//ignore
		case *slack.ReconnectUrlEvent:
			//ignore
		case *slack.LatencyReport:
			log.Debug(fmt.Sprintf("Current latency: %v\n", ev.Value))
		case *slack.RTMError:
			log.Error(fmt.Sprintf("RTM Error: %v", ev.Error()))
		case *slack.InvalidAuthEvent:
			log.Error("Invalid credentials")
			return
		default:
			log.WithFields(log.Fields{
				"event": ev,
			}).Debug("Unhandled event type")
		}
	}
}
