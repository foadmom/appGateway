package service

import (
	"fmt"
	"testing"
	"time"

	m "github.com/foadmom/appGateway/heartBeat"
)

func Test_addNewService(t *testing.T) {

	var _delay time.Duration = 2 * time.Second
	var _service_A m.Service = m.Service{Name: "Service-A", Host: "localHost", Port: 7666, LastUpdated: time.Now()}
	time.Sleep(_delay)
	var _service_B m.Service = m.Service{Name: "Service-B", Host: "localHost", Port: 7555, LastUpdated: time.Now()}
	time.Sleep(_delay)
	var _service_C m.Service = m.Service{Name: "Service-C", Host: "localHost", Port: 7444, LastUpdated: time.Now()}
	time.Sleep(_delay)
	var _service_B2 m.Service = m.Service{Name: "Service-B", Host: "localHost", Port: 7555, LastUpdated: time.Now()}
	time.Sleep(_delay)
	var _service_A2 m.Service = m.Service{Name: "Service-A", Host: "localHost", Port: 5555, LastUpdated: time.Now()}
	var _service_A3 m.Service = m.Service{Name: "Service-A", Host: "localHost", Port: 5556, LastUpdated: time.Now()}
	var _service_A4 m.Service = m.Service{Name: "Service-A", Host: "localHost", Port: 5557, LastUpdated: time.Now()}
	var _service_A5 m.Service = m.Service{Name: "Service-A", Host: "localHost", Port: 5558, LastUpdated: time.Now()}
	var _service_A6 m.Service = m.Service{Name: "Service-A", Host: "localHost", Port: 5565, LastUpdated: time.Now()}
	var _service_A7 m.Service = m.Service{Name: "Service-A", Host: "localHost", Port: 5575, LastUpdated: time.Now()}
	var _service_A8 m.Service = m.Service{Name: "Service-A", Host: "localHost", Port: 5585, LastUpdated: time.Now()}
	var _service_A9 m.Service = m.Service{Name: "Service-A", Host: "localHost", Port: 5595, LastUpdated: time.Now()}

	UpdateServiceInfo(_service_A)
	UpdateServiceInfo(_service_B)
	UpdateServiceInfo(_service_C)
	UpdateServiceInfo(_service_A2)
	UpdateServiceInfo(_service_B2)
	UpdateServiceInfo(_service_A3)
	UpdateServiceInfo(_service_A4)
	UpdateServiceInfo(_service_A5)
	UpdateServiceInfo(_service_A6)
	UpdateServiceInfo(_service_A7)
	UpdateServiceInfo(_service_A8)
	UpdateServiceInfo(_service_A9)

	PrintCache()

	var _staleTimeDelay time.Duration = 500 * time.Millisecond
	var i int = 12
	for {
		time.Sleep(_staleTimeDelay)
		checkForStaleServices()
		printCache(t)

		fmt.Println()
		i--
		if i == 0 {
			break
		}
	}

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
