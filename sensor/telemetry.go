// Package sensor provides access to sensors and telemetry data
package sensor

import (
	"math/rand"
	"time"
	"github.com/satori/go.uuid"
)

// Telemetry type holding all sensory data
type Telemetry struct {
	DeviceId string `json:"deviceId"`
	Timestamp string `json:"timestamp"`
	Value     int    `json:"value"`
}

// Sensor type for reading out values
type Sensor struct {
	maximum int
	deviceId string
	rand    *rand.Rand
}

// Create a new sensor which allows reading out of values
func CreateSensor(maximum int, source rand.Source) Sensor {
	s := Sensor{maximum, uuid.NewV4().String(), rand.New(source)}
	return s
}

// Read out the sensor
func (s Sensor) ReadOut() (t Telemetry) {
	t = Telemetry{
		DeviceId: s.deviceId,
		Timestamp: time.Now().String(),
		Value:     s.rand.Intn(s.maximum),
	}
	return
}
