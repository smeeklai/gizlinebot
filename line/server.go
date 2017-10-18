package line

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/smeeklai/gizlinebot/storage"
	"github.com/line/line-bot-sdk-go/linebot"
)

type LineServer struct {
	Port    string
	Bot     *linebot.Client
	Storage storage.Storage
}

func NewLineServer(port string, storage storage.Storage, secret, token string) (server *LineServer, err error) {
	bot, err := linebot.New(secret, token)
	if err != nil {
		return server, err
	}

	return &LineServer{
		Port:    port,
		Bot:     bot,
		Storage: storage,
	}, nil
}

func (ls *LineServer) Serve() error {

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/linewebhook", func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("\nReqqqq: %s\n", req)
		events, err := ls.Bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		fmt.Printf("\nevents: %+v", events)
		for _, event := range events {
			eventString, err := json.Marshal(event)
			if err != nil {
				log.Printf("[err] Could not marshal event: %+v; err: %s", event, err)
				continue
			}
			err = ls.Storage.AddRawLineEvent(string(event.Type), event.ReplyToken, string(eventString))
			if err != nil {
				log.Printf("[err] Could not store event: %+v; err: %s", event, err)
				continue
			}

			if event.Type == linebot.EventTypeFollow {

			}

			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = ls.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})

	log.Printf("Starting http server on port %s", ls.Port)
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+ls.Port, nil); err != nil {
		log.Fatal(err)
	}

	return nil
}
