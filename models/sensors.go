package models

type SensorData struct {
	Name     string
	Switch   bool
	Power_mA float64
	Power_W  float64
	Power_V  float64
	State    string
}
