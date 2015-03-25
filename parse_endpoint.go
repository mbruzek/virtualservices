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
	relation := strings.TrimSpace(str[:relation_index])
	if relation == "" {
		return "", "", "", fmt.Errorf("No relation name found in %q", str)
	}

	splits := strings.Split(str[relation_index+1:], "=")
	if len(splits) < 2 {
		return relation, "", "", fmt.Errorf("No interface name found in %q", str)
	}
	iface := strings.TrimSpace(splits[0])
	if iface == "" {
		return relation, "", "", fmt.Errorf("No interface name found in %q", str)
	}

	json_data := strings.TrimSpace(splits[1])

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(json_data), &data); err != nil {
		return relation, iface, "", fmt.Errorf("Invalid JSON: %q", json_data)
	}

	return relation, iface, json_data, nil
}
