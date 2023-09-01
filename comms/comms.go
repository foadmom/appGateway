package comms

import comms "github.com/foadmom/appGateway/comms/nats"

type CommsInterface interface {
	GetMessage() ([]byte, error)
	PutMessage(b []byte) error
}

var implementation CommsInterface

func init() {
	implementation = comms.Instance()
}

func GetMessage() ([]byte, error) {
	var _message []byte
	var _err error

	_message, _err = implementation.GetMessage()

	return []byte(_message), _err
}

func PutMessage(b []byte) error {
	return nil
}
