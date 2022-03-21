package exporter

import (
	"log"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	logPower_mAGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "power",
			Name:      "amps",
			Help:      "Power Amps.",
		},
		[]string{
			"sensor",
			"name",
		},
	)
	logPower_WGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "power",
			Name:      "watts",
			Help:      "Power Watts.",
		},
		[]string{
			"sensor",
			"name",
		},
	)
	logPower_VGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "power",
			Name:      "voltage",
			Help:      "Power Voltage.",
		},
		[]string{
			"sensor",
			"name",
		},
	)
)

const Sensor = "TUYA"

const ExpiryAtc = 2.5 * 10 * time.Second
const ExpiryStock = 2.5 * 10 * time.Minute
const ExpiryConn = 2.5 * 10 * time.Second

var expirers = make(map[string]*time.Timer)
var expirersLock sync.Mutex

func bump(name string, expiry time.Duration) {
	expirersLock.Lock()
	if t, ok := expirers[name]; ok {
		t.Reset(expiry)
	} else {
		expirers[name] = time.AfterFunc(expiry, func() {
			log.Println("[exporter] Device expiring:", name)
			logPower_mAGauge.DeleteLabelValues(Sensor, name)
			logPower_WGauge.DeleteLabelValues(Sensor, name)
			logPower_VGauge.DeleteLabelValues(Sensor, name)

			expirersLock.Lock()
			delete(expirers, name)
			expirersLock.Unlock()
		})
	}
	expirersLock.Unlock()
}

func LogPrometheusData(name string, sw bool, Power_mA, Power_W, Power_V float64) {

	bump(name, ExpiryAtc)

	logPower_mA(name, Power_mA)
	logPower_W(name, Power_W)
	logPower_V(name, Power_V)
}

func logPower_mA(name string, Power_mA float64) {
	logPower_mAGauge.WithLabelValues(Sensor, name).Set(Power_mA)
	log.Printf("[exporter] Publishing %.1f mA for %s sensor\n", Power_mA, name)
}

func logPower_W(name string, Power_W float64) {
	logPower_WGauge.WithLabelValues(Sensor, name).Set(Power_W)
	log.Printf("[exporter] Publishing %.1f W for %s sensor\n", Power_W, name)
}

func logPower_V(name string, Power_V float64) {
	logPower_VGauge.WithLabelValues(Sensor, name).Set(Power_V)
	log.Printf("[exporter] Publishing %.1f V for %s sensor\n", Power_V, name)
}
