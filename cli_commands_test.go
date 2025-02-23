package main

import (
	"testing"
)

func TestExplore(t *testing.T) {
	conf := config{}
	err := commandExplore(&conf)

	if err == nil {
		t.Errorf("Expected error message for missing location, got nil")
	}
}
