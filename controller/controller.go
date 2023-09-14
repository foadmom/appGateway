package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	c "github.com/foadmom/appGateway/comms"
	h "github.com/foadmom/appGateway/heartBeat"
	m "github.com/foadmom/common/message"
)

// ========================================================
// GetMessage
// Check Message type
// Convert to struct
// Pass it to the appropriate func(struct) for processing
// ========================================================

func MainLoop() {
	var _err error
	var _message m.GenericMessage
	var _jsonMessage []byte

	go sendHearBeat()
	for {
		_jsonMessage, _err = c.GetMessage("")
		if _err != nil {
			log.Printf("controller.MainLoop: Error returned from comms.GetMessage(). Error=%v\n", _err)
		} else {
			_err = json.Unmarshal(_jsonMessage, &_message)
			if _err != nil {
				log.Printf("controller.MainLoop: Error returned from json.Unmarshall. Error=%v\n", _err)
			} else {
				switch _message.Payload.MessageCode {
				case "HeartBeat":
					{
						fmt.Printf("Message received=%v\n", _jsonMessage)
					}
				}

			}
		}
	}
}

func sendHearBeat() {
	var _err error
	_ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-_ticker.C:
			// send a heartbeat to the comms channel
			var _message h.ServiceMessage = h.ServiceMessage{}
			_err = _message.ServiceInstance("name", "V1.0", "ACTIVE", 6661)
			if _err == nil {
				_err = c.SendHearBeat(_message)
				if _err == nil {
					fmt.Printf("message sent=%v\n", _message)
				} else {
					log.Printf("error sending heartBeat. Error = %v\n", _err)
				}
			}
		}

	}
}
