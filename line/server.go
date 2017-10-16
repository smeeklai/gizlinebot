package line

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VagabondDataNinjas/gizlinebot/storage"
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

func (s *LineServer) Serve() error {

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/linewebhook", func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("\nReqqqq: %s\n", req)
		events, err := s.Bot.ParseRequest(req)
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
			if event.Type == linebot.EventTypeFollow {

			}

			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = s.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})

	log.Printf("Starting http server on port %s", s.Port)
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+s.Port, nil); err != nil {
		log.Fatal(err)
	}

	return nil
}
