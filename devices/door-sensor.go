package devices

import (
	"encoding/json"
	"github.com/PhilGruber/dimmy/core"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

type DoorSensor struct {
	Device

	state        string
	triggerState string
}

type DoorSensorMessage struct {
	core.Zigbee2MqttMessage
	Contact bool `json:"contact"`
}

func MakeDoorSensor(config core.DeviceConfig) DoorSensor {
	log.Println("Creating new door sensor with topic " + config.Topic)
	s := DoorSensor{}
	s.setBaseConfig(config)
	s.MqttState = config.Topic

	s.Type = "door-sensor"
	s.Triggers = []string{"sensor"}
	s.state = ""
	s.triggerState = ""

	return s
}

func NewDoorSensor(config core.DeviceConfig) *DoorSensor {
	s := MakeDoorSensor(config)
	return &s
}

func (s *DoorSensor) GetMessageHandler(channel chan core.SwitchRequest, sw DeviceInterface) mqtt.MessageHandler {
	return func(client mqtt.Client, mqttMessage mqtt.Message) {
		log.Println("Door sensor message received")
		payload := mqttMessage.Payload()
		var data DoorSensorMessage
		err := json.Unmarshal(payload, &data)
		if err != nil {
			log.Println("Error: " + err.Error())
			return
		}
		if data.Contact {
			s.state = "closed"
		} else {
			s.state = "open"
		}
		log.Printf("Door is %s\n", s.state)

		s.triggerState = s.state
	}
}

func (s *DoorSensor) GetTriggerValue(trigger string) interface{} {
	if trigger == "sensor" {
		return s.triggerState
	}
	return nil
}

func (s *DoorSensor) ClearTrigger(trigger string) {
	if trigger == "sensor" {
		s.triggerState = ""
	}
}

func (s *DoorSensor) UpdateValue() (float64, bool) {
	return 0, false
}
func (s *DoorSensor) GetCurrent() float64 {
	return 1
}
