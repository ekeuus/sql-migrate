//go:build oracle
// +build oracle

package main

import (
	_ "github.com/mattn/go-oci8"
	migrate "github.com/ekeuus/sql-migrate"
)

func init() {
	dialects["oci8"] = migrate.OracleDialect{}
}
