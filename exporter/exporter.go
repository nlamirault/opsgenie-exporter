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
	"strings"
	// "sync"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"github.com/nlamirault/opsgenie-exporter/opsgenie"
)

const (
	namespace = "opsgenie"
)

var (
	up = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "up"),
		"Was the last query of Opsgenie successful.",
		nil, nil,
	)

	alertCountDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "alert_count"),
		"Number of alert",
		[]string{"tags"}, nil,
	)

	alertStatusDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "alert_status"),
		"The current status of the alert (1: up, 0: down)",
		[]string{"id", "name", "priority", "status", "source", "owner", "tags"}, nil,
	)
)

// Exporter collects Opsgenie stats from the given server and exports them using
// the prometheus metrics package.
type Exporter struct {
	Opsgenie *opsgenie.Client
}

// NewExporter returns an initialized Exporter.
func NewExporter(apikey string, tags string) (*Exporter, error) {
	log.Infof("Setup Opsgenie exporter")
	client, err := opsgenie.NewClient(apikey, tags)
	if err != nil {
		return nil, err
	}
	return &Exporter{
		Opsgenie: client,
	}, nil
}

// Describe describes all the metrics ever exported by the Opsgenie exporter.
// It implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
	ch <- alertCountDesc
	ch <- alertStatusDesc
}

// Collect the stats from channel and delivers them as Prometheus metrics.
// It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	log.Infof("Opsgenie exporter starting")

	metrics, err := e.Opsgenie.GetMetrics()
	if err != nil {
		log.Errorf("Opsgenie API error: %s", err)
		ch <- prometheus.MustNewConstMetric(
			up, prometheus.GaugeValue, 0,
		)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		alertCountDesc,
		prometheus.GaugeValue,
		float64(metrics.Count),
		metrics.Tags,
	)

	for _, alert := range metrics.Alerts {
		log.Infof("ID: %s", alert.Id)
		log.Infof("Name: %s", alert.Alias)
		log.Infof("Priority: %s", alert.Priority)
		log.Infof("Status: %s", alert.Status)
		log.Infof("Owner: %s", alert.Owner)
		log.Infof("Tags: %s", strings.Join(alert.Tags, ","))

		var status float64
		if alert.Status == "closed" {
			status = 0
		} else {
			status = 1
		}

		ch <- prometheus.MustNewConstMetric(
			alertStatusDesc,
			prometheus.GaugeValue,
			status,
			alert.Id,
			alert.Alias,
			fmt.Sprintf("%s", alert.Priority),
			alert.Status,
			alert.Source,
			alert.Owner,
			strings.Join(alert.Tags, ","),
		)
	}

	ch <- prometheus.MustNewConstMetric(
		up, prometheus.GaugeValue, 1,
	)
	log.Infof("Opsgenie exporter finished")
}

// func storeMetric(ch chan<- prometheus.Metric, value float64, desc *prometheus.Desc, labels ...string) {
// 	ch <- prometheus.MustNewConstMetric(
// 		desc, prometheus.GaugeValue, value, labels...)
// }
