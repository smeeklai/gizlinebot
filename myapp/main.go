// Copyright 2016 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// +build appengine

package myapp

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

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
	var s *Sql
	var err error
	if localDB {
		s, err = storage.NewSql(user + ":" + password + "@(" + connectionName + ":" + dbPort + ")/" + dbName)
		if err != nil {
			fmt.Sprintf("\nCould not open db: %v", err)
		}
	} else {
		s, err = storage.NewGoogleSql(user, password, connectionName, dbName)
		if err != nil {
			fmt.Sprintf("\nCould not open db: %v", err)
		}
	}

	err = line.ServeAppEngine(s, mustGetenv("CHANNEL_SECRET"), mustGetenv("CHANNEL_TOKEN"))
	if err != nil {
		fmt.Sprintf("\nError: %s\n", err)
	}
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}
