package entity

type Measurement struct {
	AirTemperature float64 `json:"air_temperature"`
	Humidity       float64 `json:"humidity"`
	WaterLevel     uint8   `json:"water_level"`
}
