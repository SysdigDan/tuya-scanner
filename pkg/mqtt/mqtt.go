package mqtt

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/SysdigDan/tuya-scanner/models"
	"github.com/SysdigDan/tuya-scanner/pkg/application"
)

func Publish(app *application.Application, data *models.SensorData) {
	topic18 := fmt.Sprintf("%s/%s/dps/%s/state", app.Cfg.BrokerTopic, data.Name, "18")
	payload18, _ := json.Marshal(data.Power_mA)
	log.Printf("[mqtt] Publishing %s mA to %s\n", payload18, topic18)
	_ = app.MQTT.Publish(topic18, false, payload18)

	topic19 := fmt.Sprintf("%s/%s/dps/%s/state", app.Cfg.BrokerTopic, data.Name, "19")
	payload19, _ := json.Marshal(data.Power_W)
	log.Printf("[mqtt] Publishing %s W to %s\n", payload19, topic19)
	_ = app.MQTT.Publish(topic19, false, payload19)

	topic20 := fmt.Sprintf("%s/%s/dps/%s/state", app.Cfg.BrokerTopic, data.Name, "20")
	payload20, _ := json.Marshal(data.Power_V)
	log.Printf("[mqtt] Publishing %s V to %s\n", payload20, topic20)
	_ = app.MQTT.Publish(topic20, false, payload20)

	// topic := fmt.Sprintf("%s/%s/state", app.Cfg.BrokerTopic, data.Name)
	// payload, _ := json.Marshal(data.State)
	// log.Printf("[mqtt] Publishing %s state to %s\n", payload, topic)
	// _ = app.MQTT.Publish(topic, false, payload)
}
