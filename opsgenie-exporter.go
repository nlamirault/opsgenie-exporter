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

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"

	"github.com/nlamirault/opsgenie-exporter/exporter"
	"github.com/nlamirault/opsgenie-exporter/version"
)

const (
	banner = "opsgenie_exporter - %s\n"
)

var (
	debug       bool
	vrs         bool
	port        int
	metricsPath string
	apikey      string
	tags        string
)

func init() {
	// parse flags
	flag.BoolVar(&vrs, "version", false, "print version and exit")
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.IntVar(&port, "port", 9158, "port to listen on")
	flag.StringVar(&metricsPath, "metrics-path", "/metrics", "Path under which to expose metrics.")
	flag.StringVar(&apikey, "apikey", "", "Opsgenie API key")
	flag.StringVar(&tags, "tags", "", "tag list separated by commas")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(banner, version.Version))
		flag.PrintDefaults()
	}

	flag.Parse()
	if vrs {
		fmt.Printf("%s", version.Version)
		os.Exit(0)
	}

}

func main() {
	if len(apikey) == 0 {
		usageAndExit("Opsgenie API key cannot be empty.", 1)
	}

	exporter, err := exporter.NewExporter(apikey, tags)
	if err != nil {
		log.Errorf("Can't create exporter : %s", err)
		os.Exit(1)
	}

	log.Infoln("Register exporter")
	registry := prometheus.NewPedanticRegistry()
	registry.MustRegister(
		exporter,
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
		prometheus.NewGoCollector(),
	)

	http.Handle(metricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Opsgenie Exporter</title></head>
             <body>
             <h1>Opsgenie Exporter</h1>
             <p><a href='` + metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})

	log.Infof("Opsgenie Exporter %s listening on http://0.0.0.0:%d\n", version.Version, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n")
	}
	flag.Usage()
	os.Exit(exitCode)
}
