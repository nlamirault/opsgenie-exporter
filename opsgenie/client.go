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
	"fmt"

	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/prometheus/common/log"

	"github.com/nlamirault/opsgenie-exporter/version"
)

const (
	apiURL = "https://api.opsgenie.com/v2"
)

var (
	application = "opsgenie-exporter"
	userAgent   = fmt.Sprintf("%s/%s", application, version.Version)
)

// Metrics define Opsgenie Prometheus metrics
type Metrics struct {
	Count  int           `json:"count"`
	Tags   string        `json:"tags"`
	Alerts []alert.Alert `json:"alerts"`
}

// Client defines an Opsgenie API client
type Client struct {
	alert *alert.Client
	tags  string
}

// NewClient creates an Opsgenie client
func NewClient(apikey string, tags string) (*Client, error) {
	alertClient, err := alert.NewClient(&client.Config{
		ApiKey: apikey,
	})
	if err != nil {
		return nil, err
	}
	return &Client{
		alert: alertClient,
		tags:  tags,
	}, nil
}

// GetMetrics retrieve available metrics for the API Router
func (client *Client) GetMetrics() (*Metrics, error) {
	log.Infof("Get metrics")

	var metrics Metrics

	tagsFilter := ""
	if len(client.tags) > 0 {
		tagsFilter = fmt.Sprintf("tag=%s", client.tags)
		log.Infof("Tag used to filter queries: %s", tagsFilter)
	}

	size, err := client.Count(tagsFilter)
	log.Infof("Number of alerts: %v", size)
	metrics.Count = size.Count

	alerts, err := client.List(tagsFilter)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving alerts : %s", err)
	}
	log.Infof("Alerts: %v", alerts)
	for _, alert := range alerts.Alerts {
		log.Infof("Alert: %s %s %s %#v", alert.Status, alert.Message, alert.Priority, alert.Tags)
		for _, tag := range alert.Tags {
			log.Infof("- : %s", tag)
		}
		// log.Infof("Alert: %#v", alert)
	}
	metrics.Alerts = alerts.Alerts
	metrics.Tags = client.tags

	return &metrics, nil
}
