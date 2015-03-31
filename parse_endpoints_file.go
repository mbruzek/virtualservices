package virtualservices

import "encoding/json"
import "gopkg.in/yaml.v1"
import "io/ioutil"

// The VirtualEndpointConfig is a data structure to use to validate when reading
// in the JSON or YAML file that contains virtual endpoint definition.
// Example:
// {"endpoints": [{"relation":"db", "interface":"mysql", "values":{"database":"juju", "user": "root", "password": "root", "private-address": "172.17.42.1"}}]}
type VirtualEndpointConfig struct {
	Endpoints []struct {
		Interface string                 `json:"interface" yaml:"interface"`
		Relation  string                 `json:"relation" yaml:"relation"`
		Values    map[string]interface{} `json:"values" yaml:"values"`
	} `json:"endpoints" yaml:"endpoints"`
}

// ParseGenericJSONFile reads in the JSON file converting the content to a
// map, slice of maps.
func ParseGenericJSONFile(filepath string) (map[string][]map[string]interface{}, error) {
	var genericData map[string][]map[string]interface{}

	contents, err := ioutil.ReadFile(filepath)
	if err != nil {
		return genericData, err
	}
	if err := json.Unmarshal([]byte(contents), &genericData); err != nil {
		return genericData, err
	}
	return genericData, nil
}

// parseVirtualEndpointsJSONFile reads in the JSON file converting the content
// to the VirtualEndpoint struct.
func ParseVirtualEndpointsJSONFile(filepath string) ([]VirtualEndpoint, error) {
	var jsonData VirtualEndpointConfig
	contents, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(contents, &jsonData); err != nil {
		return nil, err
	}
	var virtualEndpoints []VirtualEndpoint
	for a := range jsonData.Endpoints {
		var endpoint VirtualEndpoint
		endpoint.Relation = jsonData.Endpoints[a].Relation
		endpoint.Interface = jsonData.Endpoints[a].Interface
		endpoint.Values = jsonData.Endpoints[a].Values
		virtualEndpoints = append(virtualEndpoints, endpoint)
	}
	return virtualEndpoints, nil
}

// parseVirtualEndpointsYAMLFile reads in the YAML file converting the content
// to the VirtualEndpointConfig struct.
func ParseVirtualEndpointsYAMLFile(filepath string) ([]VirtualEndpoint, error) {
	var yamlData VirtualEndpointConfig
	contents, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(contents, &yamlData); err != nil {
		return nil, err
	}
	var virtualEndpoints []VirtualEndpoint
	for a := range yamlData.Endpoints {
		var endpoint VirtualEndpoint
		endpoint.Relation = yamlData.Endpoints[a].Relation
		endpoint.Interface = yamlData.Endpoints[a].Interface
		endpoint.Values = yamlData.Endpoints[a].Values
		virtualEndpoints = append(virtualEndpoints, endpoint)
	}
	return virtualEndpoints, nil
}
