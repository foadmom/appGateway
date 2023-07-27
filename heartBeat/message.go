package heartBeat

import (
	"time"

	m "github.com/foadmom/common/message"
)

type HostStats struct {
	CPU         int `json:"CPU"`
	ThreadCount int `json:"threadCount"`
	Memory      int `json:"memory"`
	Disk        int `json:"disk"`
	Network     int `json:"networkIO"`
}

type Service struct {
	Name        string    `json:"name"`
	Host        string    `json:"Host"`
	IP          string    `json:"IP"` // optional
	Port        int       `json:"port"`
	Sessions    int       `json:"sessions"`
	Status      string    `json:"status"`
	LastUpdated time.Time `json:"lastUpdated"`
}

type HealthStatus struct {
	HostStats HostStats `json:"hostStats"`
	Services  []Service `json:"services"`
}

type HeartBeatPayload struct {
	MessageCode string
	Data        HealthStatus
}

type HeartbeatMessage struct {
	Header  m.MessageHeader
	Payload HeartBeatPayload
}

func (h *HeartbeatMessage) Instance() error {
	var _header m.MessageHeader
	var _hearBeat HeartbeatMessage

	_err := _header.Instance()
	if _err == nil {
		_hearBeat.Header = _header
		_hearBeat.Payload.MessageCode = m.HEART_BEAT_MESSAGE_CODE
	}

	return _err
}

func HealthMessageRecieved(message m.GenericMessage) {

}
