package virtualservices

import "fmt"
import "testing"

// TestParseVirtualEndpointsJSONFile checks the file can be read and converted
// to the VirtualEndpointConfig structure.
func TestParseVirtualEndpointsJSONFile(t *testing.T) {
	filepath := "endpoints.json"
	jsonData, err := ParseVirtualEndpointsJSONFile(filepath)
	if err != nil {
		t.Error(err)
	}
	// Verify the go structure was filled in properly.
	fmt.Printf("Testing JSON: %#v\n", jsonData)
	if jsonData.Endpoints == nil || len(jsonData.Endpoints) == 0 {
		t.Error("the endpoints array should not be empty")
	}
	fmt.Printf("Endpoints: %q\n", jsonData.Endpoints)
	for a := range jsonData.Endpoints {
		endpoint := jsonData.Endpoints[a]
		if endpoint.Relation == "" {
			t.Error("the Relation cannot be empty string")
		}
		fmt.Printf("Relation: %q\n", endpoint.Relation)
		if endpoint.Interface == "" {
			t.Error("the Interface cannot be empty string")
		}
		fmt.Printf("Interface: %q\n", endpoint.Interface)
		if len(endpoint.Values) == 0 {
			t.Error("the Values cannot be empty")
		}
		fmt.Printf("Values: %#v\n", endpoint.Values)
		for b := range endpoint.Values {
			for key, value := range endpoint.Values[b] {
				fmt.Printf("key: %q, value: %q\n", key, value)
			}
		}
	}
	fmt.Println()
}

// TestParseVirtualEndpointsYAMLFile checks the file can be read and converted
// to the VirtualEndpointConfig structure.
func TestParseVirtualEndpointsYAMLFile(t *testing.T) {
	filepath := "endpoints.yaml"
	jsonData, err := ParseVirtualEndpointsYAMLFile(filepath)
	if err != nil {
		t.Error(err)
	}
	// Verify the go structure was filled in properly.
	fmt.Printf("Testing YAML: %#v\n", jsonData)
	if jsonData.Endpoints == nil || len(jsonData.Endpoints) == 0 {
		t.Error("the endpoints array should not be empty")
	}
	fmt.Printf("Endpoints: %q\n", jsonData.Endpoints)
	for a := range jsonData.Endpoints {
		endpoint := jsonData.Endpoints[a]
		if endpoint.Relation == "" {
			t.Error("the Relation cannot be empty string")
		}
		fmt.Printf("Relation: %q\n", endpoint.Relation)
		if endpoint.Interface == "" {
			t.Error("the Interface cannot be empty string")
		}
		fmt.Printf("Interface: %q\n", endpoint.Interface)
		if len(endpoint.Values) == 0 {
			t.Error("the Values cannot be empty")
		}
		fmt.Printf("Values: %#v\n", endpoint.Values)
		for b := range endpoint.Values {
			for key, value := range endpoint.Values[b] {
				fmt.Printf("key: %q, value: %q\n", key, value)
			}
		}
	}
	fmt.Println()
}
