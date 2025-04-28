package main

// NOTE: To run go generate .

//go:generate go tool oapi-codegen -config api/oapi-codegen.yaml api/openapi.yaml
//go:generate go tool sqlc generate -f sql/sqlc.yaml
