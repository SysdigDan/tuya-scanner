package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SysdigDan/tuya-scanner/pkg/tuya"
	"github.com/SysdigDan/tuya-scanner/pkg/application"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	app, err := application.Get()
	if err != nil {
		panic("unable to initialize application config: " + err.Error())
	}

	// load prometheus metrics endpoint and connect broker
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<html><head><title>tuya-scanner</title></head><body><h1>tuya-scanner</h1><p><a href="/metrics">Metrics</a></p></body></html>`))
		})
		http.Handle("/metrics", promhttp.Handler())
		log.Println("Prometheus metrics listening on", app.Cfg.ListeningAddress)
		err := http.ListenAndServe(app.Cfg.ListeningAddress, nil)
		if err != http.ErrServerClosed {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	// connect with broker
	err = app.MQTT.Connect()
	if err != nil {
		log.Fatal("oops: ", err)
	}

	// starts device scan
	tuya.TuyaScanner(app)

}
