package main

import (
	"fmt"
	"testing"
)

func TestParseIni(t *testing.T) {
	actual, err := parseInI("PHP.ini")
	if err != nil {
		t.Error(fmt.Sprintf("Error in parsing: '%v'", err))
	}
	expected := make(map[string]map[string]string)
	expected["owner"] = make(map[string]string)
	expected["owner"]["name"] = "John Doe"
	expected["owner"]["organization"] = "Acme Widgets Inc"
	expected["database"] = make(map[string]string)
	expected["database"]["server"] = "42"
	expected["database"]["port"] = "143"
	expected["database"]["file"] = "payroll.dat"
	if !compareMapMap(actual, expected) {
		t.Error(fmt.Sprintf("Expected '%v', Actual '%v'", expected, actual))
	}
}

func compareMapMap(a, b map[string]map[string]string) bool {

	if len(a) != len(b) {
		return false
	}
	for key, value := range a {
		if !compareMap(value, b[key]) {
			return false
		}
	}
	for key, value := range b {
		if !compareMap(value, a[key]) {
			return false
		}
	}
	return true
}

func compareMap(a, b map[string]string) bool {

	if len(a) != len(b) {
		return false
	}
	for key, value := range a {
		if value != b[key] {
			return false
		}
	}
	for key, value := range b {
		if value != a[key] {
			return false
		}
	}
	return true
}
