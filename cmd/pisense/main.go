// Package main provides the pisense client command
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/amenzhinsky/iothub/iotdevice"
	iotmqtt "github.com/amenzhinsky/iothub/iotdevice/transport/mqtt"
	"github.com/pommestheke/pisense/sensor"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	// create our sensor
	s := sensor.CreateSensor(20, rand.NewSource(time.Now().UnixNano()))

	// connect to the iot hub
	c, err := iotdevice.NewFromConnectionString(
		iotmqtt.New(), os.Getenv("IOTHUB_DEVICE_CONNECTION_STRING"),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err = c.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}

	// loop and generate events
	for {
		// read sensor
		t := s.ReadOut()
		res, _ := json.Marshal(t)

		fmt.Println("sending: " + string(res))

		// send a device-to-cloud message
		if err = c.SendEvent(context.Background(), []byte(res)); err != nil {
			log.Fatal(err)
			return
		}

		// wait until next readout
		time.Sleep(60 * time.Second)
	}
}
