package controller

import (
	"encoding/json"
	"fmt"
	"log"

	c "github.com/foadmom/appGateway/comms"
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

	for {
		_jsonMessage, _err = c.GetMessage()
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
