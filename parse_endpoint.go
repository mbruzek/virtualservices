package virtualservices

import "encoding/json"
import "fmt"
import "strings"

type Endpoint struct {
	relation string
	iface    string
	data     string
}

// ParseEndpoint takes a sigle endpoint string and parses out the relation, interface, and JSON data.
// relation:interface=JSON
func ParseEndpoint(str string) (Endpoint, error) {
	var endpoint Endpoint
	relation_index := strings.Index(str, ":")
	if relation_index == -1 {
		return endpoint, fmt.Errorf("no relation index found in %q", str)
	}
	endpoint.relation = strings.TrimSpace(str[:relation_index])
	if endpoint.relation == "" {
		return endpoint, fmt.Errorf("no relation name found in %q", str)
	}

	interface_index := strings.Index(str, "=")
	if interface_index == -1 {
		return endpoint, fmt.Errorf("no interface name found in %q", str)
	}
	endpoint.iface = strings.TrimSpace(str[relation_index+1:interface_index])
	if endpoint.iface == "" {
		return endpoint, fmt.Errorf("no interface name found in %q", str)
	}
	// The JSON data to return is the string, not the map.
	endpoint.data = strings.TrimSpace(str[interface_index+1:])

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(endpoint.data), &data); err != nil {
		return endpoint, fmt.Errorf("invalid JSON: %q", endpoint.data)
	}

	return endpoint, nil
}
