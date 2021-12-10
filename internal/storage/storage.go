package storage

import (
	"encoding/csv"
	"fmt"
	"github.com/KaiserWerk/Greenhouse-Manager/internal/entity"
	"os"
	"time"
)

var (
	f      *os.File
	writer *csv.Writer
	err    error
	year   int
)

func init() {
	t := time.Now()
	year = t.Year()
	err = openFile(year)
	if err != nil {
		panic("could not open CSV file for writing: " + err.Error())
	}
}

func openFile(year int) error {
	f, err = os.OpenFile(fmt.Sprintf("measurements_%d.csv", year), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0744)
	if err != nil {
		return err
	}
	writer = csv.NewWriter(f)
	return nil
}

func WriteMeasurement(m *entity.Measurement) error {
	t := time.Now()
	if year != t.Year() {
		Close()
		if err := openFile(t.Year()); err != nil {
			panic("could not re-open CSV file for writing: " + err.Error())
		}
	}
	return writer.Write(stringifyMeasurement(m))
}

func stringifyMeasurement(m *entity.Measurement) []string {
	s := make([]string, 3)
	s[0] = fmt.Sprintf("%d", &m.AirTemperature)
	s[1] = fmt.Sprintf("%d", &m.Humidity)
	s[2] = fmt.Sprintf("%d", &m.WaterLevel)
	return s
}

func Close() {
	writer.Flush()
	f.Close()
}
