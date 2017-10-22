package storage

type Storage interface {
	// used for after-the-fact debugging
	AddRawLineEvent(eventType, rawMsg string) error
	AddUserProfile(userId, displayName string) error
	Close() error
}
