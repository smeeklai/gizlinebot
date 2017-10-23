package myapp

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/smeeklai/gizlinebot/line"
	"github.com/smeeklai/gizlinebot/storage"
)

func init() {
	connectionName := mustGetenv("GIZLB_SQL_HOST")
	user := mustGetenv("GIZLB_SQL_USER")
	password := os.Getenv("GIZLB_SQL_PASS") // NOTE: password may be empty
	dbName := mustGetenv("GIZLB_SQL_DB")
	dbPort := os.Getenv("GIZLB_SQL_PORT") // NOTE: dbPort may be empty
	localDB, err := strconv.ParseBool(mustGetenv("LOCAL_DB"))
	if err != nil {
		fmt.Sprintf("\nLOCAL_DB value error")
	}

	if localDB {
		s, err := storage.NewSql(user + ":" + password + "@(" + connectionName + ":" + dbPort + ")/" + dbName)
		if err != nil {
			fmt.Sprintf("\nCould not open db: %v", err)
		}
		err = line.ServeAppEngine(s, mustGetenv("CHANNEL_SECRET"), mustGetenv("CHANNEL_TOKEN"))
		if err != nil {
			fmt.Sprintf("\nError: %s\n", err)
		}
	} else {
		s, err := storage.NewGoogleSql(user, password, connectionName, dbName)
		if err != nil {
			fmt.Sprintf("\nCould not open db: %v", err)
		}
		err = line.ServeAppEngine(s, mustGetenv("CHANNEL_SECRET"), mustGetenv("CHANNEL_TOKEN"))
		if err != nil {
			fmt.Sprintf("\nError: %s\n", err)
		}
	}
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}
