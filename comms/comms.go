package comms

import (
	"encoding/json"

	s "github.com/foadmom/appGateway/heartBeat"
	comms "github.com/foadmom/common/comms"
)

var implementation comms.CommsInterface
var Channel string = "appGateway"

func init() {
	implementation = comms.Instance()
}

func GetMessage(channel string) ([]byte, error) {
	var _message []byte
	var _err error

	_message, _err = implementation.GetMessage(channel)

	return []byte(_message), _err
}

func SendHearBeat(message s.ServiceMessage) error {
	b, _err := json.Marshal(message)
	if _err == nil {
		_err = implementation.PutMessage(Channel, b)
	}
	return _err
}
