// Package grpc/schema provides generator for GRPC gateway's schema
package schema

import (
	"github.com/SysUtils/1c-gateway/shared"
	"os"
)

type Generator struct {
	TypeMap      map[string]string
	NameMap      map[string]string
	Associations map[string]map[string]string
	schema       shared.Schema
}

// NewGenerator returns initialized generator
func NewGenerator(schema shared.Schema) *Generator {
	return &Generator{schema: schema, TypeMap: make(map[string]string), NameMap: make(map[string]string), Associations: map[string]map[string]string{}}
}

// Generate generates the grpc schema and writes it to ./odata folder
func (g *Generator) Generate() {
	f, _ := os.Create("odata/grpc.proto")
	f.WriteString(`syntax = "proto3";

package protobuf;
option go_package = "odata";

message Timestamp {
	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	int64 seconds = 1;

	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	int32 nanos = 2;
}

message BaseWhere {
	int32 Top = 1;
	int32 Skip = 2;
	string Orderby = 3;
}
`)
	f.WriteString(g.genMessages(g.schema.Entities))
	f.WriteString("\n")
	f.WriteString(g.genMessages(g.schema.Complexes))
	f.WriteString("\n")
	f.WriteString(g.genService(g.schema.Entities))
	f.Close()

}
