package service

import (
	"errors"
	"fmt"
	"sync"
	"time"

	m "github.com/foadmom/appGateway/heartBeat"
)

type cacheType struct {
	// the key for the map is service name
	cache map[string]ServiceMapElement
	lock  sync.Mutex
}

var ServiceCache cacheType = cacheType{}

type ServiceMapElement struct {
	// Name  string    `json:"name"`
	List  []m.Service `json:"list"`
	Index int         `json:"index"` // for load balancing. used to indicate the last service used
}

const checkInterval = 500 * time.Millisecond
const (
	ACTIVE = "active"
)

func init() {
	ServiceCache.cache = make(map[string]ServiceMapElement)
	go periodicCacheCheck()
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
			// this is an existing service, so update
			// the service info
			_list[_index] = service
			return
		}
	}
	// we did not find a match for the same host and port
	// so it must be a new instance of the same service
	e.List = e.addSeviceInstance(service)
	ServiceCache.cache[service.Name] = e
}

// ========================================================
// removes the last item in the slice and copies it into
// the spot for the service you wish to delete, then
// returns the slice from top to end -1
// ========================================================
func removeService(s []m.Service, index int) []m.Service {
	var _len int = len(s)
	s[index] = s[_len-1]
	return s[:_len-1]
}

// ========================================================
// every so often we check for stale service entries
// if the timeStamp is older than a threshold
// remove the entry. If the service starts again
// it will register again
// ========================================================
func periodicCacheCheck() {
	fmt.Println("******** periodicCacheCheck started")
	for {
		checkCache()
		time.Sleep(checkInterval)
	}
}

func checkCache() {
	ServiceCache.lock.Lock()
	defer ServiceCache.lock.Unlock()

	for _name, _serviceMapElement := range ServiceCache.cache {
		checkForStaleServices(_serviceMapElement)
		if len(ServiceCache.cache[_name].List) == 0 {
			delete(ServiceCache.cache, _name)
		}
	}
}

// ========================================================
// every so often we check for stale service entries
// if the timeStamp is older than a threshold
// remove the entry. If the service starts again
// it will register again
// ========================================================
func checkForStaleServices(serviceMapElem ServiceMapElement) bool {
	var _staleServiceFound bool = false

	var _now time.Time = time.Now()
	if len(serviceMapElem.List) > 0 {
		for _index, _elem := range serviceMapElem.List {
			_diff := _now.Sub(_elem.LastUpdated).Seconds()
			if _diff > 5 {
				var _serviceName string = _elem.Name
				serviceMapElem.List = removeService(serviceMapElem.List, _index)
				ServiceCache.cache[_serviceName] = serviceMapElem
				_staleServiceFound = true
				break
			}
		}
	}
	return _staleServiceFound
}

func nextAvailableService(serviceMap ServiceMapElement) m.Service {
	var _service m.Service
	var _index int

	if serviceMap.Index == len(serviceMap.List)-1 {
		// wrap the index around
		serviceMap.Index = 0
	}
	for _index, _service = range serviceMap.List {
		if _service.Status == ACTIVE && _index > serviceMap.Index {
			return _service
		}
	}

	return _service
}

func getNextAvailableService(name string) (m.Service, error) {
	var _err error = errors.New("No service found for " + name)
	var _service m.Service

	_serviceMap, _found := ServiceCache.cache[name]
	if _found == true {
		_len := len(_serviceMap.List)
		if _len > 0 {
			if _serviceMap.Index == -1 {
				//                _service =
			}
		} else {
			_err = fmt.Errorf("No Active service available for %s", name)
		}
	}

	return _service, _err
}

// ============================================================================
// ============== service (External) functions
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
		_err = addNewService(service)
	}
	return _err
}

func ServiceDiscovery(name string) (m.Service, error) {
	var _err error
	var _service m.Service

	ServiceCache.lock.Lock()
	defer ServiceCache.lock.Unlock()

	_service, _err = getNextAvailableService(name)

	return _service, _err
}

// ============================================================================
// ============== debug functions
// ============================================================================
func PrintCache() {
	ServiceCache.lock.Lock()
	defer ServiceCache.lock.Unlock()
	for _name, _service := range ServiceCache.cache {
		fmt.Printf("service name=%s  Index=%d\n", _name, _service.Index)
		for _, _elem := range _service.List {
			fmt.Printf("    %v\n", _elem)
		}
	}
}
