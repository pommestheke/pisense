// Package sensor provides access to sensors and telemetry data
package sensor

import (
	"testing"
	"math/rand"
	"time"
)

// Test creation of sensor
func TestCreateSensor(t *testing.T) {
	s := CreateSensor(10, rand.NewSource(time.Now().UnixNano()))
	if s.maximum != 10 {
		t.Errorf("Expected maximum of 10, received %d", s.maximum)
	}
}
