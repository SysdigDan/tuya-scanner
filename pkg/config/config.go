package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	ListeningAddress string `mapstructure:"LISTENING_ADDRESS"`
	BrokerAddress    string `mapstructure:"BROKER_ADDRESS"`
	BrokerPort       int    `mapstructure:"BROKER_PORT"`
	BrokerUser       string `mapstructure:"BROKER_USER"`
	BrokerPassword   string `mapstructure:"BROKER_PASSWORD"`
	BrokerTopic      string `mapstructure:"BROKER_TOPIC"`
	ClientID         string `mapstructure:"CLIENT_ID"`
	Frequency        int    `mapstructure:"FREQ"`
}

type DeviceConfig []struct {
	GwID string `json:"gwId"`
	Key  string `json:"key"`
	Type string `json:"type"`
	Name string `json:"name"`
}

// LoadConfig reads configuration from file environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("scanner")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; using env variables
			viper.BindEnv("LISTENING_ADDRESS")
			viper.BindEnv("BROKER_ADDRESS")
			viper.BindEnv("BROKER_PORT")
			viper.BindEnv("BROKER_USER")
			viper.BindEnv("BROKER_PASSWORD")
			viper.BindEnv("BROKER_TOPIC")
			viper.BindEnv("CLIENT_ID")
			viper.BindEnv("FREQ")
		}
	}

	err = viper.Unmarshal(&config)
	return
}

func LoadDevices(path string) (devices DeviceConfig, err error) {
	devicefile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening devices.json file: ", err)
	}

	err = json.Unmarshal([]byte(devicefile), &devices)
	return
}
