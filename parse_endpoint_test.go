package virtualservices

import "fmt"
import "testing"

// TestParseEndpointPostiive checks valid input strings generate expected results.
func TestParseEndpointPositive(t *testing.T) {
	// Create a structure of endpoint strings that should pass.
	positives := []struct {
		input, rel, iface, json string
	}{
		{`relation-name:interface-name={"key":"value"}`, "relation-name", "interface-name", `{"key":"value"}`},
		{`db:postgresql={"user": "matt", "password":"GO_CODE"}`, "db", "postgresql", `{"user": "matt", "password":"GO_CODE"}`},
		{`db:mysql={"user":"mysqladmin", "password":"mysql'pass{}"}`, "db", "mysql", `{"user":"mysqladmin", "password":"mysql'pass{}"}`},
		{`relation : interface ={"user": "ubuntu", "password":"Kr@(ken][>}#"}`, "relation", "interface", `{"user": "ubuntu", "password":"Kr@(ken][>}#"}`},
		{` rel : iface = {"valid": "JSON", "more":"JSON"}`, "rel", "iface", `{"valid": "JSON", "more":"JSON"}`},
	}
	fmt.Println("Starting positive ParseVirtualEnpoint tests.")
	for _, a := range positives {
		fmt.Printf("Input: %s\n", a.input)
		endpoint, err := ParseVirtualEndpoint(a.input)
		fmt.Printf("Endpoint: %#v\n", endpoint)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("Payload: %#v\n", endpoint.Payload)
		if endpoint.Relation != a.rel || endpoint.Interface != a.iface || endpoint.Payload == nil {
			t.Errorf("incorrect output returned %q for input %q", endpoint, a.input)
		}
	}
	fmt.Println()
}

// TestParseEndpointNegative tests input strings that should fail and return error.
func TestParseEndpointNegative(t *testing.T) {
	// Create a structure of endpoint strings that should generate errors.
	negatives := []struct {
		input, rel, iface, json string
	}{
		{``, "", "", ""},
		{`relation-name:={"key": "value" }`, "relation-name", "", `{"key": "value" }`},
		{`:interface-name={"key":value}`, "", "interface-name", `{"key":value}`},
		{`relation:interface={"key":"value"`, "relation", "interface", ""},
		{`relation:interface = {"key : value"}`, "relation", "interface", ""},
	}
	fmt.Println("Staring negative ParseVirtualEndpoint tests.")
	for _, b := range negatives {
		fmt.Printf("Input: %s\n", b.input)
		endpoint, err := ParseVirtualEndpoint(b.input)
		fmt.Printf("Endpoint: %#v\n", endpoint)
		if err == nil {
			t.Errorf("no error generated for bad input %q", b.input)
		}
	}
	fmt.Println()
}
