package bot

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/nlopes/slack"
)

// Parse handles a given message event
func Parse(self *slack.UserDetails, ev *slack.MessageEvent, rtm *slack.RTM, api *slack.Client, db *SQLite) {
	var evText string
	var evUser string
	log.Debug(ev)
	if ev.Msg.Text != "" {
		evText = ev.Msg.Text
		evUser = ev.Msg.User
	} else if ev.Msg.DeletedTimestamp == "" && ev.Msg.Attachments == nil { //not a delete event
		// get thread tex
		evText = ev.SubMessage.Text
		evUser = ev.SubMessage.User
	}
	if strings.EqualFold(evText, "!addme") {
		log.Debug()
		db.InsertUser(evUser)
	} else if strings.EqualFold(evText, "!score") {
		score := db.GetKudos(evUser)
		userInfo, err := api.GetUserInfo(evUser)
		if err != nil {
			log.Error("could not get user data for %s", evUser)
		}
		output := fmt.Sprintf("%s, you have %d kudos", userInfo.Profile.RealName, score)
		rtm.SendMessage(rtm.NewOutgoingMessage(output, ev.Msg.Channel))
	} else if strings.EqualFold(evText, "!vote") {
		log.Info("vote!")
	}
	// rtm.SendMessage(rtm.NewOutgoingMessage("hello", ev.Msg.Channel))
}

// PlusKudo adds kudos to a user
func PlusKudo(self *slack.UserDetails, ev *slack.ReactionAddedEvent, rtm *slack.RTM, db *SQLite) {
	log.Debug(ev)
	db.PlusKudo(ev.ItemUser)
}

// MinusKudo subtracts kudos from a user
func MinusKudo(self *slack.UserDetails, ev *slack.ReactionRemovedEvent, rtm *slack.RTM, db *SQLite) {
	log.Debug(ev)
	db.MinusKudo(ev.ItemUser)
}
