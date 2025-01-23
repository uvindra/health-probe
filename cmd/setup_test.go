package main

import (
	"testing"
)

func TestReadYaml(t *testing.T) {
	config := readYaml()

	if len(config.Services) != 4 {
		t.Errorf("Expected 3 services, got %d", len(config.Services))
	}

	for _, service := range config.Services {
		if service.Name == "" {
			t.Errorf("Service name is empty")
		}

		if service.Host == "" {
			t.Errorf("Service host is empty")
		}

		if service.Port == 0 {
			t.Errorf("Service port is 0")
		}
	}
}

func TestSetup(t *testing.T) {
	setup()

	if len(serviceConfigLookup) != 3 {
		t.Errorf("Expected 3 services, got %d", len(serviceConfigLookup))
	}

	if len(services) != 3 {
		t.Errorf("Expected 3 services, got %d", len(services))
	}
}
