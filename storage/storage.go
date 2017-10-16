package storage

type Storage interface {
	AddRawLineEvent(eventType, replyToken, rawMsg string) error
	Close() error
}
