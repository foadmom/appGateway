package service

import (
	"testing"
	"time"

	m "github.com/foadmom/appGateway/heartBeat"
)

func Test_addNewService(t *testing.T) {
	// var _err error
	var _service_A m.Service = m.Service{Name: "Service-A", Host: "localHost", Port: 7666, LastUpdated: time.Now()}
	time.Sleep(1 * time.Second)
	var _service_B m.Service = m.Service{Name: "Service-B", Host: "localHost", Port: 7555, LastUpdated: time.Now()}
	time.Sleep(1 * time.Second)
	var _service_C m.Service = m.Service{Name: "Service-C", Host: "localHost", Port: 7444, LastUpdated: time.Now()}
	time.Sleep(1 * time.Second)
	var _service_B2 m.Service = m.Service{Name: "Service-B", Host: "localHost", Port: 7555, LastUpdated: time.Now()}
	time.Sleep(1 * time.Second)
	var _service_A2 m.Service = m.Service{Name: "Service-A", Host: "localHost", Port: 5555, LastUpdated: time.Now()}

	UpdateServiceInfo(_service_A)
	UpdateServiceInfo(_service_B)
	UpdateServiceInfo(_service_C)
	UpdateServiceInfo(_service_A2)
	UpdateServiceInfo(_service_B2)

	printCache(t)
	PrintCache()

}

// ============================================================================
// ============== debug functions
// ============================================================================
func printCache(t *testing.T) {
	for _name, _service := range ServiceCache.cache {
		t.Logf("service name=%s  Index=%d\n", _name, _service.Index)
		for _, _elem := range _service.List {
			t.Logf("    %v\n", _elem)
		}
	}
}
