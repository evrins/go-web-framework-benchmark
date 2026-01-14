package main

import (
	"fmt"
	"testing"

	"github.com/smallnest/go-web-framework-benchmark/server"
)

func TestGenAllNames(t *testing.T) {
	for key := range server.BFMap {
		fmt.Printf("\"%s\" \\\n", key)
	}
}
