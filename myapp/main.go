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

	_ "github.com/go-sql-driver/mysql"

	"github.com/smeeklai/gizlinebot/storage"
	"github.com/smeeklai/gizlinebot/line"

)

func init() {
	connectionName := mustGetenv("CLOUDSQL_CONNECTION_NAME")
	user := mustGetenv("CLOUDSQL_USER")
	password := os.Getenv("CLOUDSQL_PASSWORD") // NOTE: password may be empty
	dbName := mustGetenv("CLOUDSQL_DB")

	s, err := storage.NewGoogleSql(user, password, connectionName, dbName)
	if err != nil {
		fmt.Sprintf("\nCould not open db: %v", err)
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