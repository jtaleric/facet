// Copyright 2019, Red Hat
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"github.com/metalkube/facet/pkg/common"
	"github.com/metalkube/facet/pkg/integration"
	"log"
	"net/http"
)

func HostsHandler(w http.ResponseWriter, r *http.Request) {
	hostList, err := integration.GetHosts()
	// TODO: Implement a real JSON error handler
	if err != nil {
		log.Print(err)
	}
	RespondWithJson(w, hostList)
}

// This is an example of a REST API endpoint handler which will trigger a long
// running task.  It takes a Notification channel as an argument, and returns a
// standard HTTP handler.  The idea is to create a closure over the notification
// channel.  When registering your handler with the router, you might do
// something like:
//
//   router.HandleFunc("/some-endpoint", LongRunningTaskHandler(notificationChannel))
//
//  See pkg/server/server.go for an example.
//
// In the body itself, you are expected to start the actual long-running task in
// a go routine, and quickly respond with a 2xx.
func LongRunningTaskHandler(notificationChannel chan common.Notification) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := "OK"
		go integration.PerformLongTask(notificationChannel)
		RespondWithJson(w, response)
	}
}

func BootstrapVMHandler(notificationChannel chan common.Notification) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "", http.StatusMethodNotAllowed)
			return
		}

		go integration.CreateBootstrapVM(notificationChannel)

		RespondWithJson(w, "Request accepted")

	}

}
