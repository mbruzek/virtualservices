package virtualservices

import "fmt"
import "testing"

func TestParseEndpoint(t *testing.T) {
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
		relation, iface, json, err := ParseEndpoint(a.input)
		fmt.Println(relation, iface, json, err)
    if err != nil {
      t.Error(err)
    }
		if relation != a.rel || iface != a.iface || json != a.json {
			t.Errorf("ParseEndpoint did not return correct output %q, %q, %q", relation, iface, json)
		}
	}
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
    relation, iface, json, err := ParseEndpoint(b.input)
    fmt.Println(relation, iface, json, err)
    if err == nil {
      t.Errorf("ParseEndpoint() returned no error for bad input %q", b.input)
    }
  }
}
