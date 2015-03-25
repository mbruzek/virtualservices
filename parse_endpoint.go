package virtualservices

import "fmt"
import "encoding/json"
import "strings"

// Take a sigle enpoint string and parse out the relation, interface, and JSON data.
// relation:interface=JSON
func ParseEndpoint(str string) (string, string, string, error) {
	relation_index := strings.Index(str, ":")
	if relation_index == -1 {
		return "", "", "", fmt.Errorf("No relation index found in %q", str)
	}
	relation := str[:relation_index]
	fmt.Print("Relation = ")
	fmt.Println(relation)

	splits := strings.Split(str[relation_index+1:], "=")
	if len(splits) < 2 || splits[0] == "" {
		return relation, "", "", fmt.Errorf("No interface name found in %q", str)
	}
	iface := splits[0]
	fmt.Print("Interface = ")
	fmt.Println(iface)

	json_data := splits[1]

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(json_data), &data); err != nil {
		return relation, iface, "", fmt.Errorf("Invalid JSON: %q", json_data)
	}
	fmt.Print("JSON data = ")
	fmt.Println(data)

	return relation, iface, json_data, nil
}
