package service

import (
	m "github.com/foadmom/appGateway/heartBeat"
)

// the key for the map is service name
var ServiceCache map[string]ServiceMapElement

func init() {
	ServiceCache = make(map[string]ServiceMapElement)
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
	var _mapElement ServiceMapElement = ServiceMapElement{_list, -1}

	_list = append(_list, service)

	ServiceCache[service.Name] = _mapElement

	return _err
}

// ============================================================================
// ============== service functions
// ============================================================================
// ========================================================
// update/add a service
// ========================================================
func Update(service m.Service) error {
	var _err error

	return _err
}

// ============================================================================
// ============== debug functions
// ============================================================================
func PrintCache () {
	for _, _service 
}