package virtualservices

import "encoding/json"
import "fmt"
import "strings"

// VirtualEndpoint
type VirtualEndpoint struct {
	Interface string
	Relation  string
	Values    map[string]interface{}
}

// ParseVirtualEndpoint takes a single endpoint string and converts it to a
// VirtualEndpoint strcture that contains Relation, Interface, and Payload data.
// Expected format: relation:interface=JSON
// Example: website:http={"private-address":"10.0.3.1", "hostname":"10.0.3.1", "port":"6543"}
func ParseVirtualEndpoint(data string) (VirtualEndpoint, error) {
	var endpoint VirtualEndpoint

	relation_index := strings.Index(data, ":")
	if relation_index == -1 {
		return endpoint, fmt.Errorf("no relation index found in %q", data)
	}
	endpoint.Relation = strings.TrimSpace(data[:relation_index])
	if endpoint.Relation == "" {
		return endpoint, fmt.Errorf("no relation name found in %q", data)
	}

	interface_index := strings.Index(data, "=")
	if interface_index == -1 {
		return endpoint, fmt.Errorf("no interface name found in %q", data)
	}

	endpoint.Interface = strings.TrimSpace(data[relation_index+1 : interface_index])
	if endpoint.Interface == "" {
		return endpoint, fmt.Errorf("no interface name found in %q", data)
	}

	json_data := strings.TrimSpace(data[interface_index+1:])
	if err := json.Unmarshal([]byte(json_data), &endpoint.Values); err != nil {
		return endpoint, fmt.Errorf("invalid JSON: %+v", json_data)
	}

	return endpoint, nil
}
