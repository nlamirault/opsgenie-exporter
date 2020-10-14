// Copyright (C) 2020 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package opsgenie

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/prometheus/common/log"
)

// Count return number of alerts
func (client *Client) Count(query string) (*alert.CountAlertResult, error) {
	log.Info("List alerts")
	return client.alert.CountAlerts(nil, &alert.CountAlertsRequest{
		Query: query,
	})
}

// List return a list of available alerts
func (client *Client) List(query string) (*alert.ListAlertResult, error) {
	log.Info("List alerts")
	return client.alert.List(nil, &alert.ListAlertRequest{
		Query: query,
	})
}
