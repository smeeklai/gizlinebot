package storage

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Sql struct {
	Db *sql.DB
}

func NewSql(conDsn string) (s *Sql, err error) {
	db, err := sql.Open("mysql", conDsn)
	if err != nil {
		return s, err
	}

	return &Sql{
		Db: db,
	}, nil
}

func NewGoogleSql(user, password, connectionName, dbName string) (s *Sql, err error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/%s", user, password, connectionName, dbName))
	if err != nil {
		return s, err
	}
	fmt.Println("Connected to Google Sql database")
	return &Sql{
		Db: db,
	}, nil
}

func (s *Sql) Close() error {
	return s.Db.Close()
}

func (s *Sql) AddRawLineEvent(eventType, replyToken, rawevent string) error {
	stmt, err := s.Db.Prepare("INSERT INTO linebot_raw_events(eventtype, replytoken, rawevent, timestamp) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(eventType, replyToken, rawevent, int32(time.Now().UTC().Unix()))
	if err != nil {
		return err
	}

	return nil
}
