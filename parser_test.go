package main

import (
	"fmt"
	"testing"
)

func TestParseIni(t *testing.T) {
	actual, err := parse_ini("PHP.ini")
	if err != nil {
		t.Error(fmt.Sprintf("Error in parsing: '%v'", err))
	}
	expected := make(map[string]map[string]string)
	expected["owner"] = make(map[string]string)
	expected["owner"]["name"] = "John Doe"
	expected["owner"]["organization"] = "Acme Widgets Inc."
	expected["database"] = make(map[string]string)
	expected["database"]["server"] = "192.0.2.62"
	expected["database"]["port"] = "143"
	expected["database"]["file"] = "payroll.dat"
	if !compareMapMap(actual, expected) {
		t.Error(fmt.Sprintf("Expected : '%v', Actual : '%v'", expected, actual))
	}
}

func getSectionsTest(t *testing.T) {
	nested := map[string]map[string]string{
		"owner": {
			"name":         "John Doe",
			"organization": "Acme Widgets Inc.",
		},
		"database": {
			"file":   "payroll.dat",
			"port":   "143",
			"server": "129.0.2.62",
		},
	}

	expected := map[string]map[string]string{
		"database": {
			"server": "129.0.2.62",
			"port":   "143",
			"file":   "payroll.dat",
		},
	}

	actual, err := getSections("database", nested)

	if err != nil {
		t.Error(fmt.Sprintf("Expected '%v', Actual '%v'", expected, actual))
	}
}

func setTest(t *testing.T) {
	//map[file:payroll.dat port:500 server:192.0.2.62]
	//set("database", "port", "500", x)
	nested := map[string]map[string]string{
		"owner": {
			"name":         "John Doe",
			"organization": "Acme Widgets Inc.",
		},
		"database": {
			"file":   "payroll.dat",
			"port":   "143",
			"server": "129.0.2.62",
		},
	}

	expected := map[string]string{
		"file":   "payroll.dat",
		"port":   "500",
		"server": "192.0.2.62",
	}

	actual := set("database", "port", "500", nested)

	if !compareMap(actual, expected) {
		t.Error(fmt.Sprintf("Expected '%v', Actual '%v'", expected, actual))
	}

}

func getKeysTest(t *testing.T) {

	expected := map[string]string{
		"server": "192.0.2.62",
		"port":   "143",
		"file":   "payroll.dat",
	}

	nested := map[string]map[string]string{
		"owner": {
			"name":         "John Doe",
			"organization": "Acme Widgets Inc.",
		},
		"database": {
			"file":   "payroll.dat",
			"port":   "143",
			"server": "129.0.2.62",
		},
	}

	err := getKeys("database", nested)

	if err != nil {
		t.Error(fmt.Sprintf("Expected '%v' ", expected))
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
