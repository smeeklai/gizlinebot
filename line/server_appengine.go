package line

import (
	"encoding/json"
	"log"
	// "log"
	"net/http"

	"google.golang.org/appengine"
	aelog "google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"

	// "github.com/VagabondDataNinjas/gizlinebot/storage"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/linebot/httphandler"
	"github.com/smeeklai/gizlinebot/storage"
)

func ServeAppEngine(storage storage.Storage, secret, token string) error {
	handler, err := httphandler.New(secret, token)
	if err != nil {
		return err
	}

	// Setup HTTP Server for receiving requests from LINE platform
	handler.HandleEvents(func(events []*linebot.Event, r *http.Request) {
		ctx := appengine.NewContext(r)
		bot, err := handler.NewClient(linebot.WithHTTPClient(urlfetch.Client(ctx)))
		if err != nil {
			log.Printf("\nError: %s\n", err)
			aelog.Errorf(ctx, "%v", err)
			return
		}
		for _, event := range events {
			eventString, err := json.Marshal(event)
			if err != nil {
				aelog.Infof(ctx, "[err] Could not marshal err: %s", err)
				return
			}
			err = storage.AddRawLineEvent(string(event.Type), event.ReplyToken, string(eventString))
			if err != nil {
				aelog.Infof(ctx, "[err] Could store event err: %s", err)
				return
			}

			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).WithContext(ctx).Do(); err != nil {
						log.Printf("\nError: %s\n", err)
						aelog.Errorf(ctx, "%v", err)
					}
				}
			}
		}
	})
	http.Handle("/linewebhook", handler)

	return nil
}
