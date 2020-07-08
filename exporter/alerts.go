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

package exporter

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	// "github.com/nlamirault/opsgenie-exporter/opsgenie"
)

var (
	numberOfAlerts = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "alerts_number_of"),
		"Number of alerts",
		[]string{"customer", "environment"}, nil,
	)
)

// type AlertMetric struct {
// 	Count              float64            `json:"domains_being_blocked"`
// 	DNSQueriesToday    float64            `json:"dns_queries_today"`
// 	AdsBlockedToday    float64            `json:"ads_blocked_today"`
// 	AdsPercentageToday float64            `json:"ads_percentage_today"`
// 	DomainsOverTime    DomainsOverTime    `json:"domains_over_time"`
// 	AdsOverTime        AdsOverTime        `json:"ads_over_time"`
// 	TopQueries         map[string]float64 `json:"top_queries"`
// }

// func describeAlertsMetrics(ch chan<- *prometheus.Desc) {
// 	ch <- alerts
// }

func storeAlertsMetrics(ch chan<- prometheus.Metric, result alert.ListAlertResult) {
	log.Info("Store Alerts metrics")
	for _, alert := range result.Alerts {
		log.Infof("Alert: %s %s %s %v", alert.Status, alert.Message, alert.Priority, alert.Tags)
		for _, tag := range alert.Tags {
			log.Infof("Tag: %s", tag)
		}
	}
	// return nil
}
