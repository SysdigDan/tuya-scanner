package application

import (
	"log"

	"github.com/SysdigDan/tuya-scanner/pkg/agent"
	"github.com/SysdigDan/tuya-scanner/pkg/config"
)

// Application holds commonly used app wide data, for ease of use
type Application struct {
	MQTT    *agent.Agent
	Cfg     config.Config
	Devices config.DeviceConfig
}

// Get captures env vars, establishes broker connection and keeps/returns
func Get() (*Application, error) {
	env, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
		return nil, err
	}

	devices, err := config.LoadDevices("./config/devices.json")
	if err != nil {
		log.Fatal("cannot load config:", err)
		return nil, err
	}

	client := agent.NewAgent(env.BrokerAddress, env.BrokerPort, env.BrokerUser, env.BrokerPassword, env.ClientID)

	return &Application{
		MQTT:    client,
		Cfg:     env,
		Devices: devices,
	}, nil
}
