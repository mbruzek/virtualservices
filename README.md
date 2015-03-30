# virtualservices

This is a go package for the Juju virtual services work.

To define a virtual service you must have the relation name and interface
name with the key value pairs that you want to set.

The `parse_endpoint.go` file was the first attempt to parse this information
on the command line, in a format of `relation:interface={"key", "value"}`

The plan changed to deploy the virtual service with a "config" file.

The configuration file could be YAML or JSON format. See `endpoints.yaml`
and `endpoints.json` for more details on the file format.

