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
	for _, endpoint := range jsonData {
		if endpoint.Relation == "" {
			t.Error("the Relation cannot be empty string")
		}
		fmt.Printf("Relation: %q\n", endpoint.Relation)
		if endpoint.Interface == "" {
			t.Error("the Interface cannot be empty string")
		}
		fmt.Printf("Interface: %q\n", endpoint.Interface)

		fmt.Printf("Values: %#v\n", endpoint.Values)
		if endpoint.Values == nil || len(endpoint.Values) == 0 {
			t.Error("there are no values set.")
		}
		for key, value := range endpoint.Values {
			fmt.Printf("key: %q, value: %q\n", key, value)
		}
	}
	fmt.Println()
}

// TestParseVirtualEndpointsYAMLFile checks the file can be read and converted
// to the VirtualEndpointConfig structure.
func TestParseVirtualEndpointsYAMLFile(t *testing.T) {
	filepath := "endpoints.yaml"
	yamlData, err := ParseVirtualEndpointsYAMLFile(filepath)
	if err != nil {
		t.Error(err)
	}
	// Verify the go structure was filled in properly.
	fmt.Printf("Testing YAML: %#v\n", yamlData)
	for _, endpoint := range yamlData {
		if endpoint.Relation == "" {
			t.Error("the Relation cannot be empty string")
		}
		fmt.Printf("Relation: %q\n", endpoint.Relation)
		if endpoint.Interface == "" {
			t.Error("the Interface cannot be empty string")
		}
		fmt.Printf("Interface: %q\n", endpoint.Interface)

		fmt.Printf("Values: %#v\n", endpoint.Values)
		if endpoint.Values == nil || len(endpoint.Values) == 0 {
			t.Error("there are no values set.")
		}
		for key, value := range endpoint.Values {
			fmt.Printf("key: %q, value: %q\n", key, value)
		}
	}
	fmt.Println()
}
