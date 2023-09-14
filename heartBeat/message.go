package heartBeat

import (
	"time"

	m "github.com/foadmom/common/message"
	u "github.com/foadmom/common/utils"
)

// ========================================================
// these are server/host statistics
// ========================================================
type HostStats struct {
	CPU         int `json:"CPU"`
	ThreadCount int `json:"threadCount"`
	Memory      int `json:"memory"`
	Disk        int `json:"disk"`
	Network     int `json:"networkIO"`
}

const HEART_BEAT_MESSAGE_CODE = "HeartBeat"

const ACTIVE = "active"
const DRAINING = "draining"
const DRAINED = "drained"
const DISABLED = "disabled"
const ERROR = "error"
const UNKNOWN = "unknown"

// ========================================================
// These are service heartbeat indicating that a service
// is alive and indicating it's status
// ========================================================
type Service struct {
	Name        string    `json:"name"`        // Service unique name or id
	Version     string    `json:"version"`     // Version. multiple version may be active
	Host        string    `json:"host"`        // FQN of the host sending the message
	IP          string    `json:"IP"`          // optional if host is not specified
	Port        int       `json:"port"`        // if not specified then a default will be used
	Sessions    int       `json:"sessions"`    // ?
	Status      string    `json:"status"`      //
	LastUpdated time.Time `json:"lastUpdated"` // timestamp
}

type HealthStatus struct {
	HostStats HostStats `json:"hostStats"`
	Services  Service   `json:"services"`
}

type ServicePayload struct {
	MessageCode      string  `json:"messageCode"` // should be "HeartBeat"
	HeartBeatService Service `json:"Serivce`
}

type ServiceMessage struct {
	Header  m.MessageHeader
	Payload ServicePayload
}

// ============================================================================
//
// ============================================================================
var HostName string

// ============================================================================
//
// ============================================================================
func init() {
	HostName = u.HostName()
}

// ============================================================================
//
// ============================================================================
func (message *ServiceMessage) Instance() error {
	var _header m.MessageHeader
	var _hearBeat ServiceMessage

	_err := _header.Instance()
	if _err == nil {
		_hearBeat.Header = _header
		_hearBeat.Payload.MessageCode = HEART_BEAT_MESSAGE_CODE
	}

	return _err
}

func (message *ServiceMessage) ServiceInstance(name, version, status string, port int) error {
	_ = message.Instance()
	message.Payload.HeartBeatService.serviceInstance(name, version, status, port)

	return nil
}

func (s *Service) serviceInstance(name, version, status string, port int) error {
	s.Name = name
	s.Version = version
	s.Host = HostName
	s.Port = port
	s.Status = status
	s.LastUpdated = time.Now()

	return nil
}

func HealthMessageRecieved(message m.GenericMessage) {

}
