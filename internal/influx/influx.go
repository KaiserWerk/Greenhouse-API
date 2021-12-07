package influx

import (
	"github.com/KaiserWerk/Greenhouse-Manager/internal/config"
	"github.com/KaiserWerk/Greenhouse-Manager/internal/entity"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"time"
)

var (
	client   influxdb2.Client
	writeAPI api.WriteAPI
	ticker   *time.Ticker
)

func init() {
	client = influxdb2.NewClient(config.GetInfluxUrl(), config.GetInfluxKey())
	writeAPI = client.WriteAPI(config.GetInfluxOrg(), config.GetInfluxBucket())

	ticker = time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			flush()
		}
	}
}

func flush() {
	writeAPI.Flush()
}

func Close() {
	client.Close()
}

func InsertMeasurement(m entity.Measurement) {
	writeAPI.WritePoint(influxdb2.NewPoint(
		"greenhouse_sensor_stats",
		map[string]string{
			"app":    "GreenhouseManager",
			"vendor": "KaiserWerk",
		},
		map[string]interface{}{
			"air_temperature": m.AirTemperature,
			"humidity":        m.Humidity,
			"water_level":     m.WaterLevel,
		},
		time.Now(),
	))
}
