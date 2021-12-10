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
	if err := openFile(year); err != nil {
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

func WriteMeasurement(m entity.Measurement) error {
	t := time.Now()
	if year != t.Year() {
		Close()
		if err := openFile(t.Year()); err != nil {
			panic("could not re-open CSV file for writing: " + err.Error())
		}
	}
	err = writer.Write(stringifyMeasurement(m))
	if err != nil {
		return fmt.Errorf("error while writing: %s", err.Error())
	}
	writer.Flush()
	err = writer.Error()
	if err != nil {
		return fmt.Errorf("error while flushing: %s", err.Error())
	}

	return nil
}

func stringifyMeasurement(m entity.Measurement) []string {
	s := make([]string, 4)
	s[0] = time.Now().Format(time.RFC3339)
	s[1] = fmt.Sprintf("%.1f", m.AirTemperature)
	s[2] = fmt.Sprintf("%.1f", m.Humidity)
	s[3] = fmt.Sprintf("%d", m.WaterLevel)
	return s
}

func Close() error {
	writer.Flush()
	err := writer.Error()
	if err != nil {
		return fmt.Errorf("error while writing/flushing: %s", err.Error())
	}
	err = f.Close()
	if err != nil {
		return fmt.Errorf("error while closing file: %s", err.Error())
	}
	return nil
}
