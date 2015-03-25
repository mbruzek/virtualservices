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
	for _, a := range positives {
		fmt.Println(a.input)
		endpoint, err := ParseEndpoint(a.input)
		fmt.Println(endpoint, err)
		if err != nil {
			t.Error(err)
		}
		if endpoint.relation != a.rel || endpoint.iface != a.iface || endpoint.data != a.json {
			t.Errorf("incorrect output returned %q for input %q", endpoint, a.input)
		}
	}
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
	for _, b := range negatives {
		fmt.Println(b.input)
		endpoint, err := ParseEndpoint(b.input)
		fmt.Println(endpoint, err)
		if err == nil {
			t.Errorf("no error generated for bad input %q", b.input)
		}
	}
}
