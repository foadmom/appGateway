package service

import (
	"testing"
	"time"

	m "github.com/foadmom/appGateway/heartBeat"
)

func Test_addNewService(t *testing.T) {
	// var _err error
	var _service_A m.Service = m.Service{Name: "Service-A", LastUpdated: time.Now()}
	time.Sleep(1 * time.Second)
	var _service_B m.Service = m.Service{Name: "Service-B", LastUpdated: time.Now()}
	time.Sleep(1 * time.Second)
	var _service_C m.Service = m.Service{Name: "Service-C", LastUpdated: time.Now()}

	addNewService(_service_A)
	addNewService(_service_B)
	addNewService(_service_C)

}
