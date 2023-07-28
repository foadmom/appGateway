package service

import (
	"fmt"
	"sync"

	m "github.com/foadmom/appGateway/heartBeat"
)

type cacheType struct {
	// the key for the map is service name
	cache map[string]ServiceMapElement
	lock  sync.Mutex
}

var ServiceCache cacheType = cacheType{}

func init() {
	ServiceCache.cache = make(map[string]ServiceMapElement)
}

type ServiceMapElement struct {
	// Name  string    `json:"name"`
	List  []m.Service `json:"list"`
	Index int         `json:"index"` // for load balancing. used to indicate the last service used
}

// ============================================================================
// ============  Cache functions
// ============================================================================
// ========================================================
// Add a new service to the map
// ========================================================
func addNewService(service m.Service) error {
	var _err error
	var _list []m.Service = make([]m.Service, 0, 6)

	_list = append(_list, service)
	var _mapElement ServiceMapElement = ServiceMapElement{_list, -1}

	ServiceCache.cache[service.Name] = _mapElement

	return _err
}

// ========================================================
// add another instance to a service list
// ========================================================
func (e ServiceMapElement) addSeviceInstance(service m.Service) []m.Service {
	_list := e.List
	_list = append(_list, service)

	return _list
}

// ========================================================
// add another instance to a service list
// ========================================================
func (e ServiceMapElement) updateService(service m.Service) {
	_list := e.List
	for _index, _service := range _list {
		if (_service.Host == service.Host) &&
			(_service.Port == service.Port) {
			// this is an existing service there we should update
			// the same service status
			_list[_index] = service
			e.List = _list
			return
		}
	}
	// we did not find a match for the same host and port
	// so it must be a new instance of the same service
	e.List = e.addSeviceInstance(service)
	ServiceCache.cache[service.Name] = e
}

// ========================================================
// every so often we check for stale service entries
// if the timeStamp is older than a threshold
// remove the entry. If the service starts again
// it will register again
// ========================================================
func checkForStaleServices() {
	ServiceCache.lock.Lock()
	defer ServiceCache.lock.Unlock()

}

// ============================================================================
// ============== service functions
// ============================================================================
// ========================================================
// update/add a service
// ========================================================
func UpdateServiceInfo(service m.Service) error {
	var _err error

	ServiceCache.lock.Lock()
	defer ServiceCache.lock.Unlock()
	_elem, _found := ServiceCache.cache[service.Name]
	if _found {
		_elem.updateService(service)
	} else {
		addNewService(service)
	}
	return _err
}

func ServiceDiscovery() (m.Service, error) {
	var _err error
	var _service m.Service

	ServiceCache.lock.Lock()
	defer ServiceCache.lock.Unlock()

	return _service, _err
}

// ============================================================================
// ============== debug functions
// ============================================================================
func PrintCache() {
	for _name, _service := range ServiceCache.cache {
		fmt.Printf("service name=%s  Index=%d\n", _name, _service.Index)
		for _, _elem := range _service.List {
			fmt.Printf("    %v\n", _elem)
		}
	}
}
